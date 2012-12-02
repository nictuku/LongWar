package longwar

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/jasondelponte/golib/seqgen/perlinnoise"
	mrand "math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	lastTerrain = 5

	// not sure if these are what I named them for.
	interpolations = 3
	amplitude      = 10
	z              = 0.5
)

func init() {
	http.HandleFunc("/rnd", RandomMap)
	http.HandleFunc("/createmap", CreateMap)
	mrand.Seed(time.Now().UnixNano())
}

// tileMap for tile generation, to control usage of the tiles without having to change the image file directly. Terrains in the middle are more common.
var tileMap []int64 = []int64{
	3, // sand
	3,
	3,
	3,
	3,
	5, // savannah
	5,
	0, // dirt
	0,
	0,
	0,
	4, // forest
	4,
	4,
	// 6, // water
	2, // dark-forest
	2,
	1, // mountain
	1,
	1,
	1,
}

func perlinTiles(x, y float64, seed int64) []int64 {

	h := 30
	w := 30
	tiles := make([]int64, h*w)

	var pn *perlinnoise.PerlinNoise
	if seed > 0 {
		pn = perlinnoise.New(seed)
	} else {
		pn = perlinnoise.NewDefault()
	}

	for i := 0; i < h; i++ { // y
		y := float64(i) / float64(h)
		for j := 0; j < w; j++ { // x
			x := float64(j) / float64(w)

			n := 0.
			for ix := 0; ix < interpolations; ix++ {
				n += pn.Noise(amplitude*x, amplitude*y, z+interpolations) // [-1, 1]
			}
			n = n / interpolations
			f := float64(len(tileMap)) * (n + 1) / 2
			tiles[(i*w)+j] = tileMap[Round(f)]
		}
	}
	return tiles
}

// Generate a randomized 30*30 map.
// I could save space by returning a more compressed result (e.g: multiple tiles per byte), but this will do.
func CreateMap(w http.ResponseWriter, r *http.Request) {

	seed, _ := strconv.ParseInt(r.URL.Query().Get("seed"), 10, 0)
	if seed == 0 {
		seed = mrand.Int63()
	}

	// jsonp is merely a json file with a wrapping function.
	callback := r.URL.Query().Get("jsoncallback")
	if callback == "" {
		callback = "jsonpCallback"
	}

	var err error
	tiles := perlinTiles(30, 30, seed)

	b, err := json.Marshal(tiles)
	if err != nil {
		http.Error(w, "json:"+err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/javascript")
	fmt.Fprintf(w, "%v(%v)", callback, string(b))
}

// Generate a randomized 30*30 map.
// I could save space by returning a more compressed result (e.g: multiple tiles per byte), but this will do.
func RandomMap(w http.ResponseWriter, r *http.Request) {
	rnd := make([]byte, 30*30)
	var err error

	for items := 0; items < 30*30-1; {
		if c, err := rand.Read(rnd[items:]); err != nil {
			http.Error(w, "rand:"+err.Error(), http.StatusInternalServerError)
		} else {
			items += c
		}
	}
	// Slow copy. I have to convert this to int because a slice of byte is a binary blob for json.
	tiles := make([]int, 30*30)
	for i, v := range rnd {
		tiles[i] = int(v) % lastTerrain
	}

	b, err := json.Marshal(tiles)
	if err != nil {
		http.Error(w, "json:"+err.Error(), http.StatusInternalServerError)
	}

	// jsonp is merely a json file with a wrapping function.
	callback := r.URL.Query().Get("jsoncallback")
	if callback == "" {
		callback = "jsonpCallback"
	}
	w.Header().Set("Content-Type", "text/javascript")
	fmt.Fprintf(w, "%v(%v)", callback, string(b))
}

func Round(value float64) int64 {
	if value < 0.0 {
		value -= 0.5
	} else {
		value += 0.5
	}
	return int64(value)
}

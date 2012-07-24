package longwar

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// Valid tiles range between zero and `lastTerrain`.
const lastTerrain = 5

func init() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/rnd", RandomMap)

}

func Home(w http.ResponseWriter, r *http.Request) {
	err := homeTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

var homeTemplate = template.Must(template.New("home").Parse(homeTemplateHTML))

const homeTemplateHTML = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
	<script type="text/javascript" src="https://ajax.googleapis.com/ajax/libs/jquery/1.7.2/jquery.min.js"></script>
	<script type="text/javascript" src="../js/thirdparty/crafty.js"></script>
	<script type="text/javascript" src="../js/game.js"></script>
	<title>LongWar</title>
	<style>
	body, html { margin:0; padding: 0; overflow:hidden }
	</style>
</head>
<body>
</body>
</html>
`

// Generate a randomized 30*30 map.
// I could save space by returning a more compressed result (e.g: multiple tiles per byte), but this will do.
func RandomMap(w http.ResponseWriter, r *http.Request) {
	rnd := make([]byte, 30*30)
	var err error
	// Fill up all 30*30 tiles with random numbers.
	// I could speed up things by using a single byte for two tiles (4 bits each), but this will do.
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
		tiles[i] = int(v)%lastTerrain + 1
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

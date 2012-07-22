package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Valid tiles range between zero and `lastTerrain`.
const lastTerrain = 5

func main() {

	// Generate a randomized 30*30 map.
	// I could save space by returning a more compressed result (e.g: multiple tiles per byte), but this will do.
	http.HandleFunc("/rnd", func(w http.ResponseWriter, r *http.Request) {
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
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

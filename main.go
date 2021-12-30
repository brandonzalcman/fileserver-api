package main

import (
	"net/http"
)

func urlHandler(w http.ResponseWriter, r *http.Request) {
	back := r.URL.Query().Get("back")
	song := r.URL.Query().Get("song")
	if len(back) != 0 {
		http.ServeFile(w, r, "assets/backgrounds/"+back)
	} else if len(song) != 0 {
		http.ServeFile(w, r, "assets/songs/"+song)
	} else {
		http.NotFound(w, r)
		return
	}
}

func main() {
	http.HandleFunc("/assets", urlHandler)
	http.ListenAndServe(":1000", nil)
}

//test

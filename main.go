package main

import (
	"net/http"
)

//Handler to get the url values, using the value to serve the files
func urlHandler(w http.ResponseWriter, r *http.Request) {
	back := r.URL.Query().Get("back") //Allocated for backgrounds
	song := r.URL.Query().Get("song") //Allocated for songs
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
	http.HandleFunc("/assets", urlHandler) //Set /assets as route
	http.ListenAndServe(":1000", nil)      //Listens on port 1000 | CUSTOMISE HERE
}

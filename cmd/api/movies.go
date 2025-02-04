package main

import (
	"fmt"
	"net/http"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "creating a movie")
}

func (app *application) getMovieDetailHandler(w http.ResponseWriter, r *http.Request) {
	id := app.readIDParam(r)
	fmt.Fprintf(w, "the movie's id is: %d", id)
}

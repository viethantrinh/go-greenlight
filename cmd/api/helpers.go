package main

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) int64 {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
	return id
}
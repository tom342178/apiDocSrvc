package apiDocSrvc

import (
	"embed"
	"github.com/gorilla/mux"
	"io/fs"
	"net/http"
)

//TODO: REMOVE THIS TODO

//go:embed embed
var swagfs embed.FS

func byteHandler(b []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Write(b)
	}
}

func addSwaggerRoute(spec []byte, r *mux.Router) {
	// render the index template with the proper spec name inserted
	swFS, _ := fs.Sub(swagfs, "embed")
	r.HandleFunc("/swagger_spec", byteHandler(spec))
	// MUST BE LAST ROUTE
	r.PathPrefix("/").Handler(http.FileServer(http.FS(swFS)))

	return
}

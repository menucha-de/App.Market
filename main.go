/*
 * App market
 *
 * API version: 0.0.1
 * Contact: support@peraMIC.io
 */

package main

import (
	"flag"
	"fmt"
	"html"
	"net/http"

	"github.com/peramic/logging"
	"github.com/peramic/utils"

	"github.com/peramic/App.Market/market"
)

var log *logging.Logger = logging.GetLogger("market")

func main() {
	var port = flag.Int("p", 8080, "port")
	flag.Parse()

	market.AddRoutes(logging.LogRoutes)
	market.AddRoutes(market.Routes)
	router := market.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(notFound)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: router,
	}

	log.Info("Server Started on port ", *port)

	log.Fatal(srv.ListenAndServe())
}

func notFound(w http.ResponseWriter, r *http.Request) {
	if r.Method == "HEAD" {
		w.WriteHeader(204)
		return
	}
	if !(r.Method == "GET") {
		w.WriteHeader(404)
	}
	file := "./www" + html.EscapeString(r.URL.Path)
	if file == "./www/" {
		file = "./www/index.html"
	}
	if utils.FileExists(file) {
		http.ServeFile(w, r, file)
	} else {
		w.WriteHeader(404)
	}
}

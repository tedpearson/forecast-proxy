package web

import (
	"flag"
	"github.com/tedpearson/forecast-proxy/forecast"
	"log"
	"net/http"
	"net/http/httputil"
)

type Config struct {
	Port      int
	BingToken string
}

func Main() {
	port := flag.String("port", "8080", "Port to listen on")
	bingMapsToken := flag.String("token", "", "bing maps token")
	flag.Parse()
	// start http server
	// handler that will do a prometheus format
	// handler will call forecast functions
	h := &Handler{
		forcaster: &forecast.Forecaster{
			BingToken: *bingMapsToken,
		},
	}
	err := http.ListenAndServe(":"+*port, h)
	log.Println(err)
}

type Handler struct {
	forcaster *forecast.Forecaster
}

func (h *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	rs, err := httputil.DumpRequest(req, true)
	if err == nil {
		log.Println(string(rs))
	}
	_, err = resp.Write([]byte("hello"))
	if err != nil {
		log.Println("Error writing response")
	}

	loc, err := h.forcaster.GetLocation("Washington Monument")
	log.Println(loc)
}

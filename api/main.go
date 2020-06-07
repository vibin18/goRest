package main

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/vibin18/goRest/mirrors"
	"github.com/vibin18/goRest/weather"
	log "github.com/sirupsen/logrus"


)



type response struct {
	FastestURL string `json:"fastest_url"`
	Latency time.Duration `json:"latency"`
}

func main() {
	log.SetFormatter(&log.TextFormatter{DisableColors: true,FullTimestamp: true,})
	http.HandleFunc("/fastest-mirror", func(w http.ResponseWriter,
		r *http.Request) {
		log.Info("Request for /fastest-mirror")
		response := findFastest(mirrors.MirrorList)
		respJSON, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
	})


	http.HandleFunc("/weather", func(w http.ResponseWriter,
		r *http.Request) {
		log.Info("Request for /weather")
		Wresponse := weather.GetWeatherAndLocation()
		WrespJSON, _ := json.Marshal(Wresponse)
		w.Header().Set("Content-Type", "application/json")
		w.Write(WrespJSON)
	})


	port := ":8000"
	server := &http.Server{
		Addr: port,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Info("Starting server on port",port)
	log.Fatal(server.ListenAndServe())
}

func findFastest(urls []string) response {
	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)
	for _, url := range urls {
		mirrorURL := url
		go func() {
			start := time.Now()
			_, err := http.Get(mirrorURL + "/README")
			latency := time.Now().Sub(start) / time.Millisecond
			if err == nil {
				urlChan <- mirrorURL
				latencyChan <- latency
			}
		}()
	}
	return response{<-urlChan, <-latencyChan}
}


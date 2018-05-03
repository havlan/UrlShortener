package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"syscall"
)

const (
	PORT        = 8080
	BASEURL     = "localhost:8080/"
	HTTPSPREFIX = "https://" // use this to force https.
	HTTPPREFIX  = "http://"
)

func UrlHandler(storage Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			fmt.Printf("Method %s attempted.", r.Method)
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		} else {
			var ret urlEntity

			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&ret)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			shortUrl, err := storage.New(ret.URL)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write([]byte(BASEURL + shortUrl))
			return
		}
	}
}

func WildHandler(storage Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var urlStripped = r.URL.Path[1:]
		data, err := storage.Get(urlStripped)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if strings.HasPrefix("http://", data) || strings.HasPrefix("https://", data) {
			http.Redirect(w, r, data, 301)
		} else {
			http.Redirect(w, r, HTTPSPREFIX+data, 301)
		}
		return
	}
}

func main() {
	/*
		c := cron.New()

		c.AddFunc("@every 1m", func() { fmt.Println("Helu") })
		go c.Start()
	*/

	var cfg, err = NewConfig("config.json")
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}

	dbStruct, err := NewDb(cfg)
	if err != nil {
		fmt.Printf("%s", err.Error())
		syscall.Exit(1)
	}
	//dbStruct.Init("WEB_URL.sql")

	s := http.NewServeMux()
	s.HandleFunc("/url", UrlHandler(dbStruct))
	s.HandleFunc("/", WildHandler(dbStruct))

	logger := log.New(os.Stdout, "", log.LstdFlags)
	logger.Printf("Starting webserver and listen on %d", PORT)

	// Start HTTP Server with request logging
	loggingHandler := handlers.LoggingHandler(os.Stdout, s)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), loggingHandler))
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/hoisie/redis"
	"github.com/julienschmidt/httprouter"
)

func main() {
	port := os.Getenv("MetaPort")
	if port == "" {
		port = "9820"
	}

	r := httprouter.New()

	r.GET("/v2/repositories/:library/:name", GetInfoHandler)
	r.GET("/v2/repositories/:library/:name/tags", GetInfoHandler)
	r.OPTIONS("/v2/repositories/:library/:name", OptInfoHandler)
	r.OPTIONS("/v2/repositories/:library/:name/tags", OptInfoHandler)

	http.ListenAndServe(":"+port, r)
}

func GetInfoHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
        rw.Header().Set("Access-Control-Allow-Origin", "*")
        rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
        rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With , yourHeaderFeild")
        rw.Header().Set("Content-Type", "application/json")
	var client redis.Client
	body, err := client.Get(r.RequestURI)
	if err != nil {
		resp, err := http.Get("https://hub.docker.com/" + r.RequestURI)
		defer resp.Body.Close()
		if err != nil {
			fmt.Println(r.RequestURI + " failed to get")
			rw.WriteHeader(http.StatusNotFound)
		}

                rw.Header().Set("Access-Control-Allow-Origin", "*")
                rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
                rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With , yourHeaderFeild")
		rw.Header().Set("Content-Type", "application/json")
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(r.RequestURI + " failed to read")
			rw.WriteHeader(http.StatusInternalServerError)
		}
		client.Setex(r.RequestURI, 3600*24*7, body)
	}

	rw.Write(body)
}
func OptInfoHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
        rw.Header().Set("Access-Control-Allow-Headers", "X-PINGOTHER")
}

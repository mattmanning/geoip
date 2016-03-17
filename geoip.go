package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oschwald/geoip2-golang"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/city/{ip}", Lookup)
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":80", nil))
}

func Lookup(w http.ResponseWriter, r *http.Request) {
	ipStr := mux.Vars(r)["ip"]
	ip := net.ParseIP(ipStr)

	if ip == nil {
		http.Error(w, "IP address malformed", 400)
		return
	}

	db, err := geoip2.Open("GeoLite2-City.mmdb")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	record, err := db.City(ip)

	if err != nil {
		log.Fatal(err)
	}

	j, err := json.Marshal(record)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(j))
}

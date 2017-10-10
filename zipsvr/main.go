package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/ezhai24/info344-in-class/zipsvr/handlers"
	"github.com/ezhai24/info344-in-class/zipsvr/models"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	// w.Write([]byte("Hello, World!"))
	fmt.Fprintf(w, "Hello %s!", name)
}

func memoryHandler(w http.ResponseWriter, r *http.Request) {
	// runtime.GC()
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(stats)
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	tlskey := os.Getenv("TLSKEY")
	tlscert := os.Getenv("TLSCERT")
	if len(tlscert) == 0 || len(tlskey) == 0 {
		log.Fatal("please set TLSKEY and TLSCERT")
	}

	zips, err := models.LoadZips("/data/zips.csv")
	if err != nil {
		log.Fatalf("error loading zips: %v", err)
	}
	log.Printf("loaded %d zips", len(zips))

	cityIndex := models.ZipIndex{}
	for _, z := range zips {
		cityLower := strings.ToLower(z.City) // make case insensitive
		cityIndex[cityLower] = append(cityIndex[cityLower], z)
	}

	// fmt.Println("Hello World!")
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/memory", memoryHandler)

	cityHandler := &handlers.CityHandler{
		Index:      cityIndex,
		PathPrefix: "/zips/",
	}
	mux.Handle("/zips/", cityHandler)

	fmt.Printf("server is listening at https://%s\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlscert, tlskey, mux))
}

package main

import "fmt"
import "net/http"
import "log"
import "runtime"
import "encoding/json"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Header().Add("Content-Type", "text/plain")
	// w.Write([]byte("Hello, World!"))
	fmt.Fprintf(w, "Hello %s!", name)
}

func memoryHandler(w http.ResponseWriter, r *http.Request) {
	// runtime.GC()
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func main() {
	// fmt.Println("Hello World!")
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/memory", memoryHandler)
	fmt.Printf("server is listening at http://localhost:4000\n")
	log.Fatal(http.ListenAndServe("localhost:4000", mux))
}

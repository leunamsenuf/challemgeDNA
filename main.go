/*
* APP: challengeDNA
* AUTHOR: manuelfunes@yahoo.com.br
* VERSION 0.0 - 2022-DEZ-01
 */
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/valyala/fastjson"
)

var (
	// flagPort is the open port the application listens on
	flagPort = flag.String("port", "9000", "Port to listen on")
)

var results []string

// GetHandler handles the index route
func GetHandler(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error converting results to json",
			http.StatusInternalServerError)
	}

	//fmt.Println(results[1])
	w.Write(jsonBody)

}

// PostHandler converts post request body to string
func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		results = append(results, string(body))
		log.Printf(results[1])
		fmt.Fprint(w, "POST done")

		//begin process
		s := []byte(results[1])

		fmt.Printf("foo.0=%s\n", fastjson.GetString(s, "letters", "0"))
		//end process

		for i := 1; i < 6; i++ {
			strIndx := strconv.Itoa(i)
			fmt.Printf("foo.0=%s\n", fastjson.GetString(s, "letters", strIndx))
		}

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	flag.Parse()
}

func main() {
	results = append(results, time.Now().Format(time.RFC3339))

	mux := http.NewServeMux()
	mux.HandleFunc("/", GetHandler)
	mux.HandleFunc("/sequence", PostHandler)

	log.Printf("listening on port %s", *flagPort)
	log.Fatal(http.ListenAndServe(":"+*flagPort, mux))
}

func processDNA(jsonDNA []string) {

}

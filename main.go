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
	"strings"
	"time"

	"github.com/valyala/fastjson"
)

var (
	// flagPort is the open port the application listens on
	flagPort = flag.String("port", "9000", "Port to listen on")
)

var results []string

type myJSON struct {
	Dna []string
}

var totalSequence int

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

		contaLinhas(results)

		log.Println("totalSequence: ", totalSequence)

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

//----------------------FUNCTIOS

// Conta ocurrencis 4 letras igauis
func conta(jsonDNA string) int {
	check := 1

	//split
	spl := strings.Split(jsonDNA, "")
	//log.Println(spl[0])
	for i := 0; i < 4; i++ {
		if spl[i] == spl[i+1] {
			check = check + 1
		} else {
			if check < 4 {
				check = 0
			}
		}
	}
	return check
}

// Conta linhas
func contaLinhas(results []string) {
	//Get results
	s := []byte(results[1])

	//Conta linhas
	for i := 0; i < 6; i++ {
		strIndx := strconv.Itoa(i)
		seq := fastjson.GetString(s, "letters", strIndx)
		log.Println(seq)
		conta := conta(seq)
		log.Println(conta)
		if conta == 4 {
			totalSequence += 1
		}
	}

}

// Conta colunas
func contaColunas(results []string) {
	//Get results
	s := []byte(results[1])

	for i := 0; i < 6; i++ {
		strIndx := strconv.Itoa(i)
		seq := fastjson.GetString(s, "letters", strIndx)
		log.Println(seq)
		conta := conta(seq)
		log.Println(conta)
		if conta == 4 {
			totalSequence += 1
		}
	}

}

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func dieRoll(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	die, _ := strconv.Atoi(params["die"])
	sides, _ := strconv.Atoi(params["sides"])

	total := 0
	for i := 0; i < die; i++ {
		total += rand.Intn(sides-1) + 1
	}

	fmt.Fprintf(w, strconv.Itoa(total))
}

func main() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/health", helloworld)
	rtr.HandleFunc("/roll/{die}/{sides}", dieRoll)

	http.Handle("/", rtr)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

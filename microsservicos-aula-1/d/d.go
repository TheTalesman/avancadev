package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type Result struct {
	Status string
}

func main() {

	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	//coupon := r.PostFormValue("coupon")
	min := 1
	max := 10
	ran := rand.Intn(max-min) + min
	var result Result
	if ran > 7 {
		fmt.Println("valid")
		result = Result{Status: "valid"}
	} else {
		fmt.Println("invalid")
		result = Result{Status: "invalid"}
		w.WriteHeader(http.StatusInternalServerError)

	}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, string(jsonResult))

}

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/vocationnations/algorithm/algorithm"
	"github.com/vocationnations/algorithm/config_parser"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to the HomePage!")
	if err != nil {
		fmt.Printf("ERROR: Couldn't write response %v", err)
	}
	fmt.Println("Endpoint Hit: homePage")
}

func calculateScore(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)

	config := config_parser.Config{}

	err := yaml.Unmarshal(reqBody, &config)
	if err != nil {
		_, err2 := fmt.Fprintf(w, "ERROR: Couldn't marshal YAML: %v", err)
		if err2 != nil {
			fmt.Printf("ERROR: Couldn't write response %v", err)
		}
		return
	}

	// run the algorithm
	res, err := algorithm.Run(&config)
	if err != nil {
		panic(fmt.Sprintf("ERROR: The algorithm failed to run, erro %v", err))
	}

	// print the results
	_, err = fmt.Fprintf(w, algorithm.Print(res))
	if err != nil {
		fmt.Printf("ERROR: Couldn't write response %v", err)
	}

}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/calculateScore", calculateScore).Methods("POST")
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	handler := c.Handler(myRouter)

	log.Fatal(http.ListenAndServe(":10000", handler))
}

func main() {
	handleRequests()
}

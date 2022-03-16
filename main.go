package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Mobile    int64  `json:"MobileNumber"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", HomePageHandler).Methods(http.MethodGet)
	router.HandleFunc("/Signup", SignUpHandler).Methods(http.MethodPost)
	fmt.Println("Server at 8080")
	http.ListenAndServe(":8080", router)
}
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	SearchKey := r.URL.Query().Get("q")
	fmt.Fprint(w, SearchKey)
}
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var user *User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	errJson := json.Unmarshal(body, &user)
	if errJson != nil {
		log.Fatalln(errJson)
	}
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, user)
}

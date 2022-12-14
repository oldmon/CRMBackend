package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Customer struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

var customers = map[string]Customer{
	"35d6cf6e-731c-11ed-a1eb-0242ac120002": {
		Name:      "Aye",
		Role:      "AyeAye",
		Email:     "Aye@local.host",
		Phone:     "(123) 456-7893",
		Contacted: true,
	},
	"44f6f776-731c-11ed-a1eb-0242ac120002": {
		Name:      "Bay",
		Role:      "BayBay",
		Email:     "Bay@local.host",
		Phone:     "(123)456-7892",
		Contacted: false,
	},
	"44f6r27324-731c-11ed-a1eb-0242ac120002": {
		Name:      "Cey",
		Role:      "CeyCey",
		Email:     "Cey@local.host",
		Phone:     "(123) 456-7891",
		Contacted: false,
	},
	"41234r27324-731c-11ed-a1eb-0242ac120002": {
		Name:      "Dey",
		Role:      "DeyDey",
		Email:     "Dey@local.host",
		Phone:     "(123) 456-7899",
		Contacted: false,
	},
}

func homePage(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./static/")).ServeHTTP(w, r)
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if _, ok := customers[mux.Vars(r)["id"]]; ok {
		json.NewEncoder(w).Encode(customers[mux.Vars(r)["id"]])
		w.WriteHeader(http.StatusOK)
	} else {
		doesNotExist := map[string]string{}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(doesNotExist)
	}
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	customer.Id = uuid.New().String()
	reqBody, _ := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")
	if _, ok := customers[mux.Vars(r)["id"]]; ok {
		doesNotExist := map[string]string{}
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(doesNotExist)
	} else {
		json.Unmarshal(reqBody, &customer)
		customers[customer.Id] = customer
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(customers[customer.Id])
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	// delete an existing customer
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	if _, ok := customers[id]; ok {
		delete(customers, id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)
	} else {
		doesNotExist := map[string]string{}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(doesNotExist)
	}
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	// update an existing customer
	w.Header().Set("Content-Type", "application/json")
	var customer Customer
	reqBody, _ := ioutil.ReadAll(r.Body)
	id := mux.Vars(r)["id"]
	if _, ok := customers[id]; ok {
		json.Unmarshal(reqBody, &customer)
		customers[id] = customer
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers[id])
	} else {
		doesNotExist := map[string]string{}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(doesNotExist)
	}
}
func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	http.ServeFile(w, r, "./static/404.html")
}
func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	if port == "" {
		log.Printf("$PORT variable not set. Setting to default 8080")
		port = "8080"
	}
	log.Printf("Starting the Server on %s...", port)
	router.HandleFunc("/", homePage).Methods("GET")              // home page (static page)
	router.HandleFunc("/customers", getCustomers).Methods("GET") // get all customers
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")    // update a customer
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE") // delete a customer
	// not found page
	router.NotFoundHandler = http.HandlerFunc(pageNotFound)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

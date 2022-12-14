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
	"b656c55f-4ea5-404c-8904-16c269053c63": {Name: "Krista Povele", Role: "Estimator", Email: "kpovele0@businessinsider.com", Phone: "(825) 1713311", Contacted: false},
	"a1dd84eb-e8c1-4d31-adf4-bb8a445707ff": {Name: "Tam Ridesdale", Role: "Project Manager", Email: "tridesdale1@cbslocal.com", Phone: "(275) 5507051", Contacted: true},
	"8c3e8d04-df17-4cd4-ba98-8c25d5a780a0": {Name: "Alberto Reuss", Role: "Construction Foreman", Email: "areuss2@salon.com", Phone: "(479) 5656120", Contacted: false},
	"72fa2fd0-e173-48b6-bfb6-4fb979291f75": {Name: "Marysa Bye", Role: "Subcontractor", Email: "mbye3@kickstarter.com", Phone: "(439) 7598244", Contacted: false},
	"c43155c2-a9e0-4626-bf96-0c837b7f0d46": {Name: "Karney Hurdidge", Role: "Construction Foreman", Email: "khurdidge4@blogtalkradio.com", Phone: "(120) 4285006", Contacted: false},
	"045fb7d6-e216-4b32-9b8b-ec1f637d052e": {Name: "Alick Willcott", Role: "Architect", Email: "awillcott5@abc.net.au", Phone: "(723) 1637142", Contacted: true},
	"4c515efa-03d6-4f15-acdc-8287a1f1160e": {Name: "Nickie Dmitrievski", Role: "Construction Foreman", Email: "ndmitrievski6@marriott.com", Phone: "(989) 5708009", Contacted: true},
	"98b69cef-7167-4d70-acaf-29d50a191d69": {Name: "Arleyne Luca", Role: "Construction Foreman", Email: "aluca7@dropbox.com", Phone: "(474) 3193118", Contacted: false},
	"bf4770ca-3a37-492b-ba56-211a1f0df9c5": {Name: "Mallissa Ferruzzi", Role: "Project Manager", Email: "mferruzzi8@slate.com", Phone: "(791) 4083959", Contacted: false},
	"062bcd48-cb3a-4313-aea3-237872ed1698": {Name: "Ferdinande Maberley", Role: "Surveyor", Email: "fmaberley9@china.com.cn", Phone: "(733) 9679131", Contacted: true},
	"839f39c1-9a99-41df-bb82-492f797e8f63": {Name: "Pablo Drewitt", Role: "Subcontractor", Email: "pdrewitta@reference.com", Phone: "(684) 4068614", Contacted: true},
	"2ce962ca-eee5-462f-9743-9101066e0c03": {Name: "Gael Georgescu", Role: "Project Manager", Email: "ggeorgescub@latimes.com", Phone: "(298) 4077513", Contacted: true},
	"543e6ebe-2119-4a19-99f6-f4029c401a42": {Name: "Carina Shearstone", Role: "Engineer", Email: "cshearstonec@alexa.com", Phone: "(660) 4971377", Contacted: true},
	"dc267f78-b2a9-40c0-a0f7-1bc59cf2d3e1": {Name: "Caria Moultrie", Role: "Surveyor", Email: "cmoultried@dagondesign.com", Phone: "(230) 8326629", Contacted: true},
	"7c4474c0-065b-40e2-8866-e75180aea421": {Name: "Ginnifer Pentlow", Role: "Subcontractor", Email: "gpentlowe@macromedia.com", Phone: "(413) 7144098", Contacted: false},
	"8c85ab40-a3e6-4528-baf6-0bee7233e717": {Name: "Magda Bailiss", Role: "Construction Manager", Email: "mbailissf@latimes.com", Phone: "(432) 2299707", Contacted: true},
}
var doesNotExist = map[string]string{}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	http.ServeFile(w, r, "usage.html")
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
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	if _, ok := customers[id]; ok {
		delete(customers, id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(doesNotExist)
	}
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(doesNotExist)
	}
}
func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	if port == "" {
		log.Printf("System $PORT not set. Setting to 3000")
		port = "3000"
	}
	log.Printf("Server listen on %s", port)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	//Not found
	router.NotFoundHandler = http.HandlerFunc(pageNotFound)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

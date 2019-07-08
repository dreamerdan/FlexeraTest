
// main.go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
)

 func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New Person Entry!")
	fmt.Println ("Endpoint Hit: homePage")
}

//Returns all data - call to sort function if passed as parameter
func returnAllPeople (w http.ResponseWriter, r *http.Request) {
	log.Println ("Endpoint Hit: returnAllPeople")
	sort, err := r.URL.Query()["sort"]
  
	if !err || len(sort[0]) < 1 {
		log.Printf("unable to sort - return all, %v", err)
				json.NewEncoder(w).Encode(People)
	} else if (sort[0] == "Name"){
		log.Println("sort by:" + sort[0])
				sortByName(People)
				json.NewEncoder(w).Encode(People)
	}  else if (sort[0] == "Email"){
		log.Println("sort by:" + sort[0])
				sortByEmail(People)
				json.NewEncoder(w).Encode(People)
	}
}

func sortByName(people []Person) {
    sort.SliceStable(people, func(i, j int) bool {                    
        return people[i].Name < people[j].Name                         
    })
}

func sortByEmail(people []Person) {
    sort.SliceStable(people, func(i, j int) bool {                                        
		return people[i].Email < people[j].Email       
    })
}

//return one single record using the id as endpoint
func returnPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	
	for _, person := range People {
		if person.Id == key {
			json.NewEncoder(w).Encode(person)
		}
	}
}

//create a new record - add to People array
func createNewPerson(w http.ResponseWriter, r *http.Request) {
	log.Println ("Endpoint Hit: create")        
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
				log.Printf("body read error, %v", err)
				w.WriteHeader(500)
	} else {
				var person Person
				json.Unmarshal(reqBody, &person)
				People = append(People, person)
	}  
}

//delete a person from the People array
func deletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	for index, person := range People {
		if person.Id == id {
			People = append(People[:index], People[index+1:]...)
		}
	}
	
	fmt.Println ("Endpoint Hit: delete")
}

//set the person record to be checked to persist state of checkbox on frontend
func checkPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	for index, person := range People {
		if person.Id == id {		
			if person.Checked == 1 {
				person.Checked = 0
			} else {
				person.Checked =1
			}
         People[index] = person			
		}
	}
	fmt.Println(People)
	fmt.Println ("Endpoint Hit: checkPerson")
}
		
		
//handle endpoint requests		
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/app/people", returnAllPeople).Methods("GET")
	myRouter.HandleFunc("/app/people", createNewPerson).Methods("POST")
	myRouter.HandleFunc("/app/people/{id}", checkPerson).Methods("PUT")
	myRouter.HandleFunc("/app/people/{id}", deletePerson).Methods("DELETE")
    myRouter.HandleFunc("/app/people/{id}", returnPerson)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}


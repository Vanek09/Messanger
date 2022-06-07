package main

import (
	mongoapp "mongo_service/mongo_managment"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	// "strconv"

	"github.com/gorilla/mux"
)

func put(w http.ResponseWriter, r *http.Request) {
	var tList mongoapp.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(reqBody, &tList)
	if tList.Nickname == "" {
		log.Println("Nickname is empty. Skipping")
		return
	}
	_, err = mongoapp.CreateUser(tList)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	log.Printf("Item {%s:%v} was added to Mongo DB\n", tList.Nickname, tList.Hashed_pwd)
}


func getAll(w http.ResponseWriter, r *http.Request) {
	lists, err := mongoapp.GetUsers()
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("{}")
		log.Println(err)
		return
	}
	json.NewEncoder(w).Encode(lists)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	item, err := mongoapp.GetUser(id)
	if err != nil {
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func getMessagesList(w http.ResponseWriter, r *http.Request) {
	from_id := mux.Vars(r)["from_id"]
	to_id := mux.Vars(r)["to_id"]
	messages := mongoapp.GetUserMessages(from_id, to_id)
	json.NewEncoder(w).Encode(messages)
}

func send(w http.ResponseWriter, r *http.Request) {
	var msg mongoapp.Message
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(reqBody, &msg)

	go mongoapp.SaveMessage(msg)
	return
}

func main() {
	mongoapp.Setup()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/put", put).Methods("POST")
	router.HandleFunc("/sendMessage", send).Methods("POST")
	router.HandleFunc("/getAll", getAll).Methods("GET")
	router.HandleFunc("/get/{id}", getItem).Methods("GET")
	router.HandleFunc("/getMessages/{from_id}-{to_id}", getMessagesList).Methods("GET")
	log.Println("API is listening on port 3000")
	http.ListenAndServe("0.0.0.0:3000", router)
}
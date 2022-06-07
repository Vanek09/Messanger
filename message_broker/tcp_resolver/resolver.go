package tcpresolver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	models "tcp/forms"
)
const (
	host = "http://api:3000"
)
func Recieve(b []byte, response *[]byte) {
	var r models.Request
	//We need to exclude 0 from buf in order to parse JSON
	err := json.Unmarshal(b[:bytes.IndexByte(b, 0)], &r)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(r)

	switch r["command"]{
	case "addUser":
		populateUser(r["metadata"])
		*response = []byte("{\"status\": \"OK\"}")
	case "getUsers":
		*response = getUsers()
	case "getMessages":
		*response = getMessages(r)
	case "sendMessage":
		sendMessage(r["metadata"])
		*response = []byte("{\"status\": \"OK\"}")
	}
}

func populateUser(u interface{}) {
	log.Println("Recieved populateUser command")
	var user models.User
	log.Printf("DEBUG: %v", u)
	u1, _ := u.(string)
	json.Unmarshal([]byte(u1), &user)
	fmt.Println(user)
	//Encode the data
	postBody, _ := json.Marshal(user)
	responseBody := bytes.NewBuffer(postBody)
  //Leverage Go's HTTP Post function to make request
	 resp, err := http.Post(fmt.Sprintf("%s/put", host), "application/json", responseBody)
	 if err != nil {
		fmt.Println(err)
	 }
	 defer resp.Body.Close()
}

func getUsers() []byte {
	log.Println("Recieved getUsers command")
	resp, err := http.Get(fmt.Sprintf("%s/getAll", host))
	if err != nil {
		fmt.Println(err)
	}
 //We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
 //Convert the body to type string
	return body
}

func getMessages(r models.Request) []byte {
	log.Println("Recieved getMessages command")
	resp, err := http.Get(fmt.Sprintf("%s/getMessages/%s-%s", host, r["from"], r["to"]))
	if err != nil {
		fmt.Println(err)
	}
    //We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Message body: %s", body)
	return body
}

func sendMessage(m interface{}) {
	log.Println("Recieved sendMessage command")
	var message models.Message
	log.Printf("DEBUG: %v", m)
	m1, _ := m.(string)
	log.Println(m1)
	json.Unmarshal([]byte(m1), &message)
	//Encode the data
	postBody, _ := json.Marshal(message)
	responseBody := bytes.NewBuffer(postBody)
  //Leverage Go's HTTP Post function to make request
	 resp, err := http.Post(fmt.Sprintf("%s/sendMessage", host), "application/json", responseBody)
	 if err != nil {
		fmt.Println(err)
	 }
	 defer resp.Body.Close()
}
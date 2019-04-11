package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)


type User struct {

	Name	string		`json:"name"`		
	Email   string		`json:"email"`
	Pass    string		`json:"pass"`
	Adress	Adress		`json:"adress"`
}

type UserUpdate struct {

	Name	string		`json:"name"`
	Email   string		`json:"email"`
	Adress	Adress		`json:"adress"`
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	body := r.Body

	var user User

	// parsing io.ReadCLoser to slice of bytes []byte
	bytes, _ := ioutil.ReadAll(body)

	err := json.Unmarshal(bytes, &user)

	// validating decoding error
	if err != nil{
		// Invalid Information Format HTTP response
		w.Write([]byte(`{"code": 400, "message": "Invalid json!"}`))
		return
	}

	err = SaveUser(user)

	// validation in database interaction
	if err != nil{

		//HTTP response from server failure
		w.Write([]byte("500"))
		return
	}

	// HTTP response success
	w.Write([]byte(`{"code": 200, "message": "success!"}`))


}

func SearchUserEmail(w http.ResponseWriter, r *http.Request){

	q := r.URL.Query()

	value := q["email"][0]

	user, err := SearchUser(value)

	if err != nil  || user == nil  {
		log.Printf("[ERROR] user not found: %v", err)
		w.Write([]byte("404"))
		return
	}

	userJson, err:= json.Marshal(user)
	if err != nil {
		log.Printf("[ERROR] Invalid email: %v", err)
		w.Write([]byte("400"))
		return
	}

	w.Write([]byte(userJson))

	return
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	body := r.Body

	var user UserUpdate

	// parsing io.ReadCLoser to slice of bytes []byte
	bytes, _ := ioutil.ReadAll(body)

	err := json.Unmarshal(bytes, &user)

	// validating decoding error
	if err != nil{
		// Invalid Information Format HTTP response
		w.Write([]byte(`{"code": 400, "message": "Invalid json!"}`))
		return
	}

	err = EditUser(user)

	// validation in database interaction
	if err != nil{

		//HTTP response from server failure
		w.Write([]byte("500"))
		return
	}

	// HTTP response success
	w.Write([]byte(`{"code": 200, "message": "success!"}`))


}


package main

import (
	"github.com/globalsign/mgo/bson"
	"log"
)

func SaveUser(user User) (err error) {

	conn, err := GetConnection()
	if err != nil{
		log.Printf("[ERROR] could not get BD connection", err)
	}

	collection := conn.DB("marmitasfit2").C("users")


	err = collection.Insert(user)

	if err != nil{
		log.Printf("[ERROR] could not save user")
		return
	}

	return
}

func SearchUser(s string) (n []User, err error) {
	conn, err := GetConnection()
	collection := conn.DB("marmitasfit2").C("users")

	err = collection.Find(bson.M{"email": s}).All(&n)

	return n, err

}

func EditUser (user UserUpdate) (err error){
	conn, err := GetConnection()

	if err != nil{
		log.Printf("[ERROR] could not get BD connection", err)
	}

	collection := conn.DB("marmitasfit2").C("users")

	err = collection.Update(bson.M{"email": user.Email}, bson.M{"$set": bson.M{"name": user.Name,
	"adress": user.Adress}})

	if err != nil{
		log.Printf("[ERROR] could not edit user %v", err)
		return
	}
	return
}

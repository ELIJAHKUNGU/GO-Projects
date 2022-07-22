package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	// "lulu/dbconnection"
)

type User struct {
	gorm.Model
	Name        string `gorm:"json:"Name"`
	ID          string `gorm:"json:"ID"`
	county      string `gorm:"json:"county"`
	Sessionid   string `gorm:"json:"sessionid"`
	servicecode string `gorm:"json:"servicecode"`
	phonenumber string `gorm:"json:"phonenumber"`
	text        string `gorm:"json:"text"`
	items       string `gorm:"json:"items"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	// s := "ccc"
	// fmt.Println(strings.Count(s, "cc"))

	cm := user.text

	cnt := strings.Count(cm, "*")

	response1 := ""
	response2 := ""

	if cnt == 0 {
		// This is the first request. Note how we start the response with CON
		response1 = "CON Welcome to jameniplaform Kenya \n"
		response2 = "1. Please enter your Name \n"
	} else if cnt == 1 {
		// Business logic for first levelresponse
		response1 = "CON 2. Please enter your ID Number \n"
	} else if cnt == 2 {
		// Business logic for first levelresponse
		response1 = "CON 3. Please enter your county code \n"
		response2 = "e.g 033 for Narok \n"
	} else if cnt == 3 {
		// Business logic for first levelresponse
		response1 = "CON 4. Please enter the name of your co operative \n"
		response2 = "e.g for Narok co operative \n"
	} else if cnt == 4 {
		// Business logic for first levelresponse
		response1 = "CON 5. Please enter the items that you sell separated by a comma \n"
		response2 = "e.g Belts, bags,earings \n"

	} else if cnt == 5 {
		// Business logic for first levelresponse
		// This is a terminal request. Note how we start theresponse with END
		response1 = "END Thank you very much for registering with us"
		// x = explode("*",user.text);
		x := strings.Split(user.text, "*")

		//to make sure there are used
		println(x)
		//"kadzo*kadzo*33334009*045*narok coop";
	}

	//to make sure there are used
	println(response1, response2)

	//Create a db connection

	//Insert into mysql

}
func main() {
	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints

	r.HandleFunc("/books", createUser).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
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

	if cnt == 0 {
		// This is the first request. Note how we start the response with CON
		// response  := "CON Welcome to jameniplaform Kenya \n";
		var response string = "1. Please enter your Name \n"
	} else if cnt == 1 {
		// Business logic for first levelresponse
		response := "CON 2. Please enter your ID Number \n"
	} else if cnt == 2 {
		// Business logic for first levelresponse
		response := "CON 3. Please enter your county code \n"
		response := "e.g 033 for Narok \n"
	} else if cnt == 3 {
		// Business logic for first levelresponse
		response := "CON 4. Please enter the name of your co operative \n"
		response := "e.g for Narok co operative \n"
	} else if cnt == 4 {
		// Business logic for first levelresponse
		response := "CON 5. Please enter the items that you sell separated by a comma \n"
		response := "e.g Belts, bags,earings \n"

	} else if cnt == 5 {
		// Business logic for first levelresponse
		// This is a terminal request. Note how we start theresponse with END
		response := "END Thank you very much for registering with us"
		// x := explode("*",user.text);
		//"kadzo*kadzo*33334009*045*narok coop";
	}

}
func main() {
	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints

	r.HandleFunc("/books", createUser).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", getUsers)
	fmt.Println("API está rodando em :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

type User struct {
	Id   int
	Name string
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode([]User{
		{
			Id:   1,
			Name: "Nori",
		},

		{
			Id:   2,
			Name: "Thaise",
		},

		{
			Id:   3,
			Name: "Maria",
		},
	})

}

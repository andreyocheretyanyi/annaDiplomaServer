package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
)

func main() {
	create_table()
	http.HandleFunc("/getAll", getDataHandler)
	http.HandleFunc("/add", addHandler)
	log.Println("Listening...")
	port := ":"+os.Getenv("PORT")
	//port := ":5500"
	fmt.Print(port)
	http.ListenAndServe(port, nil)

}






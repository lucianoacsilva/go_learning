package main

import (
	"net/http"
	"log"
)

func main()  {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	log.Print("Server running at port 3000!", nil)
	
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	
}
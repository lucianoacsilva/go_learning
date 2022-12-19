package main

import (
	"fmt"
	"net/http"
)

func ola(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Olá Mundo!\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main()  {
	http.HandleFunc("/ola", ola)
	http.HandleFunc("/header", headers)

	fmt.Println("Servidor rodando na porta 8000!!\n")
	http.ListenAndServe(":8000", nil)
}
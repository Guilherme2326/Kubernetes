package main

import (
	"fmt"
	"log"
	"net/http"
)

// func main() {
// 	err := http.ListenAndServe(":80", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	http.HandleFunc("/", Hello)
// 	http.ListenAndServe(":8080", nil)
// }

// func Hello(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("<h1>Hello, World!</h1>"))
// }

func main() {
    // Registra um handler para a rota raiz ("/")
	
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Olá, Mundo!")
		w.Write([]byte("<h1>Hello, World!!!!</h1>"))
// }
    })

    // Inicia o servidor na porta 8080 (ou outra porta não privilegiada)
    fmt.Println("Servidor iniciado na porta 8080")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
	log.Fatal(err)
	}
}

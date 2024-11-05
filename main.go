package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not found", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "hello human")
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "form parsing error : %v", err)
		return
	}

	fmt.Fprintf(w, "post request successfull\n")
	name := r.FormValue("name")
	adress := r.FormValue("adress")

	fmt.Fprintf(w, "name = %s\n", name)
	fmt.Fprintf(w, "adress = %s\n", adress)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("server running at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

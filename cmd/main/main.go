package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/ndegealbert/go-bookstore/pkg/routes"
	"log"
	"net/http"
)

func main() {
	//https://github.com/gorilla/mux
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	fmt.Println("Listening to port 9010")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}

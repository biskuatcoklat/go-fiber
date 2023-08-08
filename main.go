package main

import (
	"GO-FIBER/config"
	"GO-FIBER/controllers/categorycontroller"
	"GO-FIBER/controllers/homecontroller"
	"GO-FIBER/controllers/productcontroller"
	"log"
	"net/http"
)

func main(){
	config.ConnectDB()

	//homepage
	http.HandleFunc("/", homecontroller.Welcome)

	//categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/update", categorycontroller.Update)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// Products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server running in port 3030")
	http.ListenAndServe(":3030", nil)
}
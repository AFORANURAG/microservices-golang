package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	schemas "github.com/AFORANURAG/microservices-golang/schemas"
)

type ProductHandler struct{}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}
func (ph *ProductHandler) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	if h.Method == "GET" {
		ph.getProducts(rw)
		return
	}
	if h.Method == "POST" {
		body, err := io.ReadAll(h.Body)
		if err != nil {
			http.Error(rw, "Malformed request body", http.StatusBadRequest)
		}
		var product schemas.Product
		json.Unmarshal(body, &product)
		fmt.Println("Product unmarshalled is", product)
		ph.addProduct(&product)
		log.Printf("Product added successfully")
		return
	}
	// fmt.Println("Hello product")
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (ph *ProductHandler) getProducts(rw http.ResponseWriter) {
	productList := schemas.GetProduct()

	if productMarshallingError := productList.ToJSON(rw); productMarshallingError != nil {
		log.Println("productMarshallingError is", productMarshallingError)
		http.Error(rw, "Some json encoding error", http.StatusInternalServerError)
		return
	}
}
func (ph *ProductHandler) addProduct(product *schemas.Product) {
	schemas.ProductList = append(schemas.ProductList, product)
	fmt.Println("List is", schemas.ProductList)
}

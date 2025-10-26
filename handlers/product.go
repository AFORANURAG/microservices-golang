package handlers

import (
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

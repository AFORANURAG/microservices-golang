package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type GoodByeStruct struct {
	Hello string `json:"hello"`
}

type GoodByeHandler struct{}

func (h *GoodByeHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Printf("Goodbye world, I will soon return")
	time.Sleep(5 * time.Second)
	d, err := io.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Goodbye! you killed me"))
	}
	var goodbyestruct GoodByeStruct
	jsonMarshallingError := json.Unmarshal(d, &goodbyestruct)
	if jsonMarshallingError != nil {
		http.Error(rw, "Goodbye! you killed me", http.StatusBadRequest)
		return
	}
	var s GoodByeStruct = GoodByeStruct{Hello: "world"}
	jsonned, _ := json.Marshal(s)
	// errors.Is(jsonMarshallingError,context.Canceled)
	fmt.Printf("jsonned is : %s", jsonned)

	log.Println("Response body is: ", s)
	fmt.Fprintf(rw, "Hello from base handler")
	defer r.Body.Close()
}
func NewGoodByeHandler() *GoodByeHandler {
	return &GoodByeHandler{}
}

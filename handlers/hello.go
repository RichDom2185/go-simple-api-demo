package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RichDom2185/go-simple-api-demo/models"
)

func HandleGetHello(w http.ResponseWriter, r *http.Request) {
	var responseData models.GetHelloAPIResponse

	responseData = models.GetHelloAPIResponse{
		Message: "Hello, ",
		Name:    "Bob!",
	}

	output, err := json.Marshal(responseData)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	fmt.Fprintln(w, string(output))
	// Or use this instead:
	// w.Write(output)
}

func HandlePostHello(w http.ResponseWriter, r *http.Request) {
	var params models.PostHelloAPIRequestParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(w, "Hello %s, you are %d years old!\n", params.Name, params.Age)
}

package main

//each and every file in go must have a package name
import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Message    string `json:"message"`
	Technology string `json:"technology"`
}

func main() {
	mymux := SetupMyHandlers()
	http.ListenAndServe(":8080", mymux)
}

func SetupMyHandlers() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/test", HelloHandler)

	return mux
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	//notice how this function takes two parameters
	//the first parameter is a ResponseWriter writer and
	//this is where we write the response we want to send back to the user
	//in this case Hello world
	//the second parameter is a pointer of type  http.Request this holds
	//all information of the request sent by the user
	//this may include query parameters,path parameters and many more
	result := Result{Message: "Hello", Technology: "go"}

	json.NewEncoder(w).Encode(result)
}

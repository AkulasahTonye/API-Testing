package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hello", helloHendler)
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
func helloHendler(w http.ResponseWriter, r *http.Request) {
	visitorName := r.URL.Query().Get("visitorName")
	clientIP := r.RemoteAddr             //Get the requester's IP Address
	location := r.Header.Get("New York") // You can use an IP geolocation service to get the actual location

	greeting := fmt.Sprintf(`Hello, %s!, The temperature is 11 degree celsius in %s.`, visitorName, location)

	response := map[string]interface{}{
		clientIP: clientIP,
		location: location,
		greeting: greeting,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//http://localhost:8080/api/hello?visitor_name=Mark

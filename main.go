package main

import "net/http"

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/api/v1/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("OK"))
	})

	http.ListenAndServe(":8080", router)
}

package routers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kamaal111/container-registry-play/utils"
)

func Start() {
	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		serverAddress = fmt.Sprintf(":%s", utils.UnwrapEnvironment("PORT"))
	}

	mux := http.NewServeMux()

	mux.Handle("/ping", loggerMiddleware(http.HandlerFunc(ping)))
	mux.Handle("/", loggerMiddleware(http.HandlerFunc(notFound)))

	log.Printf("Listening on %s...", serverAddress)

	err := http.ListenAndServe(serverAddress, mux)
	log.Fatal(err)
}

func ping(writer http.ResponseWriter, request *http.Request) {
	output, err := json.Marshal(struct {
		Message string `json:"message"`
	}{Message: "pong"})
	if err != nil {
		utils.ErrorHandler(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("content-type", "application/json")
	writer.Write(output)
}

package routers

import (
	"net/http"

	"github.com/kamaal111/container-registry-play/utils"
)

func notFound(writer http.ResponseWriter, request *http.Request) {
	utils.ErrorHandler(writer, "Not found", http.StatusNotFound)
}

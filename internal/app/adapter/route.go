package adapter

import (
	"encoding/json"
	"net/http"
	"simple-graphql-server/internal/app/adapter/controller"
)

func Routes(service *controller.Service) {
	http.HandleFunc("/graphql", func(writer http.ResponseWriter, request *http.Request) {
		res := controller.InitGraphQL(request.Context(), request.URL.Query().Get("query"), service)
		err := json.NewEncoder(writer).Encode(res)
		if err != nil {
			return
		}
	})
}

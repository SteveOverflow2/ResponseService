package rest

import (
	"net/http"

	response "response-service/pkg/http/rest/handlers/response"
)

func (s *server) routes() {
	// Example handlers subset
	responseHandler := s.Router.PathPrefix("/response").Subrouter()

	//POST
	responseHandler.HandleFunc("", response.CreateResponseHandler(s.ResponseService)).Methods(http.MethodPost)

	//GET
	responseHandler.HandleFunc("/getAll", response.GetAllResponsesHandler(s.ResponseService)).Methods(http.MethodGet)
	responseHandler.HandleFunc("/{uuid}", response.GetResponseHandler(s.ResponseService)).Methods(http.MethodGet)
}

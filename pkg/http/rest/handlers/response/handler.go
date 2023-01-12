package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"response-service/pkg/http/rest/handlers"
	"response-service/pkg/response"
	"response-service/pkg/util"

	"github.com/gorilla/mux"
)

func CreateResponseHandler(responseService response.ResponseService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var response response.CreateResponse
		if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
			handlers.RenderErrorResponse(w, "Invalid request payload", r.URL.Path, util.WrapErrorf(err, util.ErrorCodeInvalid, "json decoder"))
			return
		}

		responseId, err := responseService.CreateResponse(r.Context(), response)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			handlers.RenderErrorResponse(w, "Invalid request payload", r.URL.Path, util.WrapErrorf(err, util.ErrorCodeInvalid, err.Error()))
			return
		}
		handlers.RenderResponse(w, http.StatusOK, responseId)

	}
}

func GetResponseHandler(responseService response.ResponseService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := mux.Vars(r)["uuid"]
		if len(uuid) == 0 {
			err := util.NewErrorf(util.ErrorCodeInternal, "Query parameters are invalid")
			handlers.RenderErrorResponse(w, err.Error(), r.URL.Path, err)
			return
		}

		response, err := responseService.GetResponse(r.Context(), uuid)
		if err != nil {
			handlers.RenderErrorResponse(w, "internal server error", r.URL.Path, util.WrapErrorf(err, util.ErrorCodeInternal, err.Error()))
			return
		}

		handlers.RenderResponse(w, http.StatusOK, response)
	}
}

func GetAllResponsesHandler(responseService response.ResponseService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		response, err := responseService.GetAllResponses(r.Context())
		if err != nil {
			handlers.RenderErrorResponse(w, "internal server error", r.URL.Path, util.WrapErrorf(err, util.ErrorCodeInternal, err.Error()))
		}

		handlers.RenderResponse(w, http.StatusOK, response)
	}
}

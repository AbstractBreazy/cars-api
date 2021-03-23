package response

import (
	"encoding/json"
	"net/http"
)

type StatusResponse struct {
	HTTPStatusCode       int    `json:"http_status_code"`
	Status               bool   `json:"status"`
	ResponseMessage      string `json:"response_message"`
	ResponseInternalCode int    `json:"response_internal_code"`
}

func (s *StatusResponse) GetJSON() []byte {
	context, _ := json.Marshal(struct {
		StatusResponse *StatusResponse `json:"status_response"`
	}{
		StatusResponse: s,
	})
	return context
}

func New() *StatusResponse {
	resp := &StatusResponse{
		HTTPStatusCode:       200,
		Status:               true,
		ResponseMessage:      "status:ok",
		ResponseInternalCode: 0,
	}
	return resp
}

func Check(err error, w http.ResponseWriter, resp *StatusResponse, message string, code int) bool {
	if err != nil {
		resp.Status = false
		if len(message) == 0 {
			resp.ResponseMessage = err.Error()
		} else {
			resp.ResponseMessage = message + ": " + err.Error()
		}
		resp.ResponseInternalCode = code
		w.Write(resp.GetJSON())
		return true
	}
	return false
}

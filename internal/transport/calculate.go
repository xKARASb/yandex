package transport

import (
	"calculate/pkg/rpn"
	"encoding/json"
	"net/http"
	"strconv"
)

type calcRequest struct {
	Expression string `json:"expression"`
}

type calcResponse struct {
	Result string `json:"result"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	var exp calcRequest
	err := json.NewDecoder(r.Body).Decode(&exp)
	if err != nil {
		http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		return
	}
	res, err := rpn.Calc(exp.Expression)

	if err != nil {
		if err.Error() == "Not correct input" {
			http.Error(w, `{"error": "Expression is not valid"}`, http.StatusUnprocessableEntity)
			return
		}
		http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(calcResponse{Result: strconv.FormatFloat(res, 'f', -1, 64)})
}

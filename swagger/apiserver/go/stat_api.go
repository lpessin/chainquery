/*
 * Chain Query
 *
 * The LBRY blockchain is read into SQL where important structured information can be extracted through the Chain Query API.
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"net/http"
    "encoding/json"

	"github.com/lbryio/chainquery/apis"
)

func HandleAddressSummary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response, err := apis.HandleAction("AddressSummary", w, r)
	if err != nil {
    		w.WriteHeader(http.StatusInternalServerError)
    		w.Write([]byte(err.Error()))
    }
    process(w,response)
}

func HandleStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response, err := apis.HandleAction("Status", w, r)
	if err != nil {
    		w.WriteHeader(http.StatusInternalServerError)
    		w.Write([]byte(err.Error()))
    }
    process(w,response)
}

// Processes the response information and sends it back.
func process(w http.ResponseWriter, response *apis.Response) {
	jsonBytes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error encoding response to json"))
	}
	_, err = w.Write(jsonBytes) //Ignore bytes written
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error encoding response to json"))
	}
}
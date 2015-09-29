package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RootVersionOK struct {
	Res    string `json:"result"`
	VerAPI string `json:"apiVersion"`
	VerDb  string `json:"dbVersion"`
	Rid    int    `json:"rid"`
}

func (ct *Controller) RootHandler(c *gin.Context) ([]byte, int) {

	rOK := RootVersionOK{"OK", c.Keys["ApiVersion"].(string), c.Keys["DbVersion"].(string), 0}

	body, err := json.MarshalIndent(rOK, "", "\t")

	if err != nil {
		return []byte("Json error"), http.StatusInternalServerError
	}

	return body, http.StatusOK
}

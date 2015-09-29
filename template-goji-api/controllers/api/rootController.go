package api

import (
	"encoding/json"
	"net/http"

	"github.com/zenazn/goji/web"
)

type RootVersionOK struct {
	Res    string `json:"result"`
	VerAPI string `json:"apiVersion"`
	VerDb  string `json:"dbVersion"`
	Rid    int    `json:"rid"`
}

func (ct *Controller) RootVersion(c web.C, r *http.Request) ([]byte, int) {

	rOK := RootVersionOK{"OK", c.Env["ApiVersion"].(string), c.Env["DbVersion"].(string), 0}

	body, err := json.MarshalIndent(rOK, "", "\t")

	if err != nil {
		return []byte("Json error"), http.StatusInternalServerError
	}

	return body, http.StatusOK
}

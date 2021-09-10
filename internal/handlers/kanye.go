package handlers

import (
	"encoding/json"
	"fmt"
	"gofirstapp/internal/model"
	"io/ioutil"
	"net/http"
)

func (h *Handlers) Kanye(w http.ResponseWriter, r *http.Request) {
	resp, err := h.client.Get("https://api.kanye.rest/")

	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprintf("something went wrong. %s", err.Error())))
		return
	}

	if resp == nil || resp.StatusCode != http.StatusOK {
		_, _ = w.Write([]byte("api responded incorrectly"))	
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		_, _ = w.Write([]byte("bad kanye body"))	
		return
	}


	var quote model.Quote

	err = json.Unmarshal(body, &quote)

	if err != nil {
		_, _ = w.Write([]byte("failed to parse body"))	
		return
	}

	_, _ = w.Write([]byte(quote.Quote))
}
package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/graph-gophers/graphql-go"
)

type handler struct {
	Schema *graphql.Schema
}

type request struct {
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
	OperationName string                 `json:"operationName"`
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := &request{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result := h.Schema.Exec(r.Context(), req.Query, req.OperationName, req.Variables)
	if req.OperationName == "doLogin" {
		sessionID := string(result.Data)[12:24]
		cookie := http.Cookie{
			Name:     "robofunding_session",
			Value:    sessionID,
			Path:     "/",
			Expires:  time.Now().Add(time.Hour * 24 * 30),
			HttpOnly: true,
			Domain:   "robo.app.org",
		}
		http.SetCookie(w, &cookie)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

/*cookie := http.Cookie{
	Name:     "robofunding_session",
	Value:    session.Id,
	Path:     "/",
	MaxAge:   int(session.Expires - time.Now().Unix()),
	HttpOnly: true,
}
fmt.Println("\nW BEFORE:", ctx.Value("ctx").(models.ContextValue).ResponseWriter)
http.SetCookie(ctx.Value("ctx").(models.ContextValue).ResponseWriter, &cookie)
fmt.Println("\nW AFTER:", ctx.Value("ctx").(models.ContextValue).ResponseWriter)*/

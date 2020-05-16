package auth

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/zgordan-vv/robofunding/backend/db"
	"github.com/zgordan-vv/robofunding/backend/models"
)

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, POST, DELETE, OPTIONS")
			w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			ctxValue := models.ContextValue{}
			log.Println("COOKIES:", r.Cookies())
			cookie, err := r.Cookie("robofunding_session")
			log.Println("COOKIE:", cookie, err)
			if cookie != nil {
				sessionId := cookie.Value
				session, err := db.GetDB().GetSession(sessionId)
				if err == nil {
					if session.Expires <= time.Now().Unix() {
						err := db.GetDB().DeleteSession(sessionId)
						if err != nil {
							log.Println("Session error:", err)
						}
					} else {
						currentUser, err := db.GetDB().GetUser(session.UserId)
						if err != nil {
							log.Println("User not found:", err)
						} else {
							ctxValue.SessionId = sessionId
							ctxValue.CurrentUser = currentUser
						}
					}
				}
			}
			ctx := context.WithValue(r.Context(), "ctx", ctxValue)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

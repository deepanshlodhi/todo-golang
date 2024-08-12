package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/machinebox/graphql"
)

const (
	userSessionDataQuery = "query UserSessionData {userSessionData {id email}}"
	url                  = "https://alpha-api.adalat.ai/user-service/graphql"
)

type UserSessionData struct {
	ID    string
	Email string
}

type ResponseData struct {
	UserSessionData *UserSessionData
}

// Define a custom type for the context key
type contextKey string

// Define a constant for the user key
const UserKey contextKey = "userSessionData"

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("authorization")
		// fmt.Println(authHeader)
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}
		client := graphql.NewClient(url)
		req := graphql.NewRequest(userSessionDataQuery)

		var respData ResponseData

		req.Header.Set("Authorization", authHeader)
		ctx := context.Background()
		if err := client.Run(ctx, req, &respData); err != nil {
			fmt.Println("Error executing GraphQL request:", err)
			http.Error(w, "Failed to retrieve user data", http.StatusInternalServerError)
			return
		}

		// r.Header.Set("X-USER-ID", respData.UserSessionData.ID)
		// fmt.Println("----->>>", respData.UserSessionData)

		ctx = context.WithValue(r.Context(), UserKey, respData.UserSessionData)
		next.ServeHTTP(w, r.WithContext(ctx))
		// fmt.Println("Middleware execution after response")
	})
}

package middleware

import (
	"fmt"
	"net/http"
)

type GlobalMiddleware struct {
}

func NewGlobalAuthMiddleware() *GlobalMiddleware {
	return &GlobalMiddleware{}
}

func (m *GlobalMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("global middleware start")
		next(w, r)
		fmt.Println("global middleware end")

	}
}

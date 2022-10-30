package middleware

import (
	"fmt"
	"net/http"
)

// CommonJwtAuthMiddleware : with jwt on the verification, no jwt on the verification
type CommonMiddleware struct {
}

func NewCommonMiddleware() *CommonMiddleware {
	return &CommonMiddleware{}
}

func (m *CommonMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("global before...")
		next(w, r)
		fmt.Println("global after...")
	}
}

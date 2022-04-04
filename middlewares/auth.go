package middlewares

import (
	"fmt"
	"net/http"
)

type AuthMiddleware struct {
	Next http.Handler
}

func (am *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if am.Next == nil {
		am.Next = http.DefaultServeMux
	}

	fmt.Println("middleware in")
	am.Next.ServeHTTP(writer, request)
	fmt.Println("middleware out")
}

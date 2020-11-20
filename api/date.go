// Ejemplo sacado de https://vercel.com/docs/serverless-functions/supported-languages#go

package handler

import (
	"fmt"
	"net/http"
	"time"
)

// Handler func prints the current time on w
func Handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, currentTime)
}

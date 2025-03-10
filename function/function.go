package function

import (
	"net/http"
)

// HelloWorld returns a simple "Hello World" message
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
} 
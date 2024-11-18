package tmuxsheets

import (
	"fmt"
	"io"
	"net/http"
)

func handleOAuthCallback(w http.ResponseWriter, r *http.Request, pw *io.PipeWriter) {
	code := r.URL.Query().Get("code")
	if code == "" {
		fmt.Println("Auth code not found in callback request")
		return
	}


	fmt.Fprintf(pw, "AUTHORIZATION_CODE=%s\n", code)
	http.ServeFile(w, r, "success.html")
}

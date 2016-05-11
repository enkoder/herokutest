package api

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func OneHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		f := filepath.Join(os.Getenv("GOPATH"), "src", "herokutest", "python.py")
		out, err := exec.Command("python", f).Output()
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		fmt.Fprintf(w, "Hello from one! Here is the output from python %s", string(out))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

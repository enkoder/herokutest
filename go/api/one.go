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
		cwd, err := os.Getwd()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		gp := os.Getenv("GOPATH")
		fmt.Fprintf(w, "Current Working DIR: %s\n", cwd)
		fmt.Fprintf(w, "GOPATH = %s\n", cwd)
		f := filepath.Join(gp, "src", "herokutest", "python.py")
		out, err := exec.Command("python", f).Output()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprintf(w, "Hello from one! Here is the output from python %s", string(out))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

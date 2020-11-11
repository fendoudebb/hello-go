package main

import (
	"errors"
	"io"
	"log"
	"net/http"
)

func MockPost(w http.ResponseWriter, req *http.Request) {
	err := func() error {
		if req.Method != http.MethodPost {
			log.Println("Error req method: ", req.Method)
			return errors.New("bad req method")
			//return errors.New("expected GET")
		}
		return nil
	}()
	if err != nil {
		w.WriteHeader(400)
		_, _ = io.WriteString(w, err.Error())
		return
	}
}

func main() {
	http.HandleFunc("/mock/post", MockPost)
	if err := http.ListenAndServe(":8080", nil); err!= nil {
		log.Fatal(err)
	}
}

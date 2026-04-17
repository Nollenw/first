package httpservice

import (
	"errors"
	"fmt"
	"net/http"
)

func StartHttpService() error {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Обработка")
		w.Write([]byte("Hello from docker\n"))
	})
	err := http.ListenAndServe(":9091", nil)
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	} else {
		return err
	}
}

package handlers

import (
	"encoding/json"
	"net/http"
	"study/sql"
	"study/task"
	"time"

	"github.com/jackc/pgx/v5"
)

func CreateTaskHandler(conn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var task task.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		task.TimeCreated = time.Now()
		task.IsCompleted = false

		if err := sql.CreateTask(*conn, r.Context(), &task); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

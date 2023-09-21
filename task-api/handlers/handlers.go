package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	models "task-api/Models"
	"task-api/database"

	"task-api/config"
)

var db *sql.DB

func init() {
	db = database.ConnectDB()
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(r.URL.Path)
	var tasks []models.Task

	getQuery := "SELECT id,title,completed,created_at FROM tasks"

	rows, err := db.Query(getQuery)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var task models.Task

		err := rows.Scan(&task.ID, &task.Title, &task.Completed, &task.CreatedAt)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := config.APIResponse{
		Status:  http.StatusOK,
		Message: "Fetched all tasks successfully",
		Data:    tasks,
	}

	jsonResponse, err := json.Marshal(&response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf(r.URL.Path)

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid HTTP method Only POST is allow", http.StatusMethodNotAllowed)
		return
	}

	var task models.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	insertQuery := "INSERT INTO tasks (title,completed,created_at) VALUES ($1,$2,NOW()) RETURNING id"

	if err := db.QueryRow(insertQuery, task.Title, task.Completed).Scan(&task.ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Return the created task as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)

}

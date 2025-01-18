package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}

var tasks []Task
var currentID = 1

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func createTask(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

	// Обработка предварительного запроса OPTIONS
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = currentID
	currentID++
	tasks = append(tasks, task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

	// Обработка предварительного запроса OPTIONS
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

	// Обработка предварительного запроса OPTIONS
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, task := range tasks {
		if task.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	http.NotFound(w, r)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

	// Обработка предварительного запроса OPTIONS
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var updatedTask Task
	_ = json.NewDecoder(r.Body).Decode(&updatedTask)

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = updatedTask.Title
			tasks[i].Body = updatedTask.Body
			tasks[i].Completed = updatedTask.Completed
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tasks[i])
			return
		}
	}
	http.NotFound(w, r)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

	// Обработка предварительного запроса OPTIONS
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/tasks", createTask).Methods("POST", "OPTIONS")
	r.HandleFunc("/tasks", getTasks).Methods("GET", "OPTIONS")
	r.HandleFunc("/tasks/{id}", getTask).Methods("GET", "OPTIONS")
	r.HandleFunc("/tasks/{id}", updateTask).Methods("PUT", "OPTIONS")
	r.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE", "OPTIONS")

	// Start server
	http.ListenAndServe(":8080", r)
}

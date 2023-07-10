package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Task struct {
	ID         int    `json:"id"`
	NumberTask string `json:"number"`
	NameTask   string `json:"name"`
	Quantity   int    `json:"quantity"`
	Reps       int    `json:"reps"`
}

var tasks = []Task{
	{ID: 1, NumberTask: "Task1", NameTask: "Push-ups", Quantity: 25, Reps: 4},
	{ID: 2, NumberTask: "Task2", NameTask: "Pull-ups", Quantity: 15, Reps: 4},
	{ID: 3, NumberTask: "Task3", NameTask: "Hanging leg raises", Quantity: 20, Reps: 4},
	{ID: 4, NumberTask: "Task4", NameTask: "Sit-ups", Quantity: 30, Reps: 4},
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/tasks", GetTasks).Methods(http.MethodGet)

	fmt.Println("Server is starting at port 8080")
	http.ListenAndServe(":8080", r)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	number := r.URL.Query().Get("number")

	var filteredTasks []Task
	for _, task := range tasks {
		if task.NumberTask == number {
			filteredTasks = append(filteredTasks, task)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredTasks)
}

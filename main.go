package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Task struct {
	ID   int    `json:"id"`
	Date string `json:"date"`
	Name string `json:"name"`
}

var tasks = []Task{
	{ID: 1, Name: "Task1", Date: "07.11.2010"},
	{ID: 2, Name: "Task2", Date: "21.08.2005"},
	{ID: 3, Name: "Task3", Date: "10.06.2002"},
	{ID: 4, Name: "Task4", Date: "16.01.2023"},
	{ID: 5, Name: "Task5", Date: "07.11.2010"},
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/tasks", GetTasks).Methods(http.MethodGet)

	fmt.Println("Server is starting at port 8080")
	http.ListenAndServe(":8080", r)

}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")

	for _, task := range tasks {
		if task.Date == date {
			json.NewEncoder(w).Encode(task)
		}

	}

	w.Header().Set("Content-Type", "application/json")
}

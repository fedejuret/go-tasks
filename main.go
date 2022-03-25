package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type task struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type allTasks []task

var tasks = allTasks{
	{
		Id:        1,
		Title:     "Task 1",
		Completed: false,
	},
	{
		Id:        2,
		Title:     "Task 2",
		Completed: true,
	},
	{
		Id:        3,
		Title:     "Task 3",
		Completed: false,
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!!!!")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {

	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.Id = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tasks)

}

func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	taskId, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	for _, item := range tasks {
		if item.Id == taskId {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&task{})
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	taskId, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	for index, item := range tasks {
		if item.Id == taskId {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(tasks)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	taskId, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	var updatedTask task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	json.Unmarshal(reqBody, &updatedTask)

	for index, item := range tasks {
		if item.Id == taskId {
			tasks = append(tasks[:index], tasks[index+1:]...)
			updatedTask.Id = taskId
			tasks = append(tasks, updatedTask)
			break
		}
	}

	json.NewEncoder(w).Encode(tasks)
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}

package main

import (
	"api-02/tasks"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", tasks.IndexRoute)
	router.HandleFunc("/tasks", tasks.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", tasks.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", tasks.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", tasks.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", tasks.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}

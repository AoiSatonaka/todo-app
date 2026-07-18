package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/AoiSatonaka/todo-app/backend/internal/application"
	"github.com/AoiSatonaka/todo-app/backend/internal/domain/task"
	"github.com/AoiSatonaka/todo-app/backend/internal/infrastructure/mongodb"
	"github.com/AoiSatonaka/todo-app/backend/internal/presentation/grpc"
	"github.com/joho/godotenv"
)

func Hello(name string) string {
	return "hello " + name
}

type T struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	User        string `json:"user"`
}

func decodeT(b io.Reader) (*T, error) {
	var t T
	if err := json.NewDecoder(b).Decode(&t); err != nil {
		return nil, fmt.Errorf("decode error: %v\n", err)
	}
	return &t, nil
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	t, err := decodeT(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tr := mongodb.NewTaskRepository()
	err = tr.Delete(t.Id, t.User)
	if err != nil {
		http.Error(w, fmt.Sprintf("create error: %v\n", err), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "delete task")
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	t, err := decodeT(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tt, err := task.New(t.Id, t.Title, t.Description, time.Now(), int32(2), int32(3), []string{"hoge", "test", "fuga"})
	if err != nil {
		http.Error(w, fmt.Sprintf("task.New error: %v\n", err), http.StatusBadRequest)
	}

	tr := mongodb.NewTaskRepository()
	err = tr.Update(tt, t.User)
	if err != nil {
		http.Error(w, fmt.Sprintf("create error: %v\n", err), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "update task")
}

func getTask(w http.ResponseWriter, r *http.Request) {
	t, err := decodeT(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tr := mongodb.NewTaskRepository()
	ts, err := tr.FindById(t.Id, t.User)
	if err != nil {
		http.Error(w, fmt.Sprintf("create error: %v\n", err), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, ts)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	t, err := decodeT(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tt, err := task.New("", t.Title, t.Description, time.Now(), int32(2), int32(3), []string{"hoge", "test", "fuga"})
	if err != nil {
		http.Error(w, fmt.Sprintf("task.New error: %v\n", err), http.StatusBadRequest)
	}

	tr := mongodb.NewTaskRepository()
	err = tr.Create(tt, t.User)
	if err != nil {
		http.Error(w, fmt.Sprintf("create error: %v\n", err), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "create task")
}

func main() {
	// load env file
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "local"
	}

	err := godotenv.Load(fmt.Sprintf("backend/env/.env.%s", env))
	if err != nil {
		panic(err)
	}

	// server setup
	//  http.HandleFunc("/create", createTask)
	//  http.HandleFunc("/update", updateTask)
	//  http.HandleFunc("/delete", deleteTask)
	//  http.HandleFunc("/get", getTask)
	//  http.ListenAndServe(":3000", nil)
	ta := application.NewToDoService(mongodb.NewTaskRepository())
	s, err := grpc.New(ta)
	if err != nil {
		panic(err)
	}
	if err = s.Serve(); err != nil {
		panic(err)
	}
}

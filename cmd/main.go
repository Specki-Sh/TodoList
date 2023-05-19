package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"todolist"
	"todolist/db"
	"todolist/handlers"
	"todolist/repository"
	"todolist/service"
)

func main() {
	db.StartDbConnection()
	defer db.CloseDbConnection()

	storage := repository.NewTaskRepository(db.GetDBConn())
	todoListService := service.NewTodoList(storage)
	app := handlers.NewWebApp(todoListService)
	srv := new(todolist.Server)
	go func() {
		if err := srv.Run("9191", app.Route()); err != nil {
			log.Fatalf("Error occured while running http server: %s", err.Error())
			return
		}

	}()

	fmt.Println("App Started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	fmt.Println("Shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Errorf("error occurred on server shutting down: %s", err.Error())
	}
}

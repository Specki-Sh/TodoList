package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"todolist"
	"todolist/service"
	"todolist/db"
	"todolist/handlers"
	"todolist/repository"
)

func main() {
	db.StartDbConnection()
	defer db.CloseDbConnection()

	TaskStorage := repository.NewTaskRepository(db.GetDBConn())
	TaskService := service.NewTaskService(TaskStorage)
	TaskHandler := handlers.NewTaskHandler(TaskService)

	UserStorage := repository.NewUserRepository(db.GetDBConn())
	UserService := service.NewUserService(UserStorage)
	UserHandler := handlers.NewUserHandlers(UserService)

	app := handlers.NewWebApp(TaskHandler, UserHandler)

	srv := new(todolist.Server)
	go func() {
		if err := srv.Run("9191", app.SetupRoutes()); err != nil {
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

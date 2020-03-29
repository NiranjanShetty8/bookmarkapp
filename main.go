package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/NiranjanShetty8/bookmarkapp/controllers"
	"github.com/NiranjanShetty8/bookmarkapp/repository"
	"github.com/NiranjanShetty8/bookmarkapp/services"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:4040)/swabhav?charset=utf8&parseTime=True")
	defer db.Close()
	if err != nil {
		fmt.Print(err)
	}
	router := mux.NewRouter()
	if router == nil {
		log.Fatal("No router Created")
	}
	fmt.Println("Server Started")
	repos := repository.NewGormRepository()
	initialize(db, repos, router)
	headers := handlers.AllowedHeaders([]string{"Content-Type", "token"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origin := handlers.AllowedOrigins([]string{"*"})
	server := &http.Server{
		Handler:      handlers.CORS(headers, methods, origin)(router),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Addr:         ":8080",
	}

	var wait time.Duration

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt)

	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	server.Shutdown(ctx)
	func() {
		fmt.Println("Closing DB")
		db.Close()
	}()
	fmt.Println("Server ShutDown....")
	os.Exit(0)
}

// Create instance of services & controllers & register routes
func initialize(db *gorm.DB, repos *repository.GormRepository, router *mux.Router) {
	userService := services.NewUserService(db, repos)
	categoryService := services.NewCategoryService(db, repos)
	bookmarkService := services.NewBookmarkService(db, repos)

	userController := controllers.NewUserController(userService)
	categoryController := controllers.NewCategoryController(categoryService)
	bookmarkController := controllers.NewBookmarkController(bookmarkService)

	userController.RegisterRoutes(router)
	categoryController.RegisterRoutes(router)
	bookmarkController.RegisterRoutes(router)
}

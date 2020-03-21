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
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
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
	// db.Close()
	fmt.Println("Server ShutDown....")
	os.Exit(0)
}

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

// repos := repository.NewGormRepository()
// us := services.NewUserService(db, repos)
// cs := services.NewCategoryService(db, repos)
// bms := services.NewBookmarkService(db, repos)

// userGenId := uuid.NewV4()
// categoryGenId := uuid.NewV4()

// bookmark1 := model.Bookmark{
// 	Description: "google",
// 	URL:         "www.google.com",
// 	CategoryID:  categoryGenId,
// }

// bookmark2 := model.Bookmark{
// 	Description: "yahoo",
// 	URL:         "www.yahoo.com",
// 	CategoryID:  categoryGenId,
// }
// // allBookMarks := []model.Bookmark{bookmark1, bookmark2}

// category1 := model.Category{
// 	Base: model.Base{
// 		ID: categoryGenId,
// 	},
// 	Name: "Social",
// 	// Bookmarks: allBookMarks,
// 	UserID: userGenId,
// }
// // allCategories := []model.Category{category1}
// user := model.User{
// 	Base: model.Base{
// 		ID: userGenId,
// 	},
// 	Username: "Niranjan",
// 	Password: "12345678",
// Categories: allCategories,

// Categories: []model.Category{
// 	{
// 		Base: model.Base{
// 			ID: uuid.NewV4(),
// 		},
// 		Name: "Social",
// 		Bookmarks: []model.Bookmark{
// 			{
// 				Base: model.Base{
// 					ID: uuid.NewV4(),
// 				},
// 				Description: "google",
// 				URL:         "www.google.com",
// 			},
// {
// 	Base: model.Base{
// 		ID: uuid.NewV4(),
// 	},
// 	Description: "yahoo",
// 	URL:         "www.yahoo.com",
// },
// },
// }}
// }
// us.Register(&user)
// cs.AddCategory(&category1)
// bms.AddBookmark(&bookmark1)
// bms.AddBookmark(&bookmark2)
// fmt.Println("main", user.ID)
// fmt.Println("this", cs, bms)
// checkUser := &model.User{
// 	Username: "Niranjan",
// 	Password: "12345678",
// }
// eror := us.Login(checkUser)
// fmt.Println(eror)
// bookmark := model.Bookmark{
// 	Description: "facebook",
// 	URL: "www.facebook.com" ,

// }
// bms.AddBookmark()

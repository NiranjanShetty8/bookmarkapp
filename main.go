package main

import (
	"fmt"

	"github.com/NiranjanShetty8/bookmarkapp/model"
	"github.com/NiranjanShetty8/bookmarkapp/repository"
	"github.com/NiranjanShetty8/bookmarkapp/services"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:4040)/swabhav?charset=utf8&parseTime=True")
	defer db.Close()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Hello")
	repos := repository.NewGormRepository()
	bms := services.NewBookmarkService(db, repos)
	cs := services.NewCategoryService(db, repos)
	us := services.NewUserService(db, repos)
	user := model.User{
		Username: "Niranjan",
		Password: "12345678",
		Bookmarks: []model.Bookmark{
			{
				Base: model.Base{
					ID: uuid.NewV4(),
				},
				Description: "google",
				URL:         "www.google.com",
			}, {
				Base: model.Base{
					ID: uuid.NewV4(),
				},
				Description: "yahoo",
				URL:         "www.yahoo.com",
			}}}
	us.Register(&user)
	fmt.Println("main", user.ID)
	fmt.Println("this", cs, bms)
	fmt.Println(user)
	checkUser := &model.User{
		Username: "Niranjan",
		Password: "12345678",
	}
	eror := us.Login(checkUser)
	fmt.Println(eror)
	// bookmark := model.Bookmark{
	// 	Description: "facebook",
	// 	URL: "www.facebook.com" ,

	// }
	// bms.AddBookmark()

}

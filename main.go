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
	us := services.NewUserService(db, repos)
	cs := services.NewCategoryService(db, repos)
	bms := services.NewBookmarkService(db, repos)

	userGenId := uuid.NewV4()
	categoryGenId := uuid.NewV4()

	bookmark1 := model.Bookmark{
		Description: "google",
		URL:         "www.google.com",
		CategoryID:  categoryGenId,
	}

	bookmark2 := model.Bookmark{
		Description: "yahoo",
		URL:         "www.yahoo.com",
		CategoryID:  categoryGenId,
	}
	// allBookMarks := []model.Bookmark{bookmark1, bookmark2}

	category1 := model.Category{
		Base: model.Base{
			ID: categoryGenId,
		},
		Name: "Social",
		// Bookmarks: allBookMarks,
		UserID: userGenId,
	}
	// allCategories := []model.Category{category1}
	user := model.User{
		Base: model.Base{
			ID: userGenId,
		},
		Username: "Niranjan",
		Password: "12345678",
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
	}
	us.Register(&user)
	cs.AddCategory(&category1)
	bms.AddBookmark(&bookmark1)
	bms.AddBookmark(&bookmark2)
	fmt.Println("main", user.ID)
	fmt.Println("this", cs, bms)
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

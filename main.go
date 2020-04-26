package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// User is ~~~
type User struct {
	gorm.Model        // gorm.Model is a basic GoLang struct which includes the following fields: ID, CreatedAt, UpdatedAt, DeletedAt.
	Username   string `form:"username" binding:"required" gorm:"unique;not null"`
	Password   string `form:"password" binding:"required"`
}

// Tweet is ~~~
type Tweet struct {
	gorm.Model // gorm.Model is a basic GoLang struct which includes the following fields: ID, CreatedAt, UpdatedAt, DeletedAt.
	Name       string
	Age        int
	Content    string
}

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("view/*.html")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	dbInit()
	//db_addRecord("aaa",11,"aaa")
	//db_editRecord(61, "bbb", 22, "bbb")
	//db_deleteRecord(42)
	//tweet := db_getRecord(1)

	router.GET("/", func(c *gin.Context) {
		tweet := dbGetAllRecords()
		c.HTML(200, "index.html", gin.H{
			"tweet":   tweet,
			"message": "pong",
		})
	})
	router.GET("/signup", func(c *gin.Context) {
		c.HTML(200, "signup.html", gin.H{
			"message": "pong",
		})
	})

	router.GET("/signin", func(c *gin.Context) {
		c.HTML(200, "signin.html", gin.H{
			"message": "pong",
		})
	})

	router.POST("/add", func(c *gin.Context) {
		name := c.PostForm("name")
		age, _ := strconv.Atoi(c.PostForm("age"))
		content := c.PostForm("content")
		dbAddRecord(name, age, content)
		c.Redirect(302, "/")

	})
	router.POST("/signup", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		println("OneHere!")
		if err := dbCreatUser(username, password); err != nil {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"err": err})
			println("Here!")
			c.Abort()

		}
		println("wneHere!")
		c.Redirect(302, "/")

	})
	router.Run() // listen and serve on 0.0.0.0:8080
}

// データベースの初期化
func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Failed to connect database\n")
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Tweet{})
	defer db.Close()
}

// データベースにレコード追加
func dbCreatUser(username string, password string) []error {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Failed to connect database\n")
	}
	db.Create(&User{
		Username: username,
		Password: password,
	})
	defer db.Close()
	return nil
}

// データベースにレコード追加
func dbAddRecord(name string, age int, content string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Failed to connect database\n")
	}
	db.Create(&Tweet{
		Name:    name,
		Age:     age,
		Content: content,
	})
	defer db.Close()
}

// データベースからid番目のレコードを取得
func dbGetRecord(id int) Tweet {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Failed to connect database\n")
	}
	var tweet Tweet
	db.First(&tweet, id)
	// db.Close()
	return tweet
}

// データベースからすべてのレコードを取得
func dbGetAllRecords() []Tweet {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Failed to connect database\n")
	}
	var tweet []Tweet
	// レコードの並び替え と 取得
	db.Order("created_at desc").Find(&tweet)
	// db.Close()
	return tweet
}

// データベースからid番目のレコードを取得し、データを更新
func dbEditRecord(id int, name string, age int, content string) Tweet {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Failed to connect database\n")
	}
	var tweet Tweet
	db.First(&tweet, id)
	tweet.Name = name
	tweet.Age = age
	tweet.Content = content
	db.Save(&tweet)
	// db.Close()
	return tweet
}

// データベースからid番目のレコードを取得し、データを削除
func dbDeleteRecord(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Failed to connect database\n")
	}
	var tweet Tweet
	db.Delete(&tweet, id)
	// db.Close()
}

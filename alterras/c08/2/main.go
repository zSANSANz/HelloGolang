package main

// import library
import (
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// create a new echo instance
	e := echo.New()

	// Route / to handler function
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)

	// start the server, and log if it fails
	e.Start(":7000")
}

func init() {
	InitDB()
	InitialMigration()
}

// database connection
func InitDB() {

	//  declare struct config and variabel connection string
	connectionString :=
		"root:@tcp(127.0.0.1:3306)/db_name?charset=utf8&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

var (
	DB *gorm.DB
)

type User struct {
	gorm.Model
	Name     string `json:"name" form: "name"`
	Email    string `json:"email" form: "email"`
	Password string `json:"password" form: "password"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

// get all users
func GetUsersController(c echo.Context) error {
	var users []User

	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// func allUsers(db *gorm.DB) func(echo.Context) error {
// 	return func(c echo.Context) error {
// 		var users []User
// 		db.Find(&users)
// 		fmt.Println("{}", users)

// 		return c.JSON(http.StatusOK, users)
// 	}
// }

// func newUser(db *gorm.DB) func(echo.Context) error {
// 	return func(c echo.Context) error {
// 		name := c.Param("name")
// 		email := c.Param("email")
// 		db.Create(&User{Name: name, Email: email})
// 		return c.String(http.StatusOK, name+" user successfully created")
// 	}
// }

// func deleteUser(db *gorm.DB) func(echo.Context) error {
// 	return func(c echo.Context) error {
// 		name := c.Param("name")

// 		var user User
// 		db.Where("name = ?", name).Find(&user)
// 		db.Delete(&user)

// 		return c.String(http.StatusOK, name+" user successfully deleted")
// 	}
// }

// func updateUser(db *gorm.DB) func(echo.Context) error {
// 	return func(c echo.Context) error {
// 		name := c.Param("name")
// 		email := c.Param("email")
// 		var user User
// 		db.Where("name=?", name).Find(&user)
// 		user.Email = email
// 		db.Save(&user)
// 		return c.String(http.StatusOK, name+" user successfully updated")
// 	}
// }

// func usersByPage(db *gorm.DB) func(echo.Context) error {
// 	return func(c echo.Context) error {
// 		limit, _ := strconv.Atoi(c.QueryParam("limit"))
// 		page, _ := strconv.Atoi(c.QueryParam("page"))
// 		var result []User
// 		db.Limit(limit).Offset(limit * (page - 1)).Find(&result)
// 		return c.JSON(http.StatusOK, result)
// 	}
// }
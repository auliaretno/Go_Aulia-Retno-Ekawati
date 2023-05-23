package routes

import (
	"code_competence/constants"
	"code_competence/controllers"
	m "code_competence/middlewares"

	mid "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)

	// users routes
	eUser := e.Group("/users")
	eUser.POST("/register", controllers.CreateUserController)
	eUser.POST("/login", controllers.LoginUserController)
	// 	Authenticated with JWT
	eUserJwt := eUser.Group("")
	eUserJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eUserJwt.GET("", controllers.GetUsersController)
	eUserJwt.GET("/:id", controllers.GetUserController)
	eUserJwt.PUT("/:id", controllers.UpdateUserController)
	eUserJwt.DELETE("/:id", controllers.DeleteUserController)

	// books routes
	eBook := e.Group("/books")
	// Authenticated with JWT
	eBook.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eBook.GET("", controllers.GetBooksController)
	eBook.POST("", controllers.CreateBookController)
	eBook.GET("/:id", controllers.GetBookController)
	eBook.PUT("/:id", controllers.UpdateBookController)
	eBook.DELETE("/:id", controllers.DeleteBookController)

		//categories routes
	eCategory := e.Group("/categories")
		// Authenticated with JWT
	eCategory.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eCategory.GET("", controllers.GetCategoriesController)
	eCategory.GET("/:id", controllers.GetCategoryController)
	eCategory.POST("", controllers.CreateCategoryController,)
	eCategory.PUT("/:id", controllers.UpdateCategoryController)
	eCategory.DELETE("/:id", controllers.DeleteCategoryController)

	return e
}
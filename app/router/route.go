package router

import (
	kd "KosKita/features/kos/data"
	kh "KosKita/features/kos/handler"
	ks "KosKita/features/kos/service"
	ud "KosKita/features/user/data"
	uh "KosKita/features/user/handler"
	us "KosKita/features/user/service"
	"KosKita/utils/cloudinary"
	"KosKita/utils/encrypts"
	"KosKita/utils/middlewares"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	hash := encrypts.New()
	cloudinaryUploader := cloudinary.New()
	// midtrans := externalapi.New()

	userData := ud.New(db)
	userService := us.New(userData, hash)
	userHandlerAPI := uh.New(userService, cloudinaryUploader)

	kosData := kd.New(db)
	kosService := ks.New(kosData)
	kosHandlerAPI := kh.New(kosService, cloudinaryUploader)

	// define routes/ endpoint USER
	e.POST("/login", userHandlerAPI.Login)
	e.POST("/users", userHandlerAPI.RegisterUser)
	e.GET("/users", userHandlerAPI.GetUser, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.UpdateUser, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.DeleteUser, middlewares.JWTMiddleware())

	// define routes/ endpoint KOS
	e.POST("/kos", kosHandlerAPI.CreateKos, middlewares.JWTMiddleware())
}

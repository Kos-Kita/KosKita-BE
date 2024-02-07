package handler

import (
	"KosKita/features/user"
	"KosKita/utils/cloudinary"
	"KosKita/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
	cld         cloudinary.CloudinaryUploaderInterface
}

func New(service user.UserServiceInterface, cloudinaryUploader cloudinary.CloudinaryUploaderInterface) *UserHandler {
	return &UserHandler{
		userService: service,
		cld:         cloudinaryUploader,
	}
}

func (handler *UserHandler) RegisterUser(c echo.Context) error {
	newUser := UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data, data not valid", nil))
	}

	userCore := RequestToCore(newUser)
	errInsert := handler.userService.Create(userCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data. "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Successful Operation", nil))
}

func (handler *UserHandler) Login(c echo.Context) error {
	var reqData = LoginRequest{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data, data not valid", nil))
	}
	result, token, err := handler.userService.Login(reqData.Email, reqData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error login. "+err.Error(), nil))
	}
	responseData := map[string]any{
		"token": token,
		"nama":  result.Name,
		"role":  result.Role,
	}
	return c.JSON(http.StatusOK, responses.WebResponse("success login", responseData))
}

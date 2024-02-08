package handler

import (
	"KosKita/features/booking"
	"KosKita/utils/middlewares"
	"KosKita/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookService booking.BookServiceInterface
}

func New(bs booking.BookServiceInterface) *BookHandler {
	return &BookHandler{
		bookService: bs,
	}
}

func (handler *BookHandler) CreateBook(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	newBook := BookRequest{}
	errBind := c.Bind(&newBook)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data order not valid", nil))
	}

	bookCore := RequestToCoreBook(newBook, uint(userIdLogin))
	payment, errInsert := handler.bookService.Create(userIdLogin, bookCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse(errInsert.Error(), nil))
	}

	result := CoreToResponseBook(payment)

	return c.JSON(http.StatusOK, responses.WebResponse("success booking kos", result))
}

package handler

import (
	"KosKita/features/kos"
	"KosKita/utils/cloudinary"
	"KosKita/utils/middlewares"
	"KosKita/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type KosHandler struct {
	kosService kos.KosServiceInterface
	cld        cloudinary.CloudinaryUploaderInterface
}

func New(ks kos.KosServiceInterface, cloudinaryUploader cloudinary.CloudinaryUploaderInterface) *KosHandler {
	return &KosHandler{
		kosService: ks,
		cld:        cloudinaryUploader,
	}
}

func (handler *KosHandler) CreateKos(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	newKos := KosRequest{}
	errBind := c.Bind(&newKos)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	var imageUrls []string
	photoFields := []string{"main_kos_photo", "front_kos_photo", "back_kos_photo", "front_room_photo", "inside_room_photo"}

	for _, field := range photoFields {
		fileHeader, err := c.FormFile(field)
		if err != nil {
			return c.JSON(http.StatusBadRequest, responses.WebResponse("error retrieving the file", nil))
		}

		imageURL, err := handler.cld.UploadImage(fileHeader)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.WebResponse("error uploading the image", nil))
		}

		// Menambahkan URL gambar ke slice
		imageUrls = append(imageUrls, imageURL)
	}

	kosCore := RequestToCore(newKos, imageUrls, uint(userIdLogin))

	errInsert := handler.kosService.Create(userIdLogin, kosCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error create kos - "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert kos", nil))
}

func (handler *KosHandler) UpdateKos(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	kosID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error parsing product id", nil))
	}

	updateKos := KosRequest{}
	errBind := c.Bind(&updateKos)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	var imageUrls []string
	photoFields := []string{"main_kos_photo", "front_kos_photo", "back_kos_photo", "front_room_photo", "inside_room_photo"}

	for _, field := range photoFields {
		fileHeader, err := c.FormFile(field)
		if err != nil {
			return c.JSON(http.StatusBadRequest, responses.WebResponse("error retrieving the file", nil))
		}

		imageURL, err := handler.cld.UploadImage(fileHeader)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.WebResponse("error uploading the image", nil))
		}

		// Menambahkan URL gambar ke slice
		imageUrls = append(imageUrls, imageURL)
	}

	kosCore := RequestToCore(updateKos, imageUrls, uint(userIdLogin))
	kosCore.ID = uint(kosID)
	
	errUpdate := handler.kosService.Put(userIdLogin, kosCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error create kos - "+errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success update kos", nil))
}
package checker

import (
	"context"
	"mime/multipart"
	"net/http"

	"api.kalbe.crm/config/constants"
	"api.kalbe.crm/config/logger"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gofiber/fiber/v2"
)

func ValidateFileHeaderMime(fileHeader *multipart.FileHeader, mime ...string) (resp bool, err error) {

	if fileHeader == nil {
		return false, nil
	}

	file, err := fileHeader.Open()
	if err != nil {
		logger.LogError(context.Background(), err, http.StatusBadRequest)
		return false, fiber.NewError(http.StatusBadRequest, constants.MIME_TYPE_NOT_VALID)
	}

	mimes, err := mimetype.DetectReader(file)
	if err != nil {
		logger.LogError(context.Background(), err, http.StatusBadRequest)
		return false, fiber.NewError(http.StatusBadRequest, constants.MIME_TYPE_NOT_VALID)
	}

	for _, excelMimeType := range mime {
		if mimes.Is(excelMimeType) {
			return true, nil
		}
	}

	return false, fiber.NewError(http.StatusBadRequest, constants.MIME_TYPE_NOT_VALID)
}

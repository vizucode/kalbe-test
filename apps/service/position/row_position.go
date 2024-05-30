package position

import (
	"context"
	"log"
	"net/http"

	"api.kalbe.crm/apps/domain"
	"github.com/gofiber/fiber/v2"
)

func (uc *position) GetPosition(ctx context.Context, id int) (resp domain.Position, err error) {

	positionModel, err := uc.db.GetPosition(ctx, id)
	if err != nil {
		log.Println(err)
		return resp, fiber.NewError(http.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	resp = domain.Position{
		Id:            int(positionModel.ID),
		DepartementId: positionModel.DepartementId,
		PositionName:  positionModel.PositionName,
		CreatedBy:     positionModel.CreatedBy,
		CreatedAt:     positionModel.CreatedAt,
		UpdatedBy:     positionModel.UpdatedBy,
		UpdatedAt:     positionModel.UpdatedAt,
	}

	return resp, nil
}

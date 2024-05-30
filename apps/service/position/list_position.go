package position

import (
	"context"
	"log"
	"net/http"

	"api.kalbe.crm/apps/domain"
	"github.com/gofiber/fiber/v2"
)

func (uc *position) GetAllPosition(ctx context.Context) (resp []domain.Position, err error) {

	positionModelList, err := uc.db.GetAllPosition(ctx)
	if err != nil {
		log.Println(err)
		return resp, fiber.NewError(http.StatusInternalServerError, fiber.ErrInternalServerError.Message)
	}

	for _, v := range positionModelList {
		resp = append(resp, domain.Position{
			Id:            int(v.ID),
			DepartementId: v.DepartementId,
			PositionName:  v.PositionName,
			CreatedBy:     v.CreatedBy,
			CreatedAt:     v.CreatedAt,
			UpdatedBy:     v.UpdatedBy,
			UpdatedAt:     v.UpdatedAt,
		})
	}

	return resp, nil
}

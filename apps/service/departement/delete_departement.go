package departement

import (
	"context"
	"log"
)

func (uc *departement) DeleteDepartement(ctx context.Context, id int) (err error) {
	if err = uc.db.DeleteDepartement(ctx, id); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

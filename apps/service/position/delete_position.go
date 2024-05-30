package position

import (
	"context"
	"log"
)

func (uc *position) DeletePosition(ctx context.Context, id int) (err error) {
	err = uc.db.DeletePosition(ctx, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

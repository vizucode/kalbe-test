package location

import (
	"context"
	"log"
)

func (s *location) DeleteLocation(ctx context.Context, id int) (err error) {
	if err = s.db.DeleteLocation(ctx, id); err != nil {
		log.Println(err)
		return err
	}
	return
}

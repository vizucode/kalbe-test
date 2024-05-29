package timestamp

import (
	"context"
	"net/http"
	"time"

	"api.kalbe.crm/config/logger"
)

func GenerateExpired(expiredDuration time.Duration) int64 {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		logger.LogError(context.TODO(), err, http.StatusInternalServerError)
		return 0
	}

	return time.Now().In(location).Add(expiredDuration).Unix()
}

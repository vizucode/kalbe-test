package timestamp

import (
	"context"
	"fmt"
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

func GenerateCode(runningNumber int) string {
	now := time.Now()
	yymm := now.Format("0601")

	runningNumberStr := fmt.Sprintf("%04d", runningNumber)

	result := yymm + runningNumberStr

	return result
}

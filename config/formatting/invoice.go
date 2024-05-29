package formatting

import (
	"fmt"
	"time"
)

func GenerateInvoice(prefix string) string {
	return fmt.Sprintf("INV%s%d", prefix, GetEpoch())
}

func GetEpoch() int {
	return int(time.Now().Unix())
}

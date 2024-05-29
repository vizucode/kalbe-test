package formatting

import (
	"strings"

	"api.kalbe.crm/config/constants"
)

func ConvertPaymentStatus(target string) (resp string) {
	switch strings.ToLower(target) {
	case strings.ToLower(constants.TRIPAY_PAYMENT_PAID):
		resp = constants.PAID
	case strings.ToLower(constants.TRIPAY_PAYMENT_PENDING):
		resp = constants.PENDING
	case strings.ToLower(constants.TRIPAY_PAYMENT_SUCCESS):
		resp = constants.SUCCESS
	case strings.ToLower(constants.TRIPAY_PAYMENT_FAILED):
		resp = constants.FAILED
	case strings.ToLower(constants.TRIPAY_PAYMENT_UNPAID):
		resp = constants.UNPAID
	case strings.ToLower(constants.TRIPAY_PAYMENT_EXPIRED):
		resp = constants.EXPIRED
	case strings.ToLower(constants.DIGI_SUCCESS):
		resp = constants.SUCCESS
	case strings.ToLower(constants.DIGI_DIPROSES):
		resp = constants.PENDING
	case strings.ToLower(constants.DIGI_PENDING):
		resp = constants.PENDING
	case strings.ToLower(constants.DIGI_FAILED):
		resp = constants.FAILED
	case strings.ToLower(constants.TRANSACTION_REFUND):
		resp = constants.REFUND
	}
	return resp
}

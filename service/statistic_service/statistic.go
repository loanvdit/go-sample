package statistic_service

import (
	"fujitech/web/web-go/go-challenge/models"

	"github.com/jinzhu/gorm"
)

type countReport struct {
	Success int
	Waiting int
}

type amountReport struct {
	Vnd int64
	Jpy int64
	Bth int64
}

type feeReport struct {
	Vnd int64
	Jpy int64
	Bth int64
}

type statisticReport struct {
	Count  countReport
	Amount amountReport
	Fee    feeReport
}

var (
	T_SUCCESS = 2
	C_VND     = 0
	C_JPY     = 1
	C_BTH     = 2
	M_SEND    = 1
	M_RECEIVE = 2
)

func Get(user_id string) (statisticReport, int) {
	result, err := models.GetStatistic(user_id)

	if err != nil && err == gorm.ErrRecordNotFound {
		return statisticReport{}, 0
	}

	var (
		count  countReport
		amount amountReport
		fee    feeReport
	)

	amount.Vnd = 0
	amount.Jpy = 0
	amount.Bth = 0

	fee.Vnd = 0
	fee.Jpy = 0
	fee.Bth = 0

	for _, transaction := range result {
		// Count status
		if transaction.Status == T_SUCCESS {
			count.Success += 1
		} else {
			count.Waiting += 1
		}

		// Count amount, fee
		if transaction.Currency == C_VND { // VND
			if transaction.Mode == M_RECEIVE {
				amount.Vnd += transaction.Amount
				fee.Vnd += transaction.Fee
			} else if transaction.Mode == M_SEND {
				amount.Vnd -= transaction.Amount
				fee.Vnd -= transaction.Fee
			}

		} else if transaction.Currency == C_JPY { // JPY
			if transaction.Mode == M_RECEIVE {
				amount.Jpy += transaction.Amount
				fee.Jpy += transaction.Fee
			} else if transaction.Mode == M_SEND {
				amount.Jpy -= transaction.Amount
				fee.Jpy -= transaction.Fee
			}
		} else if transaction.Currency == C_BTH { // BTH
			if transaction.Mode == M_RECEIVE {
				amount.Bth += transaction.Amount
				fee.Bth += transaction.Fee
			} else if transaction.Mode == M_SEND {
				amount.Bth -= transaction.Amount
				fee.Bth -= transaction.Fee
			}
		}
	}

	return statisticReport{
		Count:  count,
		Amount: amount,
		Fee:    fee,
	}, len(result)
}

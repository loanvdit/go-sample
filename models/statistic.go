package models

type Statistic struct {
	ID       string
	Email    string
	Amount   int64
	Fee      int64
	Status   int
	Mode     int
	Currency int
}

// Get report details
func GetStatistic(user_id string) ([]*Statistic, error) {
	var st []*Statistic

	err := db.Table("fuji_user").Select("fuji_transaction.id, fuji_user.email, fuji_transaction.amount, fuji_transaction.fee, fuji_transaction.status, fuji_transaction.mode, fuji_transaction.currency").
		Joins("join fuji_transaction on fuji_transaction.user_id = fuji_user.id").
		Where("fuji_user.id=?", user_id).Find(&st).Error

	return st, err
}

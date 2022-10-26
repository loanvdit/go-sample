package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "Invalid params",

	ERROR_SEND_FAIL:        "Send money fail",
	ERROR_RECEIVE_FAIL:     "Receive money fail",
	ERROR_TRANSACTION_FAIL: "Transaction fail",
	ERROR_UPLOAD_FILE:      "Can not upload file",

	ERROR_GET_HISTORY_FAIL:               "Get transaction history fail",
	ERROR_COUNT_TRANSACTION_HISTORY_FAIL: "Count transaction history fail",

	ERROR_EXIST_FAIL: "Not exist",
	ERROR_GET_FAIL:   "Get fail",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token fail",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token timeout",
	ERROR_AUTH_TOKEN:               "Token generation failed",
	ERROR_AUTH:                     "Token error",
	ERROR_AUTH_INVALID_PARAM:       "Email or password was invalid",

	SUCCESS_ACTIVE_USER: "Active user was successful",

	ERROR_ACCOUNT_EXIST:     "Account is existing",
	ERROR_PARSE_DATE:        "Can not parse date time",
	ERROR_TOKEN_TIMEOUT:     "Token timeout! Please re-try.",
	ERROR_ACTIVE_USER:       "Can not active this user",
	ERROR_ACCOUNT_NOT_EXIST: "Account not exist",

	SUCCESS_DELETE_USER: "Delete partner was successful",

	ERROR_PARTNER_EXIST:     "Partner is existing",
	ERROR_PARTNER_NOT_EXIST: "Partner is not exist",

	SUCCESS_ADD_TRANSACTION:           "Add transaction was successful",
	SUCCESS_EDIT_TRANSACTION:          "Edit transaction was successful",
	SUCCESS_DELETE_TRANSACTION:        "Delete transaction was successful",
	SUCCESS_UPDATE_STATUS_TRANSACTION: "Update status transaction was successful",

	ERROR_ADD_TRANSACTION:           "Add transaction has error",
	ERROR_EDIT_TRANSACTION:          "Edit transaction has error",
	ERROR_DELETE_TRANSACTION:        "Delete transaction has error",
	ERROR_UPDATE_STATUS_TRANSACTION: "Update status transaction has error",
	ERROR_NOT_EXIST_TRANSACTION:     "Transaction not exist",
	ERROR_NOT_STATUS_TRANSACTION:    "Transaction status not exist",

	ERROR_STATISTIC_NO_RESULT: "Not found any transaction",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

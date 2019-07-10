package errutil

import "errors"

var (
	ErrIllegalParameter      = errors.New("illegal parameter")
	ErrDBOperation           = errors.New("database operation failed")
	ErrPermissionDenied      = errors.New("permission denied")
	ErrInvalidParameter      = errors.New("invalid parameter")
	ErrThirdAccountNotFound  = errors.New("third account not found")
	ErrIllegalDeskStatus     = errors.New("illegal desk status")
	ErrPlayerNotFound        = errors.New("player not found")
	ErrUserNotFound          = errors.New("user not found")
	ErrRequestPrepayIDFailed = errors.New("request prepay id failed")
	ErrOrderNotFound         = errors.New("order not found")
	ErrTradeExisted          = errors.New("trade has existed")
	ErrHistoryNotFound       = errors.New("history not found")
	ErrDeskNotFound          = errors.New("desk not found")
)

package db

import (
	"goer/db/model"
	"goer/pkg/errutil"
)

func InsertHistory(h *model.History) error {
	if h == nil {
		return errutil.ErrInvalidParameter
	}
	_, err := database.Insert(h)
	if err != nil {
		return errutil.ErrDBOperation
	}
	return nil
}

func QueryHistory(id int64) (*model.History, error) {
	h := &model.History{Id: id}
	has, err := database.Get(h)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if !has {
		return nil, errutil.ErrHistoryNotFound
	}
	return h, nil
}

func QueryHistoriesByDeskID(deskId int64) ([]model.History, int, error) {
	result := make([]model.History, 0)
	err := database.Where("desk_id=?", deskId).Desc("begin_at").Find(&result)
	if err != nil {
		logger.Error(err.Error())
		return nil, 0, errutil.ErrDBOperation
	}
	return result, len(result), nil
}

package db

import (
	"goer/db/model"
	"goer/pkg/algoutil"
	"goer/pkg/errutil"
)

const (
	noLimitFlag  = -1 //如果count=-1则表示返回所有数据
	noTimeFilter = -1 //如果start/end == -1 则表示无时间筛选
)

func QueryOrder(orderID string) (*model.Order, error) {
	order := &model.Order{OrderId: orderID}
	has, err := database.Get(order)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errutil.ErrOrderNotFound
	}
	return order, nil
}

func InsertOrder(r *model.Order) error {
	if r == nil {
		return errutil.ErrInvalidParameter
	}
	_, err := database.Insert(r)
	if err != nil {
		return errutil.ErrDBOperation
	}
	return nil
}

func OrderList(uid int64, appid, channelID, orderID, payBy string, start, end int64, status, offset, count int) ([]model.Order, int, error) {
	order := &model.Order{
		AppId:       appid,
		ChannelId:   channelID,
		Uid:         uid,
		OrderId:     orderID,
		PayPlatform: payBy,
		Status:      status,
	}

	start, end = algoutil.TimeRange(start, end)
	total, err := database.Where("created_at BETWEEN ? AND ?", start, end).Count(order)
	if err != nil {
		logger.Error(err)
		return nil, 0, errutil.ErrDBOperation
	}
	result := make([]model.Order, 0)
	if count == noLimitFlag {
		err = database.Where("created_at BETWEEN ? AND ?", start, end).
			Desc("id").Find(&result, order)
	} else {
		err = database.Where("created_at BETWEEN ? AND ?", start, end).
			Desc("id").Limit(count, offset).Find(&result, order)
	}
	if err != nil {
		return nil, 0, errutil.ErrDBOperation
	}
	return result, int(total), nil
}

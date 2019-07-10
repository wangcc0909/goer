package db

import (
	"fmt"
	"goer/db/model"
	"goer/pkg/errutil"
)

func InsertTrade(t *model.Trade) error {
	logger.Info("insert trade, order id: ", t.OrderId)

	trade := &model.Trade{OrderId: t.OrderId}
	has, err := database.Get(trade)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if has {
		return errutil.ErrTradeExisted
	}
	order, err := QueryOrder(t.OrderId)
	if err != nil {
		return err
	}
	if order.Type == OrderTypeBuyToken {
		order.Status = OrderStatusNotified
	} else {
		order.Status = OrderStatusPayed
	}
	sess := database.NewSession()

	//开始事务
	sess.Begin()
	defer sess.Close()
	if _, err := sess.Insert(t); err != nil {
		fmt.Println(err.Error())
		sess.Rollback()
		return err
	}
	if _, err := sess.Where("order_id = ?", order.OrderId).Update(order); err != nil {
		fmt.Println(err.Error())
		sess.Rollback()
		return err
	}
	u := &model.User{}
	//todo 个人认为这里是id=? 不是uid=?
	sess.Where("uid = ?", order.Uid).Get(u)

	//添加首充时间
	if u.FirstRechargeAt == 0 {
		u.FirstRechargeAt = order.CreatedAt
		if _, err := sess.ID(u.Id).Update(u); err != nil {
			fmt.Println(err.Error())
			sess.Rollback()
			return err
		}
	}
	return sess.Commit()
}

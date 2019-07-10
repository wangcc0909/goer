package db

import (
	"github.com/go-xorm/xorm"
	log "github.com/sirupsen/logrus"
	"goer/db/model"
	"time"
)

const asyncTaskBacklog = 128

var (
	database *xorm.Engine
	logger   *log.Entry
	chWrite  chan interface{} //async write channel
	chUpdate chan interface{} //async update channel
)

type options struct {
	showSQL      bool
	maxOpenConns int
	maxIdleConns int
}

type ModelOption func(*options)

func MaxIdleConns(i int) ModelOption {
	return func(o *options) {
		o.maxIdleConns = i
	}
}

func MaxOpenConns(i int) ModelOption {
	return func(o *options) {
		o.maxOpenConns = i
	}
}

func ShowSQL(show bool) ModelOption {
	return func(o *options) {
		o.showSQL = show
	}
}

func envInit() {
	//async tack
	go func() {
		for {
			select {
			case t, ok := <-chWrite:
				if !ok {
					return
				}
				if _, err := database.Insert(t); err != nil {
					logger.Error(err)
				}
			case t, ok := <-chUpdate:
				if !ok {
					return
				}
				if _, err := database.Update(t); err != nil {
					logger.Error(err)
				}
			}
		}
	}()

	//定时ping数据库,保持连接池连接
	go func() {
		ticker := time.NewTicker(time.Minute * 5)
		for {
			select {
			case <-ticker.C:
				database.Ping()
			}
		}

	}()
}

//New create the database's connection
func MustStartup(dsn string, opts ...ModelOption) func() {
	logger = log.WithField("component", "model")
	settings := &options{
		maxOpenConns: defaultMaxConns,
		maxIdleConns: defaultMaxConns,
		showSQL:      true,
	}
	for _, opt := range opts {
		opt(settings)
	}
	logger.Infof("DSN=%s ShowSQL=%t MaxIdleConn=%v MaxOpenConn=%v", dsn, settings.showSQL, settings.maxIdleConns, settings.maxOpenConns)
	//create database instance
	if db, err := xorm.NewEngine("mysql", dsn); err != nil {
		panic(err)
	} else {
		database = db
	}
	//设置日志相关
	database.SetLogger(&Logger{Entry: logger.WithField("orm", "xorm")})

	chWrite = make(chan interface{}, asyncTaskBacklog)
	chUpdate = make(chan interface{}, asyncTaskBacklog)

	//options
	database.SetMaxIdleConns(settings.maxIdleConns)
	database.SetMaxOpenConns(settings.maxOpenConns)
	database.ShowSQL(settings.showSQL)

	syncSchema()
	envInit()
	Closer := func() {
		close(chWrite)
		close(chUpdate)
		_ = database.Close()
		logger.Info("stopped")
	}
	return Closer
}

func syncSchema() {
	_ = database.StoreEngine("InnoDB").Sync2(
		new(model.CardConsume),
		new(model.Desk),
		new(model.History),
		new(model.Login),
		new(model.Online),
		new(model.Order),
		new(model.Register),
		new(model.ThirdAccount),
		new(model.Trade),
		new(model.User),
		new(model.Club),
		new(model.UserClub),
	)
}

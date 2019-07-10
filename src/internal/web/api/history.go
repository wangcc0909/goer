package api

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/lonnng/nex"
	"goer/db"
	"goer/pkg/errutil"
	"goer/pkg/whitelist"
	"goer/protocol"
	"net/http"
	"strconv"
	"time"
)

const (
	format = "01-02 15:04:05"
)

func MakeHistoryService() http.Handler {
	router := mux.NewRouter()
	router.Handle("/v1/history/lite/{desk_id}", nex.Handler(historyList)).Methods("GET") //获取历史列表(lite),参数为desk_id
	router.Handle("/v1/history/{id}", nex.Handler(historyByID)).Methods("GET")           //获取历史记录
	return router
}

func HistoryByID(id int64) (*protocol.History, error) {
	p, err := db.QueryHistory(id)
	if err != nil {
		return nil, err
	}
	return &protocol.History{
		HistoryLite: protocol.HistoryLite{
			Id:           p.Id,
			DeskId:       p.DeskId,
			BeginAt:      p.BeginAt,
			Mode:         p.Mode,
			BeginAtStr:   time.Unix(p.BeginAt, 0).Format(format),
			EndAt:        p.EndAt,
			PlayName0:    p.PlayerName0,
			PlayName1:    p.PlayerName1,
			PlayName2:    p.PlayerName2,
			PlayName3:    p.PlayerName3,
			ScoreChange0: p.ScoreChange0,
			ScoreChange1: p.ScoreChange1,
			ScoreChange2: p.ScoreChange2,
			ScoreChange3: p.ScoreChange3,
		},
		Snapshot: p.SnapShot,
	}, nil
}

func HistoryLiteList(deskId int64) ([]protocol.HistoryLite, int64, error) {
	//默认全部
	ps, total, err := db.QueryHistoriesByDeskID(deskId)
	if err != nil {
		return nil, 0, err
	}
	list := make([]protocol.HistoryLite, total)
	for i, p := range ps {
		beginAtStr := time.Unix(p.BeginAt, 0).Format(format)
		list[i] = protocol.HistoryLite{
			Id:           p.Id,
			DeskId:       p.DeskId,
			BeginAt:      p.BeginAt,
			BeginAtStr:   beginAtStr,
			EndAt:        p.EndAt,
			PlayName0:    p.PlayerName0,
			PlayName1:    p.PlayerName1,
			PlayName2:    p.PlayerName2,
			PlayName3:    p.PlayerName3,
			ScoreChange0: p.ScoreChange0,
			ScoreChange1: p.ScoreChange1,
			ScoreChange2: p.ScoreChange2,
			ScoreChange3: p.ScoreChange3,
		}
	}
	return list, int64(len(list)), nil
}

func historyList(_ context.Context, r *http.Request) (*protocol.HistoryLiteListResponse, error) {
	if !whitelist.VerifyIP(r.RemoteAddr) {
		return nil, errutil.ErrPermissionDenied
	}

	vars := mux.Vars(r)
	idStr, ok := vars["desk_id"]
	if !ok || idStr == "" {
		return nil, errutil.ErrInvalidParameter
	}
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		return nil, errutil.ErrInvalidParameter
	}
	list, t, err := HistoryLiteList(id)
	if err != nil {
		return nil, err
	}
	return &protocol.HistoryLiteListResponse{Data: list, Total: t}, nil
}

func historyByID(r *http.Request) (interface{}, error) {
	if !whitelist.VerifyIP(r.RemoteAddr) {
		return nil, errutil.ErrPermissionDenied
	}
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok || idStr == "" {
		return nil, errutil.ErrIllegalParameter
	}

	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		return nil, errutil.ErrInvalidParameter
	}
	h, err := HistoryByID(id)
	if err != nil {
		return nil, err
	}
	return protocol.HistoryByIDResponse{Data: h}, nil
}

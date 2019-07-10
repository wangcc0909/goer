package game

import (
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
	"goer/db"
	"goer/pkg/async"
	"goer/protocol"
)

type ClubManager struct {
	component.Base
}

var defaultClubManager = NewClubManager()

func NewClubManager() *ClubManager {
	return &ClubManager{}
}

func (c *ClubManager) ApplyClub(s *session.Session, payload *protocol.ApplyClubRequest) error {
	mid := s.LastMid()
	logger.Debugf("玩家申请加入俱乐部,UID=%d, 俱乐部ID=%d", s.UID(), payload.ClubId)
	async.Run(func() {
		if err := db.ApplyClub(s.UID(), payload.ClubId); err != nil {
			s.ResponseMID(mid, &protocol.ErrorResponse{
				Code:  -1,
				Error: err.Error(),
			})
		} else {
			s.ResponseMID(mid, &protocol.SuccessResponse)
		}
	})
	return nil
}

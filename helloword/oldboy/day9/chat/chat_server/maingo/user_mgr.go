package maingo

import "learngo/learngo/helloword/oldboy/day9/chat/chat_server/model"

var mgr *model.UserMgr

func InitUserMgr(){
	mgr = model.NewUserMgr(pool)
}
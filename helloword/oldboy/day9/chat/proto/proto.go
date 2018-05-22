package proto

import "learngo/learngo/helloword/oldboy/day9/chat/model"

type Messge struct {
	Cmd		string `json:"cmd"`
	Data	string `json:"data"`
}

type LoginCmd struct {
	Id 		int		`json:"id"`
	Passwd	string	`json:"passwd"`
}

type RegisterCmd struct {
	User model.User	`json:"user"`
}

type LoginCmdRes struct {
	Code	int 	`json:"code"`
	Error	string  `json:"error"`
}

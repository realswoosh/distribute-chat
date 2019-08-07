package handler

import (
	"sync"
)


type msgHandler struct {
	handler map[string]interface{}
}

var instance *msgHandler
var once sync.Once

func GetInstance() *msgHandler {
	once.Do(func() {
		instance = &msgHandler{ handler: map[string]interface{} {
			"test" : msgHandlerTest,
		}}
	})
	return instance
}

func (handler* msgHandler) MsgHandler(msg string, msgBody []byte) {

}

func msgHandlerTest(msgBody []byte) {

}
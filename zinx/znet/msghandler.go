package znet

import (
	"fmt"
	"go_code/zinx/utils"
	"go_code/zinx/ziface"
)

type MsgHandler struct {
	Apis           map[uint32]ziface.IRouter
	TaskQueue      []chan ziface.IRequest
	WorkerPoolSize uint32
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]ziface.IRouter, 0),
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
	}
}

func (mh *MsgHandler) DoMsgHandler(req ziface.IRequest) {
	handler, ok := mh.Apis[req.GetMsgId()]
	if !ok {
		fmt.Errorf("MsgId :%v is not found", req.GetMsgId())
		return
	}
	handler.PreHandle(req)
	handler.Handle(req)
	handler.PostHandle(req)
}

func (mh *MsgHandler) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := mh.Apis[msgId]; ok {
		fmt.Errorf("MsgId :%v is exist", msgId)
		return
	}
	mh.Apis[msgId] = router
}

func (mh *MsgHandler) StartWorkerPool() {
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskSize)
		go mh.startOneWorker(i, mh.TaskQueue[i])
	}
}

func (mh *MsgHandler) startOneWorker(workerId int, taskQueue chan ziface.IRequest) {
	for {
		select {
		case req := <-taskQueue:
			fmt.Println("workerId :", workerId, "is working...")
			mh.DoMsgHandler(req)
		}
	}
}

func (mh *MsgHandler) SendReqToTaskQueue(req ziface.IRequest) {
	workerId := req.GetConnection().GetConnID() % mh.WorkerPoolSize
	mh.TaskQueue[workerId] <- req
}

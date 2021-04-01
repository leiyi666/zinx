package znet

import (
	"errors"
	"fmt"
	"go_code/zinx/ziface"
	"sync"
)

type ConnManager struct {
	Connections map[uint32]ziface.IConnection
	ConnLock    sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		Connections: make(map[uint32]ziface.IConnection),
	}
}

func (cm *ConnManager) Add(conn ziface.IConnection) {
	cm.ConnLock.Lock()
	defer cm.ConnLock.Unlock()

	cm.Connections[conn.GetConnID()] = conn
	fmt.Println("connection add to server succ, connId :", conn.GetConnID())
}

func (cm *ConnManager) Remove(conn ziface.IConnection) {
	cm.ConnLock.Lock()
	defer cm.ConnLock.Unlock()

	delete(cm.Connections, conn.GetConnID())
	fmt.Println("connection delete succ, connId :", conn.GetConnID())
}

func (cm *ConnManager) Get(connId uint32) (ziface.IConnection, error) {
	cm.ConnLock.RLock()
	defer cm.ConnLock.RUnlock()

	if conn, ok := cm.Connections[connId]; ok {
		return conn, nil
	}
	return nil, errors.New("Connection not found!")
}

func (cm *ConnManager) Len() int {
	return len(cm.Connections)
}

func (cm *ConnManager) ClearConn() {
	cm.ConnLock.Lock()
	defer cm.ConnLock.Unlock()

	for connId, conn := range cm.Connections {
		conn.Stop()
		delete(cm.Connections, connId)
	}
	fmt.Println("Clear all connections succ")
}

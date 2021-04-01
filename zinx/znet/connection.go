package znet

import (
	"errors"
	"fmt"
	"go_code/zinx/utils"
	"go_code/zinx/ziface"
	"io"
	"net"
	"sync"
)

type Connection struct {
	Server       ziface.IServer
	Conn         *net.TCPConn
	ConnID       uint32
	isClosed     bool
	MsgHandler   ziface.IMsgHandler
	MsgChan      chan []byte
	ExitChan     chan bool
	Property     map[string]interface{}
	PropertyLock sync.RWMutex
}

func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandler) *Connection {
	c := &Connection{
		Server:     server,
		Conn:       conn,
		ConnID:     connID,
		MsgHandler: msgHandler,
		isClosed:   false,
		MsgChan:    make(chan []byte),
		ExitChan:   make(chan bool, 1),
		Property:   make(map[string]interface{}),
	}
	c.Server.GetConnMgr().Add(c)
	return c

}

func (c *Connection) Start() {
	fmt.Println("Conn start().. ConnID = ", c.ConnID)
	//TODO 启动写数据的业务
	go c.Reader()
	go c.Writer()

	c.Server.CallOnConnStart(c)
}

func (c *Connection) Stop() {
	defer c.Conn.Close()
	fmt.Println("Conn stop().. ConnID = ", c.ConnID)
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	c.ExitChan <- true
	c.Server.GetConnMgr().Remove(c)
	c.Server.CallOnConnStop(c)
	close(c.ExitChan)
	close(c.MsgChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) SendMsg(MsgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("connection closed when send msg")
	}
	dp := NewDataPkg()
	binaryMsg, err := dp.Pack(NewMessage(MsgId, data))
	if err != nil {
		fmt.Errorf("connection SendMsg Pack err :%v\n", err)
		return err
	}
	c.MsgChan <- binaryMsg
	return nil
}

func (c *Connection) Reader() {
	fmt.Println("Reader goroutine is running...")
	defer c.Stop()
	for {
		dp := NewDataPkg()
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.Conn, headData); err != nil {
			fmt.Errorf("connection reader headData err :%v\n", err)
			return
		}
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Errorf("connection reader Unpack err :%v\n", err)
			return
		}
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.Conn, data); err != nil {
				fmt.Errorf("connection reader data Unpack err :%v\n", err)
				return
			}
		}
		msg.SetData(data)

		req := Request{
			Conn: c,
			Msg:  msg,
		}
		if utils.GlobalObject.WorkerPoolSize > 0 {
			go c.MsgHandler.SendReqToTaskQueue(&req)
		} else {
			go c.MsgHandler.DoMsgHandler(&req)
		}
	}
}

func (c *Connection) Writer() {
	for {
		select {
		case data := <-c.MsgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Errorf("connection writer err :%v\n", err)
				return
			}
		case <-c.ExitChan:
			return
		}
	}
}

func (c *Connection) SetProperty(key string, value interface{}) {
	c.PropertyLock.Lock()
	defer c.PropertyLock.Unlock()

	c.Property[key] = value
	fmt.Println("SetProperty key :", key, " value :", c.Property[key])
}

func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.PropertyLock.RLock()
	defer c.PropertyLock.RUnlock()

	if value, ok := c.Property[key]; ok {
		return value, nil
	}
	return nil, errors.New("GetProperty not found")
}

func (c *Connection) RemoveProperty(key string) {
	c.PropertyLock.Lock()
	defer c.PropertyLock.Unlock()

	delete(c.Property, key)
}

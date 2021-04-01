package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDatapkg_Pack(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Errorf("server create err :%v\n", err)
		return
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Errorf("server listen err :%v\n", err)
				return
			}
			go func(conn net.Conn) {
				dp := Datapkg{}
				for {
					data := make([]byte, dp.GetHeadLen())
					if _, err := io.ReadFull(conn, data); err != nil {
						fmt.Errorf("server read err :%v\n", err)
						return
					}
					msg, err := dp.Unpack(data)
					if err != nil {
						fmt.Errorf("server Unpack err :%v\n", err)
						return
					}
					if msg.GetMsgLen() > 0 {
						msgData := msg.(*Message)
						msgData.Data = make([]byte, msg.GetMsgLen())
						if _, err := io.ReadFull(conn, msgData.Data); err != nil {
							fmt.Errorf("server Unpack err :%v\n", err)
							return
						}
						fmt.Println("Id: ", msgData.Id, "Len: ", msgData.DataLen, "Data: ", string(msgData.Data))
					}
				}
			}(conn)
		}
	}()

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Errorf("client dial err :%v\n", err)
		return
	}
	dp := Datapkg{}
	msg1 := &Message{
		Id:      1,
		DataLen: 4,
		Data:    []byte{'z', 'i', 'n', 'x'},
	}
	bytes1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Errorf("client Pack err :%v\n", err)
		return
	}
	msg2 := &Message{
		Id:      2,
		DataLen: 7,
		Data:    []byte{'n', 'i', 'h', 'a', 'o', '!', '!'},
	}
	bytes2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Errorf("client Pack err :%v\n", err)
		return
	}
	bytes1 = append(bytes1, bytes2...)
	if _, err := conn.Write(bytes1); err != nil {
		fmt.Errorf("client write err :%v\n", err)
		return
	}
	select {}
}

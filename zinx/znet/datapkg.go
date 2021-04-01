package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"go_code/zinx/utils"
	"go_code/zinx/ziface"
)

type Datapkg struct {
}

func NewDataPkg() *Datapkg {
	return &Datapkg{}
}

func (dp *Datapkg) GetHeadLen() uint32 {
	return 8
}

func (dp *Datapkg) Pack(msg ziface.IMessage) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (dp *Datapkg) Unpack(binaryData []byte) (ziface.IMessage, error) {
	buf := bytes.NewReader(binaryData)
	msg := &Message{}
	if err := binary.Read(buf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too large mes to rev")
	}
	return msg, nil
}

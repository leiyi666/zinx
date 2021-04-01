package ziface

type IMessage interface {
	GetMsgId() uint32
	GetMsgLen() uint32
	GetData() []byte
	SetMsgId(id uint32)
	SetMsgLen(dataLen uint32)
	SetData(data []byte)
}

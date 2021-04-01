package ziface

type IDataPkg interface {
	GetHeadLen() uint32
	Pack(msg IMessage) ([]byte, error)
	Unpack([]byte) (IMessage, error)
}

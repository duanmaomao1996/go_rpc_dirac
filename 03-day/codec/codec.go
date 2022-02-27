package codec

import "io"

//格式
type Header struct {
	ServiceMeth  string //方法
	Seq			 uint64
	Err          string
}

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

var NewCodecFuncMap map[Type]NewCodecFunc

const (
	GobType Type = "application/gob"
	JsonType Type = "application/json"
)

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}



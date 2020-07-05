package mircodec

import (
	"encoding/binary"
	"errors"
	"reflect"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	"github.com/davyxu/golog"
)

var log = golog.New("codec.mirCodec")

var (
	ErrInvalidType = errors.New("invalid type")
	ErrOutOfData   = errors.New("out of data")
)

func init() {
	codec.RegisterCodec(new(mirCodec))
}

// mirCodec 编码解码
type mirCodec struct{}

// Name 返回名字
func (m *mirCodec) Name() string {
	return "mirCodec"
}

// MimeType 我也不知道是干嘛的
func (m *mirCodec) MimeType() string {
	return "application/binary"
}

// Encode 将结构体转换为字节数组
// 作用是将 server.Packet 转换为字节发送给客户端
func (*mirCodec) Encode(msgObj interface{}, ctx cellnet.ContextSet) (data interface{}, err error) {

	v := reflect.Indirect(reflect.ValueOf(msgObj))
	size := dataSize(v, nil)
	if size < 0 {
		return nil, ErrInvalidType
	}

	buf := make([]byte, size)

	e := &encoder{order: binary.LittleEndian, buf: buf}
	e.value(v)

	return buf, nil
}

// Decode 将字节数组转换为结构体
// 作用是将客户端发来的字节转换为 client.Packet
func (*mirCodec) Decode(bytes interface{}, msgObj interface{}) error {
	data := bytes.([]byte)

	if len(data) == 0 {
		return nil
	}

	v := reflect.ValueOf(msgObj)

	switch v.Kind() {
	case reflect.Ptr:
		v = v.Elem()
	}

	size := dataSize(v, nil)
	if size < 0 {
		return ErrInvalidType
	}

	if len(data) < size {
		return ErrOutOfData
	}

	d := &decoder{order: binary.LittleEndian, buf: data}
	d.value(v)

	return nil
}

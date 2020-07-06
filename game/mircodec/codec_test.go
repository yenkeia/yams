package mircodec

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/yenkeia/yams/game/proto/client"
	"github.com/yenkeia/yams/game/ut"
)

func TestCodec(t *testing.T) {
	buf := make([]byte, 12)
	e := &encoder{order: binary.LittleEndian, buf: buf}
	fmt.Println(buf)

	e.uint32(uint32(123))
	fmt.Println(buf)

	e.uint32(uint32(321))
	fmt.Println(buf)

	fmt.Println(ut.StringToBytes("123"))

	// e.int32(int32(123))
	// fmt.Println(buf)

	e.string("123")
	fmt.Println(buf)
}

func TestDecodeClientPacket(t *testing.T) {
	bytes := []byte{
		// 54, 0, 3, 0,
		7, 97, 99, 99, 111, 117, 110, 116, 8, 112, 97, 115, 115, 119, 111, 114, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 113, 117, 101, 115, 116, 105, 111, 110, 6, 97, 110, 115, 119, 101, 114, 7, 97, 100, 100, 114, 101, 115, 115}
	msg := new(client.NewAccount)
	codec := new(mirCodec)
	if err := codec.Decode(bytes, msg); err != nil {
		panic(err)
	}
	t.Log(msg.SecretQuestion)
	t.Log(msg.SecretAnswer)
	t.Log(msg.EMailAddress)
}

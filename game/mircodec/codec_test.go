package mircodec

import (
	"encoding/binary"
	"fmt"
	"testing"

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

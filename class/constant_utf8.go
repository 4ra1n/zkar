package class

import (
	"encoding/binary"
	"fmt"
	"github.com/phith0n/zkar/commons"
)

type ConstantUTF8 struct {
	Data string
}

func (c *ConstantUTF8) ToBytes() []byte {
	var bs = []byte{ConstantUtf8Info}

	// integer overflow
	bs = append(bs, commons.NumberToBytes(uint16(len(c.Data)))...)
	bs = append(bs, []byte(c.Data)...)
	return bs
}

func (cf *ClassFile) readConstantUTF8(stream *commons.Stream) (*ConstantUTF8, error) {
	_, _ = stream.ReadN(1)
	bs, err := stream.ReadN(2)
	if err != nil {
		return nil, fmt.Errorf("read constant utf8 size failed, no enough data in the stream")
	}

	length := binary.BigEndian.Uint16(bs)
	data, err := stream.ReadN(int(length))
	if err != nil {
		return nil, fmt.Errorf("read constant utf8 failed, no enough data in the stream")
	}

	c := &ConstantUTF8{
		Data: string(data),
	}
	return c, nil
}

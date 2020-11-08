package iso8583

import (
	"bytes"
	"strconv"
	"strings"
)

type buffer struct {
	byteBuff *bytes.Buffer
}

func newBuffer(data []byte) *buffer {
	buffer := &buffer{}
	buffer.byteBuff = bytes.NewBuffer(data)
	return buffer
}

//Read ...
func (b *buffer) Read(length int) []byte {
	if b.byteBuff == nil {
		b.byteBuff = &bytes.Buffer{}
	}
	data := make([]byte, length)
	if _, e := b.byteBuff.Read(data); e != nil {
		println(e.Error())
	}
	return data
}

//ReadString ...
func (b *buffer) ReadString(length int) string {
	return strings.TrimSpace(string(b.Read(length)))
}

//ReadInt
func (b *buffer) ReadInt(length int) int {
	str := b.ReadString(length)
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return i
}

func (b *buffer) writeString(data string) *buffer {
	if b.byteBuff == nil {
		b.byteBuff = &bytes.Buffer{}
	}
	b.byteBuff.WriteString(data)
	return b
}

func (b *buffer) write(data []byte) *buffer {
	if b.byteBuff == nil {
		b.byteBuff = &bytes.Buffer{}
	}
	b.byteBuff.Write(data)
	return b
}

func (b *buffer) bytes() []byte {
	return b.byteBuff.Bytes()
}

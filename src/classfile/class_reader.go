package classfile

import "encoding/binary"

// 类似于工具类，读取各种各样长度块的数据
type ClassReader struct {
	data []byte
}

// u1
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

// u2
func (self *ClassReader) readUint16() uint16 {
	// 读一位的时候不存在大端小端问题，两位以上则存在高位在前还是在后的问题
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

// u4
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

func (self *ClassReader) readUint16s() []uint16 {
	// 读取一个short作为数组长度
	n := self.readUint16()
	// 然后读取指定长度个short
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

func (self *ClassReader) readBytes(n uint32) []byte {
	// 读取指定长度
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}

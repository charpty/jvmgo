package classfile

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/
// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.1
type ConstantClassInfo struct {
	cp        ConstantPool
	// 存储class存储的位置索引
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}
func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

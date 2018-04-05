package classfile

import "fmt"

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	// 常量池第一个short存储了常量池元素的个数
	// 每一个常量都有一个byte标志位和具体的数据组成
	cpCount := int(reader.readUint16())
	// 创建一个动态数组，因为有特殊情况
	cp := make([]ConstantInfo, cpCount)

	// The constant_pool table is indexed from 1 to constant_pool_count - 1.
	for i := 1; i < cpCount-1; i++ {
		fmt.Println(i)
		fmt.Println(cpCount)
		cp[i] = readConstantInfo(reader, cp)
		// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.5
		// 这就是个数的特殊情况，读到long和double时，必须下一个元素是个空，以兼容老版本
		// 这是由于一个byte占常量池2个位置
		switch cp[i].(type) {
		//case *ConstantLongInfo, *ConstantDoubleInfo:
		//	i++
		}
	}

	return cp
}

func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic(fmt.Errorf("Invalid constant pool index: %v!", index))
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

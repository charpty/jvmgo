package classfile

import "fmt"

/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type ClassFile struct {
	//magic      uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		// 函数返回前判断是否有异常发生
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	// 构建一个classReader
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader) {
	// 读取版本信息
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	// 读取常量池，动长
	self.constantPool = readConstantPool(reader)
	// 访问标志，是一个位图标记，记录了类的访问级别，类是否为final，是否是注解类型等等
	self.accessFlags = reader.readUint16()
	// 当前类名在常量池中的索引
	self.thisClass = reader.readUint16()
	// 当前类父类名在常量池中的索引
	self.superClass = reader.readUint16()
	// 读取该类实现的所有的接口
	self.interfaces = reader.readUint16s()
	// 读取当前类的属性，包括静态属性
	self.fields = readMembers(reader, self.constantPool)
	// 读取当前类的方法信息，包括静态方法
	self.methods = readMembers(reader, self.constantPool)
	// 读取剩余的不包含在方法或者字段里的其它属性表信息
	self.attributes = readAttributes(reader, self.constantPool)
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	// 读取编译JDK版本信息
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	// 1.1
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		// 到JDK8止
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

package classfile

/*
MethodParameters_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 parameters_count;
    {   u2 name_index;
        u2 access_flags;
    } parameters[parameters_count];
}
 */

// JDK8以后可以指定编译器保留形参的名称
type MethodParameters struct {
	cp         ConstantPool
	parameters []*MethodParameter
}

type MethodParameter struct {
	nameIndex   uint16
	accessFlags uint16
}

func (self *MethodParameters) readInfo(reader *ClassReader) {
	parametersCount := reader.readUint8()
	parameters := make([]*MethodParameter, parametersCount)
	for i := range self.parameters {
		parameters[i] = readMethodParameter(reader)
	}
	self.parameters = parameters
}

func readMethodParameter(reader *ClassReader) *MethodParameter {
	return &MethodParameter{
		nameIndex:   reader.readUint16(),
		accessFlags: reader.readUint16(),
	}
}

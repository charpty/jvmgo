package classfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7
// 一共23种属性表
/*
SourceFile
InnerClasses
EnclosingMethod
SourceDebugExtension
BootstrapMethods
ConstantValue
Code
Exceptions
RuntimeVisibleParameterAnnotations, RuntimeInvisibleParameterAnnotations
AnnotationDefault
MethodParameters
Synthetic
// Deprecated
Signature
RuntimeVisibleAnnotations, RuntimeInvisibleAnnotations
LineNumberTable
LocalVariableTable
LocalVariableTypeTable
StackMapTable
RuntimeVisibleTypeAnnotations, RuntimeInvisibleTypeAnnotations
 */

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	case "MethodParameters":
		return &MethodParameters{cp: cp}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}

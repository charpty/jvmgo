package heap

type Constant interface{}

type ConstantPool struct {
	class     *Class
	constants []Constant
}

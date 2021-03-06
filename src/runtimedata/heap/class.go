package heap

import (
	"strings"
	"classfile"
)

type Class struct {
	accessFlags       uint16
	name              string // thisClassName
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initStarted       bool
	jClass            *Object
}

func newClass(cf *classfile.ClassFile) *Class {
	r := &Class{
		accessFlags:    cf.AccessFlags(),
		name:           cf.ClassName(),
		superClassName: cf.SuperClassName(),
	}
	r.constantPool = newConstantPool(r, cf.ConstantPool())
	r.fields = newFields(r, cf.Fields())
	r.methods = newMethods(r, cf.Methods())
	r.interfaceNames = cf.InterfaceNames()
	return r
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

func (self *Class) isJlObject() bool {
	return self.name == "java/lang/Object"
}
func (self *Class) isJlCloneable() bool {
	return self.name == "java/lang/Cloneable"
}
func (self *Class) isJioSerializable() bool {
	return self.name == "java/io/Serializable"
}

func (self *Class) Name() string {
	return self.name
}

func (self *Class) SuperClass() *Class {
	return self.superClass
}

func (self *Class) Loader() *ClassLoader {
	return self.loader
}

// getters
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
}

// jvms 5.4.4
func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() ||
		self.getPackageName() == other.getPackageName()
}

func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) GetField(name string, descriptor string) *Field {
	for c := self; c != nil; c = c.superClass {
		for _, f := range c.fields {
			if f.name == name && f.descriptor == descriptor && !f.IsStatic() {
				return f
			}
		}
	}
	return nil
}

func (self *Class) GetStaticField(name string, descriptor string) *Field {
	for c := self; c != nil; c = c.superClass {
		for _, f := range c.fields {
			if f.name == name && f.descriptor == descriptor && f.IsStatic() {
				return f
			}
		}
	}
	return nil
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

func (self *Class) GetMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if !method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

func (self *Class) isAssignableFrom(child *Class) bool {
	p, c := self, child
	if p == c {
		return true
	}
	if c.IsArray() {
		if p.IsArray() {
			cc := c.ComponentClass()
			pc := p.ComponentClass()
			return cc == pc || pc.isAssignableFrom(cc)
		} else {
			// 如果父不是数组存在两种特殊情况，父是Object或者是cloneable接口或序列化接口
			if p.IsInterface() {
				return p.isJlCloneable() || p.isJioSerializable()
			} else {
				return p.isJlObject()
			}
		}
	} else {
		if p.IsInterface() {
			return c.IsImplements(p)
		} else {
			return c.IsSubClassOf(p)
		}
	}
}

func (self *Class) IsSubClassOf(parent *Class) bool {
	if self.IsInterface() {
		return parent.isJlObject()
	}
	for c := self; c != nil; c = c.superClass {
		if c == parent {
			return true
		}
	}
	return false
}

func (self *Class) IsImplements(parent *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, intfc := range c.interfaces {
			if intfc == parent || intfc.IsSubInterfaceOf(parent) {
				return true
			}
		}
	}
	return false
}

func (self *Class) IsSubInterfaceOf(parent *Class) bool {
	for _, c := range self.interfaces {
		if c == parent || c.IsSubInterfaceOf(parent) {
			return true;
		}
	}
	return false
}

func (self *Class) JClass() *Object {
	return self.jClass
}

func (self *Class) JavaName() string {
	return strings.Replace(self.name, "/", ".", -1)
}

func (self *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[self.name]
	return ok
}

func (self *Class) GetRefVar(fieldName, fieldDescriptor string) *Object {
	field := self.GetStaticField(fieldName, fieldDescriptor)
	return self.staticVars.GetRef(field.slotId)
}
func (self *Class) SetRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := self.GetStaticField(fieldName, fieldDescriptor)
	self.staticVars.SetRef(field.slotId, ref)
}

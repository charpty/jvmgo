package heap

import (
	"unicode/utf16"
	"util"
)

var internedStrings = map[string]*Object{}

func JString(loader *ClassLoader, goStr string) *Object {
	if result, ok := internedStrings[goStr]; ok {
		return result
	}
	// stringToUtf16
	chars := utf16.Encode([]rune(goStr))
	charArrClass := loader.LoadClass("[C")
	strValue := &Object{charArrClass, chars}
	javaStringObject := loader.LoadClass("java/lang/String").NewObject()
	javaStringObject.SetRefValue("value", "[C", strValue)
	util.Debug("success create string: %s", goStr)
	internedStrings[goStr] = javaStringObject
	return javaStringObject
}

func GoString(javaStringObject *Object) string {
	strValue := javaStringObject.GetRefValue("value", "[C")
	return string(utf16.Decode(strValue.data.([]uint16)))
}

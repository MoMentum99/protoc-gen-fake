package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func getDefaultValueForField(field *protogen.Field) string {
	switch field.Desc.Kind() {
	case protoreflect.StringKind:
		return `"fake_` + field.GoName + `"`
	case protoreflect.Int32Kind, protoreflect.Int64Kind:
		return "42"
	case protoreflect.BoolKind:
		return "true"
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return "3.14"
	case protoreflect.EnumKind:
		return "0"
	case protoreflect.MessageKind:
		if field.Message.Desc.FullName() == "google.protobuf.Timestamp" {
			return "nil" // 또는 현재 시간으로 설정
		}
		return "nil"
	case protoreflect.BytesKind:
		return "[]byte(\"fake_bytes\")"
	}
	return ""
}

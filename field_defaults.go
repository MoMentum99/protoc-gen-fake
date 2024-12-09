package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func getDefaultValueForField(field *protogen.Field) string {
	// Proto3 optional field check - this is the key change
	if field.Desc.HasPresence() || field.Desc.Syntax() == protoreflect.Proto3 && field.Desc.HasOptionalKeyword() {
		return "nil"
	}

	switch field.Desc.Kind() {
	case protoreflect.StringKind:
		return fmt.Sprintf("\"%s\"", "fake_"+field.GoName)
	case protoreflect.Int32Kind, protoreflect.Int64Kind,
		protoreflect.Uint32Kind, protoreflect.Uint64Kind,
		protoreflect.Sint32Kind, protoreflect.Sint64Kind,
		protoreflect.Fixed32Kind, protoreflect.Fixed64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
		return "42"
	case protoreflect.BoolKind:
		return "true"
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return "3.14"
	case protoreflect.EnumKind:
		return "0"
	case protoreflect.MessageKind:
		return "nil"
	case protoreflect.BytesKind:
		return "[]byte(\"fake_bytes\")"
	}
	return ""
}

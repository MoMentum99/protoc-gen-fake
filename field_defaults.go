package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func getDefaultValueForField(field *protogen.Field) string {
	// Proto3 optional field check
	if field.Desc.HasPresence() || field.Desc.Syntax() == protoreflect.Proto3 && field.Desc.HasOptionalKeyword() {
		return "nil"
	}

	// Repeated field check
	if field.Desc.IsList() {
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			return fmt.Sprintf("[]string{\"fake_%s_1\", \"fake_%s_2\"}", field.GoName, field.GoName)
		case protoreflect.Int32Kind, protoreflect.Int64Kind,
			protoreflect.Uint32Kind, protoreflect.Uint64Kind,
			protoreflect.Sint32Kind, protoreflect.Sint64Kind,
			protoreflect.Fixed32Kind, protoreflect.Fixed64Kind,
			protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
			return "[]int32{42, 43}"
		case protoreflect.BoolKind:
			return "[]bool{true, false}"
		case protoreflect.FloatKind, protoreflect.DoubleKind:
			return "[]float64{3.14, 2.71}"
		case protoreflect.BytesKind:
			return "[][]byte{[]byte(\"fake_bytes_1\"), []byte(\"fake_bytes_2\")}"
		case protoreflect.EnumKind:
			return fmt.Sprintf("[]%s{0, 1}", field.Enum.GoIdent)
		case protoreflect.MessageKind:
			return "nil"
		}
		return "nil"
	}

	// Regular field check
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
		if field.Message.Desc.FullName() == "google.protobuf.Empty" {
			return "&emptypb.Empty{}"
		}
		if field.Message.Desc.FullName() == "google.protobuf.Timestamp" {
			return "nil"
		}
		return "nil"
	case protoreflect.BytesKind:
		return "[]byte(\"fake_bytes\")"
	}
	return ""
}

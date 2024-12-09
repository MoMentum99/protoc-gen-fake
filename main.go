package main

import (
	"flag"
	"google.golang.org/protobuf/types/pluginpb"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var flags flag.FlagSet

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			GenerateFake(gen, f)
		}
		return nil
	})
}

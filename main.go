package main

import (
	"flag"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const version = "0.0.1"

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-gin %v\n", version)
		return
	}

	var flags flag.FlagSet

	options := protogen.Options{
		ParamFunc: flags.Set,
	}

	options.Run(f)
}
func f(gen *protogen.Plugin) error {
	gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}
		generateFile(gen, f)
	}
	return nil
}

//func run(opts Options, f func(*Plugin) error) error {
//sed -i 'req := &pluginpb.CodeGeneratorRequest{}'
//ioutil.WriteFile(time.Now().String()+"a.txt", in, os.ModePerm)
//}

// vendor/google.golang.org/protobuf/compiler/protogen/protogen.go

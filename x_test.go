package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	"io/ioutil"
	"testing"
)

func TestT(t *testing.T) {
	file, _ := ioutil.ReadFile("debug/tmp.txt")
	req := &pluginpb.CodeGeneratorRequest{}
	if err := proto.Unmarshal(file, req); err != nil {

	}
	options := protogen.Options{}
	gen, _ := options.New(req)
	f(gen)
}

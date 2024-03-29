// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpcar/options/annotations.proto

package options

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

var E_Service = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*Service)(nil),
	Field:         11042,
	Name:          "wxio.grpcar.options.service",
	Tag:           "bytes,11042,opt,name=service",
	Filename:      "grpcar/options/annotations.proto",
}

var E_Method = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*Method)(nil),
	Field:         11042,
	Name:          "wxio.grpcar.options.method",
	Tag:           "bytes,11042,opt,name=method",
	Filename:      "grpcar/options/annotations.proto",
}

var E_Message = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*Message)(nil),
	Field:         11042,
	Name:          "wxio.grpcar.options.message",
	Tag:           "bytes,11042,opt,name=message",
	Filename:      "grpcar/options/annotations.proto",
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*Field)(nil),
	Field:         11042,
	Name:          "wxio.grpcar.options.field",
	Tag:           "bytes,11042,opt,name=field",
	Filename:      "grpcar/options/annotations.proto",
}

func init() {
	proto.RegisterExtension(E_Service)
	proto.RegisterExtension(E_Method)
	proto.RegisterExtension(E_Message)
	proto.RegisterExtension(E_Field)
}

func init() { proto.RegisterFile("grpcar/options/annotations.proto", fileDescriptor_99ca6c4fffd4c911) }

var fileDescriptor_99ca6c4fffd4c911 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4f, 0xc5, 0x20,
	0x14, 0x85, 0xe3, 0xe0, 0x33, 0xa9, 0x5b, 0x5d, 0xb4, 0xea, 0xb3, 0xa3, 0xd3, 0x25, 0xd1, 0xad,
	0xa3, 0x83, 0xdb, 0x8b, 0xa6, 0x1a, 0x63, 0xdc, 0x28, 0xe5, 0x51, 0x92, 0xd7, 0x5e, 0x02, 0x54,
	0xfd, 0x3d, 0xfe, 0x52, 0x53, 0x6e, 0x19, 0x6c, 0x19, 0xdc, 0x08, 0xe7, 0x3b, 0x1f, 0x27, 0x21,
	0x2b, 0x95, 0x35, 0x82, 0x5b, 0x86, 0xc6, 0x6b, 0x1c, 0x1c, 0xe3, 0xc3, 0x80, 0x9e, 0x87, 0x33,
	0x18, 0x8b, 0x1e, 0xf3, 0xb3, 0xaf, 0x6f, 0x8d, 0x40, 0x18, 0xcc, 0x58, 0x51, 0x2a, 0x44, 0x75,
	0x90, 0x2c, 0x20, 0xcd, 0xb8, 0x67, 0xad, 0x74, 0xc2, 0x6a, 0xe3, 0xd1, 0x52, 0xad, 0xb8, 0x58,
	0x8a, 0x47, 0xdf, 0x51, 0x54, 0xbd, 0x67, 0x27, 0x4e, 0xda, 0x4f, 0x2d, 0x64, 0x7e, 0x03, 0x24,
	0x82, 0x28, 0x82, 0x17, 0x4a, 0x9e, 0xa8, 0x76, 0xfe, 0xf3, 0x56, 0x1e, 0xdd, 0x9e, 0xde, 0x5d,
	0x41, 0x62, 0x45, 0x64, 0xeb, 0xa8, 0xab, 0x5e, 0xb3, 0x4d, 0x2f, 0x7d, 0x87, 0x6d, 0xbe, 0x5d,
	0x89, 0x77, 0x21, 0x58, 0x78, 0x2f, 0x93, 0x5e, 0x42, 0xeb, 0xd9, 0x35, 0xed, 0xed, 0xa5, 0x73,
	0x5c, 0xa5, 0xf6, 0xee, 0x28, 0xf9, 0xd7, 0xde, 0x99, 0xad, 0xa3, 0xae, 0x7a, 0xce, 0x8e, 0xf7,
	0x5a, 0x1e, 0xda, 0xfc, 0x7a, 0xe5, 0x7d, 0x9c, 0xee, 0x17, 0xd6, 0x22, 0x69, 0x0d, 0x64, 0x4d,
	0xa2, 0x87, 0xf2, 0x63, 0xab, 0xb4, 0xef, 0xc6, 0x06, 0x04, 0xf6, 0x6c, 0xc2, 0xd9, 0xdf, 0x8f,
	0x68, 0x36, 0xe1, 0x89, 0xfb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x21, 0xc4, 0xd4, 0xfa,
	0x01, 0x00, 0x00,
}

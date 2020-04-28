//
//Copyright 2018 Google LLC
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//https://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.6.1
// source: proto/gnmitest/gnmitest.proto

package gnmitest

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	report "github.com/openconfig/gnmitest/proto/report"
	suite "github.com/openconfig/gnmitest/proto/suite"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_proto_gnmitest_gnmitest_proto protoreflect.FileDescriptor

var file_proto_gnmitest_gnmitest_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6e, 0x6d, 0x69, 0x74, 0x65, 0x73, 0x74,
	0x2f, 0x67, 0x6e, 0x6d, 0x69, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x67, 0x6e, 0x6d, 0x69, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x67, 0x6e, 0x6d, 0x69, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x75, 0x69, 0x74, 0x65, 0x2f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x6e, 0x6d, 0x69, 0x74, 0x65, 0x73,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2f, 0x0a, 0x08, 0x47,
	0x4e, 0x4d, 0x49, 0x54, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x0c,
	0x2e, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2e, 0x53, 0x75, 0x69, 0x74, 0x65, 0x1a, 0x0e, 0x2e, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_gnmitest_gnmitest_proto_goTypes = []interface{}{
	(*suite.Suite)(nil),   // 0: suite.Suite
	(*report.Report)(nil), // 1: report.Report
}
var file_proto_gnmitest_gnmitest_proto_depIdxs = []int32{
	0, // 0: gnmitest.GNMITest.Run:input_type -> suite.Suite
	1, // 1: gnmitest.GNMITest.Run:output_type -> report.Report
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_gnmitest_gnmitest_proto_init() }
func file_proto_gnmitest_gnmitest_proto_init() {
	if File_proto_gnmitest_gnmitest_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_gnmitest_gnmitest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_gnmitest_gnmitest_proto_goTypes,
		DependencyIndexes: file_proto_gnmitest_gnmitest_proto_depIdxs,
	}.Build()
	File_proto_gnmitest_gnmitest_proto = out.File
	file_proto_gnmitest_gnmitest_proto_rawDesc = nil
	file_proto_gnmitest_gnmitest_proto_goTypes = nil
	file_proto_gnmitest_gnmitest_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GNMITestClient is the client API for GNMITest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GNMITestClient interface {
	// Run runs the given Suite proto and returns Report proto.
	Run(ctx context.Context, in *suite.Suite, opts ...grpc.CallOption) (*report.Report, error)
}

type gNMITestClient struct {
	cc grpc.ClientConnInterface
}

func NewGNMITestClient(cc grpc.ClientConnInterface) GNMITestClient {
	return &gNMITestClient{cc}
}

func (c *gNMITestClient) Run(ctx context.Context, in *suite.Suite, opts ...grpc.CallOption) (*report.Report, error) {
	out := new(report.Report)
	err := c.cc.Invoke(ctx, "/gnmitest.GNMITest/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GNMITestServer is the server API for GNMITest service.
type GNMITestServer interface {
	// Run runs the given Suite proto and returns Report proto.
	Run(context.Context, *suite.Suite) (*report.Report, error)
}

// UnimplementedGNMITestServer can be embedded to have forward compatible implementations.
type UnimplementedGNMITestServer struct {
}

func (*UnimplementedGNMITestServer) Run(context.Context, *suite.Suite) (*report.Report, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}

func RegisterGNMITestServer(s *grpc.Server, srv GNMITestServer) {
	s.RegisterService(&_GNMITest_serviceDesc, srv)
}

func _GNMITest_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(suite.Suite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GNMITestServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnmitest.GNMITest/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GNMITestServer).Run(ctx, req.(*suite.Suite))
	}
	return interceptor(ctx, in, info, handler)
}

var _GNMITest_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gnmitest.GNMITest",
	HandlerType: (*GNMITestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _GNMITest_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/gnmitest/gnmitest.proto",
}

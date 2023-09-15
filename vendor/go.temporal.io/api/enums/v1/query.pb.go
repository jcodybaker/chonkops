// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/api/enums/v1/query.proto

package enums

import (
	fmt "fmt"
	math "math"
	strconv "strconv"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryResultType int32

const (
	QUERY_RESULT_TYPE_UNSPECIFIED QueryResultType = 0
	QUERY_RESULT_TYPE_ANSWERED    QueryResultType = 1
	QUERY_RESULT_TYPE_FAILED      QueryResultType = 2
)

var QueryResultType_name = map[int32]string{
	0: "Unspecified",
	1: "Answered",
	2: "Failed",
}

var QueryResultType_value = map[string]int32{
	"Unspecified": 0,
	"Answered":    1,
	"Failed":      2,
}

func (QueryResultType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b9a616a97224ce1d, []int{0}
}

type QueryRejectCondition int32

const (
	QUERY_REJECT_CONDITION_UNSPECIFIED QueryRejectCondition = 0
	// None indicates that query should not be rejected.
	QUERY_REJECT_CONDITION_NONE QueryRejectCondition = 1
	// NotOpen indicates that query should be rejected if workflow is not open.
	QUERY_REJECT_CONDITION_NOT_OPEN QueryRejectCondition = 2
	// NotCompletedCleanly indicates that query should be rejected if workflow did not complete cleanly.
	QUERY_REJECT_CONDITION_NOT_COMPLETED_CLEANLY QueryRejectCondition = 3
)

var QueryRejectCondition_name = map[int32]string{
	0: "Unspecified",
	1: "None",
	2: "NotOpen",
	3: "NotCompletedCleanly",
}

var QueryRejectCondition_value = map[string]int32{
	"Unspecified":         0,
	"None":                1,
	"NotOpen":             2,
	"NotCompletedCleanly": 3,
}

func (QueryRejectCondition) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b9a616a97224ce1d, []int{1}
}

func init() {
	proto.RegisterEnum("temporal.api.enums.v1.QueryResultType", QueryResultType_name, QueryResultType_value)
	proto.RegisterEnum("temporal.api.enums.v1.QueryRejectCondition", QueryRejectCondition_name, QueryRejectCondition_value)
}

func init() { proto.RegisterFile("temporal/api/enums/v1/query.proto", fileDescriptor_b9a616a97224ce1d) }

var fileDescriptor_b9a616a97224ce1d = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd2, 0xcf, 0x4e, 0xe2, 0x40,
	0x1c, 0x07, 0xf0, 0x4e, 0x37, 0xd9, 0xc3, 0x5c, 0xb6, 0x69, 0x76, 0xb3, 0x04, 0x75, 0x08, 0x9a,
	0x78, 0x20, 0xa6, 0xb5, 0xf1, 0x56, 0x4f, 0xa5, 0x1d, 0x92, 0x9a, 0x3a, 0x2d, 0x65, 0xc0, 0xe0,
	0xa5, 0x41, 0x6d, 0xcc, 0x18, 0x60, 0x6a, 0x29, 0x24, 0xdc, 0x7c, 0x04, 0x1f, 0xc3, 0xf8, 0x00,
	0x3c, 0x83, 0x47, 0x8e, 0x1c, 0xa5, 0x5c, 0x8c, 0x27, 0x1e, 0xc1, 0xd0, 0x40, 0xfc, 0x83, 0x78,
	0x9b, 0xe4, 0xfb, 0x99, 0xf9, 0x25, 0xbf, 0xef, 0xc0, 0x62, 0x12, 0x76, 0x22, 0x1e, 0xb7, 0xda,
	0x6a, 0x2b, 0x62, 0x6a, 0xd8, 0xed, 0x77, 0x7a, 0xea, 0x40, 0x53, 0x6f, 0xfb, 0x61, 0x3c, 0x54,
	0xa2, 0x98, 0x27, 0x5c, 0xfe, 0xb7, 0x22, 0x4a, 0x2b, 0x62, 0x4a, 0x46, 0x94, 0x81, 0x56, 0x8a,
	0xe1, 0x9f, 0xea, 0x42, 0xf9, 0x61, 0xaf, 0xdf, 0x4e, 0xe8, 0x30, 0x0a, 0xe5, 0x22, 0xdc, 0xa9,
	0xd6, 0xb1, 0xdf, 0x0c, 0x7c, 0x5c, 0xab, 0x3b, 0x34, 0xa0, 0x4d, 0x0f, 0x07, 0x75, 0x52, 0xf3,
	0xb0, 0x69, 0x57, 0x6c, 0x6c, 0x49, 0x82, 0x8c, 0x60, 0x7e, 0x9d, 0x18, 0xa4, 0x76, 0x86, 0x7d,
	0x6c, 0x49, 0x40, 0xde, 0x86, 0xb9, 0xf5, 0xbc, 0x62, 0xd8, 0x0e, 0xb6, 0x24, 0xb1, 0x34, 0x02,
	0xf0, 0xef, 0x72, 0xe8, 0x4d, 0x78, 0x99, 0x98, 0xbc, 0x7b, 0xc5, 0x12, 0xc6, 0xbb, 0xf2, 0x3e,
	0xdc, 0x5d, 0x5d, 0x3b, 0xc1, 0x26, 0x0d, 0x4c, 0x97, 0x58, 0x36, 0xb5, 0x5d, 0xf2, 0x65, 0x7c,
	0x01, 0x6e, 0x6d, 0x70, 0xc4, 0x25, 0x58, 0x02, 0xf2, 0x1e, 0x2c, 0x6c, 0x04, 0x34, 0x70, 0x3d,
	0x4c, 0x24, 0x51, 0x3e, 0x84, 0x07, 0x3f, 0x20, 0xd3, 0x3d, 0xf5, 0x1c, 0x4c, 0xb1, 0x15, 0x98,
	0x0e, 0x36, 0x88, 0xd3, 0x94, 0x7e, 0x95, 0x47, 0x60, 0x3c, 0x45, 0xc2, 0x64, 0x8a, 0x84, 0xf9,
	0x14, 0x81, 0xbb, 0x14, 0x81, 0x87, 0x14, 0x81, 0xa7, 0x14, 0x81, 0x71, 0x8a, 0xc0, 0x73, 0x8a,
	0xc0, 0x4b, 0x8a, 0x84, 0x79, 0x8a, 0xc0, 0xfd, 0x0c, 0x09, 0xe3, 0x19, 0x12, 0x26, 0x33, 0x24,
	0xc0, 0x1c, 0xe3, 0xca, 0xb7, 0xdb, 0x2f, 0xc3, 0x6c, 0x0d, 0xde, 0xa2, 0x20, 0x0f, 0x9c, 0x17,
	0xaf, 0x3f, 0x38, 0xc6, 0x3f, 0x75, 0x79, 0x9c, 0x1d, 0x1e, 0xc5, 0xff, 0x74, 0x09, 0x18, 0x57,
	0x8c, 0x88, 0x29, 0x38, 0x7b, 0xaa, 0xa1, 0xbd, 0x8a, 0xf9, 0xf7, 0x44, 0xd7, 0x8d, 0x88, 0xe9,
	0x7a, 0x96, 0xe9, 0x7a, 0x43, 0xbb, 0xf8, 0x9d, 0xfd, 0x81, 0xa3, 0xb7, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x97, 0x88, 0x9e, 0x7b, 0x28, 0x02, 0x00, 0x00,
}

func (x QueryResultType) String() string {
	s, ok := QueryResultType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x QueryRejectCondition) String() string {
	s, ok := QueryRejectCondition_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

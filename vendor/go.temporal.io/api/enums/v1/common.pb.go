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
// source: temporal/api/enums/v1/common.proto

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

type EncodingType int32

const (
	ENCODING_TYPE_UNSPECIFIED EncodingType = 0
	ENCODING_TYPE_PROTO3      EncodingType = 1
	ENCODING_TYPE_JSON        EncodingType = 2
)

var EncodingType_name = map[int32]string{
	0: "Unspecified",
	1: "Proto3",
	2: "Json",
}

var EncodingType_value = map[string]int32{
	"Unspecified": 0,
	"Proto3":      1,
	"Json":        2,
}

func (EncodingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_768283dde248a0f8, []int{0}
}

type IndexedValueType int32

const (
	INDEXED_VALUE_TYPE_UNSPECIFIED  IndexedValueType = 0
	INDEXED_VALUE_TYPE_TEXT         IndexedValueType = 1
	INDEXED_VALUE_TYPE_KEYWORD      IndexedValueType = 2
	INDEXED_VALUE_TYPE_INT          IndexedValueType = 3
	INDEXED_VALUE_TYPE_DOUBLE       IndexedValueType = 4
	INDEXED_VALUE_TYPE_BOOL         IndexedValueType = 5
	INDEXED_VALUE_TYPE_DATETIME     IndexedValueType = 6
	INDEXED_VALUE_TYPE_KEYWORD_LIST IndexedValueType = 7
)

var IndexedValueType_name = map[int32]string{
	0: "Unspecified",
	1: "Text",
	2: "Keyword",
	3: "Int",
	4: "Double",
	5: "Bool",
	6: "Datetime",
	7: "KeywordList",
}

var IndexedValueType_value = map[string]int32{
	"Unspecified": 0,
	"Text":        1,
	"Keyword":     2,
	"Int":         3,
	"Double":      4,
	"Bool":        5,
	"Datetime":    6,
	"KeywordList": 7,
}

func (IndexedValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_768283dde248a0f8, []int{1}
}

type Severity int32

const (
	SEVERITY_UNSPECIFIED Severity = 0
	SEVERITY_HIGH        Severity = 1
	SEVERITY_MEDIUM      Severity = 2
	SEVERITY_LOW         Severity = 3
)

var Severity_name = map[int32]string{
	0: "Unspecified",
	1: "High",
	2: "Medium",
	3: "Low",
}

var Severity_value = map[string]int32{
	"Unspecified": 0,
	"High":        1,
	"Medium":      2,
	"Low":         3,
}

func (Severity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_768283dde248a0f8, []int{2}
}

func init() {
	proto.RegisterEnum("temporal.api.enums.v1.EncodingType", EncodingType_name, EncodingType_value)
	proto.RegisterEnum("temporal.api.enums.v1.IndexedValueType", IndexedValueType_name, IndexedValueType_value)
	proto.RegisterEnum("temporal.api.enums.v1.Severity", Severity_name, Severity_value)
}

func init() {
	proto.RegisterFile("temporal/api/enums/v1/common.proto", fileDescriptor_768283dde248a0f8)
}

var fileDescriptor_768283dde248a0f8 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x3d, 0x2e, 0x14, 0xf4, 0x28, 0x62, 0x18, 0xa0, 0x2d, 0xa9, 0x98, 0x8a, 0xb2, 0xcb,
	0xc2, 0x51, 0xd4, 0x9d, 0x59, 0x25, 0xf1, 0xa3, 0x1d, 0x70, 0x6c, 0x2b, 0x9e, 0xb8, 0x0d, 0x0b,
	0xac, 0xd0, 0x58, 0xd5, 0x48, 0x8d, 0xc7, 0x0a, 0x69, 0x44, 0x77, 0x1c, 0x01, 0x6e, 0x81, 0xb8,
	0x00, 0x57, 0x60, 0x99, 0x65, 0x97, 0xc4, 0xd9, 0x20, 0x56, 0x3d, 0x02, 0x8a, 0x05, 0x81, 0x20,
	0x77, 0x67, 0xbd, 0xef, 0xf7, 0xfc, 0x33, 0x4f, 0x1f, 0xec, 0x8d, 0x93, 0x61, 0xa6, 0x47, 0xfd,
	0xb3, 0x5a, 0x3f, 0x53, 0xb5, 0x24, 0x3d, 0x1f, 0xbe, 0xab, 0x4d, 0xea, 0xb5, 0x13, 0x3d, 0x1c,
	0xea, 0xd4, 0xca, 0x46, 0x7a, 0xac, 0xd9, 0xa3, 0x3f, 0x19, 0xab, 0x9f, 0x29, 0xab, 0xc8, 0x58,
	0x93, 0x7a, 0x35, 0x86, 0x0d, 0x4c, 0x4f, 0xf4, 0x40, 0xa5, 0xa7, 0xf2, 0x22, 0x4b, 0xd8, 0x13,
	0x78, 0x8c, 0x5e, 0xcb, 0x77, 0x84, 0x77, 0x10, 0xcb, 0x5e, 0x80, 0x71, 0xd7, 0x0b, 0x03, 0x6c,
	0x89, 0x17, 0x02, 0x1d, 0x6a, 0xb0, 0x6d, 0x78, 0xb8, 0x8a, 0x83, 0x8e, 0x2f, 0xfd, 0x7d, 0x4a,
	0xd8, 0x26, 0xb0, 0x55, 0xf2, 0x32, 0xf4, 0x3d, 0x6a, 0x56, 0x3f, 0x99, 0x40, 0x45, 0x3a, 0x48,
	0xde, 0x27, 0x83, 0xa8, 0x7f, 0x76, 0x9e, 0x14, 0x2d, 0x7b, 0xc0, 0x85, 0xe7, 0xe0, 0x31, 0x3a,
	0x71, 0xd4, 0x70, 0xbb, 0x58, 0x56, 0xb5, 0x03, 0x5b, 0x25, 0x19, 0x89, 0xc7, 0x92, 0x12, 0xc6,
	0xa1, 0x52, 0x02, 0x5f, 0x61, 0xef, 0xc8, 0xef, 0x38, 0xd4, 0x64, 0x15, 0xd8, 0x2c, 0xe1, 0xc2,
	0x93, 0x74, 0x6d, 0xf1, 0xc4, 0x12, 0xe6, 0xf8, 0xdd, 0xa6, 0x8b, 0xf4, 0xc6, 0x35, 0xbd, 0x4d,
	0xdf, 0x77, 0xe9, 0x4d, 0xb6, 0x0b, 0x3b, 0x65, 0xff, 0x36, 0x24, 0x4a, 0xd1, 0x46, 0xba, 0xce,
	0x9e, 0xc1, 0xee, 0xf5, 0x17, 0x8b, 0x5d, 0x11, 0x4a, 0x7a, 0xab, 0xfa, 0x06, 0x6e, 0x87, 0xc9,
	0x24, 0x19, 0xa9, 0xf1, 0xc5, 0x62, 0xa3, 0x21, 0x46, 0xd8, 0x11, 0xb2, 0xf7, 0xdf, 0x02, 0xee,
	0xc3, 0xdd, 0x25, 0x39, 0x14, 0x07, 0x87, 0x94, 0xb0, 0x07, 0x70, 0x6f, 0x39, 0x6a, 0xa3, 0x23,
	0xba, 0x6d, 0x6a, 0x32, 0x0a, 0x1b, 0xcb, 0xa1, 0xeb, 0x1f, 0xd1, 0xb5, 0xe6, 0x57, 0x32, 0x9d,
	0x71, 0xe3, 0x72, 0xc6, 0x8d, 0xab, 0x19, 0x27, 0x1f, 0x72, 0x4e, 0x3e, 0xe7, 0x9c, 0x7c, 0xcb,
	0x39, 0x99, 0xe6, 0x9c, 0x7c, 0xcf, 0x39, 0xf9, 0x91, 0x73, 0xe3, 0x2a, 0xe7, 0xe4, 0xe3, 0x9c,
	0x1b, 0xd3, 0x39, 0x37, 0x2e, 0xe7, 0xdc, 0x80, 0x6d, 0xa5, 0xad, 0x52, 0x4b, 0x9a, 0x77, 0x5a,
	0x85, 0x4a, 0xc1, 0xc2, 0xa4, 0x80, 0xbc, 0x7e, 0x7a, 0xfa, 0x4f, 0x50, 0xe9, 0x15, 0xeb, 0x9e,
	0x17, 0x1f, 0x5f, 0xcc, 0x2d, 0xf9, 0x3b, 0xa0, 0xb4, 0xd5, 0xc8, 0x94, 0x85, 0xc5, 0x59, 0x51,
	0xfd, 0xa7, 0x59, 0xf9, 0x4b, 0x6c, 0xbb, 0x91, 0x29, 0xdb, 0x2e, 0x98, 0x6d, 0x47, 0xf5, 0xb7,
	0xeb, 0x85, 0xac, 0xfb, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x93, 0xd8, 0x96, 0xd2, 0x02,
	0x00, 0x00,
}

func (x EncodingType) String() string {
	s, ok := EncodingType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x IndexedValueType) String() string {
	s, ok := IndexedValueType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x Severity) String() string {
	s, ok := Severity_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

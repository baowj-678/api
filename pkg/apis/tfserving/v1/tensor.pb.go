//
//     Copyright 2023 The Dragonfly Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: pkg/apis/tfserving/v1/tensor.proto

package tfserving

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Protocol buffer representing a tensor.
type TensorProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dtype DataType `protobuf:"varint,1,opt,name=dtype,proto3,enum=tfserving.v1.DataType" json:"dtype,omitempty"`
	// Shape of the tensor.  TODO(touts): sort out the 0-rank issues.
	TensorShape *TensorShapeProto `protobuf:"bytes,2,opt,name=tensor_shape,json=tensorShape,proto3" json:"tensor_shape,omitempty"`
	// Version number.
	//
	// In version 0, if the "repeated xxx" representations contain only one
	// element, that element is repeated to fill the shape.  This makes it easy
	// to represent a constant Tensor with a single value.
	VersionNumber int32 `protobuf:"varint,3,opt,name=version_number,json=versionNumber,proto3" json:"version_number,omitempty"`
	// Serialized raw tensor content from either Tensor::AsProtoTensorContent or
	// memcpy in tensorflow::grpc::EncodeTensorToByteBuffer. This representation
	// can be used for all tensor types. The purpose of this representation is to
	// reduce serialization overhead during RPC call by avoiding serialization of
	// many repeated small items.
	TensorContent []byte `protobuf:"bytes,4,opt,name=tensor_content,json=tensorContent,proto3" json:"tensor_content,omitempty"`
	// DT_HALF, DT_BFLOAT16. Note that since protobuf has no int16 type, we'll
	// have some pointless zero padding for each value here.
	HalfVal []int32 `protobuf:"varint,13,rep,packed,name=half_val,json=halfVal,proto3" json:"half_val,omitempty"`
	// DT_FLOAT.
	FloatVal []float32 `protobuf:"fixed32,5,rep,packed,name=float_val,json=floatVal,proto3" json:"float_val,omitempty"`
	// DT_DOUBLE.
	DoubleVal []float64 `protobuf:"fixed64,6,rep,packed,name=double_val,json=doubleVal,proto3" json:"double_val,omitempty"`
	// DT_INT32, DT_INT16, DT_UINT16, DT_INT8, DT_UINT8.
	IntVal []int32 `protobuf:"varint,7,rep,packed,name=int_val,json=intVal,proto3" json:"int_val,omitempty"`
	// DT_STRING
	StringVal [][]byte `protobuf:"bytes,8,rep,name=string_val,json=stringVal,proto3" json:"string_val,omitempty"`
	// DT_COMPLEX64. scomplex_val(2*i) and scomplex_val(2*i+1) are real
	// and imaginary parts of i-th single precision complex.
	ScomplexVal []float32 `protobuf:"fixed32,9,rep,packed,name=scomplex_val,json=scomplexVal,proto3" json:"scomplex_val,omitempty"`
	// DT_INT64
	Int64Val []int64 `protobuf:"varint,10,rep,packed,name=int64_val,json=int64Val,proto3" json:"int64_val,omitempty"`
	// DT_BOOL
	BoolVal []bool `protobuf:"varint,11,rep,packed,name=bool_val,json=boolVal,proto3" json:"bool_val,omitempty"`
	// DT_COMPLEX128. dcomplex_val(2*i) and dcomplex_val(2*i+1) are real
	// and imaginary parts of i-th double precision complex.
	DcomplexVal []float64 `protobuf:"fixed64,12,rep,packed,name=dcomplex_val,json=dcomplexVal,proto3" json:"dcomplex_val,omitempty"`
	// DT_RESOURCE
	ResourceHandleVal []*ResourceHandleProto `protobuf:"bytes,14,rep,name=resource_handle_val,json=resourceHandleVal,proto3" json:"resource_handle_val,omitempty"`
	// DT_VARIANT
	VariantVal []*VariantTensorDataProto `protobuf:"bytes,15,rep,name=variant_val,json=variantVal,proto3" json:"variant_val,omitempty"`
	// DT_UINT32
	Uint32Val []uint32 `protobuf:"varint,16,rep,packed,name=uint32_val,json=uint32Val,proto3" json:"uint32_val,omitempty"`
	// DT_UINT64
	Uint64Val []uint64 `protobuf:"varint,17,rep,packed,name=uint64_val,json=uint64Val,proto3" json:"uint64_val,omitempty"`
	// DT_FLOAT8_*, use variable-sized set of bytes
	// (i.e. the equivalent of repeated uint8, if such a thing existed).
	Float8Val []byte `protobuf:"bytes,18,opt,name=float8_val,json=float8Val,proto3" json:"float8_val,omitempty"`
}

func (x *TensorProto) Reset() {
	*x = TensorProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_tfserving_v1_tensor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TensorProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TensorProto) ProtoMessage() {}

func (x *TensorProto) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_tfserving_v1_tensor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TensorProto.ProtoReflect.Descriptor instead.
func (*TensorProto) Descriptor() ([]byte, []int) {
	return file_pkg_apis_tfserving_v1_tensor_proto_rawDescGZIP(), []int{0}
}

func (x *TensorProto) GetDtype() DataType {
	if x != nil {
		return x.Dtype
	}
	return DataType_DT_INVALID
}

func (x *TensorProto) GetTensorShape() *TensorShapeProto {
	if x != nil {
		return x.TensorShape
	}
	return nil
}

func (x *TensorProto) GetVersionNumber() int32 {
	if x != nil {
		return x.VersionNumber
	}
	return 0
}

func (x *TensorProto) GetTensorContent() []byte {
	if x != nil {
		return x.TensorContent
	}
	return nil
}

func (x *TensorProto) GetHalfVal() []int32 {
	if x != nil {
		return x.HalfVal
	}
	return nil
}

func (x *TensorProto) GetFloatVal() []float32 {
	if x != nil {
		return x.FloatVal
	}
	return nil
}

func (x *TensorProto) GetDoubleVal() []float64 {
	if x != nil {
		return x.DoubleVal
	}
	return nil
}

func (x *TensorProto) GetIntVal() []int32 {
	if x != nil {
		return x.IntVal
	}
	return nil
}

func (x *TensorProto) GetStringVal() [][]byte {
	if x != nil {
		return x.StringVal
	}
	return nil
}

func (x *TensorProto) GetScomplexVal() []float32 {
	if x != nil {
		return x.ScomplexVal
	}
	return nil
}

func (x *TensorProto) GetInt64Val() []int64 {
	if x != nil {
		return x.Int64Val
	}
	return nil
}

func (x *TensorProto) GetBoolVal() []bool {
	if x != nil {
		return x.BoolVal
	}
	return nil
}

func (x *TensorProto) GetDcomplexVal() []float64 {
	if x != nil {
		return x.DcomplexVal
	}
	return nil
}

func (x *TensorProto) GetResourceHandleVal() []*ResourceHandleProto {
	if x != nil {
		return x.ResourceHandleVal
	}
	return nil
}

func (x *TensorProto) GetVariantVal() []*VariantTensorDataProto {
	if x != nil {
		return x.VariantVal
	}
	return nil
}

func (x *TensorProto) GetUint32Val() []uint32 {
	if x != nil {
		return x.Uint32Val
	}
	return nil
}

func (x *TensorProto) GetUint64Val() []uint64 {
	if x != nil {
		return x.Uint64Val
	}
	return nil
}

func (x *TensorProto) GetFloat8Val() []byte {
	if x != nil {
		return x.Float8Val
	}
	return nil
}

// Protocol buffer representing the serialization format of DT_VARIANT tensors.
type VariantTensorDataProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the type of objects being serialized.
	TypeName string `protobuf:"bytes,1,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	// Portions of the object that are not Tensors.
	Metadata []byte `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Tensors contained within objects being serialized.
	Tensors []*TensorProto `protobuf:"bytes,3,rep,name=tensors,proto3" json:"tensors,omitempty"`
}

func (x *VariantTensorDataProto) Reset() {
	*x = VariantTensorDataProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_tfserving_v1_tensor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VariantTensorDataProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariantTensorDataProto) ProtoMessage() {}

func (x *VariantTensorDataProto) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_tfserving_v1_tensor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariantTensorDataProto.ProtoReflect.Descriptor instead.
func (*VariantTensorDataProto) Descriptor() ([]byte, []int) {
	return file_pkg_apis_tfserving_v1_tensor_proto_rawDescGZIP(), []int{1}
}

func (x *VariantTensorDataProto) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *VariantTensorDataProto) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *VariantTensorDataProto) GetTensors() []*TensorProto {
	if x != nil {
		return x.Tensors
	}
	return nil
}

var File_pkg_apis_tfserving_v1_tensor_proto protoreflect.FileDescriptor

var file_pkg_apis_tfserving_v1_tensor_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x74, 0x66, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x1a, 0x2b, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x74, 0x66, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x28, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x74, 0x66, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x73, 0x68,
	0x61, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x74, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x05, 0x0a,
	0x0b, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2c, 0x0a, 0x05,
	0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74, 0x66,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x74, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x68, 0x61, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x52, 0x0b, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x68, 0x61, 0x70, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x08, 0x68,
	0x61, 0x6c, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x05, 0x42, 0x02, 0x10,
	0x01, 0x52, 0x07, 0x68, 0x61, 0x6c, 0x66, 0x56, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x09, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x03, 0x28, 0x02, 0x42, 0x02, 0x10,
	0x01, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0a, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x03, 0x28, 0x01, 0x42,
	0x02, 0x10, 0x01, 0x52, 0x09, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x12, 0x1b,
	0x0a, 0x07, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x42,
	0x02, 0x10, 0x01, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x09, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x12, 0x25, 0x0a, 0x0c, 0x73, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x03, 0x28, 0x02,
	0x42, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x73, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x56, 0x61,
	0x6c, 0x12, 0x1f, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x03, 0x42, 0x02, 0x10, 0x01, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x08, 0x42, 0x02, 0x10, 0x01, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x12, 0x25, 0x0a, 0x0c, 0x64, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x76, 0x61,
	0x6c, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x01, 0x42, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x64, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x78, 0x56, 0x61, 0x6c, 0x12, 0x51, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x18,
	0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x12, 0x45, 0x0a, 0x0b, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x74, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0a, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x56,
	0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0a, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c,
	0x18, 0x10, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x02, 0x10, 0x01, 0x52, 0x09, 0x75, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0a, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f,
	0x76, 0x61, 0x6c, 0x18, 0x11, 0x20, 0x03, 0x28, 0x04, 0x42, 0x02, 0x10, 0x01, 0x52, 0x09, 0x75,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x38, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x38, 0x56, 0x61, 0x6c, 0x22, 0x86, 0x01, 0x0a, 0x16, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x74, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x33, 0x0a, 0x07, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74,
	0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x07, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73,
	0x42, 0x2c, 0x5a, 0x2a, 0x64, 0x37, 0x79, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x74, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_apis_tfserving_v1_tensor_proto_rawDescOnce sync.Once
	file_pkg_apis_tfserving_v1_tensor_proto_rawDescData = file_pkg_apis_tfserving_v1_tensor_proto_rawDesc
)

func file_pkg_apis_tfserving_v1_tensor_proto_rawDescGZIP() []byte {
	file_pkg_apis_tfserving_v1_tensor_proto_rawDescOnce.Do(func() {
		file_pkg_apis_tfserving_v1_tensor_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_apis_tfserving_v1_tensor_proto_rawDescData)
	})
	return file_pkg_apis_tfserving_v1_tensor_proto_rawDescData
}

var file_pkg_apis_tfserving_v1_tensor_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_apis_tfserving_v1_tensor_proto_goTypes = []interface{}{
	(*TensorProto)(nil),            // 0: tfserving.v1.TensorProto
	(*VariantTensorDataProto)(nil), // 1: tfserving.v1.VariantTensorDataProto
	(DataType)(0),                  // 2: tfserving.v1.DataType
	(*TensorShapeProto)(nil),       // 3: tfserving.v1.TensorShapeProto
	(*ResourceHandleProto)(nil),    // 4: tfserving.v1.ResourceHandleProto
}
var file_pkg_apis_tfserving_v1_tensor_proto_depIdxs = []int32{
	2, // 0: tfserving.v1.TensorProto.dtype:type_name -> tfserving.v1.DataType
	3, // 1: tfserving.v1.TensorProto.tensor_shape:type_name -> tfserving.v1.TensorShapeProto
	4, // 2: tfserving.v1.TensorProto.resource_handle_val:type_name -> tfserving.v1.ResourceHandleProto
	1, // 3: tfserving.v1.TensorProto.variant_val:type_name -> tfserving.v1.VariantTensorDataProto
	0, // 4: tfserving.v1.VariantTensorDataProto.tensors:type_name -> tfserving.v1.TensorProto
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_apis_tfserving_v1_tensor_proto_init() }
func file_pkg_apis_tfserving_v1_tensor_proto_init() {
	if File_pkg_apis_tfserving_v1_tensor_proto != nil {
		return
	}
	file_pkg_apis_tfserving_v1_resource_handle_proto_init()
	file_pkg_apis_tfserving_v1_tensor_shape_proto_init()
	file_pkg_apis_tfserving_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_apis_tfserving_v1_tensor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TensorProto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_apis_tfserving_v1_tensor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VariantTensorDataProto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_apis_tfserving_v1_tensor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_apis_tfserving_v1_tensor_proto_goTypes,
		DependencyIndexes: file_pkg_apis_tfserving_v1_tensor_proto_depIdxs,
		MessageInfos:      file_pkg_apis_tfserving_v1_tensor_proto_msgTypes,
	}.Build()
	File_pkg_apis_tfserving_v1_tensor_proto = out.File
	file_pkg_apis_tfserving_v1_tensor_proto_rawDesc = nil
	file_pkg_apis_tfserving_v1_tensor_proto_goTypes = nil
	file_pkg_apis_tfserving_v1_tensor_proto_depIdxs = nil
}

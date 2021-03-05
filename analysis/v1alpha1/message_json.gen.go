// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: analysis/v1alpha1/message.proto

// Describes the structure of messages generated by Istio analyzers.

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/golang/protobuf/jsonpb"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/struct"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for AnalysisMessageBase
func (this *AnalysisMessageBase) MarshalJSON() ([]byte, error) {
	str, err := MessageMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AnalysisMessageBase
func (this *AnalysisMessageBase) UnmarshalJSON(b []byte) error {
	return MessageUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for AnalysisMessageBase_Type
func (this *AnalysisMessageBase_Type) MarshalJSON() ([]byte, error) {
	str, err := MessageMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AnalysisMessageBase_Type
func (this *AnalysisMessageBase_Type) UnmarshalJSON(b []byte) error {
	return MessageUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for AnalysisMessageWeakSchema
func (this *AnalysisMessageWeakSchema) MarshalJSON() ([]byte, error) {
	str, err := MessageMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AnalysisMessageWeakSchema
func (this *AnalysisMessageWeakSchema) UnmarshalJSON(b []byte) error {
	return MessageUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for AnalysisMessageWeakSchema_ArgType
func (this *AnalysisMessageWeakSchema_ArgType) MarshalJSON() ([]byte, error) {
	str, err := MessageMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AnalysisMessageWeakSchema_ArgType
func (this *AnalysisMessageWeakSchema_ArgType) UnmarshalJSON(b []byte) error {
	return MessageUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GenericAnalysisMessage
func (this *GenericAnalysisMessage) MarshalJSON() ([]byte, error) {
	str, err := MessageMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GenericAnalysisMessage
func (this *GenericAnalysisMessage) UnmarshalJSON(b []byte) error {
	return MessageUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for InternalErrorAnalysisMessage
func (this *InternalErrorAnalysisMessage) MarshalJSON() ([]byte, error) {
	str, err := MessageMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for InternalErrorAnalysisMessage
func (this *InternalErrorAnalysisMessage) UnmarshalJSON(b []byte) error {
	return MessageUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	MessageMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	MessageUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{AllowUnknownFields: true}
)

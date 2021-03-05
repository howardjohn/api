// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security/v1beta1/jwt.proto

package v1beta1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for JWTRule
func (this *JWTRule) MarshalJSON() ([]byte, error) {
	str, err := JwtMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for JWTRule
func (this *JWTRule) UnmarshalJSON(b []byte) error {
	return JwtUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for JWTHeader
func (this *JWTHeader) MarshalJSON() ([]byte, error) {
	str, err := JwtMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for JWTHeader
func (this *JWTHeader) UnmarshalJSON(b []byte) error {
	return JwtUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	JwtMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	JwtUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{AllowUnknownFields: true}
)

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/k8s_http_route.proto

// $schema: istio.networking.v1alpha3.KubernetesHTTPRoute
// $title: Kubernetes HTTPRoute
// $description: Alpha
// $location: https://istio.io/docs/reference/config/networking/kubernetes-httproute.html

package v1alpha3

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for KubernetesHTTPRoute
func (this *KubernetesHTTPRoute) MarshalJSON() ([]byte, error) {
	str, err := K8SHttpRouteMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for KubernetesHTTPRoute
func (this *KubernetesHTTPRoute) UnmarshalJSON(b []byte) error {
	return K8SHttpRouteUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for KubernetesHTTPRouteRule
func (this *KubernetesHTTPRouteRule) MarshalJSON() ([]byte, error) {
	str, err := K8SHttpRouteMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for KubernetesHTTPRouteRule
func (this *KubernetesHTTPRouteRule) UnmarshalJSON(b []byte) error {
	return K8SHttpRouteUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for KubernetesHTTPRouteActionBackend
func (this *KubernetesHTTPRouteActionBackend) MarshalJSON() ([]byte, error) {
	str, err := K8SHttpRouteMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for KubernetesHTTPRouteActionBackend
func (this *KubernetesHTTPRouteActionBackend) UnmarshalJSON(b []byte) error {
	return K8SHttpRouteUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	K8SHttpRouteMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	K8SHttpRouteUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)

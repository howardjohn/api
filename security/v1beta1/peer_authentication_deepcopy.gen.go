// Code generated by protoc-gen-go-ascii. DO NOT EDIT.
package v1beta1

import (
	proto "github.com/golang/protobuf/proto"
)

// DeepCopyInto supports using PeerAuthentication within kubernetes types, where deepcopy-gen is used.
func (in *PeerAuthentication) DeepCopyInto(out *PeerAuthentication) {
	p := proto.Clone(in).(*PeerAuthentication)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerAuthentication. Required by controller-gen.
func (in *PeerAuthentication) DeepCopy() *PeerAuthentication {
	if in == nil {
		return nil
	}
	out := new(PeerAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new PeerAuthentication. Required by controller-gen.
func (in *PeerAuthentication) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using PeerAuthentication_MutualTLS within kubernetes types, where deepcopy-gen is used.
func (in *PeerAuthentication_MutualTLS) DeepCopyInto(out *PeerAuthentication_MutualTLS) {
	p := proto.Clone(in).(*PeerAuthentication_MutualTLS)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerAuthentication_MutualTLS. Required by controller-gen.
func (in *PeerAuthentication_MutualTLS) DeepCopy() *PeerAuthentication_MutualTLS {
	if in == nil {
		return nil
	}
	out := new(PeerAuthentication_MutualTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new PeerAuthentication_MutualTLS. Required by controller-gen.
func (in *PeerAuthentication_MutualTLS) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

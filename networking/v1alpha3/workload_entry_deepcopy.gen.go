// Code generated by protoc-gen-go-ascii. DO NOT EDIT.
package v1alpha3

import (
	proto "github.com/golang/protobuf/proto"
)

// DeepCopyInto supports using WorkloadEntry within kubernetes types, where deepcopy-gen is used.
func (in *WorkloadEntry) DeepCopyInto(out *WorkloadEntry) {
	p := proto.Clone(in).(*WorkloadEntry)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadEntry. Required by controller-gen.
func (in *WorkloadEntry) DeepCopy() *WorkloadEntry {
	if in == nil {
		return nil
	}
	out := new(WorkloadEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadEntry. Required by controller-gen.
func (in *WorkloadEntry) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

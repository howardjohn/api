// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/destination_rule.proto

// `DestinationRule` defines policies that apply to traffic intended for a
// service after routing has occurred. These rules specify configuration
// for load balancing, connection pool size from the sidecar, and outlier
// detection settings to detect and evict unhealthy hosts from the load
// balancing pool. For example, a simple load balancing policy for the
// ratings service would look as follows:
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: DestinationRule
// metadata:
//   name: bookinfo-ratings
// spec:
//   host: ratings.prod.svc.cluster.local
//   trafficPolicy:
//     loadBalancer:
//       simple: LEAST_CONN
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: DestinationRule
// metadata:
//   name: bookinfo-ratings
// spec:
//   host: ratings.prod.svc.cluster.local
//   trafficPolicy:
//     loadBalancer:
//       simple: LEAST_CONN
// ```
// {{</tab>}}
// {{</tabset>}}
//
// Version specific policies can be specified by defining a named
// `subset` and overriding the settings specified at the service level. The
// following rule uses a round robin load balancing policy for all traffic
// going to a subset named testversion that is composed of endpoints (e.g.,
// pods) with labels (version:v3).
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: DestinationRule
// metadata:
//   name: bookinfo-ratings
// spec:
//   host: ratings.prod.svc.cluster.local
//   trafficPolicy:
//     loadBalancer:
//       simple: LEAST_CONN
//   subsets:
//   - name: testversion
//     labels:
//       version: v3
//     trafficPolicy:
//       loadBalancer:
//         simple: ROUND_ROBIN
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: DestinationRule
// metadata:
//   name: bookinfo-ratings
// spec:
//   host: ratings.prod.svc.cluster.local
//   trafficPolicy:
//     loadBalancer:
//       simple: LEAST_CONN
//   subsets:
//   - name: testversion
//     labels:
//       version: v3
//     trafficPolicy:
//       loadBalancer:
//         simple: ROUND_ROBIN
// ```
// {{</tab>}}
// {{</tabset>}}
//
// **Note:** Policies specified for subsets will not take effect until
// a route rule explicitly sends traffic to this subset.
//
// Traffic policies can be customized to specific ports as well. The
// following rule uses the least connection load balancing policy for all
// traffic to port 80, while uses a round robin load balancing setting for
// traffic to the port 9080.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: DestinationRule
// metadata:
//   name: bookinfo-ratings-port
// spec:
//   host: ratings.prod.svc.cluster.local
//   trafficPolicy: # Apply to all ports
//     portLevelSettings:
//     - port:
//         number: 80
//       loadBalancer:
//         simple: LEAST_CONN
//     - port:
//         number: 9080
//       loadBalancer:
//         simple: ROUND_ROBIN
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: DestinationRule
// metadata:
//   name: bookinfo-ratings-port
// spec:
//   host: ratings.prod.svc.cluster.local
//   trafficPolicy: # Apply to all ports
//     portLevelSettings:
//     - port:
//         number: 80
//       loadBalancer:
//         simple: LEAST_CONN
//     - port:
//         number: 9080
//       loadBalancer:
//         simple: ROUND_ROBIN
// ```
// {{</tab>}}
// {{</tabset>}}
//

package v1alpha3

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using DestinationRule within kubernetes types, where deepcopy-gen is used.
func (in *DestinationRule) DeepCopyInto(out *DestinationRule) {
	p := proto.Clone(in).(*DestinationRule)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationRule. Required by controller-gen.
func (in *DestinationRule) DeepCopy() *DestinationRule {
	if in == nil {
		return nil
	}
	out := new(DestinationRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new DestinationRule. Required by controller-gen.
func (in *DestinationRule) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TrafficPolicy within kubernetes types, where deepcopy-gen is used.
func (in *TrafficPolicy) DeepCopyInto(out *TrafficPolicy) {
	p := proto.Clone(in).(*TrafficPolicy)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficPolicy. Required by controller-gen.
func (in *TrafficPolicy) DeepCopy() *TrafficPolicy {
	if in == nil {
		return nil
	}
	out := new(TrafficPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TrafficPolicy. Required by controller-gen.
func (in *TrafficPolicy) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TrafficPolicy_PortTrafficPolicy within kubernetes types, where deepcopy-gen is used.
func (in *TrafficPolicy_PortTrafficPolicy) DeepCopyInto(out *TrafficPolicy_PortTrafficPolicy) {
	p := proto.Clone(in).(*TrafficPolicy_PortTrafficPolicy)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficPolicy_PortTrafficPolicy. Required by controller-gen.
func (in *TrafficPolicy_PortTrafficPolicy) DeepCopy() *TrafficPolicy_PortTrafficPolicy {
	if in == nil {
		return nil
	}
	out := new(TrafficPolicy_PortTrafficPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TrafficPolicy_PortTrafficPolicy. Required by controller-gen.
func (in *TrafficPolicy_PortTrafficPolicy) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Subset within kubernetes types, where deepcopy-gen is used.
func (in *Subset) DeepCopyInto(out *Subset) {
	p := proto.Clone(in).(*Subset)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subset. Required by controller-gen.
func (in *Subset) DeepCopy() *Subset {
	if in == nil {
		return nil
	}
	out := new(Subset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Subset. Required by controller-gen.
func (in *Subset) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LoadBalancerSettings within kubernetes types, where deepcopy-gen is used.
func (in *LoadBalancerSettings) DeepCopyInto(out *LoadBalancerSettings) {
	p := proto.Clone(in).(*LoadBalancerSettings)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings. Required by controller-gen.
func (in *LoadBalancerSettings) DeepCopy() *LoadBalancerSettings {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings. Required by controller-gen.
func (in *LoadBalancerSettings) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LoadBalancerSettings_ConsistentHashLB within kubernetes types, where deepcopy-gen is used.
func (in *LoadBalancerSettings_ConsistentHashLB) DeepCopyInto(out *LoadBalancerSettings_ConsistentHashLB) {
	p := proto.Clone(in).(*LoadBalancerSettings_ConsistentHashLB)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings_ConsistentHashLB. Required by controller-gen.
func (in *LoadBalancerSettings_ConsistentHashLB) DeepCopy() *LoadBalancerSettings_ConsistentHashLB {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSettings_ConsistentHashLB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings_ConsistentHashLB. Required by controller-gen.
func (in *LoadBalancerSettings_ConsistentHashLB) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LoadBalancerSettings_ConsistentHashLB_HTTPCookie within kubernetes types, where deepcopy-gen is used.
func (in *LoadBalancerSettings_ConsistentHashLB_HTTPCookie) DeepCopyInto(out *LoadBalancerSettings_ConsistentHashLB_HTTPCookie) {
	p := proto.Clone(in).(*LoadBalancerSettings_ConsistentHashLB_HTTPCookie)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings_ConsistentHashLB_HTTPCookie. Required by controller-gen.
func (in *LoadBalancerSettings_ConsistentHashLB_HTTPCookie) DeepCopy() *LoadBalancerSettings_ConsistentHashLB_HTTPCookie {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSettings_ConsistentHashLB_HTTPCookie)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings_ConsistentHashLB_HTTPCookie. Required by controller-gen.
func (in *LoadBalancerSettings_ConsistentHashLB_HTTPCookie) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ConnectionPoolSettings within kubernetes types, where deepcopy-gen is used.
func (in *ConnectionPoolSettings) DeepCopyInto(out *ConnectionPoolSettings) {
	p := proto.Clone(in).(*ConnectionPoolSettings)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings. Required by controller-gen.
func (in *ConnectionPoolSettings) DeepCopy() *ConnectionPoolSettings {
	if in == nil {
		return nil
	}
	out := new(ConnectionPoolSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings. Required by controller-gen.
func (in *ConnectionPoolSettings) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ConnectionPoolSettings_TCPSettings within kubernetes types, where deepcopy-gen is used.
func (in *ConnectionPoolSettings_TCPSettings) DeepCopyInto(out *ConnectionPoolSettings_TCPSettings) {
	p := proto.Clone(in).(*ConnectionPoolSettings_TCPSettings)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings_TCPSettings. Required by controller-gen.
func (in *ConnectionPoolSettings_TCPSettings) DeepCopy() *ConnectionPoolSettings_TCPSettings {
	if in == nil {
		return nil
	}
	out := new(ConnectionPoolSettings_TCPSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings_TCPSettings. Required by controller-gen.
func (in *ConnectionPoolSettings_TCPSettings) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ConnectionPoolSettings_TCPSettings_TcpKeepalive within kubernetes types, where deepcopy-gen is used.
func (in *ConnectionPoolSettings_TCPSettings_TcpKeepalive) DeepCopyInto(out *ConnectionPoolSettings_TCPSettings_TcpKeepalive) {
	p := proto.Clone(in).(*ConnectionPoolSettings_TCPSettings_TcpKeepalive)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings_TCPSettings_TcpKeepalive. Required by controller-gen.
func (in *ConnectionPoolSettings_TCPSettings_TcpKeepalive) DeepCopy() *ConnectionPoolSettings_TCPSettings_TcpKeepalive {
	if in == nil {
		return nil
	}
	out := new(ConnectionPoolSettings_TCPSettings_TcpKeepalive)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings_TCPSettings_TcpKeepalive. Required by controller-gen.
func (in *ConnectionPoolSettings_TCPSettings_TcpKeepalive) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ConnectionPoolSettings_HTTPSettings within kubernetes types, where deepcopy-gen is used.
func (in *ConnectionPoolSettings_HTTPSettings) DeepCopyInto(out *ConnectionPoolSettings_HTTPSettings) {
	p := proto.Clone(in).(*ConnectionPoolSettings_HTTPSettings)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings_HTTPSettings. Required by controller-gen.
func (in *ConnectionPoolSettings_HTTPSettings) DeepCopy() *ConnectionPoolSettings_HTTPSettings {
	if in == nil {
		return nil
	}
	out := new(ConnectionPoolSettings_HTTPSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings_HTTPSettings. Required by controller-gen.
func (in *ConnectionPoolSettings_HTTPSettings) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using OutlierDetection within kubernetes types, where deepcopy-gen is used.
func (in *OutlierDetection) DeepCopyInto(out *OutlierDetection) {
	p := proto.Clone(in).(*OutlierDetection)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutlierDetection. Required by controller-gen.
func (in *OutlierDetection) DeepCopy() *OutlierDetection {
	if in == nil {
		return nil
	}
	out := new(OutlierDetection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new OutlierDetection. Required by controller-gen.
func (in *OutlierDetection) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ClientTLSSettings within kubernetes types, where deepcopy-gen is used.
func (in *ClientTLSSettings) DeepCopyInto(out *ClientTLSSettings) {
	p := proto.Clone(in).(*ClientTLSSettings)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientTLSSettings. Required by controller-gen.
func (in *ClientTLSSettings) DeepCopy() *ClientTLSSettings {
	if in == nil {
		return nil
	}
	out := new(ClientTLSSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ClientTLSSettings. Required by controller-gen.
func (in *ClientTLSSettings) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LocalityLoadBalancerSetting within kubernetes types, where deepcopy-gen is used.
func (in *LocalityLoadBalancerSetting) DeepCopyInto(out *LocalityLoadBalancerSetting) {
	p := proto.Clone(in).(*LocalityLoadBalancerSetting)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalityLoadBalancerSetting. Required by controller-gen.
func (in *LocalityLoadBalancerSetting) DeepCopy() *LocalityLoadBalancerSetting {
	if in == nil {
		return nil
	}
	out := new(LocalityLoadBalancerSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LocalityLoadBalancerSetting. Required by controller-gen.
func (in *LocalityLoadBalancerSetting) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LocalityLoadBalancerSetting_Distribute within kubernetes types, where deepcopy-gen is used.
func (in *LocalityLoadBalancerSetting_Distribute) DeepCopyInto(out *LocalityLoadBalancerSetting_Distribute) {
	p := proto.Clone(in).(*LocalityLoadBalancerSetting_Distribute)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalityLoadBalancerSetting_Distribute. Required by controller-gen.
func (in *LocalityLoadBalancerSetting_Distribute) DeepCopy() *LocalityLoadBalancerSetting_Distribute {
	if in == nil {
		return nil
	}
	out := new(LocalityLoadBalancerSetting_Distribute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LocalityLoadBalancerSetting_Distribute. Required by controller-gen.
func (in *LocalityLoadBalancerSetting_Distribute) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LocalityLoadBalancerSetting_Failover within kubernetes types, where deepcopy-gen is used.
func (in *LocalityLoadBalancerSetting_Failover) DeepCopyInto(out *LocalityLoadBalancerSetting_Failover) {
	p := proto.Clone(in).(*LocalityLoadBalancerSetting_Failover)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalityLoadBalancerSetting_Failover. Required by controller-gen.
func (in *LocalityLoadBalancerSetting_Failover) DeepCopy() *LocalityLoadBalancerSetting_Failover {
	if in == nil {
		return nil
	}
	out := new(LocalityLoadBalancerSetting_Failover)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LocalityLoadBalancerSetting_Failover. Required by controller-gen.
func (in *LocalityLoadBalancerSetting_Failover) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

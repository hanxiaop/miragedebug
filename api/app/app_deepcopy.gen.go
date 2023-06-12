// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package app

import (
	proto "github.com/golang/protobuf/proto"
)

// DeepCopyInto supports using RemoteRuntime within kubernetes types, where deepcopy-gen is used.
func (in *RemoteRuntime) DeepCopyInto(out *RemoteRuntime) {
	p := proto.Clone(in).(*RemoteRuntime)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteRuntime. Required by controller-gen.
func (in *RemoteRuntime) DeepCopy() *RemoteRuntime {
	if in == nil {
		return nil
	}
	out := new(RemoteRuntime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RemoteRuntime. Required by controller-gen.
func (in *RemoteRuntime) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using DebugToolBuilder within kubernetes types, where deepcopy-gen is used.
func (in *DebugToolBuilder) DeepCopyInto(out *DebugToolBuilder) {
	p := proto.Clone(in).(*DebugToolBuilder)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DebugToolBuilder. Required by controller-gen.
func (in *DebugToolBuilder) DeepCopy() *DebugToolBuilder {
	if in == nil {
		return nil
	}
	out := new(DebugToolBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new DebugToolBuilder. Required by controller-gen.
func (in *DebugToolBuilder) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RemoteConfig within kubernetes types, where deepcopy-gen is used.
func (in *RemoteConfig) DeepCopyInto(out *RemoteConfig) {
	p := proto.Clone(in).(*RemoteConfig)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteConfig. Required by controller-gen.
func (in *RemoteConfig) DeepCopy() *RemoteConfig {
	if in == nil {
		return nil
	}
	out := new(RemoteConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RemoteConfig. Required by controller-gen.
func (in *RemoteConfig) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LocalConfig within kubernetes types, where deepcopy-gen is used.
func (in *LocalConfig) DeepCopyInto(out *LocalConfig) {
	p := proto.Clone(in).(*LocalConfig)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalConfig. Required by controller-gen.
func (in *LocalConfig) DeepCopy() *LocalConfig {
	if in == nil {
		return nil
	}
	out := new(LocalConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LocalConfig. Required by controller-gen.
func (in *LocalConfig) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using App within kubernetes types, where deepcopy-gen is used.
func (in *App) DeepCopyInto(out *App) {
	p := proto.Clone(in).(*App)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new App. Required by controller-gen.
func (in *App) DeepCopy() *App {
	if in == nil {
		return nil
	}
	out := new(App)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new App. Required by controller-gen.
func (in *App) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Status within kubernetes types, where deepcopy-gen is used.
func (in *Status) DeepCopyInto(out *Status) {
	p := proto.Clone(in).(*Status)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Status. Required by controller-gen.
func (in *Status) DeepCopy() *Status {
	if in == nil {
		return nil
	}
	out := new(Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Status. Required by controller-gen.
func (in *Status) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using SingleAppRequest within kubernetes types, where deepcopy-gen is used.
func (in *SingleAppRequest) DeepCopyInto(out *SingleAppRequest) {
	p := proto.Clone(in).(*SingleAppRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SingleAppRequest. Required by controller-gen.
func (in *SingleAppRequest) DeepCopy() *SingleAppRequest {
	if in == nil {
		return nil
	}
	out := new(SingleAppRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new SingleAppRequest. Required by controller-gen.
func (in *SingleAppRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using AppList within kubernetes types, where deepcopy-gen is used.
func (in *AppList) DeepCopyInto(out *AppList) {
	p := proto.Clone(in).(*AppList)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppList. Required by controller-gen.
func (in *AppList) DeepCopy() *AppList {
	if in == nil {
		return nil
	}
	out := new(AppList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new AppList. Required by controller-gen.
func (in *AppList) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Empty within kubernetes types, where deepcopy-gen is used.
func (in *Empty) DeepCopyInto(out *Empty) {
	p := proto.Clone(in).(*Empty)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Empty. Required by controller-gen.
func (in *Empty) DeepCopy() *Empty {
	if in == nil {
		return nil
	}
	out := new(Empty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Empty. Required by controller-gen.
func (in *Empty) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

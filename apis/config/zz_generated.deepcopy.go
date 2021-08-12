// +build !ignore_autogenerated

/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

SPDX-License-Identifier: Apache-2.0
*/
// Code generated by deepcopy-gen. DO NOT EDIT.

package config

import (
	core "github.com/gardener/landscaper/apis/core"
	v1alpha1 "github.com/gardener/landscaper/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentConfiguration) DeepCopyInto(out *AgentConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.OCI != nil {
		in, out := &in.OCI, &out.OCI
		*out = new(OCIConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.TargetSelectors != nil {
		in, out := &in.TargetSelectors, &out.TargetSelectors
		*out = make([]v1alpha1.TargetSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentConfiguration.
func (in *AgentConfiguration) DeepCopy() *AgentConfiguration {
	if in == nil {
		return nil
	}
	out := new(AgentConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AgentConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintStore) DeepCopyInto(out *BlueprintStore) {
	*out = *in
	out.GarbageCollectionConfiguration = in.GarbageCollectionConfiguration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintStore.
func (in *BlueprintStore) DeepCopy() *BlueprintStore {
	if in == nil {
		return nil
	}
	out := new(BlueprintStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrdManagementConfiguration) DeepCopyInto(out *CrdManagementConfiguration) {
	*out = *in
	if in.DeployCustomResourceDefinitions != nil {
		in, out := &in.DeployCustomResourceDefinitions, &out.DeployCustomResourceDefinitions
		*out = new(bool)
		**out = **in
	}
	if in.ForceUpdate != nil {
		in, out := &in.ForceUpdate, &out.ForceUpdate
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrdManagementConfiguration.
func (in *CrdManagementConfiguration) DeepCopy() *CrdManagementConfiguration {
	if in == nil {
		return nil
	}
	out := new(CrdManagementConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployItemTimeouts) DeepCopyInto(out *DeployItemTimeouts) {
	*out = *in
	if in.Pickup != nil {
		in, out := &in.Pickup, &out.Pickup
		*out = new(core.Duration)
		**out = **in
	}
	if in.Abort != nil {
		in, out := &in.Abort, &out.Abort
		*out = new(core.Duration)
		**out = **in
	}
	if in.ProgressingDefault != nil {
		in, out := &in.ProgressingDefault, &out.ProgressingDefault
		*out = new(core.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployItemTimeouts.
func (in *DeployItemTimeouts) DeepCopy() *DeployItemTimeouts {
	if in == nil {
		return nil
	}
	out := new(DeployItemTimeouts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployerManagementConfiguration) DeepCopyInto(out *DeployerManagementConfiguration) {
	*out = *in
	in.Agent.DeepCopyInto(&out.Agent)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployerManagementConfiguration.
func (in *DeployerManagementConfiguration) DeepCopy() *DeployerManagementConfiguration {
	if in == nil {
		return nil
	}
	out := new(DeployerManagementConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GarbageCollectionConfiguration) DeepCopyInto(out *GarbageCollectionConfiguration) {
	*out = *in
	out.ResetInterval = in.ResetInterval
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GarbageCollectionConfiguration.
func (in *GarbageCollectionConfiguration) DeepCopy() *GarbageCollectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(GarbageCollectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LandscaperAgentConfiguration) DeepCopyInto(out *LandscaperAgentConfiguration) {
	*out = *in
	in.AgentConfiguration.DeepCopyInto(&out.AgentConfiguration)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LandscaperAgentConfiguration.
func (in *LandscaperAgentConfiguration) DeepCopy() *LandscaperAgentConfiguration {
	if in == nil {
		return nil
	}
	out := new(LandscaperAgentConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LandscaperConfiguration) DeepCopyInto(out *LandscaperConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.RepositoryContext != nil {
		in, out := &in.RepositoryContext, &out.RepositoryContext
		*out = (*in).DeepCopy()
	}
	in.Registry.DeepCopyInto(&out.Registry)
	out.BlueprintStore = in.BlueprintStore
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(MetricsConfiguration)
		**out = **in
	}
	in.CrdManagement.DeepCopyInto(&out.CrdManagement)
	in.DeployerManagement.DeepCopyInto(&out.DeployerManagement)
	if in.DeployItemTimeouts != nil {
		in, out := &in.DeployItemTimeouts, &out.DeployItemTimeouts
		*out = new(DeployItemTimeouts)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LandscaperConfiguration.
func (in *LandscaperConfiguration) DeepCopy() *LandscaperConfiguration {
	if in == nil {
		return nil
	}
	out := new(LandscaperConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LandscaperConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalRegistryConfiguration) DeepCopyInto(out *LocalRegistryConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalRegistryConfiguration.
func (in *LocalRegistryConfiguration) DeepCopy() *LocalRegistryConfiguration {
	if in == nil {
		return nil
	}
	out := new(LocalRegistryConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsConfiguration) DeepCopyInto(out *MetricsConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsConfiguration.
func (in *MetricsConfiguration) DeepCopy() *MetricsConfiguration {
	if in == nil {
		return nil
	}
	out := new(MetricsConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCICacheConfiguration) DeepCopyInto(out *OCICacheConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCICacheConfiguration.
func (in *OCICacheConfiguration) DeepCopy() *OCICacheConfiguration {
	if in == nil {
		return nil
	}
	out := new(OCICacheConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIConfiguration) DeepCopyInto(out *OCIConfiguration) {
	*out = *in
	if in.ConfigFiles != nil {
		in, out := &in.ConfigFiles, &out.ConfigFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Cache != nil {
		in, out := &in.Cache, &out.Cache
		*out = new(OCICacheConfiguration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIConfiguration.
func (in *OCIConfiguration) DeepCopy() *OCIConfiguration {
	if in == nil {
		return nil
	}
	out := new(OCIConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryConfiguration) DeepCopyInto(out *RegistryConfiguration) {
	*out = *in
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(LocalRegistryConfiguration)
		**out = **in
	}
	if in.OCI != nil {
		in, out := &in.OCI, &out.OCI
		*out = new(OCIConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryConfiguration.
func (in *RegistryConfiguration) DeepCopy() *RegistryConfiguration {
	if in == nil {
		return nil
	}
	out := new(RegistryConfiguration)
	in.DeepCopyInto(out)
	return out
}

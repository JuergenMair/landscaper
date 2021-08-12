// +build !ignore_autogenerated

/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

SPDX-License-Identifier: Apache-2.0
*/
// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	json "encoding/json"
	unsafe "unsafe"

	config "github.com/gardener/landscaper/apis/config"
	corev1alpha1 "github.com/gardener/landscaper/apis/core/v1alpha1"
	container "github.com/gardener/landscaper/apis/deployer/container"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Configuration)(nil), (*container.Configuration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Configuration_To_container_Configuration(a.(*Configuration), b.(*container.Configuration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*container.Configuration)(nil), (*Configuration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_container_Configuration_To_v1alpha1_Configuration(a.(*container.Configuration), b.(*Configuration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ContainerSpec)(nil), (*container.ContainerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ContainerSpec_To_container_ContainerSpec(a.(*ContainerSpec), b.(*container.ContainerSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*container.ContainerSpec)(nil), (*ContainerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_container_ContainerSpec_To_v1alpha1_ContainerSpec(a.(*container.ContainerSpec), b.(*ContainerSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ContainerStatus)(nil), (*container.ContainerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ContainerStatus_To_container_ContainerStatus(a.(*ContainerStatus), b.(*container.ContainerStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*container.ContainerStatus)(nil), (*ContainerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_container_ContainerStatus_To_v1alpha1_ContainerStatus(a.(*container.ContainerStatus), b.(*ContainerStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DebugOptions)(nil), (*container.DebugOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_DebugOptions_To_container_DebugOptions(a.(*DebugOptions), b.(*container.DebugOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*container.DebugOptions)(nil), (*DebugOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_container_DebugOptions_To_v1alpha1_DebugOptions(a.(*container.DebugOptions), b.(*DebugOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PodStatus)(nil), (*container.PodStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PodStatus_To_container_PodStatus(a.(*PodStatus), b.(*container.PodStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*container.PodStatus)(nil), (*PodStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_container_PodStatus_To_v1alpha1_PodStatus(a.(*container.PodStatus), b.(*PodStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ProviderConfiguration)(nil), (*container.ProviderConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ProviderConfiguration_To_container_ProviderConfiguration(a.(*ProviderConfiguration), b.(*container.ProviderConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*container.ProviderConfiguration)(nil), (*ProviderConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_container_ProviderConfiguration_To_v1alpha1_ProviderConfiguration(a.(*container.ProviderConfiguration), b.(*ProviderConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ProviderStatus)(nil), (*container.ProviderStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ProviderStatus_To_container_ProviderStatus(a.(*ProviderStatus), b.(*container.ProviderStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*container.ProviderStatus)(nil), (*ProviderStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_container_ProviderStatus_To_v1alpha1_ProviderStatus(a.(*container.ProviderStatus), b.(*ProviderStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_Configuration_To_container_Configuration(in *Configuration, out *container.Configuration, s conversion.Scope) error {
	out.Identity = in.Identity
	out.OCI = (*config.OCIConfiguration)(unsafe.Pointer(in.OCI))
	out.Namespace = in.Namespace
	out.TargetSelector = *(*[]corev1alpha1.TargetSelector)(unsafe.Pointer(&in.TargetSelector))
	if err := Convert_v1alpha1_ContainerSpec_To_container_ContainerSpec(&in.DefaultImage, &out.DefaultImage, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ContainerSpec_To_container_ContainerSpec(&in.InitContainer, &out.InitContainer, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ContainerSpec_To_container_ContainerSpec(&in.WaitContainer, &out.WaitContainer, s); err != nil {
		return err
	}
	out.DebugOptions = (*container.DebugOptions)(unsafe.Pointer(in.DebugOptions))
	return nil
}

// Convert_v1alpha1_Configuration_To_container_Configuration is an autogenerated conversion function.
func Convert_v1alpha1_Configuration_To_container_Configuration(in *Configuration, out *container.Configuration, s conversion.Scope) error {
	return autoConvert_v1alpha1_Configuration_To_container_Configuration(in, out, s)
}

func autoConvert_container_Configuration_To_v1alpha1_Configuration(in *container.Configuration, out *Configuration, s conversion.Scope) error {
	out.Identity = in.Identity
	out.OCI = (*config.OCIConfiguration)(unsafe.Pointer(in.OCI))
	out.TargetSelector = *(*[]corev1alpha1.TargetSelector)(unsafe.Pointer(&in.TargetSelector))
	out.Namespace = in.Namespace
	if err := Convert_container_ContainerSpec_To_v1alpha1_ContainerSpec(&in.DefaultImage, &out.DefaultImage, s); err != nil {
		return err
	}
	if err := Convert_container_ContainerSpec_To_v1alpha1_ContainerSpec(&in.InitContainer, &out.InitContainer, s); err != nil {
		return err
	}
	if err := Convert_container_ContainerSpec_To_v1alpha1_ContainerSpec(&in.WaitContainer, &out.WaitContainer, s); err != nil {
		return err
	}
	out.DebugOptions = (*DebugOptions)(unsafe.Pointer(in.DebugOptions))
	return nil
}

// Convert_container_Configuration_To_v1alpha1_Configuration is an autogenerated conversion function.
func Convert_container_Configuration_To_v1alpha1_Configuration(in *container.Configuration, out *Configuration, s conversion.Scope) error {
	return autoConvert_container_Configuration_To_v1alpha1_Configuration(in, out, s)
}

func autoConvert_v1alpha1_ContainerSpec_To_container_ContainerSpec(in *ContainerSpec, out *container.ContainerSpec, s conversion.Scope) error {
	out.Image = in.Image
	out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
	out.Args = *(*[]string)(unsafe.Pointer(&in.Args))
	out.ImagePullPolicy = v1.PullPolicy(in.ImagePullPolicy)
	return nil
}

// Convert_v1alpha1_ContainerSpec_To_container_ContainerSpec is an autogenerated conversion function.
func Convert_v1alpha1_ContainerSpec_To_container_ContainerSpec(in *ContainerSpec, out *container.ContainerSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ContainerSpec_To_container_ContainerSpec(in, out, s)
}

func autoConvert_container_ContainerSpec_To_v1alpha1_ContainerSpec(in *container.ContainerSpec, out *ContainerSpec, s conversion.Scope) error {
	out.Image = in.Image
	out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
	out.Args = *(*[]string)(unsafe.Pointer(&in.Args))
	out.ImagePullPolicy = v1.PullPolicy(in.ImagePullPolicy)
	return nil
}

// Convert_container_ContainerSpec_To_v1alpha1_ContainerSpec is an autogenerated conversion function.
func Convert_container_ContainerSpec_To_v1alpha1_ContainerSpec(in *container.ContainerSpec, out *ContainerSpec, s conversion.Scope) error {
	return autoConvert_container_ContainerSpec_To_v1alpha1_ContainerSpec(in, out, s)
}

func autoConvert_v1alpha1_ContainerStatus_To_container_ContainerStatus(in *ContainerStatus, out *container.ContainerStatus, s conversion.Scope) error {
	out.Message = in.Message
	out.Reason = in.Reason
	out.State = in.State
	out.Image = in.Image
	out.ImageID = in.ImageID
	out.ExitCode = (*int32)(unsafe.Pointer(in.ExitCode))
	return nil
}

// Convert_v1alpha1_ContainerStatus_To_container_ContainerStatus is an autogenerated conversion function.
func Convert_v1alpha1_ContainerStatus_To_container_ContainerStatus(in *ContainerStatus, out *container.ContainerStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ContainerStatus_To_container_ContainerStatus(in, out, s)
}

func autoConvert_container_ContainerStatus_To_v1alpha1_ContainerStatus(in *container.ContainerStatus, out *ContainerStatus, s conversion.Scope) error {
	out.Message = in.Message
	out.Reason = in.Reason
	out.State = in.State
	out.Image = in.Image
	out.ImageID = in.ImageID
	out.ExitCode = (*int32)(unsafe.Pointer(in.ExitCode))
	return nil
}

// Convert_container_ContainerStatus_To_v1alpha1_ContainerStatus is an autogenerated conversion function.
func Convert_container_ContainerStatus_To_v1alpha1_ContainerStatus(in *container.ContainerStatus, out *ContainerStatus, s conversion.Scope) error {
	return autoConvert_container_ContainerStatus_To_v1alpha1_ContainerStatus(in, out, s)
}

func autoConvert_v1alpha1_DebugOptions_To_container_DebugOptions(in *DebugOptions, out *container.DebugOptions, s conversion.Scope) error {
	out.KeepPod = in.KeepPod
	return nil
}

// Convert_v1alpha1_DebugOptions_To_container_DebugOptions is an autogenerated conversion function.
func Convert_v1alpha1_DebugOptions_To_container_DebugOptions(in *DebugOptions, out *container.DebugOptions, s conversion.Scope) error {
	return autoConvert_v1alpha1_DebugOptions_To_container_DebugOptions(in, out, s)
}

func autoConvert_container_DebugOptions_To_v1alpha1_DebugOptions(in *container.DebugOptions, out *DebugOptions, s conversion.Scope) error {
	out.KeepPod = in.KeepPod
	return nil
}

// Convert_container_DebugOptions_To_v1alpha1_DebugOptions is an autogenerated conversion function.
func Convert_container_DebugOptions_To_v1alpha1_DebugOptions(in *container.DebugOptions, out *DebugOptions, s conversion.Scope) error {
	return autoConvert_container_DebugOptions_To_v1alpha1_DebugOptions(in, out, s)
}

func autoConvert_v1alpha1_PodStatus_To_container_PodStatus(in *PodStatus, out *container.PodStatus, s conversion.Scope) error {
	out.PodName = in.PodName
	out.LastRun = (*metav1.Time)(unsafe.Pointer(in.LastRun))
	if err := Convert_v1alpha1_ContainerStatus_To_container_ContainerStatus(&in.ContainerStatus, &out.ContainerStatus, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ContainerStatus_To_container_ContainerStatus(&in.InitContainerStatus, &out.InitContainerStatus, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ContainerStatus_To_container_ContainerStatus(&in.WaitContainerStatus, &out.WaitContainerStatus, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_PodStatus_To_container_PodStatus is an autogenerated conversion function.
func Convert_v1alpha1_PodStatus_To_container_PodStatus(in *PodStatus, out *container.PodStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_PodStatus_To_container_PodStatus(in, out, s)
}

func autoConvert_container_PodStatus_To_v1alpha1_PodStatus(in *container.PodStatus, out *PodStatus, s conversion.Scope) error {
	out.PodName = in.PodName
	out.LastRun = (*metav1.Time)(unsafe.Pointer(in.LastRun))
	if err := Convert_container_ContainerStatus_To_v1alpha1_ContainerStatus(&in.ContainerStatus, &out.ContainerStatus, s); err != nil {
		return err
	}
	if err := Convert_container_ContainerStatus_To_v1alpha1_ContainerStatus(&in.InitContainerStatus, &out.InitContainerStatus, s); err != nil {
		return err
	}
	if err := Convert_container_ContainerStatus_To_v1alpha1_ContainerStatus(&in.WaitContainerStatus, &out.WaitContainerStatus, s); err != nil {
		return err
	}
	return nil
}

// Convert_container_PodStatus_To_v1alpha1_PodStatus is an autogenerated conversion function.
func Convert_container_PodStatus_To_v1alpha1_PodStatus(in *container.PodStatus, out *PodStatus, s conversion.Scope) error {
	return autoConvert_container_PodStatus_To_v1alpha1_PodStatus(in, out, s)
}

func autoConvert_v1alpha1_ProviderConfiguration_To_container_ProviderConfiguration(in *ProviderConfiguration, out *container.ProviderConfiguration, s conversion.Scope) error {
	out.Image = in.Image
	out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
	out.Args = *(*[]string)(unsafe.Pointer(&in.Args))
	out.ImportValues = *(*json.RawMessage)(unsafe.Pointer(&in.ImportValues))
	out.Blueprint = (*corev1alpha1.BlueprintDefinition)(unsafe.Pointer(in.Blueprint))
	out.ComponentDescriptor = (*corev1alpha1.ComponentDescriptorDefinition)(unsafe.Pointer(in.ComponentDescriptor))
	out.RegistryPullSecrets = *(*[]corev1alpha1.ObjectReference)(unsafe.Pointer(&in.RegistryPullSecrets))
	return nil
}

// Convert_v1alpha1_ProviderConfiguration_To_container_ProviderConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ProviderConfiguration_To_container_ProviderConfiguration(in *ProviderConfiguration, out *container.ProviderConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ProviderConfiguration_To_container_ProviderConfiguration(in, out, s)
}

func autoConvert_container_ProviderConfiguration_To_v1alpha1_ProviderConfiguration(in *container.ProviderConfiguration, out *ProviderConfiguration, s conversion.Scope) error {
	out.Image = in.Image
	out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
	out.Args = *(*[]string)(unsafe.Pointer(&in.Args))
	out.ImportValues = *(*json.RawMessage)(unsafe.Pointer(&in.ImportValues))
	out.Blueprint = (*corev1alpha1.BlueprintDefinition)(unsafe.Pointer(in.Blueprint))
	out.ComponentDescriptor = (*corev1alpha1.ComponentDescriptorDefinition)(unsafe.Pointer(in.ComponentDescriptor))
	out.RegistryPullSecrets = *(*[]corev1alpha1.ObjectReference)(unsafe.Pointer(&in.RegistryPullSecrets))
	return nil
}

// Convert_container_ProviderConfiguration_To_v1alpha1_ProviderConfiguration is an autogenerated conversion function.
func Convert_container_ProviderConfiguration_To_v1alpha1_ProviderConfiguration(in *container.ProviderConfiguration, out *ProviderConfiguration, s conversion.Scope) error {
	return autoConvert_container_ProviderConfiguration_To_v1alpha1_ProviderConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ProviderStatus_To_container_ProviderStatus(in *ProviderStatus, out *container.ProviderStatus, s conversion.Scope) error {
	out.LastOperation = in.LastOperation
	out.PodStatus = (*container.PodStatus)(unsafe.Pointer(in.PodStatus))
	return nil
}

// Convert_v1alpha1_ProviderStatus_To_container_ProviderStatus is an autogenerated conversion function.
func Convert_v1alpha1_ProviderStatus_To_container_ProviderStatus(in *ProviderStatus, out *container.ProviderStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ProviderStatus_To_container_ProviderStatus(in, out, s)
}

func autoConvert_container_ProviderStatus_To_v1alpha1_ProviderStatus(in *container.ProviderStatus, out *ProviderStatus, s conversion.Scope) error {
	out.LastOperation = in.LastOperation
	out.PodStatus = (*PodStatus)(unsafe.Pointer(in.PodStatus))
	return nil
}

// Convert_container_ProviderStatus_To_v1alpha1_ProviderStatus is an autogenerated conversion function.
func Convert_container_ProviderStatus_To_v1alpha1_ProviderStatus(in *container.ProviderStatus, out *ProviderStatus, s conversion.Scope) error {
	return autoConvert_container_ProviderStatus_To_v1alpha1_ProviderStatus(in, out, s)
}

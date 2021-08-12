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

	corev1alpha1 "github.com/gardener/landscaper/apis/core/v1alpha1"
	mock "github.com/gardener/landscaper/apis/deployer/mock"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Configuration)(nil), (*mock.Configuration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Configuration_To_mock_Configuration(a.(*Configuration), b.(*mock.Configuration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*mock.Configuration)(nil), (*Configuration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_mock_Configuration_To_v1alpha1_Configuration(a.(*mock.Configuration), b.(*Configuration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ProviderConfiguration)(nil), (*mock.ProviderConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ProviderConfiguration_To_mock_ProviderConfiguration(a.(*ProviderConfiguration), b.(*mock.ProviderConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*mock.ProviderConfiguration)(nil), (*ProviderConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_mock_ProviderConfiguration_To_v1alpha1_ProviderConfiguration(a.(*mock.ProviderConfiguration), b.(*ProviderConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_Configuration_To_mock_Configuration(in *Configuration, out *mock.Configuration, s conversion.Scope) error {
	out.Identity = in.Identity
	out.TargetSelector = *(*[]corev1alpha1.TargetSelector)(unsafe.Pointer(&in.TargetSelector))
	return nil
}

// Convert_v1alpha1_Configuration_To_mock_Configuration is an autogenerated conversion function.
func Convert_v1alpha1_Configuration_To_mock_Configuration(in *Configuration, out *mock.Configuration, s conversion.Scope) error {
	return autoConvert_v1alpha1_Configuration_To_mock_Configuration(in, out, s)
}

func autoConvert_mock_Configuration_To_v1alpha1_Configuration(in *mock.Configuration, out *Configuration, s conversion.Scope) error {
	out.Identity = in.Identity
	out.TargetSelector = *(*[]corev1alpha1.TargetSelector)(unsafe.Pointer(&in.TargetSelector))
	return nil
}

// Convert_mock_Configuration_To_v1alpha1_Configuration is an autogenerated conversion function.
func Convert_mock_Configuration_To_v1alpha1_Configuration(in *mock.Configuration, out *Configuration, s conversion.Scope) error {
	return autoConvert_mock_Configuration_To_v1alpha1_Configuration(in, out, s)
}

func autoConvert_v1alpha1_ProviderConfiguration_To_mock_ProviderConfiguration(in *ProviderConfiguration, out *mock.ProviderConfiguration, s conversion.Scope) error {
	out.Phase = (*corev1alpha1.ExecutionPhase)(unsafe.Pointer(in.Phase))
	out.InitialPhase = (*corev1alpha1.ExecutionPhase)(unsafe.Pointer(in.InitialPhase))
	out.ProviderStatus = (*runtime.RawExtension)(unsafe.Pointer(in.ProviderStatus))
	out.Export = (*json.RawMessage)(unsafe.Pointer(in.Export))
	return nil
}

// Convert_v1alpha1_ProviderConfiguration_To_mock_ProviderConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ProviderConfiguration_To_mock_ProviderConfiguration(in *ProviderConfiguration, out *mock.ProviderConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ProviderConfiguration_To_mock_ProviderConfiguration(in, out, s)
}

func autoConvert_mock_ProviderConfiguration_To_v1alpha1_ProviderConfiguration(in *mock.ProviderConfiguration, out *ProviderConfiguration, s conversion.Scope) error {
	out.Phase = (*corev1alpha1.ExecutionPhase)(unsafe.Pointer(in.Phase))
	out.InitialPhase = (*corev1alpha1.ExecutionPhase)(unsafe.Pointer(in.InitialPhase))
	out.ProviderStatus = (*runtime.RawExtension)(unsafe.Pointer(in.ProviderStatus))
	out.Export = (*json.RawMessage)(unsafe.Pointer(in.Export))
	return nil
}

// Convert_mock_ProviderConfiguration_To_v1alpha1_ProviderConfiguration is an autogenerated conversion function.
func Convert_mock_ProviderConfiguration_To_v1alpha1_ProviderConfiguration(in *mock.ProviderConfiguration, out *ProviderConfiguration, s conversion.Scope) error {
	return autoConvert_mock_ProviderConfiguration_To_v1alpha1_ProviderConfiguration(in, out, s)
}

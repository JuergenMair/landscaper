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
	core "github.com/gardener/landscaper/apis/core"
	corev1alpha1 "github.com/gardener/landscaper/apis/core/v1alpha1"
	helm "github.com/gardener/landscaper/apis/deployer/helm"
	managedresource "github.com/gardener/landscaper/apis/deployer/utils/managedresource"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ArchiveAccess)(nil), (*helm.ArchiveAccess)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ArchiveAccess_To_helm_ArchiveAccess(a.(*ArchiveAccess), b.(*helm.ArchiveAccess), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*helm.ArchiveAccess)(nil), (*ArchiveAccess)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_helm_ArchiveAccess_To_v1alpha1_ArchiveAccess(a.(*helm.ArchiveAccess), b.(*ArchiveAccess), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Chart)(nil), (*helm.Chart)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Chart_To_helm_Chart(a.(*Chart), b.(*helm.Chart), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*helm.Chart)(nil), (*Chart)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_helm_Chart_To_v1alpha1_Chart(a.(*helm.Chart), b.(*Chart), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Configuration)(nil), (*helm.Configuration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Configuration_To_helm_Configuration(a.(*Configuration), b.(*helm.Configuration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*helm.Configuration)(nil), (*Configuration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_helm_Configuration_To_v1alpha1_Configuration(a.(*helm.Configuration), b.(*Configuration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ExportConfiguration)(nil), (*helm.ExportConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ExportConfiguration_To_helm_ExportConfiguration(a.(*ExportConfiguration), b.(*helm.ExportConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*helm.ExportConfiguration)(nil), (*ExportConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_helm_ExportConfiguration_To_v1alpha1_ExportConfiguration(a.(*helm.ExportConfiguration), b.(*ExportConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ProviderConfiguration)(nil), (*helm.ProviderConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ProviderConfiguration_To_helm_ProviderConfiguration(a.(*ProviderConfiguration), b.(*helm.ProviderConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*helm.ProviderConfiguration)(nil), (*ProviderConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_helm_ProviderConfiguration_To_v1alpha1_ProviderConfiguration(a.(*helm.ProviderConfiguration), b.(*ProviderConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ProviderStatus)(nil), (*helm.ProviderStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ProviderStatus_To_helm_ProviderStatus(a.(*ProviderStatus), b.(*helm.ProviderStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*helm.ProviderStatus)(nil), (*ProviderStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_helm_ProviderStatus_To_v1alpha1_ProviderStatus(a.(*helm.ProviderStatus), b.(*ProviderStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RemoteArchiveAccess)(nil), (*helm.RemoteArchiveAccess)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RemoteArchiveAccess_To_helm_RemoteArchiveAccess(a.(*RemoteArchiveAccess), b.(*helm.RemoteArchiveAccess), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*helm.RemoteArchiveAccess)(nil), (*RemoteArchiveAccess)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_helm_RemoteArchiveAccess_To_v1alpha1_RemoteArchiveAccess(a.(*helm.RemoteArchiveAccess), b.(*RemoteArchiveAccess), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RemoteChartReference)(nil), (*helm.RemoteChartReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RemoteChartReference_To_helm_RemoteChartReference(a.(*RemoteChartReference), b.(*helm.RemoteChartReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*helm.RemoteChartReference)(nil), (*RemoteChartReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_helm_RemoteChartReference_To_v1alpha1_RemoteChartReference(a.(*helm.RemoteChartReference), b.(*RemoteChartReference), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_ArchiveAccess_To_helm_ArchiveAccess(in *ArchiveAccess, out *helm.ArchiveAccess, s conversion.Scope) error {
	out.Raw = in.Raw
	out.Remote = (*helm.RemoteArchiveAccess)(unsafe.Pointer(in.Remote))
	return nil
}

// Convert_v1alpha1_ArchiveAccess_To_helm_ArchiveAccess is an autogenerated conversion function.
func Convert_v1alpha1_ArchiveAccess_To_helm_ArchiveAccess(in *ArchiveAccess, out *helm.ArchiveAccess, s conversion.Scope) error {
	return autoConvert_v1alpha1_ArchiveAccess_To_helm_ArchiveAccess(in, out, s)
}

func autoConvert_helm_ArchiveAccess_To_v1alpha1_ArchiveAccess(in *helm.ArchiveAccess, out *ArchiveAccess, s conversion.Scope) error {
	out.Raw = in.Raw
	out.Remote = (*RemoteArchiveAccess)(unsafe.Pointer(in.Remote))
	return nil
}

// Convert_helm_ArchiveAccess_To_v1alpha1_ArchiveAccess is an autogenerated conversion function.
func Convert_helm_ArchiveAccess_To_v1alpha1_ArchiveAccess(in *helm.ArchiveAccess, out *ArchiveAccess, s conversion.Scope) error {
	return autoConvert_helm_ArchiveAccess_To_v1alpha1_ArchiveAccess(in, out, s)
}

func autoConvert_v1alpha1_Chart_To_helm_Chart(in *Chart, out *helm.Chart, s conversion.Scope) error {
	out.Ref = in.Ref
	out.FromResource = (*helm.RemoteChartReference)(unsafe.Pointer(in.FromResource))
	out.Archive = (*helm.ArchiveAccess)(unsafe.Pointer(in.Archive))
	return nil
}

// Convert_v1alpha1_Chart_To_helm_Chart is an autogenerated conversion function.
func Convert_v1alpha1_Chart_To_helm_Chart(in *Chart, out *helm.Chart, s conversion.Scope) error {
	return autoConvert_v1alpha1_Chart_To_helm_Chart(in, out, s)
}

func autoConvert_helm_Chart_To_v1alpha1_Chart(in *helm.Chart, out *Chart, s conversion.Scope) error {
	out.Ref = in.Ref
	out.FromResource = (*RemoteChartReference)(unsafe.Pointer(in.FromResource))
	out.Archive = (*ArchiveAccess)(unsafe.Pointer(in.Archive))
	return nil
}

// Convert_helm_Chart_To_v1alpha1_Chart is an autogenerated conversion function.
func Convert_helm_Chart_To_v1alpha1_Chart(in *helm.Chart, out *Chart, s conversion.Scope) error {
	return autoConvert_helm_Chart_To_v1alpha1_Chart(in, out, s)
}

func autoConvert_v1alpha1_Configuration_To_helm_Configuration(in *Configuration, out *helm.Configuration, s conversion.Scope) error {
	out.Identity = in.Identity
	out.OCI = (*config.OCIConfiguration)(unsafe.Pointer(in.OCI))
	out.TargetSelector = *(*[]corev1alpha1.TargetSelector)(unsafe.Pointer(&in.TargetSelector))
	if err := Convert_v1alpha1_ExportConfiguration_To_helm_ExportConfiguration(&in.Export, &out.Export, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Configuration_To_helm_Configuration is an autogenerated conversion function.
func Convert_v1alpha1_Configuration_To_helm_Configuration(in *Configuration, out *helm.Configuration, s conversion.Scope) error {
	return autoConvert_v1alpha1_Configuration_To_helm_Configuration(in, out, s)
}

func autoConvert_helm_Configuration_To_v1alpha1_Configuration(in *helm.Configuration, out *Configuration, s conversion.Scope) error {
	out.Identity = in.Identity
	out.OCI = (*config.OCIConfiguration)(unsafe.Pointer(in.OCI))
	out.TargetSelector = *(*[]corev1alpha1.TargetSelector)(unsafe.Pointer(&in.TargetSelector))
	if err := Convert_helm_ExportConfiguration_To_v1alpha1_ExportConfiguration(&in.Export, &out.Export, s); err != nil {
		return err
	}
	return nil
}

// Convert_helm_Configuration_To_v1alpha1_Configuration is an autogenerated conversion function.
func Convert_helm_Configuration_To_v1alpha1_Configuration(in *helm.Configuration, out *Configuration, s conversion.Scope) error {
	return autoConvert_helm_Configuration_To_v1alpha1_Configuration(in, out, s)
}

func autoConvert_v1alpha1_ExportConfiguration_To_helm_ExportConfiguration(in *ExportConfiguration, out *helm.ExportConfiguration, s conversion.Scope) error {
	out.DefaultTimeout = (*corev1alpha1.Duration)(unsafe.Pointer(in.DefaultTimeout))
	return nil
}

// Convert_v1alpha1_ExportConfiguration_To_helm_ExportConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ExportConfiguration_To_helm_ExportConfiguration(in *ExportConfiguration, out *helm.ExportConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ExportConfiguration_To_helm_ExportConfiguration(in, out, s)
}

func autoConvert_helm_ExportConfiguration_To_v1alpha1_ExportConfiguration(in *helm.ExportConfiguration, out *ExportConfiguration, s conversion.Scope) error {
	out.DefaultTimeout = (*corev1alpha1.Duration)(unsafe.Pointer(in.DefaultTimeout))
	return nil
}

// Convert_helm_ExportConfiguration_To_v1alpha1_ExportConfiguration is an autogenerated conversion function.
func Convert_helm_ExportConfiguration_To_v1alpha1_ExportConfiguration(in *helm.ExportConfiguration, out *ExportConfiguration, s conversion.Scope) error {
	return autoConvert_helm_ExportConfiguration_To_v1alpha1_ExportConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ProviderConfiguration_To_helm_ProviderConfiguration(in *ProviderConfiguration, out *helm.ProviderConfiguration, s conversion.Scope) error {
	out.Kubeconfig = in.Kubeconfig
	out.UpdateStrategy = helm.UpdateStrategy(in.UpdateStrategy)
	out.ReadinessChecks = in.ReadinessChecks
	out.DeleteTimeout = (*core.Duration)(unsafe.Pointer(in.DeleteTimeout))
	if err := Convert_v1alpha1_Chart_To_helm_Chart(&in.Chart, &out.Chart, s); err != nil {
		return err
	}
	out.Name = in.Name
	out.Namespace = in.Namespace
	out.Values = *(*json.RawMessage)(unsafe.Pointer(&in.Values))
	out.ExportsFromManifests = *(*[]managedresource.Export)(unsafe.Pointer(&in.ExportsFromManifests))
	out.Exports = (*managedresource.Exports)(unsafe.Pointer(in.Exports))
	return nil
}

// Convert_v1alpha1_ProviderConfiguration_To_helm_ProviderConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ProviderConfiguration_To_helm_ProviderConfiguration(in *ProviderConfiguration, out *helm.ProviderConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ProviderConfiguration_To_helm_ProviderConfiguration(in, out, s)
}

func autoConvert_helm_ProviderConfiguration_To_v1alpha1_ProviderConfiguration(in *helm.ProviderConfiguration, out *ProviderConfiguration, s conversion.Scope) error {
	out.Kubeconfig = in.Kubeconfig
	out.ReadinessChecks = in.ReadinessChecks
	out.DeleteTimeout = (*corev1alpha1.Duration)(unsafe.Pointer(in.DeleteTimeout))
	out.UpdateStrategy = UpdateStrategy(in.UpdateStrategy)
	if err := Convert_helm_Chart_To_v1alpha1_Chart(&in.Chart, &out.Chart, s); err != nil {
		return err
	}
	out.Name = in.Name
	out.Namespace = in.Namespace
	out.Values = *(*json.RawMessage)(unsafe.Pointer(&in.Values))
	out.ExportsFromManifests = *(*[]managedresource.Export)(unsafe.Pointer(&in.ExportsFromManifests))
	out.Exports = (*managedresource.Exports)(unsafe.Pointer(in.Exports))
	return nil
}

// Convert_helm_ProviderConfiguration_To_v1alpha1_ProviderConfiguration is an autogenerated conversion function.
func Convert_helm_ProviderConfiguration_To_v1alpha1_ProviderConfiguration(in *helm.ProviderConfiguration, out *ProviderConfiguration, s conversion.Scope) error {
	return autoConvert_helm_ProviderConfiguration_To_v1alpha1_ProviderConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ProviderStatus_To_helm_ProviderStatus(in *ProviderStatus, out *helm.ProviderStatus, s conversion.Scope) error {
	out.ManagedResources = *(*managedresource.ManagedResourceStatusList)(unsafe.Pointer(&in.ManagedResources))
	return nil
}

// Convert_v1alpha1_ProviderStatus_To_helm_ProviderStatus is an autogenerated conversion function.
func Convert_v1alpha1_ProviderStatus_To_helm_ProviderStatus(in *ProviderStatus, out *helm.ProviderStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ProviderStatus_To_helm_ProviderStatus(in, out, s)
}

func autoConvert_helm_ProviderStatus_To_v1alpha1_ProviderStatus(in *helm.ProviderStatus, out *ProviderStatus, s conversion.Scope) error {
	out.ManagedResources = *(*managedresource.ManagedResourceStatusList)(unsafe.Pointer(&in.ManagedResources))
	return nil
}

// Convert_helm_ProviderStatus_To_v1alpha1_ProviderStatus is an autogenerated conversion function.
func Convert_helm_ProviderStatus_To_v1alpha1_ProviderStatus(in *helm.ProviderStatus, out *ProviderStatus, s conversion.Scope) error {
	return autoConvert_helm_ProviderStatus_To_v1alpha1_ProviderStatus(in, out, s)
}

func autoConvert_v1alpha1_RemoteArchiveAccess_To_helm_RemoteArchiveAccess(in *RemoteArchiveAccess, out *helm.RemoteArchiveAccess, s conversion.Scope) error {
	out.URL = in.URL
	return nil
}

// Convert_v1alpha1_RemoteArchiveAccess_To_helm_RemoteArchiveAccess is an autogenerated conversion function.
func Convert_v1alpha1_RemoteArchiveAccess_To_helm_RemoteArchiveAccess(in *RemoteArchiveAccess, out *helm.RemoteArchiveAccess, s conversion.Scope) error {
	return autoConvert_v1alpha1_RemoteArchiveAccess_To_helm_RemoteArchiveAccess(in, out, s)
}

func autoConvert_helm_RemoteArchiveAccess_To_v1alpha1_RemoteArchiveAccess(in *helm.RemoteArchiveAccess, out *RemoteArchiveAccess, s conversion.Scope) error {
	out.URL = in.URL
	return nil
}

// Convert_helm_RemoteArchiveAccess_To_v1alpha1_RemoteArchiveAccess is an autogenerated conversion function.
func Convert_helm_RemoteArchiveAccess_To_v1alpha1_RemoteArchiveAccess(in *helm.RemoteArchiveAccess, out *RemoteArchiveAccess, s conversion.Scope) error {
	return autoConvert_helm_RemoteArchiveAccess_To_v1alpha1_RemoteArchiveAccess(in, out, s)
}

func autoConvert_v1alpha1_RemoteChartReference_To_helm_RemoteChartReference(in *RemoteChartReference, out *helm.RemoteChartReference, s conversion.Scope) error {
	out.ComponentDescriptorDefinition = in.ComponentDescriptorDefinition
	out.ResourceName = in.ResourceName
	return nil
}

// Convert_v1alpha1_RemoteChartReference_To_helm_RemoteChartReference is an autogenerated conversion function.
func Convert_v1alpha1_RemoteChartReference_To_helm_RemoteChartReference(in *RemoteChartReference, out *helm.RemoteChartReference, s conversion.Scope) error {
	return autoConvert_v1alpha1_RemoteChartReference_To_helm_RemoteChartReference(in, out, s)
}

func autoConvert_helm_RemoteChartReference_To_v1alpha1_RemoteChartReference(in *helm.RemoteChartReference, out *RemoteChartReference, s conversion.Scope) error {
	out.ComponentDescriptorDefinition = in.ComponentDescriptorDefinition
	out.ResourceName = in.ResourceName
	return nil
}

// Convert_helm_RemoteChartReference_To_v1alpha1_RemoteChartReference is an autogenerated conversion function.
func Convert_helm_RemoteChartReference_To_v1alpha1_RemoteChartReference(in *helm.RemoteChartReference, out *RemoteChartReference, s conversion.Scope) error {
	return autoConvert_helm_RemoteChartReference_To_v1alpha1_RemoteChartReference(in, out, s)
}

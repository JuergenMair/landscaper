// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors.
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/utils/pointer"

	"github.com/gardener/landscaper/apis/core/v1alpha1"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

// SetDefaults_LandscaperConfiguration sets the defaults for the landscaper configuration.
func SetDefaults_LandscaperConfiguration(obj *LandscaperConfiguration) {
	if obj.Registry.OCI == nil {
		obj.Registry.OCI = &OCIConfiguration{}
	}
	if obj.Registry.OCI.Cache == nil {
		obj.Registry.OCI.Cache = &OCICacheConfiguration{
			UseInMemoryOverlay: false,
		}
	}
	if len(obj.DeployerManagement.Namespace) == 0 {
		obj.DeployerManagement.Namespace = "ls-system"
	}
	if obj.DeployItemTimeouts == nil {
		obj.DeployItemTimeouts = &DeployItemTimeouts{}
	}
	if obj.DeployItemTimeouts.Pickup == nil {
		obj.DeployItemTimeouts.Pickup = &v1alpha1.Duration{Duration: 5 * time.Minute}
	}
	if obj.DeployItemTimeouts.Abort == nil {
		obj.DeployItemTimeouts.Abort = &v1alpha1.Duration{Duration: 5 * time.Minute}
	}
	if obj.DeployItemTimeouts.ProgressingDefault == nil {
		obj.DeployItemTimeouts.ProgressingDefault = &v1alpha1.Duration{Duration: 10 * time.Minute}
	}

	SetDefaults_CrdManagementConfiguration(&obj.CrdManagement)

	if obj.DeployerManagement.Disable {
		obj.DeployerManagement.Agent.Disable = true
	}
	if len(obj.DeployerManagement.Agent.Name) == 0 {
		obj.DeployerManagement.Agent.Name = "default"
	}
	if len(obj.DeployerManagement.Agent.Namespace) == 0 {
		obj.DeployerManagement.Agent.Namespace = obj.DeployerManagement.Namespace
	}
	if obj.DeployerManagement.Agent.OCI == nil {
		obj.DeployerManagement.Agent.OCI = obj.Registry.OCI
	}
}

// SetDefaults_CrdManagementConfiguration sets the defaults for the crd management configuration.
func SetDefaults_CrdManagementConfiguration(obj *CrdManagementConfiguration) {
	if obj.DeployCustomResourceDefinitions == nil {
		obj.DeployCustomResourceDefinitions = pointer.BoolPtr(true)
	}
	if obj.ForceUpdate == nil {
		obj.ForceUpdate = pointer.BoolPtr(true)
	}
}

// SetDefaults_AgentConfiguration sets the defaults for the landscaper configuration.
func SetDefaults_AgentConfiguration(obj *AgentConfiguration) {
	if len(obj.Namespace) == 0 {
		obj.Namespace = "ls-system"
	}
}

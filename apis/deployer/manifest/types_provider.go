// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors.
//
// SPDX-License-Identifier: Apache-2.0

package manifest

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/gardener/landscaper/apis/deployer/utils/managedresource"

	lscore "github.com/gardener/landscaper/apis/core"
	health "github.com/gardener/landscaper/apis/deployer/utils/readinesschecks"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProviderConfiguration is the manifest deployer configuration that is expected in a DeployItem.
type ProviderConfiguration struct {
	metav1.TypeMeta `json:",inline"`
	// Kubeconfig is the base64 encoded kubeconfig file.
	// By default the configured target is used to deploy the resources
	// +optional
	Kubeconfig string `json:"kubeconfig"`
	// UpdateStrategy defines the strategy how the manifest are updated in the cluster.
	// +optional
	UpdateStrategy UpdateStrategy `json:"updateStrategy"`
	// ReadinessChecks configures the readiness checks.
	// +optional
	ReadinessChecks health.ReadinessCheckConfiguration `json:"readiness,omitempty"`
	// DeleteTimeout is the time to wait before giving up on a resource to be deleted.
	// Defaults to 180s.
	// +optional
	DeleteTimeout *lscore.Duration `json:"deleteTimeout,omitempty"`
	// Manifests contains a list of manifests that should be applied in the target cluster
	Manifests []managedresource.Manifest `json:"manifests,omitempty"`
	// Exports describe the exports from the templated manifests that should be exported by the helm deployer.
	// +optional
	Exports *managedresource.Exports `json:"exports,omitempty"`
}

// UpdateStrategy defines the strategy that is used to apply resources to the cluster.
type UpdateStrategy string

const (
	UpdateStrategyUpdate UpdateStrategy = "update"
	UpdateStrategyPatch  UpdateStrategy = "patch"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProviderStatus is the manifest provider specific status.
type ProviderStatus struct {
	metav1.TypeMeta `json:",inline"`
	// ManagedResources contains all kubernetes resources that are deployed by the deployer.
	ManagedResources managedresource.ManagedResourceStatusList `json:"managedResources,omitempty"`
}

/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	runtimecatalog "sigs.k8s.io/cluster-api/exp/runtime/catalog"
)

// BeforeClusterCreateRequest is the request of the BeforeClusterCreate hook.
// +kubebuilder:object:root=true
type BeforeClusterCreateRequest struct {
	metav1.TypeMeta `json:",inline"`

	// Cluster is the cluster object the lifecycle hook corresponds to.
	Cluster clusterv1.Cluster `json:"cluster"`
}

var _ RetryResponseObject = &BeforeClusterCreateResponse{}

// BeforeClusterCreateResponse is the response of the BeforeClusterCreate hook.
// +kubebuilder:object:root=true
type BeforeClusterCreateResponse struct {
	metav1.TypeMeta `json:",inline"`

	// CommonRetryResponse contains Status, Message and RetryAfterSeconds fields.
	CommonRetryResponse `json:",inline"`
}

// BeforeClusterCreate is the hook that will be called right before the topology of the Cluster is created.
func BeforeClusterCreate(*BeforeClusterCreateRequest, *BeforeClusterCreateResponse) {}

// AfterControlPlaneInitializedRequest is the request of the AfterControlPlaneInitialized hook.
// +kubebuilder:object:root=true
type AfterControlPlaneInitializedRequest struct {
	metav1.TypeMeta `json:",inline"`

	// Cluster is the cluster object the lifecycle hook corresponds to.
	Cluster clusterv1.Cluster `json:"cluster"`
}

var _ ResponseObject = &AfterControlPlaneInitializedResponse{}

// AfterControlPlaneInitializedResponse is the response of the AfterControlPlaneInitialized hook.
// +kubebuilder:object:root=true
type AfterControlPlaneInitializedResponse struct {
	metav1.TypeMeta `json:",inline"`

	// CommonResponse contains Status and Message fields common to all response types.
	CommonResponse `json:",inline"`
}

// AfterControlPlaneInitialized is the hook that will be called after the control plane is initialized for the first time.
func AfterControlPlaneInitialized(*AfterControlPlaneInitializedRequest, *AfterControlPlaneInitializedResponse) {
}

// BeforeClusterUpgradeRequest is the request of the BeforeClusterUpgrade hook.
// +kubebuilder:object:root=true
type BeforeClusterUpgradeRequest struct {
	metav1.TypeMeta `json:",inline"`

	// Cluster is the cluster object the lifecycle hook corresponds to.
	Cluster clusterv1.Cluster `json:"cluster"`

	// FromKubernetesVersion is the current Kubernetes version of the cluster.
	FromKubernetesVersion string `json:"fromKubernetesVersion"`

	// ToKubernetesVersion is the target Kubernetes version of the upgrade.
	ToKubernetesVersion string `json:"toKubernetesVersion"`
}

var _ RetryResponseObject = &BeforeClusterUpgradeResponse{}

// BeforeClusterUpgradeResponse is the response of the BeforeClusterUpgrade hook.
// +kubebuilder:object:root=true
type BeforeClusterUpgradeResponse struct {
	metav1.TypeMeta `json:",inline"`

	// CommonRetryResponse contains Status, Message and RetryAfterSeconds fields.
	CommonRetryResponse `json:",inline"`
}

// BeforeClusterUpgrade is the hook that will be called after a Cluster.spec.version is upgraded and
// before the updated version is propagated to the underlying objects.
func BeforeClusterUpgrade(*BeforeClusterUpgradeRequest, *BeforeClusterUpgradeResponse) {}

// AfterControlPlaneUpgradeRequest is the request of the AfterControlPlaneUpgrade hook.
// +kubebuilder:object:root=true
type AfterControlPlaneUpgradeRequest struct {
	metav1.TypeMeta `json:",inline"`

	// Cluster is the cluster object the lifecycle hook corresponds to.
	Cluster clusterv1.Cluster `json:"cluster"`

	// KubernetesVersion is the Kubernetes version of the Control Plane after the upgrade.
	KubernetesVersion string `json:"kubernetesVersion"`
}

var _ RetryResponseObject = &AfterControlPlaneUpgradeResponse{}

// AfterControlPlaneUpgradeResponse is the response of the AfterControlPlaneUpgrade hook.
// +kubebuilder:object:root=true
type AfterControlPlaneUpgradeResponse struct {
	metav1.TypeMeta `json:",inline"`

	// CommonRetryResponse contains Status, Message and RetryAfterSeconds fields.
	CommonRetryResponse `json:",inline"`
}

// AfterControlPlaneUpgrade is the hook called after the control plane is successfully upgraded to the target
// Kubernetes version and before the target version is propagated to the workload machines.
func AfterControlPlaneUpgrade(*AfterControlPlaneUpgradeRequest, *AfterControlPlaneUpgradeResponse) {}

// AfterClusterUpgradeRequest is the request of the AfterClusterUpgrade hook.
// +kubebuilder:object:root=true
type AfterClusterUpgradeRequest struct {
	metav1.TypeMeta `json:",inline"`

	// Cluster is the cluster object the lifecycle hook corresponds to.
	Cluster clusterv1.Cluster `json:"cluster"`

	// KubernetesVersion is the Kubernetes version after upgrade.
	KubernetesVersion string `json:"kubernetesVersion"`
}

var _ ResponseObject = &AfterClusterUpgradeResponse{}

// AfterClusterUpgradeResponse is the response of the AfterClusterUpgrade hook.
// +kubebuilder:object:root=true
type AfterClusterUpgradeResponse struct {
	metav1.TypeMeta `json:",inline"`

	// CommonResponse contains Status and Message fields common to all response types.
	CommonResponse `json:",inline"`
}

// AfterClusterUpgrade is the hook that is called after the entire cluster is updated
// to the target Kubernetes version.
func AfterClusterUpgrade(*AfterClusterUpgradeRequest, *AfterClusterUpgradeResponse) {}

// BeforeClusterDeleteRequest is the request of the BeforeClusterDelete hook.
// +kubebuilder:object:root=true
type BeforeClusterDeleteRequest struct {
	metav1.TypeMeta `json:",inline"`

	// Cluster is the cluster object the lifecycle hook corresponds to.
	Cluster clusterv1.Cluster `json:"cluster"`
}

var _ RetryResponseObject = &BeforeClusterDeleteResponse{}

// BeforeClusterDeleteResponse is the response of the BeforeClusterDelete hook.
// +kubebuilder:object:root=true
type BeforeClusterDeleteResponse struct {
	metav1.TypeMeta `json:",inline"`

	// CommonRetryResponse contains Status, Message and RetryAfterSeconds fields.
	CommonRetryResponse `json:",inline"`
}

// BeforeClusterDelete is the hook that is called after delete is issued on a cluster
// and before the cluster and its underlying objects are deleted.
func BeforeClusterDelete(*BeforeClusterDeleteRequest, *BeforeClusterDeleteResponse) {}

func init() {
	catalogBuilder.RegisterHook(BeforeClusterCreate, &runtimecatalog.HookMeta{
		Tags:        []string{"Lifecycle Hooks"},
		Summary:     "Called before Cluster topology is created.",
		Description: "This blocking hook is called after the Cluster is created by the user and immediately before all the objects which are part of a Cluster topology are going to be created.",
	})

	catalogBuilder.RegisterHook(AfterControlPlaneInitialized, &runtimecatalog.HookMeta{
		Tags:        []string{"Lifecycle Hooks"},
		Summary:     "Called after the Control Plane is initialized for the first time.",
		Description: "This non-blocking hook is called after the Control Plane for the Cluster is reachable for the first time.",
	})

	catalogBuilder.RegisterHook(BeforeClusterUpgrade, &runtimecatalog.HookMeta{
		Tags:        []string{"Lifecycle Hooks"},
		Summary:     "Called before the Cluster is upgraded.",
		Description: "This blocking hook is called after the Cluster object has been updated with a new spec.topology.version by the user, and immediately before the new version is propagated to the Control Plane.",
	})

	catalogBuilder.RegisterHook(AfterControlPlaneUpgrade, &runtimecatalog.HookMeta{
		Tags:        []string{"Lifecycle Hooks"},
		Summary:     "Called after the Control Plane is upgraded.",
		Description: "This blocking hook is called after the Control Plane has been upgraded to the version specified in spec.topology.version, and immediately before the new version is propagated to the MachineDeployments.",
	})

	catalogBuilder.RegisterHook(AfterClusterUpgrade, &runtimecatalog.HookMeta{
		Tags:        []string{"Lifecycle Hooks"},
		Summary:     "Called after the Cluster is upgraded.",
		Description: "This non-blocking hook is called after the Cluster, Control Plane and MachineDeployments have been upgraded to the version specified in spec.topology.version.",
	})

	catalogBuilder.RegisterHook(BeforeClusterDelete, &runtimecatalog.HookMeta{
		Tags:        []string{"Lifecycle Hooks"},
		Summary:     "Called before the Cluster is deleted.",
		Description: "This blocking hook is called after the Cluster deletion has been triggered by the user, and immediately before objects of the Cluster are going to be deleted.",
	})
}

/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	api "kubepack.dev/lib-app/api/v1alpha1"
)

const (
	ResourceKindVaultServer = "VaultServer"
	ResourceVaultServer     = "vaultserver"
	ResourceVaultServers    = "vaultservers"
)

// VaultServer defines the schama for KubeVault Operator Installer.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=vaultservers,singular=vaultserver,categories={kubevault,appscode}
type VaultServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VaultServerSpec `json:"spec,omitempty"`
}

// VaultServerSpec is the schema for kubedb-autoscaler chart values file
type VaultServerSpec struct {
	api.Metadata `json:"metadata,omitempty"`
	Form         VaultServerSpecForm `json:"form"`
}

type VaultServerSpecForm struct {
	Alert VaultServerAlert `json:"alert"`
}

type VaultServerAlert struct {
	Enabled bool              `json:"enabled"`
	Labels  map[string]string `json:"labels"`
	// +optional
	Annotations map[string]string `json:"annotations"`
	// +optional
	AdditionalRuleLabels map[string]string      `json:"additionalRuleLabels"`
	Groups               VaultServerAlertGroups `json:"groups"`
}

type VaultServerAlertGroups struct {
	Database    VaultServerVaultAlert `json:"vault"`
	Provisioner ProvisionerAlert      `json:"provisioner"`
	OpsManager  OpsManagerAlert       `json:"opsManager"`
	Stash       StashAlert            `json:"stash"`
}

type VaultServerVaultAlert struct {
	Enabled bool                       `json:"enabled"`
	Rules   VaultServerVaultAlertRules `json:"rules"`
}

type VaultServerVaultAlertRules struct {
	VaultSealed                  FixedAlert  `json:"vaultSealed"`
	VaultAutoPilotNodeUnhealthy  FixedAlert  `json:"vaultAutoPilotNodeUnhealthy"`
	VaultLeadershipLoss          IntValAlert `json:"vaultLeadershipLoss"`
	VaultLeadershipStepsDowns    IntValAlert `json:"vaultLeadershipStepsDowns"`
	VaultLeadershipSetupFailures IntValAlert `json:"vaultLeadershipSetupFailures"`
	VaultRequestFailures         FixedAlert  `json:"vaultRequestFailures"`
	VaultResponseFailures        FixedAlert  `json:"vaultResponseFailures"`
	VaultTooManyInfinityTokens   IntValAlert `json:"vaultTooManyInfinityTokens"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VaultServerList is a list of Rediss
type VaultServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Redis CRD objects
	Items []VaultServer `json:"items,omitempty"`
}

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
	ResourceKindMongodb = "Mongodb"
	ResourceMongodb     = "mongodb"
	ResourceMongodbs    = "mongodbs"
)

// Mongodb defines the schama for KubeDB Ops Manager Operator Installer.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=mongodbs,singular=mongodb,categories={kubedb,appscode}
type Mongodb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MongodbSpec `json:"spec,omitempty"`
}

// MongodbSpec is the schema for kubedb-autoscaler chart values file
type MongodbSpec struct {
	api.Metadata `json:"metadata,omitempty"`
	Form         MongodbSpecForm `json:"form"`
}

type MongodbSpecForm struct {
	Alert MongoDBAlert `json:"alert"`
}

type MongoDBAlert struct {
	Enabled bool              `json:"enabled"`
	Labels  map[string]string `json:"labels"`
	// +optional
	Annotations map[string]string `json:"annotations"`
	// +optional
	AdditionalRuleLabels map[string]string  `json:"additionalRuleLabels"`
	Groups               MongoDBAlertGroups `json:"groups"`
}

type MongoDBAlertGroups struct {
	Database      MongoDBDatabaseAlert `json:"database"`
	Provisioner   ProvisionerAlert     `json:"provisioner"`
	OpsManager    OpsManagerAlert      `json:"opsManager"`
	Stash         StashAlert           `json:"stash"`
	SchemaManager SchemaManagerAlert   `json:"schemaManager"`
}

type MongoDBDatabaseAlert struct {
	Enabled bool                      `json:"enabled"`
	Rules   MongoDBDatabaseAlertRules `json:"rules"`
}

type MongoDBDatabaseAlertRules struct {
	MongodbVirtualMemoryUsage        IntValAlert `json:"mongodbVirtualMemoryUsage"`
	MongodbReplicationLag            IntValAlert `json:"mongodbReplicationLag"`
	MongodbNumberCursorsOpen         IntValAlert `json:"mongodbNumberCursorsOpen"`
	MongodbCursorsTimeouts           IntValAlert `json:"mongodbCursorsTimeouts"`
	MongodbTooManyConnections        IntValAlert `json:"mongodbTooManyConnections"`
	MongoDBPhaseCritical             FixedAlert  `json:"mongoDBPhaseCritical"`
	MongoDBDown                      FixedAlert  `json:"mongoDBDown"`
	MongodbHighLatency               IntValAlert `json:"mongodbHighLatency"`
	MongodbHighTicketUtilization     IntValAlert `json:"mongodbHighTicketUtilization"`
	MongodbRecurrentCursorTimeout    IntValAlert `json:"mongodbRecurrentCursorTimeout"`
	MongodbRecurrentMemoryPageFaults IntValAlert `json:"mongodbRecurrentMemoryPageFaults"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MongodbList is a list of Mongodbs
type MongodbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Mongodb CRD objects
	Items []Mongodb `json:"items,omitempty"`
}

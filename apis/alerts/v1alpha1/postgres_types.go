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
	ResourceKindPostgres = "Postgres"
	ResourcePostgres     = "postgres"
	ResourcePostgress    = "postgress"
)

// Postgres defines the schama for KubeDB Ops Manager Operator Installer.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=postgress,singular=postgres,categories={kubedb,appscode}
type Postgres struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresSpec `json:"spec,omitempty"`
}

// PostgresSpec is the schema for kubedb-autoscaler chart values file
type PostgresSpec struct {
	api.Metadata `json:"metadata,omitempty"`
	Spec         PostgresSpecSpec `json:"spec"`
}

type PostgresSpecSpec struct {
	Alert PostgresAlert `json:"alert"`
}

type PostgresAlert struct {
	Enabled              bool                `json:"enabled"`
	Labels               map[string]string   `json:"labels"`
	Annotations          map[string]string   `json:"annotations"`
	AdditionalRuleLabels map[string]string   `json:"additionalRuleLabels"`
	Groups               PostgresAlertGroups `json:"groups"`
}

type PostgresAlertGroups struct {
	Database      PostgresDatabaseAlert `json:"database"`
	Provisioner   ProvisionerAlert      `json:"provisioner"`
	OpsManager    OpsManagerAlert       `json:"opsManager"`
	Stash         StashAlert            `json:"stash"`
	SchemaManager SchemaManagerAlert    `json:"schemaManager"`
}

type PostgresDatabaseAlert struct {
	Enabled bool                       `json:"enabled"`
	Rules   PostgresDatabaseAlertRules `json:"rules"`
}

type PostgresDatabaseAlertRules struct {
	PostgresInstanceDown           FixedAlert          `json:"postgresInstanceDown"`
	PostgresRestarted              FixedAlert          `json:"postgresRestarted"`
	PostgresTooManyConnections     IntValAlert         `json:"postgresTooManyConnections"`
	PostgresqlNotEnoughConnections IntValAlert         `json:"postgresqlNotEnoughConnections"`
	PostgresSlowQueries            FixedAlert          `json:"postgresSlowQueries"`
	PostgresqlReplicationLag       StringValAlert      `json:"postgresqlReplicationLag"`
	PostgresqlHighRollbackRate     FloatValAlertConfig `json:"postgresqlHighRollbackRate"`
	PostgresqlSplitBrain           FixedAlert          `json:"postgresqlSplitBrain"`
	PostgresqlTooManyLocksAcquired FloatValAlertConfig `json:"postgresqlTooManyLocksAcquired"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PostgresList is a list of Postgress
type PostgresList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Postgres CRD objects
	Items []Postgres `json:"items,omitempty"`
}

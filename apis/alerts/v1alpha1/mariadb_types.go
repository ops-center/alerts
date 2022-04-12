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
	ResourceKindMariadb = "Mariadb"
	ResourceMariadb     = "mariadb"
	ResourceMariadbs    = "mariadbs"
)

// Mariadb defines the schama for KubeDB Ops Manager Operator Installer.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=mariadbs,singular=mariadb,categories={kubedb,appscode}
type Mariadb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MariadbSpec `json:"spec,omitempty"`
}

// MariadbSpec is the schema for kubedb-autoscaler chart values file
type MariadbSpec struct {
	api.Metadata `json:"metadata,omitempty"`
	Spec         MariadbSpecSpec `json:"spec"`
}

type MariadbSpecSpec struct {
	Alert MariaDBAlert `json:"alert"`
}

type MariaDBAlert struct {
	Enabled              bool               `json:"enabled"`
	Labels               map[string]string  `json:"labels"`
	Annotations          map[string]string  `json:"annotations"`
	AdditionalRuleLabels map[string]string  `json:"additionalRuleLabels"`
	Groups               MariaDBAlertGroups `json:"groups"`
}

type MariaDBAlertGroups struct {
	Database      MariaDBDatabaseAlert `json:"database"`
	Cluster       MariaDBClusterAlert  `json:"cluster"`
	Provisioner   ProvisionerAlert     `json:"provisioner"`
	OpsManager    OpsManagerAlert      `json:"opsManager"`
	Stash         StashAlert           `json:"stash"`
	SchemaManager SchemaManagerAlert   `json:"schemaManager"`
}

type MariaDBDatabaseAlert struct {
	Enabled bool                      `json:"enabled"`
	Rules   MariaDBDatabaseAlertRules `json:"rules"`
}

type MariaDBDatabaseAlertRules struct {
	MySQLInstanceDown       FixedAlert  `json:"mySQLInstanceDown"`
	MySQLServiceDown        FixedAlert  `json:"mySQLServiceDown"`
	MySQLTooManyConnections IntValAlert `json:"mySQLTooManyConnections"`
	MySQLHighThreadsRunning IntValAlert `json:"mySQLHighThreadsRunning"`
	MySQLSlowQueries        FixedAlert  `json:"mySQLSlowQueries"`
	MySQLInnoDBLogWaits     IntValAlert `json:"mySQLInnoDBLogWaits"`
	MySQLRestarted          IntValAlert `json:"mySQLRestarted"`
	MySQLHighQPS            IntValAlert `json:"mySQLHighQPS"`
	MySQLHighIncomingBytes  IntValAlert `json:"mySQLHighIncomingBytes"`
	MySQLHighOutgoingBytes  IntValAlert `json:"mySQLHighOutgoingBytes"`
	MySQLTooManyOpenFiles   IntValAlert `json:"mySQLTooManyOpenFiles"`
}

type MariaDBClusterAlert struct {
	Enabled bool                     `json:"enabled"`
	Rules   MariaDBClusterAlertRules `json:"rules"`
}

type MariaDBClusterAlertRules struct {
	GaleraReplicationLatencyTooLong FloatValAlertConfig `json:"galeraReplicationLatencyTooLong"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MariadbList is a list of Mariadbs
type MariadbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Mariadb CRD objects
	Items []Mariadb `json:"items,omitempty"`
}

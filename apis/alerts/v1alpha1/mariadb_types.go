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
	Namespace string       `json:"namespace"`
	DbName    string       `json:"dbName"`
	Alert     MariaDBAlert `json:"alert"`
}

type MariaDBAlert struct {
	RuleSelector AlertRuleSelector  `json:"ruleSelector"`
	Groups       MariaDBAlertGroups `json:"groups"`
}

type MariaDBAlertGroups struct {
	Database   MariaDBDatabase   `json:"database"`
	Cluster    MariaDBCluster    `json:"cluster"`
	Kubedb     MariaDBKubedb     `json:"kubedb"`
	Opsrequest MariaDBOpsrequest `json:"opsrequest"`
	Stash      MariaDBStash      `json:"stash"`
	Schema     MariaDBSchema     `json:"schema"`
}

type MariaDBDatabase struct {
	Enabled bool                 `json:"enabled"`
	Rules   MariaDBDatabaseRules `json:"rules"`
}

type MariaDBDatabaseRules struct {
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

type MariaDBCluster struct {
	Enabled bool                `json:"enabled"`
	Rules   MariaDBClusterRules `json:"rules"`
}

type MariaDBClusterRules struct {
	GaleraReplicationLatencyTooLong FloatValAlertConfig `json:"galeraReplicationLatencyTooLong"`
}

type MariaDBKubedb struct {
	Enabled bool               `json:"enabled"`
	Rules   MariaDBKubedbRules `json:"rules"`
}

type MariaDBKubedbRules struct {
	KubeDBMariaDBPhaseNotReady FixedAlert `json:"kubeDBMariaDBPhaseNotReady"`
	KubeDBMariaDBPhaseCritical FixedAlert `json:"kubeDBMariaDBPhaseCritical"`
}

type MariaDBOpsrequest struct {
	Enabled bool                   `json:"enabled"`
	Rules   MariaDBOpsrequestRules `json:"rules"`
}

type MariaDBOpsrequestRules struct {
	KubeDBMariaDBOpsRequestOnProgress              FixedAlert `json:"kubeDBMariaDBOpsRequestOnProgress"`
	KubeDBMariaDBOpsRequestStatusProgressingToLong FixedAlert `json:"kubeDBMariaDBOpsRequestStatusProgressingToLong"`
	KubeDBMariaDBOpsRequestFailed                  FixedAlert `json:"kubeDBMariaDBOpsRequestFailed"`
}

type MariaDBStash struct {
	Enabled bool              `json:"enabled"`
	Rules   MariaDBStashRules `json:"rules"`
}

type MariaDBStashRules struct {
	MariaDBStashBackupSessionFailed         FixedAlert  `json:"mariaDBStashBackupSessionFailed"`
	MariaDBStashRestoreSessionFailed        FixedAlert  `json:"mariaDBStashRestoreSessionFailed"`
	MariaDBStashNoBackupSessionForTooLong   IntValAlert `json:"mariaDBStashNoBackupSessionForTooLong"`
	MariaDBStashRepositoryCorrupted         FixedAlert  `json:"mariaDBStashRepositoryCorrupted"`
	MariaDBStashRepositoryStorageRunningLow IntValAlert `json:"mariaDBStashRepositoryStorageRunningLow"`
}

type MariaDBSchema struct {
	Enabled bool               `json:"enabled"`
	Rules   MariaDBSchemaRules `json:"rules"`
}

type MariaDBSchemaRules struct {
	KubeDBMariaDBSchemaPendingForTooLong     FixedAlert `json:"kubeDBMariaDBSchemaPendingForTooLong"`
	KubeDBMariaDBSchemaInProgressForTooLong  FixedAlert `json:"kubeDBMariaDBSchemaInProgressForTooLong"`
	KubeDBMariaDBSchemaTerminatingForTooLong FixedAlert `json:"kubeDBMariaDBSchemaTerminatingForTooLong"`
	KubeDBMariaDBSchemaFailed                FixedAlert `json:"kubeDBMariaDBSchemaFailed"`
	KubeDBMariaDBSchemaExpired               FixedAlert `json:"kubeDBMariaDBSchemaExpired"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MariadbList is a list of Mariadbs
type MariadbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Mariadb CRD objects
	Items []Mariadb `json:"items,omitempty"`
}

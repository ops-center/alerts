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
	ResourceKindMysql = "Mysql"
	ResourceMysql     = "mysql"
	ResourceMysqls    = "mysqls"
)

// Mysql defines the schama for KubeDB Ops Manager Operator Installer.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=mysqls,singular=mysql,categories={kubedb,appscode}
type Mysql struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MysqlSpec `json:"spec,omitempty"`
}

// MysqlSpec is the schema for kubedb-autoscaler chart values file
type MysqlSpec struct {
	Namespace string     `json:"namespace"`
	DbName    string     `json:"dbName"`
	Alert     MySQLAlert `json:"alert"`
}

type MySQLAlert struct {
	RuleSelector AlertRuleSelector `json:"ruleSelector"`
	Groups       MySQLAlertGroups  `json:"groups"`
}

type MySQLAlertGroups struct {
	Database   MySQLDatabase   `json:"database"`
	Group      MySQLGroup      `json:"group"`
	Kubedb     MySQLKubedb     `json:"kubedb"`
	Opsrequest MySQLOpsrequest `json:"opsrequest"`
	Stash      MySQLStash      `json:"stash"`
	Schema     MySQLSchema     `json:"schema"`
}

type MySQLDatabase struct {
	Enabled bool               `json:"enabled"`
	Rules   MySQLDatabaseRules `json:"rules"`
}

type MySQLDatabaseRules struct {
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

type MySQLGroup struct {
	Enabled bool            `json:"enabled"`
	Rules   MySQLGroupRules `json:"rules"`
}

type MySQLGroupRules struct {
	MySQLHighReplicationDelay           FloatValAlertConfig `json:"mySQLHighReplicationDelay"`
	MySQLHighReplicationTransportTime   FloatValAlertConfig `json:"mySQLHighReplicationTransportTime"`
	MySQLHighReplicationApplyTime       FloatValAlertConfig `json:"mySQLHighReplicationApplyTime"`
	MySQLReplicationHighTransactionTime FloatValAlertConfig `json:"mySQLReplicationHighTransactionTime"`
}

type MySQLKubedb struct {
	Enabled bool             `json:"enabled"`
	Rules   MySQLKubedbRules `json:"rules"`
}

type MySQLKubedbRules struct {
	KubeDBMySQLPhaseNotReady FixedAlert `json:"kubeDBMySQLPhaseNotReady"`
	KubeDBMySQLPhaseCritical FixedAlert `json:"kubeDBMySQLPhaseCritical"`
}

type MySQLOpsrequest struct {
	Enabled bool                 `json:"enabled"`
	Rules   MySQLOpsrequestRules `json:"rules"`
}

type MySQLOpsrequestRules struct {
	KubeDBMySQLOpsRequestOnProgress              FixedAlert `json:"kubeDBMySQLOpsRequestOnProgress"`
	KubeDBMySQLOpsRequestStatusProgressingToLong FixedAlert `json:"kubeDBMySQLOpsRequestStatusProgressingToLong"`
	KubeDBMySQLOpsRequestFailed                  FixedAlert `json:"kubeDBMySQLOpsRequestFailed"`
}

type MySQLStash struct {
	Enabled bool            `json:"enabled"`
	Rules   MySQLStashRules `json:"rules"`
}

type MySQLStashRules struct {
	MySQLStashBackupSessionFailed         FixedAlert  `json:"mySQLStashBackupSessionFailed"`
	MySQLStashRestoreSessionFailed        FixedAlert  `json:"mySQLStashRestoreSessionFailed"`
	MySQLStashNoBackupSessionForTooLong   IntValAlert `json:"mySQLStashNoBackupSessionForTooLong"`
	MySQLStashRepositoryCorrupted         FixedAlert  `json:"mySQLStashRepositoryCorrupted"`
	MySQLStashRepositoryStorageRunningLow IntValAlert `json:"mySQLStashRepositoryStorageRunningLow"`
	MySQLStashBackupSessionPeriodTooLong  IntValAlert `json:"mySQLStashBackupSessionPeriodTooLong"`
	MySQLStashRestoreSessionPeriodTooLong IntValAlert `json:"mySQLStashRestoreSessionPeriodTooLong"`
}

type MySQLSchema struct {
	Enabled bool             `json:"enabled"`
	Rules   MySQLSchemaRules `json:"rules"`
}

type MySQLSchemaRules struct {
	KubeDBMySQLSchemaPendingForTooLong     FixedAlert `json:"kubeDBMySQLSchemaPendingForTooLong"`
	KubeDBMySQLSchemaInProgressForTooLong  FixedAlert `json:"kubeDBMySQLSchemaInProgressForTooLong"`
	KubeDBMySQLSchemaTerminatingForTooLong FixedAlert `json:"kubeDBMySQLSchemaTerminatingForTooLong"`
	KubeDBMySQLSchemaFailed                FixedAlert `json:"kubeDBMySQLSchemaFailed"`
	KubeDBMySQLSchemaExpired               FixedAlert `json:"kubeDBMySQLSchemaExpired"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MysqlList is a list of Mysqls
type MysqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Mysql CRD objects
	Items []Mysql `json:"items,omitempty"`
}

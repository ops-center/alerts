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
	api.Metadata `json:"metadata,omitempty"`
	Form         MysqlSpecForm `json:"form"`
}

type MysqlSpecForm struct {
	Alert MySQLAlert `json:"alert"`
}

type MySQLAlert struct {
	Enabled bool              `json:"enabled"`
	Labels  map[string]string `json:"labels"`
	// +optional
	Annotations map[string]string `json:"annotations"`
	// +optional
	AdditionalRuleLabels map[string]string `json:"additionalRuleLabels"`
	Groups               MySQLAlertGroups  `json:"groups"`
}

type MySQLAlertGroups struct {
	Database      MySQLDatabaseAlert `json:"database"`
	Group         MySQLGroupAlert    `json:"group"`
	Provisioner   ProvisionerAlert   `json:"provisioner"`
	OpsManager    OpsManagerAlert    `json:"opsManager"`
	Stash         StashAlert         `json:"stash"`
	SchemaManager SchemaManagerAlert `json:"schemaManager"`
}

type MySQLDatabaseAlert struct {
	Enabled bool                    `json:"enabled"`
	Rules   MySQLDatabaseAlertRules `json:"rules"`
}

type MySQLDatabaseAlertRules struct {
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

type MySQLGroupAlert struct {
	Enabled bool                 `json:"enabled"`
	Rules   MySQLGroupAlertRules `json:"rules"`
}

type MySQLGroupAlertRules struct {
	MySQLHighReplicationDelay           FloatValAlertConfig `json:"mySQLHighReplicationDelay"`
	MySQLHighReplicationTransportTime   FloatValAlertConfig `json:"mySQLHighReplicationTransportTime"`
	MySQLHighReplicationApplyTime       FloatValAlertConfig `json:"mySQLHighReplicationApplyTime"`
	MySQLReplicationHighTransactionTime FloatValAlertConfig `json:"mySQLReplicationHighTransactionTime"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MysqlList is a list of Mysqls
type MysqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Mysql CRD objects
	Items []Mysql `json:"items,omitempty"`
}

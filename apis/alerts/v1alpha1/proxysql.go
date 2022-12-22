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
	ResourceKindProxysql = "Proxysql"
	ResourceProxysql     = "proxysql"
	ResourceProxysqls    = "proxysqls"
)

// Proxysql defines the schama for KubeDB Ops Manager Operator Installer.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=proxysqls,singular=proxysql,categories={kubedb,appscode}
type Proxysql struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProxysqlSpec `json:"spec,omitempty"`
}

// ProxysqlSpec is the schema for kubedb-autoscaler chart values file
type ProxysqlSpec struct {
	api.Metadata `json:"metadata,omitempty"`
	Form         ProxysqlSpecForm `json:"form"`
}

type ProxysqlSpecForm struct {
	Alert ProxySQLAlert `json:"alert"`
}

type ProxySQLAlert struct {
	Enabled bool              `json:"enabled"`
	Labels  map[string]string `json:"labels"`
	// +optional
	Annotations map[string]string `json:"annotations"`
	// +optional
	AdditionalRuleLabels map[string]string   `json:"additionalRuleLabels"`
	Groups               ProxySQLAlertGroups `json:"groups"`
}

type ProxySQLAlertGroups struct {
	Database    ProxySQLDatabaseAlert `json:"database"`
	Cluster     ProxySQLClusterAlert  `json:"cluster"`
	Provisioner ProvisionerAlert      `json:"provisioner"`
	OpsManager  OpsManagerAlert       `json:"opsManager"`
}

type ProxySQLDatabaseAlert struct {
	Enabled bool                       `json:"enabled"`
	Rules   ProxySQLDatabaseAlertRules `json:"rules"`
}

type ProxySQLDatabaseAlertRules struct {
	ProxySQLInstanceDown       FixedAlert  `json:"proxySQLInstanceDown"`
	ProxySQLServiceDown        FixedAlert  `json:"proxySQLServiceDown"`
	ProxySQLTooManyConnections IntValAlert `json:"proxySQLTooManyConnections"`
	ProxySQLHighThreadsRunning IntValAlert `json:"proxySQLHighThreadsRunning"`
	ProxySQLSlowQueries        FixedAlert  `json:"proxySQLSlowQueries"`
	ProxySQLRestarted          IntValAlert `json:"proxySQLRestarted"`
	ProxySQLHighQPS            IntValAlert `json:"proxySQLHighQPS"`
	ProxySQLHighIncomingBytes  IntValAlert `json:"proxySQLHighIncomingBytes"`
	ProxySQLHighOutgoingBytes  IntValAlert `json:"proxySQLHighOutgoingBytes"`
}

type ProxySQLClusterAlert struct {
	Enabled bool                      `json:"enabled"`
	Rules   ProxySQLClusterAlertRules `json:"rules"`
}

type ProxySQLClusterAlertRules struct {
	ProxysqlCLusterSyncFailure FloatValAlertConfig `json:"proxysqlCLusterSyncFailure"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProxysqlList is a list of Proxysqls
type ProxysqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Proxysql CRD objects
	Items []Proxysql `json:"items,omitempty"`
}

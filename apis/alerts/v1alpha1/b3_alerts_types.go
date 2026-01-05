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
	mona "kmodules.xyz/monitoring-agent-api/api/v1"
	api "x-helm.dev/apimachinery/apis/releases/v1alpha1"
)

const (
	ResourceKindB3Alerts = "B3Alerts"
	ResourceB3Alerts     = "b3alerts"
	ResourceB3Alertss    = "b3alertss"
)

// B3Alerts defines the schema for B3 Platform Alerting Rules

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=b3alertss,singular=b3alerts,categories={b3,appscode}
type B3Alerts struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              B3AlertsSpec `json:"spec,omitempty"`
}

// B3AlertsSpec is the schema for B3 alerts chart values file
type B3AlertsSpec struct {
	api.Metadata `json:"metadata,omitempty"`
	Form         B3AlertsSpecForm `json:"form"`
	Grafana      Grafana          `json:"grafana"`
}

type B3AlertsSpecForm struct {
	Alert B3Alert `json:"alert"`
}

type B3Alert struct {
	Enabled mona.SeverityFlag `json:"enabled"`
	Labels  map[string]string `json:"labels"`
	// +optional
	Annotations map[string]string `json:"annotations"`
	// +optional
	AdditionalRuleLabels map[string]string `json:"additionalRuleLabels"`
	Groups               B3AlertGroups     `json:"groups"`
}

type B3AlertGroups struct {
	Process   B3ProcessAlert   `json:"process"`
	HTTP      B3HTTPAlert      `json:"http"`
	Database  B3DatabaseAlert  `json:"database"`
	Messaging B3MessagingAlert `json:"messaging"`
	Auth      B3AuthAlert      `json:"auth"`
	License   B3LicenseAlert   `json:"license"`
	Storage   B3StorageAlert   `json:"storage"`
	Proxy     B3ProxyAlert     `json:"proxy"`
}

// Process & Runtime alerts
type B3ProcessAlert struct {
	Enabled mona.SeverityFlag   `json:"enabled"`
	Rules   B3ProcessAlertRules `json:"rules"`
}

type B3ProcessAlertRules struct {
	B3HighCpuUsage            IntValAlert         `json:"b3HighCpuUsage"`
	B3CriticalCpuUsage        IntValAlert         `json:"b3CriticalCpuUsage"`
	B3HighMemoryUsage         IntValAlert         `json:"b3HighMemoryUsage"`
	B3CriticalMemoryUsage     IntValAlert         `json:"b3CriticalMemoryUsage"`
	B3ProcessRestarted        FixedAlert          `json:"b3ProcessRestarted"`
	B3HighGoroutineCount      IntValAlert         `json:"b3HighGoroutineCount"`
	B3HighGcPauseDuration     FloatValAlertConfig `json:"b3HighGcPauseDuration"`
	B3HighFileDescriptorUsage IntValAlert         `json:"b3HighFileDescriptorUsage"`
}

// HTTP Server alerts
type B3HTTPAlert struct {
	Enabled mona.SeverityFlag `json:"enabled"`
	Rules   B3HTTPAlertRules  `json:"rules"`
}

type B3HTTPAlertRules struct {
	B3HighHttpErrorRate        IntValAlert `json:"b3HighHttpErrorRate"`
	B3HighHttpErrorRateWarning IntValAlert `json:"b3HighHttpErrorRateWarning"`
	B3HighHttpLatency          IntValAlert `json:"b3HighHttpLatency"`
	B3CriticalHttpLatency      IntValAlert `json:"b3CriticalHttpLatency"`
	B3HighActiveRequests       IntValAlert `json:"b3HighActiveRequests"`
}

// Database alerts
type B3DatabaseAlert struct {
	Enabled mona.SeverityFlag    `json:"enabled"`
	Rules   B3DatabaseAlertRules `json:"rules"`
}

type B3DatabaseAlertRules struct {
	B3HighDatabaseErrorRate           IntValAlert `json:"b3HighDatabaseErrorRate"`
	B3HighDatabaseLatency             IntValAlert `json:"b3HighDatabaseLatency"`
	B3DatabaseConnectionErrors        IntValAlert `json:"b3DatabaseConnectionErrors"`
	B3HighDatabaseConnectionPoolUsage IntValAlert `json:"b3HighDatabaseConnectionPoolUsage"`
	B3NoIdleDatabaseConnections       FixedAlert  `json:"b3NoIdleDatabaseConnections"`
}

// Messaging/NATS alerts
type B3MessagingAlert struct {
	Enabled mona.SeverityFlag     `json:"enabled"`
	Rules   B3MessagingAlertRules `json:"rules"`
}

type B3MessagingAlertRules struct {
	B3HighMessagePublishErrorRate  IntValAlert `json:"b3HighMessagePublishErrorRate"`
	B3HighMessageProcessingLatency IntValAlert `json:"b3HighMessageProcessingLatency"`
	B3NatsDisconnections           IntValAlert `json:"b3NatsDisconnections"`
	B3NatsSubscriptionErrors       IntValAlert `json:"b3NatsSubscriptionErrors"`
	B3NoActiveNatsConnections      FixedAlert  `json:"b3NoActiveNatsConnections"`
}

// Authentication/Authorization alerts
type B3AuthAlert struct {
	Enabled mona.SeverityFlag `json:"enabled"`
	Rules   B3AuthAlertRules  `json:"rules"`
}

type B3AuthAlertRules struct {
	B3HighAuthFailureRate     IntValAlert `json:"b3HighAuthFailureRate"`
	B3CriticalAuthFailureRate IntValAlert `json:"b3CriticalAuthFailureRate"`
	B3HighAuthzDenialRate     IntValAlert `json:"b3HighAuthzDenialRate"`
}

// License alerts
type B3LicenseAlert struct {
	Enabled mona.SeverityFlag   `json:"enabled"`
	Rules   B3LicenseAlertRules `json:"rules"`
}

type B3LicenseAlertRules struct {
	B3LicenseIssuanceFailures   IntValAlert `json:"b3LicenseIssuanceFailures"`
	B3LicenseValidationFailures IntValAlert `json:"b3LicenseValidationFailures"`
}

// Storage/BlobFS alerts
type B3StorageAlert struct {
	Enabled mona.SeverityFlag   `json:"enabled"`
	Rules   B3StorageAlertRules `json:"rules"`
}

type B3StorageAlertRules struct {
	B3HighBlobStorageErrorRate IntValAlert `json:"b3HighBlobStorageErrorRate"`
}

// Proxy/K8s Operations alerts
type B3ProxyAlert struct {
	Enabled mona.SeverityFlag `json:"enabled"`
	Rules   B3ProxyAlertRules `json:"rules"`
}

type B3ProxyAlertRules struct {
	B3HighProxyErrorRate IntValAlert `json:"b3HighProxyErrorRate"`
	B3HighProxyLatency   IntValAlert `json:"b3HighProxyLatency"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// B3AlertsList is a list of B3Alerts
type B3AlertsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of B3Alerts CRD objects
	Items []B3Alerts `json:"items,omitempty"`
}

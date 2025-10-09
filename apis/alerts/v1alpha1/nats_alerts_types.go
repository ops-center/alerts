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
	ResourceKindNATSAlerts = "NATSAlerts"
	ResourceNATSAlerts     = "natsalerts"
	ResourceNATSAlertss    = "natsalertss"
)

// NATSAlerts defines the schema for NATS Alerting Rules

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=natsalertss,singular=natsalerts,categories={nats,appscode}
type NATSAlerts struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NATSAlertsSpec `json:"spec,omitempty"`
}

// NATSAlertsSpec is the schema for NATS alerts chart values file
type NATSAlertsSpec struct {
	api.Metadata `json:"metadata,omitempty"`
	Form         NATSAlertsSpecForm `json:"form"`
	Grafana      Grafana            `json:"grafana"`
}

type NATSAlertsSpecForm struct {
	Alert NATSAlert `json:"alert"`
}

type NATSAlert struct {
	Enabled mona.SeverityFlag `json:"enabled"`
	Labels  map[string]string `json:"labels"`
	// +optional
	Annotations map[string]string `json:"annotations"`
	// +optional
	AdditionalRuleLabels map[string]string `json:"additionalRuleLabels"`
	Groups               NATSAlertGroups   `json:"groups"`
}

type NATSAlertGroups struct {
	Database      NATSDatabaseAlert  `json:"database"`
	Provisioner   ProvisionerAlert   `json:"provisioner"`
	OpsManager    OpsManagerAlert    `json:"opsManager"`
	Stash         StashAlert         `json:"stash"`
	KubeStash     KubeStashAlert     `json:"kubeStash"`
	SchemaManager SchemaManagerAlert `json:"schemaManager"`
}

type NATSDatabaseAlert struct {
	Enabled mona.SeverityFlag      `json:"enabled"`
	Rules   NATSDatabaseAlertRules `json:"rules"`
}

type NATSDatabaseAlertRules struct {
	NatsJetStreamHighMemoryUsage  IntValAlert `json:"natsJetStreamHighMemoryUsage"`
	NatsJetStreamHighStorageUsage IntValAlert `json:"natsJetStreamHighStorageUsage"`

	NatsJetStreamHighPendingMessages        IntValAlert `json:"natsJetStreamHighPendingMessages"`
	NatsJetStreamHighPendingMessagesWarning IntValAlert `json:"natsJetStreamHighPendingMessagesWarning"`
	NatsJetStreamHighAckPending             IntValAlert `json:"natsJetStreamHighAckPending"`
	NatsJetStreamConsumerStalled            FixedAlert  `json:"natsJetStreamConsumerStalled"`

	NatsJetStreamHighMessageCount IntValAlert `json:"natsJetStreamHighMessageCount"`
	NatsJetStreamHighStreamSize   IntValAlert `json:"natsJetStreamHighStreamSize"`
	NatsJetStreamNoIngestion      FixedAlert  `json:"natsJetStreamNoIngestion"`

	NatsSuddenConnectionDrop  IntValAlert `json:"natsSuddenConnectionDrop"`
	NatsHighActiveConnections IntValAlert `json:"natsHighActiveConnections"`

	NatsSuddenConsumerDrop IntValAlert `json:"natsSuddenConsumerDrop"`
	NatsHighTotalConsumers IntValAlert `json:"natsHighTotalConsumers"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NATSAlertsList is a list of NATSAlerts
type NATSAlertsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NATSAlerts CRD objects
	Items []NATSAlerts `json:"items,omitempty"`
}

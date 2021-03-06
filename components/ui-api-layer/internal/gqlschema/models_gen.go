// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlschema

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type API struct {
	Name                   string                 `json:"name"`
	Hostname               string                 `json:"hostname"`
	Service                Service                `json:"service"`
	AuthenticationPolicies []AuthenticationPolicy `json:"authenticationPolicies"`
}

type ApplicationEntry struct {
	Type        string  `json:"type"`
	GatewayURL  *string `json:"gatewayUrl"`
	AccessLabel *string `json:"accessLabel"`
}

type ApplicationEvent struct {
	Type        SubscriptionEventType `json:"type"`
	Application Application           `json:"application"`
}

type ApplicationMapping struct {
	Environment string `json:"environment"`
	Application string `json:"application"`
}

type ApplicationMutationOutput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Labels      Labels `json:"labels"`
}

type ApplicationService struct {
	ID                  string             `json:"id"`
	DisplayName         string             `json:"displayName"`
	LongDescription     string             `json:"longDescription"`
	ProviderDisplayName string             `json:"providerDisplayName"`
	Tags                []string           `json:"tags"`
	Entries             []ApplicationEntry `json:"entries"`
}

type AuthenticationPolicy struct {
	Type    AuthenticationPolicyType `json:"type"`
	Issuer  string                   `json:"issuer"`
	JwksURI string                   `json:"jwksURI"`
}

type BackendModule struct {
	Name string `json:"name"`
}

type BindableResourcesOutputItem struct {
	Kind        string              `json:"kind"`
	DisplayName string              `json:"displayName"`
	Resources   []UsageKindResource `json:"resources"`
}

type ClusterServiceBroker struct {
	Name              string              `json:"name"`
	Status            ServiceBrokerStatus `json:"status"`
	CreationTimestamp time.Time           `json:"creationTimestamp"`
	URL               string              `json:"url"`
	Labels            Labels              `json:"labels"`
}

type ClusterServiceBrokerEvent struct {
	Type                 SubscriptionEventType `json:"type"`
	ClusterServiceBroker ClusterServiceBroker  `json:"clusterServiceBroker"`
}

type ClusterServicePlan struct {
	Name                           string  `json:"name"`
	DisplayName                    *string `json:"displayName"`
	ExternalName                   string  `json:"externalName"`
	Description                    string  `json:"description"`
	RelatedClusterServiceClassName string  `json:"relatedClusterServiceClassName"`
	InstanceCreateParameterSchema  *JSON   `json:"instanceCreateParameterSchema"`
	BindingCreateParameterSchema   *JSON   `json:"bindingCreateParameterSchema"`
}

type ConnectorService struct {
	URL string `json:"url"`
}

type Container struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type ContainerState struct {
	State   ContainerStateType `json:"state"`
	Reason  string             `json:"reason"`
	Message string             `json:"message"`
}

type CreateServiceBindingOutput struct {
	Name                string `json:"name"`
	ServiceInstanceName string `json:"serviceInstanceName"`
	Environment         string `json:"environment"`
}

type CreateServiceBindingUsageInput struct {
	Name              *string                             `json:"name"`
	Environment       string                              `json:"environment"`
	ServiceBindingRef ServiceBindingRefInput              `json:"serviceBindingRef"`
	UsedBy            LocalObjectReferenceInput           `json:"usedBy"`
	Parameters        *ServiceBindingUsageParametersInput `json:"parameters"`
}

type DeleteApplicationOutput struct {
	Name string `json:"name"`
}

type DeleteServiceBindingOutput struct {
	Name        string `json:"name"`
	Environment string `json:"environment"`
}

type DeleteServiceBindingUsageOutput struct {
	Name        string `json:"name"`
	Environment string `json:"environment"`
}

type Deployment struct {
	Name                      string           `json:"name"`
	Environment               string           `json:"environment"`
	CreationTimestamp         time.Time        `json:"creationTimestamp"`
	Status                    DeploymentStatus `json:"status"`
	Labels                    Labels           `json:"labels"`
	Containers                []Container      `json:"containers"`
	BoundServiceInstanceNames []string         `json:"boundServiceInstanceNames"`
}

type DeploymentCondition struct {
	Status                  string    `json:"status"`
	Type                    string    `json:"type"`
	LastTransitionTimestamp time.Time `json:"lastTransitionTimestamp"`
	LastUpdateTimestamp     time.Time `json:"lastUpdateTimestamp"`
	Message                 string    `json:"message"`
	Reason                  string    `json:"reason"`
}

type DeploymentStatus struct {
	Replicas          int                   `json:"replicas"`
	UpdatedReplicas   int                   `json:"updatedReplicas"`
	ReadyReplicas     int                   `json:"readyReplicas"`
	AvailableReplicas int                   `json:"availableReplicas"`
	Conditions        []DeploymentCondition `json:"conditions"`
}

type EnvPrefix struct {
	Name string `json:"name"`
}

type EnvPrefixInput struct {
	Name string `json:"name"`
}

type EventActivationEvent struct {
	EventType   string `json:"eventType"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

type ExceededQuota struct {
	QuotaName         string   `json:"quotaName"`
	ResourceName      string   `json:"resourceName"`
	AffectedResources []string `json:"affectedResources"`
}

type Function struct {
	Name              string    `json:"name"`
	Trigger           string    `json:"trigger"`
	CreationTimestamp time.Time `json:"creationTimestamp"`
	Labels            Labels    `json:"labels"`
	Environment       string    `json:"environment"`
}

type IDPPreset struct {
	Name    string `json:"name"`
	Issuer  string `json:"issuer"`
	JwksURI string `json:"jwksUri"`
}

type InputTopic struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type LimitRange struct {
	Name   string           `json:"name"`
	Limits []LimitRangeItem `json:"limits"`
}

type LimitRangeItem struct {
	LimitType      LimitType    `json:"limitType"`
	Max            ResourceType `json:"max"`
	Default        ResourceType `json:"default"`
	DefaultRequest ResourceType `json:"defaultRequest"`
}

type LocalObjectReference struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

type LocalObjectReferenceInput struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

type Pod struct {
	Name              string           `json:"name"`
	NodeName          string           `json:"nodeName"`
	Namespace         string           `json:"namespace"`
	RestartCount      int              `json:"restartCount"`
	CreationTimestamp time.Time        `json:"creationTimestamp"`
	Labels            Labels           `json:"labels"`
	Status            PodStatusType    `json:"status"`
	ContainerStates   []ContainerState `json:"containerStates"`
	JSON              JSON             `json:"json"`
}

type ResourceQuota struct {
	Name     string         `json:"name"`
	Pods     *string        `json:"pods"`
	Limits   ResourceValues `json:"limits"`
	Requests ResourceValues `json:"requests"`
}

type ResourceQuotasStatus struct {
	Exceeded       bool            `json:"exceeded"`
	ExceededQuotas []ExceededQuota `json:"exceededQuotas"`
}

type ResourceType struct {
	Memory *string `json:"memory"`
	CPU    *string `json:"cpu"`
}

type ResourceValues struct {
	Memory *string `json:"memory"`
	CPU    *string `json:"cpu"`
}

type Secret struct {
	Name        string `json:"name"`
	Environment string `json:"environment"`
	Data        JSON   `json:"data"`
}

type Service struct {
	Name string `json:"name"`
	Port int    `json:"port"`
}

type ServiceBindingEvent struct {
	Type           SubscriptionEventType `json:"type"`
	ServiceBinding ServiceBinding        `json:"serviceBinding"`
}

type ServiceBindingRefInput struct {
	Name string `json:"name"`
}

type ServiceBindingStatus struct {
	Type    ServiceBindingStatusType `json:"type"`
	Reason  string                   `json:"reason"`
	Message string                   `json:"message"`
}

type ServiceBindingUsageEvent struct {
	Type                SubscriptionEventType `json:"type"`
	ServiceBindingUsage ServiceBindingUsage   `json:"serviceBindingUsage"`
}

type ServiceBindingUsageParameters struct {
	EnvPrefix *EnvPrefix `json:"envPrefix"`
}

type ServiceBindingUsageParametersInput struct {
	EnvPrefix *EnvPrefixInput `json:"envPrefix"`
}

type ServiceBindingUsageStatus struct {
	Type    ServiceBindingUsageStatusType `json:"type"`
	Reason  string                        `json:"reason"`
	Message string                        `json:"message"`
}

type ServiceBindings struct {
	Items []ServiceBinding     `json:"items"`
	Stats ServiceBindingsStats `json:"stats"`
}

type ServiceBindingsStats struct {
	Ready   int `json:"ready"`
	Failed  int `json:"failed"`
	Pending int `json:"pending"`
	Unknown int `json:"unknown"`
}

type ServiceBroker struct {
	Name              string              `json:"name"`
	Environment       string              `json:"environment"`
	Status            ServiceBrokerStatus `json:"status"`
	CreationTimestamp time.Time           `json:"creationTimestamp"`
	URL               string              `json:"url"`
	Labels            Labels              `json:"labels"`
}

type ServiceBrokerEvent struct {
	Type          SubscriptionEventType `json:"type"`
	ServiceBroker ServiceBroker         `json:"serviceBroker"`
}

type ServiceBrokerStatus struct {
	Ready   bool   `json:"ready"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

type ServiceInstanceCreateInput struct {
	Name            string                                `json:"name"`
	Environment     string                                `json:"environment"`
	ClassRef        ServiceInstanceCreateInputResourceRef `json:"classRef"`
	PlanRef         ServiceInstanceCreateInputResourceRef `json:"planRef"`
	Labels          []string                              `json:"labels"`
	ParameterSchema *JSON                                 `json:"parameterSchema"`
}

type ServiceInstanceCreateInputResourceRef struct {
	ExternalName string `json:"externalName"`
	ClusterWide  bool   `json:"clusterWide"`
}

type ServiceInstanceEvent struct {
	Type            SubscriptionEventType `json:"type"`
	ServiceInstance ServiceInstance       `json:"serviceInstance"`
}

type ServiceInstanceResourceRef struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	ClusterWide bool   `json:"clusterWide"`
}

type ServiceInstanceStatus struct {
	Type    InstanceStatusType `json:"type"`
	Reason  string             `json:"reason"`
	Message string             `json:"message"`
}

type ServicePlan struct {
	Name                          string  `json:"name"`
	Environment                   string  `json:"environment"`
	DisplayName                   *string `json:"displayName"`
	ExternalName                  string  `json:"externalName"`
	Description                   string  `json:"description"`
	RelatedServiceClassName       string  `json:"relatedServiceClassName"`
	InstanceCreateParameterSchema *JSON   `json:"instanceCreateParameterSchema"`
	BindingCreateParameterSchema  *JSON   `json:"bindingCreateParameterSchema"`
}

type UsageKind struct {
	Name        string `json:"name"`
	Group       string `json:"group"`
	Kind        string `json:"kind"`
	Version     string `json:"version"`
	DisplayName string `json:"displayName"`
}

type UsageKindResource struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type ApplicationStatus string

const (
	ApplicationStatusServing              ApplicationStatus = "SERVING"
	ApplicationStatusNotServing           ApplicationStatus = "NOT_SERVING"
	ApplicationStatusGatewayNotConfigured ApplicationStatus = "GATEWAY_NOT_CONFIGURED"
)

func (e ApplicationStatus) IsValid() bool {
	switch e {
	case ApplicationStatusServing, ApplicationStatusNotServing, ApplicationStatusGatewayNotConfigured:
		return true
	}
	return false
}

func (e ApplicationStatus) String() string {
	return string(e)
}

func (e *ApplicationStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ApplicationStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ApplicationStatus", str)
	}
	return nil
}

func (e ApplicationStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type AuthenticationPolicyType string

const (
	AuthenticationPolicyTypeJwt AuthenticationPolicyType = "JWT"
)

func (e AuthenticationPolicyType) IsValid() bool {
	switch e {
	case AuthenticationPolicyTypeJwt:
		return true
	}
	return false
}

func (e AuthenticationPolicyType) String() string {
	return string(e)
}

func (e *AuthenticationPolicyType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AuthenticationPolicyType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AuthenticationPolicyType", str)
	}
	return nil
}

func (e AuthenticationPolicyType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ContainerStateType string

const (
	ContainerStateTypeWaiting    ContainerStateType = "WAITING"
	ContainerStateTypeRunning    ContainerStateType = "RUNNING"
	ContainerStateTypeTerminated ContainerStateType = "TERMINATED"
)

func (e ContainerStateType) IsValid() bool {
	switch e {
	case ContainerStateTypeWaiting, ContainerStateTypeRunning, ContainerStateTypeTerminated:
		return true
	}
	return false
}

func (e ContainerStateType) String() string {
	return string(e)
}

func (e *ContainerStateType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContainerStateType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContainerStateType", str)
	}
	return nil
}

func (e ContainerStateType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type InstanceStatusType string

const (
	InstanceStatusTypeRunning        InstanceStatusType = "RUNNING"
	InstanceStatusTypeProvisioning   InstanceStatusType = "PROVISIONING"
	InstanceStatusTypeDeprovisioning InstanceStatusType = "DEPROVISIONING"
	InstanceStatusTypePending        InstanceStatusType = "PENDING"
	InstanceStatusTypeFailed         InstanceStatusType = "FAILED"
)

func (e InstanceStatusType) IsValid() bool {
	switch e {
	case InstanceStatusTypeRunning, InstanceStatusTypeProvisioning, InstanceStatusTypeDeprovisioning, InstanceStatusTypePending, InstanceStatusTypeFailed:
		return true
	}
	return false
}

func (e InstanceStatusType) String() string {
	return string(e)
}

func (e *InstanceStatusType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = InstanceStatusType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid InstanceStatusType", str)
	}
	return nil
}

func (e InstanceStatusType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type LimitType string

const (
	LimitTypeContainer LimitType = "Container"
	LimitTypePod       LimitType = "Pod"
)

func (e LimitType) IsValid() bool {
	switch e {
	case LimitTypeContainer, LimitTypePod:
		return true
	}
	return false
}

func (e LimitType) String() string {
	return string(e)
}

func (e *LimitType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LimitType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LimitType", str)
	}
	return nil
}

func (e LimitType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PodStatusType string

const (
	PodStatusTypePending   PodStatusType = "PENDING"
	PodStatusTypeRunning   PodStatusType = "RUNNING"
	PodStatusTypeSucceeded PodStatusType = "SUCCEEDED"
	PodStatusTypeFailed    PodStatusType = "FAILED"
	PodStatusTypeUnknown   PodStatusType = "UNKNOWN"
)

func (e PodStatusType) IsValid() bool {
	switch e {
	case PodStatusTypePending, PodStatusTypeRunning, PodStatusTypeSucceeded, PodStatusTypeFailed, PodStatusTypeUnknown:
		return true
	}
	return false
}

func (e PodStatusType) String() string {
	return string(e)
}

func (e *PodStatusType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PodStatusType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PodStatusType", str)
	}
	return nil
}

func (e PodStatusType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ServiceBindingStatusType string

const (
	ServiceBindingStatusTypeReady   ServiceBindingStatusType = "READY"
	ServiceBindingStatusTypePending ServiceBindingStatusType = "PENDING"
	ServiceBindingStatusTypeFailed  ServiceBindingStatusType = "FAILED"
	ServiceBindingStatusTypeUnknown ServiceBindingStatusType = "UNKNOWN"
)

func (e ServiceBindingStatusType) IsValid() bool {
	switch e {
	case ServiceBindingStatusTypeReady, ServiceBindingStatusTypePending, ServiceBindingStatusTypeFailed, ServiceBindingStatusTypeUnknown:
		return true
	}
	return false
}

func (e ServiceBindingStatusType) String() string {
	return string(e)
}

func (e *ServiceBindingStatusType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ServiceBindingStatusType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ServiceBindingStatusType", str)
	}
	return nil
}

func (e ServiceBindingStatusType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ServiceBindingUsageStatusType string

const (
	ServiceBindingUsageStatusTypeReady   ServiceBindingUsageStatusType = "READY"
	ServiceBindingUsageStatusTypePending ServiceBindingUsageStatusType = "PENDING"
	ServiceBindingUsageStatusTypeFailed  ServiceBindingUsageStatusType = "FAILED"
	ServiceBindingUsageStatusTypeUnknown ServiceBindingUsageStatusType = "UNKNOWN"
)

func (e ServiceBindingUsageStatusType) IsValid() bool {
	switch e {
	case ServiceBindingUsageStatusTypeReady, ServiceBindingUsageStatusTypePending, ServiceBindingUsageStatusTypeFailed, ServiceBindingUsageStatusTypeUnknown:
		return true
	}
	return false
}

func (e ServiceBindingUsageStatusType) String() string {
	return string(e)
}

func (e *ServiceBindingUsageStatusType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ServiceBindingUsageStatusType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ServiceBindingUsageStatusType", str)
	}
	return nil
}

func (e ServiceBindingUsageStatusType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SubscriptionEventType string

const (
	SubscriptionEventTypeAdd    SubscriptionEventType = "ADD"
	SubscriptionEventTypeUpdate SubscriptionEventType = "UPDATE"
	SubscriptionEventTypeDelete SubscriptionEventType = "DELETE"
)

func (e SubscriptionEventType) IsValid() bool {
	switch e {
	case SubscriptionEventTypeAdd, SubscriptionEventTypeUpdate, SubscriptionEventTypeDelete:
		return true
	}
	return false
}

func (e SubscriptionEventType) String() string {
	return string(e)
}

func (e *SubscriptionEventType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SubscriptionEventType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SubscriptionEventType", str)
	}
	return nil
}

func (e SubscriptionEventType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

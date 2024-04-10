// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AlertpolicyAggregations struct {
	/* The alignment period for per-time
	series alignment. If present,
	alignmentPeriod must be at least
	60 seconds. After per-time series
	alignment, each time series will
	contain data points only on the
	period boundaries. If
	perSeriesAligner is not specified
	or equals ALIGN_NONE, then this
	field is ignored. If
	perSeriesAligner is specified and
	does not equal ALIGN_NONE, then
	this field must be defined;
	otherwise an error is returned. */
	// +optional
	AlignmentPeriod *string `json:"alignmentPeriod,omitempty"`

	/* The approach to be used to combine
	time series. Not all reducer
	functions may be applied to all
	time series, depending on the
	metric type and the value type of
	the original time series.
	Reduction may change the metric
	type of value type of the time
	series.Time series data must be
	aligned in order to perform cross-
	time series reduction. If
	crossSeriesReducer is specified,
	then perSeriesAligner must be
	specified and not equal ALIGN_NONE
	and alignmentPeriod must be
	specified; otherwise, an error is
	returned. Possible values: ["REDUCE_NONE", "REDUCE_MEAN", "REDUCE_MIN", "REDUCE_MAX", "REDUCE_SUM", "REDUCE_STDDEV", "REDUCE_COUNT", "REDUCE_COUNT_TRUE", "REDUCE_COUNT_FALSE", "REDUCE_FRACTION_TRUE", "REDUCE_PERCENTILE_99", "REDUCE_PERCENTILE_95", "REDUCE_PERCENTILE_50", "REDUCE_PERCENTILE_05"]. */
	// +optional
	CrossSeriesReducer *string `json:"crossSeriesReducer,omitempty"`

	/* The set of fields to preserve when
	crossSeriesReducer is specified.
	The groupByFields determine how
	the time series are partitioned
	into subsets prior to applying the
	aggregation function. Each subset
	contains time series that have the
	same value for each of the
	grouping fields. Each individual
	time series is a member of exactly
	one subset. The crossSeriesReducer
	is applied to each subset of time
	series. It is not possible to
	reduce across different resource
	types, so this field implicitly
	contains resource.type. Fields not
	specified in groupByFields are
	aggregated away. If groupByFields
	is not specified and all the time
	series have the same resource
	type, then the time series are
	aggregated into a single output
	time series. If crossSeriesReducer
	is not defined, this field is
	ignored. */
	// +optional
	GroupByFields []string `json:"groupByFields,omitempty"`

	/* The approach to be used to align
	individual time series. Not all
	alignment functions may be applied
	to all time series, depending on
	the metric type and value type of
	the original time series.
	Alignment may change the metric
	type or the value type of the time
	series.Time series data must be
	aligned in order to perform cross-
	time series reduction. If
	crossSeriesReducer is specified,
	then perSeriesAligner must be
	specified and not equal ALIGN_NONE
	and alignmentPeriod must be
	specified; otherwise, an error is
	returned. Possible values: ["ALIGN_NONE", "ALIGN_DELTA", "ALIGN_RATE", "ALIGN_INTERPOLATE", "ALIGN_NEXT_OLDER", "ALIGN_MIN", "ALIGN_MAX", "ALIGN_MEAN", "ALIGN_COUNT", "ALIGN_SUM", "ALIGN_STDDEV", "ALIGN_COUNT_TRUE", "ALIGN_COUNT_FALSE", "ALIGN_FRACTION_TRUE", "ALIGN_PERCENTILE_99", "ALIGN_PERCENTILE_95", "ALIGN_PERCENTILE_50", "ALIGN_PERCENTILE_05", "ALIGN_PERCENT_CHANGE"]. */
	// +optional
	PerSeriesAligner *string `json:"perSeriesAligner,omitempty"`
}

type AlertpolicyAlertStrategy struct {
	/* If an alert policy that was active has no data for this long, any open incidents will close. */
	// +optional
	AutoClose *string `json:"autoClose,omitempty"`

	/* Control over how the notification channels in 'notification_channels'
	are notified when this alert fires, on a per-channel basis. */
	// +optional
	NotificationChannelStrategy []AlertpolicyNotificationChannelStrategy `json:"notificationChannelStrategy,omitempty"`

	/* Required for alert policies with a LogMatch condition.
	This limit is not implemented for alert policies that are not log-based. */
	// +optional
	NotificationRateLimit *AlertpolicyNotificationRateLimit `json:"notificationRateLimit,omitempty"`
}

type AlertpolicyConditionAbsent struct {
	/* Specifies the alignment of data points in
	individual time series as well as how to
	combine the retrieved time series together
	(such as when aggregating multiple streams
	on each resource to a single stream for each
	resource or when aggregating streams across
	all members of a group of resources).
	Multiple aggregations are applied in the
	order specified. */
	// +optional
	Aggregations []AlertpolicyAggregations `json:"aggregations,omitempty"`

	/* The amount of time that a time series must
	fail to report new data to be considered
	failing. Currently, only values that are a
	multiple of a minute--e.g. 60s, 120s, or 300s
	--are supported. */
	Duration string `json:"duration"`

	/* A filter that identifies which time series
	should be compared with the threshold.The
	filter is similar to the one that is
	specified in the
	MetricService.ListTimeSeries request (that
	call is useful to verify the time series
	that will be retrieved / processed) and must
	specify the metric type and optionally may
	contain restrictions on resource type,
	resource labels, and metric labels. This
	field may not exceed 2048 Unicode characters
	in length. */
	// +optional
	Filter *string `json:"filter,omitempty"`

	/* The number/percent of time series for which
	the comparison must hold in order for the
	condition to trigger. If unspecified, then
	the condition will trigger if the comparison
	is true for any of the time series that have
	been identified by filter and aggregations. */
	// +optional
	Trigger *AlertpolicyTrigger `json:"trigger,omitempty"`
}

type AlertpolicyConditionMatchedLog struct {
	/* A logs-based filter. */
	Filter string `json:"filter"`

	/* A map from a label key to an extractor expression, which is used to
	extract the value for this label key. Each entry in this map is
	a specification for how data should be extracted from log entries that
	match filter. Each combination of extracted values is treated as
	a separate rule for the purposes of triggering notifications.
	Label keys and corresponding values can be used in notifications
	generated by this condition. */
	// +optional
	LabelExtractors map[string]string `json:"labelExtractors,omitempty"`
}

type AlertpolicyConditionMonitoringQueryLanguage struct {
	/* The amount of time that a time series must
	violate the threshold to be considered
	failing. Currently, only values that are a
	multiple of a minute--e.g., 0, 60, 120, or
	300 seconds--are supported. If an invalid
	value is given, an error will be returned.
	When choosing a duration, it is useful to
	keep in mind the frequency of the underlying
	time series data (which may also be affected
	by any alignments specified in the
	aggregations field); a good duration is long
	enough so that a single outlier does not
	generate spurious alerts, but short enough
	that unhealthy states are detected and
	alerted on quickly. */
	Duration string `json:"duration"`

	/* A condition control that determines how
	metric-threshold conditions are evaluated when
	data stops arriving. Possible values: ["EVALUATION_MISSING_DATA_INACTIVE", "EVALUATION_MISSING_DATA_ACTIVE", "EVALUATION_MISSING_DATA_NO_OP"]. */
	// +optional
	EvaluationMissingData *string `json:"evaluationMissingData,omitempty"`

	/* Monitoring Query Language query that outputs a boolean stream. */
	Query string `json:"query"`

	/* The number/percent of time series for which
	the comparison must hold in order for the
	condition to trigger. If unspecified, then
	the condition will trigger if the comparison
	is true for any of the time series that have
	been identified by filter and aggregations,
	or by the ratio, if denominator_filter and
	denominator_aggregations are specified. */
	// +optional
	Trigger *AlertpolicyTrigger `json:"trigger,omitempty"`
}

type AlertpolicyConditionPrometheusQueryLanguage struct {
	/* The alerting rule name of this alert in the corresponding Prometheus
	configuration file.

	Some external tools may require this field to be populated correctly
	in order to refer to the original Prometheus configuration file.
	The rule group name and the alert name are necessary to update the
	relevant AlertPolicies in case the definition of the rule group changes
	in the future.

	This field is optional. If this field is not empty, then it must be a
	valid Prometheus label name. */
	// +optional
	AlertRule *string `json:"alertRule,omitempty"`

	/* Alerts are considered firing once their PromQL expression evaluated
	to be "true" for this long. Alerts whose PromQL expression was not
	evaluated to be "true" for long enough are considered pending. The
	default value is zero. Must be zero or positive. */
	// +optional
	Duration *string `json:"duration,omitempty"`

	/* How often this rule should be evaluated. Must be a positive multiple
	of 30 seconds or missing. The default value is 30 seconds. If this
	PrometheusQueryLanguageCondition was generated from a Prometheus
	alerting rule, then this value should be taken from the enclosing
	rule group. */
	// +optional
	EvaluationInterval *string `json:"evaluationInterval,omitempty"`

	/* Labels to add to or overwrite in the PromQL query result. Label names
	must be valid.

	Label values can be templatized by using variables. The only available
	variable names are the names of the labels in the PromQL result, including
	"__name__" and "value". "labels" may be empty. This field is intended to be
	used for organizing and identifying the AlertPolicy. */
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	/* The PromQL expression to evaluate. Every evaluation cycle this
	expression is evaluated at the current time, and all resultant time
	series become pending/firing alerts. This field must not be empty. */
	Query string `json:"query"`

	/* The rule group name of this alert in the corresponding Prometheus
	configuration file.

	Some external tools may require this field to be populated correctly
	in order to refer to the original Prometheus configuration file.
	The rule group name and the alert name are necessary to update the
	relevant AlertPolicies in case the definition of the rule group changes
	in the future.

	This field is optional. If this field is not empty, then it must be a
	valid Prometheus label name. */
	// +optional
	RuleGroup *string `json:"ruleGroup,omitempty"`
}

type AlertpolicyConditionThreshold struct {
	/* Specifies the alignment of data points in
	individual time series as well as how to
	combine the retrieved time series together
	(such as when aggregating multiple streams
	on each resource to a single stream for each
	resource or when aggregating streams across
	all members of a group of resources).
	Multiple aggregations are applied in the
	order specified.This field is similar to the
	one in the MetricService.ListTimeSeries
	request. It is advisable to use the
	ListTimeSeries method when debugging this
	field. */
	// +optional
	Aggregations []AlertpolicyAggregations `json:"aggregations,omitempty"`

	/* The comparison to apply between the time
	series (indicated by filter and aggregation)
	and the threshold (indicated by
	threshold_value). The comparison is applied
	on each time series, with the time series on
	the left-hand side and the threshold on the
	right-hand side. Only COMPARISON_LT and
	COMPARISON_GT are supported currently. Possible values: ["COMPARISON_GT", "COMPARISON_GE", "COMPARISON_LT", "COMPARISON_LE", "COMPARISON_EQ", "COMPARISON_NE"]. */
	Comparison string `json:"comparison"`

	/* Specifies the alignment of data points in
	individual time series selected by
	denominatorFilter as well as how to combine
	the retrieved time series together (such as
	when aggregating multiple streams on each
	resource to a single stream for each
	resource or when aggregating streams across
	all members of a group of resources).When
	computing ratios, the aggregations and
	denominator_aggregations fields must use the
	same alignment period and produce time
	series that have the same periodicity and
	labels.This field is similar to the one in
	the MetricService.ListTimeSeries request. It
	is advisable to use the ListTimeSeries
	method when debugging this field. */
	// +optional
	DenominatorAggregations []AlertpolicyDenominatorAggregations `json:"denominatorAggregations,omitempty"`

	/* A filter that identifies a time series that
	should be used as the denominator of a ratio
	that will be compared with the threshold. If
	a denominator_filter is specified, the time
	series specified by the filter field will be
	used as the numerator.The filter is similar
	to the one that is specified in the
	MetricService.ListTimeSeries request (that
	call is useful to verify the time series
	that will be retrieved / processed) and must
	specify the metric type and optionally may
	contain restrictions on resource type,
	resource labels, and metric labels. This
	field may not exceed 2048 Unicode characters
	in length. */
	// +optional
	DenominatorFilter *string `json:"denominatorFilter,omitempty"`

	/* The amount of time that a time series must
	violate the threshold to be considered
	failing. Currently, only values that are a
	multiple of a minute--e.g., 0, 60, 120, or
	300 seconds--are supported. If an invalid
	value is given, an error will be returned.
	When choosing a duration, it is useful to
	keep in mind the frequency of the underlying
	time series data (which may also be affected
	by any alignments specified in the
	aggregations field); a good duration is long
	enough so that a single outlier does not
	generate spurious alerts, but short enough
	that unhealthy states are detected and
	alerted on quickly. */
	Duration string `json:"duration"`

	/* A condition control that determines how
	metric-threshold conditions are evaluated when
	data stops arriving. Possible values: ["EVALUATION_MISSING_DATA_INACTIVE", "EVALUATION_MISSING_DATA_ACTIVE", "EVALUATION_MISSING_DATA_NO_OP"]. */
	// +optional
	EvaluationMissingData *string `json:"evaluationMissingData,omitempty"`

	/* A filter that identifies which time series
	should be compared with the threshold.The
	filter is similar to the one that is
	specified in the
	MetricService.ListTimeSeries request (that
	call is useful to verify the time series
	that will be retrieved / processed) and must
	specify the metric type and optionally may
	contain restrictions on resource type,
	resource labels, and metric labels. This
	field may not exceed 2048 Unicode characters
	in length. */
	// +optional
	Filter *string `json:"filter,omitempty"`

	/* When this field is present, the 'MetricThreshold'
	condition forecasts whether the time series is
	predicted to violate the threshold within the
	'forecastHorizon'. When this field is not set, the
	'MetricThreshold' tests the current value of the
	timeseries against the threshold. */
	// +optional
	ForecastOptions *AlertpolicyForecastOptions `json:"forecastOptions,omitempty"`

	/* A value against which to compare the time
	series. */
	// +optional
	ThresholdValue *float64 `json:"thresholdValue,omitempty"`

	/* The number/percent of time series for which
	the comparison must hold in order for the
	condition to trigger. If unspecified, then
	the condition will trigger if the comparison
	is true for any of the time series that have
	been identified by filter and aggregations,
	or by the ratio, if denominator_filter and
	denominator_aggregations are specified. */
	// +optional
	Trigger *AlertpolicyTrigger `json:"trigger,omitempty"`
}

type AlertpolicyConditions struct {
	/* A condition that checks that a time series
	continues to receive new data points. */
	// +optional
	ConditionAbsent *AlertpolicyConditionAbsent `json:"conditionAbsent,omitempty"`

	/* A condition that checks for log messages matching given constraints.
	If set, no other conditions can be present. */
	// +optional
	ConditionMatchedLog *AlertpolicyConditionMatchedLog `json:"conditionMatchedLog,omitempty"`

	/* A Monitoring Query Language query that outputs a boolean stream. */
	// +optional
	ConditionMonitoringQueryLanguage *AlertpolicyConditionMonitoringQueryLanguage `json:"conditionMonitoringQueryLanguage,omitempty"`

	/* A Monitoring Query Language query that outputs a boolean stream

	A condition type that allows alert policies to be defined using
	Prometheus Query Language (PromQL).

	The PrometheusQueryLanguageCondition message contains information
	from a Prometheus alerting rule and its associated rule group. */
	// +optional
	ConditionPrometheusQueryLanguage *AlertpolicyConditionPrometheusQueryLanguage `json:"conditionPrometheusQueryLanguage,omitempty"`

	/* A condition that compares a time series against a
	threshold. */
	// +optional
	ConditionThreshold *AlertpolicyConditionThreshold `json:"conditionThreshold,omitempty"`

	/* A short name or phrase used to identify the
	condition in dashboards, notifications, and
	incidents. To avoid confusion, don't use the same
	display name for multiple conditions in the same
	policy. */
	DisplayName string `json:"displayName"`

	/* The unique resource name for this condition.
	Its syntax is:
	projects/[PROJECT_ID]/alertPolicies/[POLICY_ID]/conditions/[CONDITION_ID]
	[CONDITION_ID] is assigned by Stackdriver Monitoring when
	the condition is created as part of a new or updated alerting
	policy. */
	// +optional
	Name *string `json:"name,omitempty"`
}

type AlertpolicyDenominatorAggregations struct {
	/* The alignment period for per-time
	series alignment. If present,
	alignmentPeriod must be at least
	60 seconds. After per-time series
	alignment, each time series will
	contain data points only on the
	period boundaries. If
	perSeriesAligner is not specified
	or equals ALIGN_NONE, then this
	field is ignored. If
	perSeriesAligner is specified and
	does not equal ALIGN_NONE, then
	this field must be defined;
	otherwise an error is returned. */
	// +optional
	AlignmentPeriod *string `json:"alignmentPeriod,omitempty"`

	/* The approach to be used to combine
	time series. Not all reducer
	functions may be applied to all
	time series, depending on the
	metric type and the value type of
	the original time series.
	Reduction may change the metric
	type of value type of the time
	series.Time series data must be
	aligned in order to perform cross-
	time series reduction. If
	crossSeriesReducer is specified,
	then perSeriesAligner must be
	specified and not equal ALIGN_NONE
	and alignmentPeriod must be
	specified; otherwise, an error is
	returned. Possible values: ["REDUCE_NONE", "REDUCE_MEAN", "REDUCE_MIN", "REDUCE_MAX", "REDUCE_SUM", "REDUCE_STDDEV", "REDUCE_COUNT", "REDUCE_COUNT_TRUE", "REDUCE_COUNT_FALSE", "REDUCE_FRACTION_TRUE", "REDUCE_PERCENTILE_99", "REDUCE_PERCENTILE_95", "REDUCE_PERCENTILE_50", "REDUCE_PERCENTILE_05"]. */
	// +optional
	CrossSeriesReducer *string `json:"crossSeriesReducer,omitempty"`

	/* The set of fields to preserve when
	crossSeriesReducer is specified.
	The groupByFields determine how
	the time series are partitioned
	into subsets prior to applying the
	aggregation function. Each subset
	contains time series that have the
	same value for each of the
	grouping fields. Each individual
	time series is a member of exactly
	one subset. The crossSeriesReducer
	is applied to each subset of time
	series. It is not possible to
	reduce across different resource
	types, so this field implicitly
	contains resource.type. Fields not
	specified in groupByFields are
	aggregated away. If groupByFields
	is not specified and all the time
	series have the same resource
	type, then the time series are
	aggregated into a single output
	time series. If crossSeriesReducer
	is not defined, this field is
	ignored. */
	// +optional
	GroupByFields []string `json:"groupByFields,omitempty"`

	/* The approach to be used to align
	individual time series. Not all
	alignment functions may be applied
	to all time series, depending on
	the metric type and value type of
	the original time series.
	Alignment may change the metric
	type or the value type of the time
	series.Time series data must be
	aligned in order to perform cross-
	time series reduction. If
	crossSeriesReducer is specified,
	then perSeriesAligner must be
	specified and not equal ALIGN_NONE
	and alignmentPeriod must be
	specified; otherwise, an error is
	returned. Possible values: ["ALIGN_NONE", "ALIGN_DELTA", "ALIGN_RATE", "ALIGN_INTERPOLATE", "ALIGN_NEXT_OLDER", "ALIGN_MIN", "ALIGN_MAX", "ALIGN_MEAN", "ALIGN_COUNT", "ALIGN_SUM", "ALIGN_STDDEV", "ALIGN_COUNT_TRUE", "ALIGN_COUNT_FALSE", "ALIGN_FRACTION_TRUE", "ALIGN_PERCENTILE_99", "ALIGN_PERCENTILE_95", "ALIGN_PERCENTILE_50", "ALIGN_PERCENTILE_05", "ALIGN_PERCENT_CHANGE"]. */
	// +optional
	PerSeriesAligner *string `json:"perSeriesAligner,omitempty"`
}

type AlertpolicyDocumentation struct {
	/* The text of the documentation, interpreted according to mimeType.
	The content may not exceed 8,192 Unicode characters and may not
	exceed more than 10,240 bytes when encoded in UTF-8 format,
	whichever is smaller. */
	// +optional
	Content *string `json:"content,omitempty"`

	/* The format of the content field. Presently, only the value
	"text/markdown" is supported. */
	// +optional
	MimeType *string `json:"mimeType,omitempty"`
}

type AlertpolicyForecastOptions struct {
	/* The length of time into the future to forecast
	whether a timeseries will violate the threshold.
	If the predicted value is found to violate the
	threshold, and the violation is observed in all
	forecasts made for the Configured 'duration',
	then the timeseries is considered to be failing. */
	ForecastHorizon string `json:"forecastHorizon"`
}

type AlertpolicyNotificationChannelStrategy struct {
	/* The notification channels that these settings apply to. Each of these
	correspond to the name field in one of the NotificationChannel objects
	referenced in the notification_channels field of this AlertPolicy. The format is
	'projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID]'. */
	// +optional
	NotificationChannelNames []string `json:"notificationChannelNames,omitempty"`

	/* The frequency at which to send reminder notifications for open incidents. */
	// +optional
	RenotifyInterval *string `json:"renotifyInterval,omitempty"`
}

type AlertpolicyNotificationRateLimit struct {
	/* Not more than one notification per period. */
	// +optional
	Period *string `json:"period,omitempty"`
}

type AlertpolicyTrigger struct {
	/* The absolute number of time series
	that must fail the predicate for the
	condition to be triggered. */
	// +optional
	Count *int `json:"count,omitempty"`

	/* The percentage of time series that
	must fail the predicate for the
	condition to be triggered. */
	// +optional
	Percent *float64 `json:"percent,omitempty"`
}

type MonitoringAlertPolicySpec struct {
	/* Control over how this alert policy's notification channels are notified. */
	// +optional
	AlertStrategy *AlertpolicyAlertStrategy `json:"alertStrategy,omitempty"`

	/* How to combine the results of multiple conditions to
	determine if an incident should be opened. Possible values: ["AND", "OR", "AND_WITH_MATCHING_RESOURCE"]. */
	Combiner string `json:"combiner"`

	/* A list of conditions for the policy. The conditions are combined by
	AND or OR according to the combiner field. If the combined conditions
	evaluate to true, then an incident is created. A policy can have from
	one to six conditions. */
	Conditions []AlertpolicyConditions `json:"conditions"`

	/* A short name or phrase used to identify the policy in
	dashboards, notifications, and incidents. To avoid confusion, don't use
	the same display name for multiple policies in the same project. The
	name is limited to 512 Unicode characters. */
	DisplayName string `json:"displayName"`

	/* Documentation that is included with notifications and incidents related
	to this policy. Best practice is for the documentation to include information
	to help responders understand, mitigate, escalate, and correct the underlying
	problems detected by the alerting policy. Notification channels that have
	limited capacity might not show this documentation. */
	// +optional
	Documentation *AlertpolicyDocumentation `json:"documentation,omitempty"`

	/* Whether or not the policy is enabled. The default is true. */
	// +optional
	Enabled *bool `json:"enabled,omitempty"`

	// +optional
	NotificationChannels []v1alpha1.ResourceRef `json:"notificationChannels,omitempty"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type AlertpolicyCreationRecordStatus struct {
	/* When the change occurred. */
	// +optional
	MutateTime *string `json:"mutateTime,omitempty"`

	/* The email address of the user making the change. */
	// +optional
	MutatedBy *string `json:"mutatedBy,omitempty"`
}

type MonitoringAlertPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   MonitoringAlertPolicy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* A read-only record of the creation of the alerting policy.
	If provided in a call to create or update, this field will
	be ignored. */
	// +optional
	CreationRecord []AlertpolicyCreationRecordStatus `json:"creationRecord,omitempty"`

	/* The unique resource name for this policy.
	Its syntax is: projects/[PROJECT_ID]/alertPolicies/[ALERT_POLICY_ID]. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmonitoringalertpolicy;gcpmonitoringalertpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MonitoringAlertPolicy is the Schema for the monitoring API
// +k8s:openapi-gen=true
type MonitoringAlertPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitoringAlertPolicySpec   `json:"spec,omitempty"`
	Status MonitoringAlertPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MonitoringAlertPolicyList contains a list of MonitoringAlertPolicy
type MonitoringAlertPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringAlertPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MonitoringAlertPolicy{}, &MonitoringAlertPolicyList{})
}

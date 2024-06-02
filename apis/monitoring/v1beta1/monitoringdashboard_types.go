// Copyright 2024 Google LLC
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

package v1beta1

import (
	"reflect"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// SchemeBuilder is used to add go types to the GroupVersionKind scheme.
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// AddToScheme is a global function that registers this API group & version to a scheme
	AddToScheme = SchemeBuilder.AddToScheme

	MonitoringDashboardGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(MonitoringDashboard{}).Name(),
	}
)

// +kcc:proto=google.monitoring.dashboard.v1.AlertChart
type AlertChart struct {
	// Required. The resource name of the alert policy. The format is:
	//
	//      projects/[PROJECT_ID_OR_NUMBER]/alertPolicies/[ALERT_POLICY_ID]
	//
	// +required
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.ChartOptions
type ChartOptions struct {
	// The chart mode.
	Mode *string `json:"mode,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.XyChart
type XyChart struct {
	// Required. The data displayed in this chart.
	// +required
	DataSets []XyChart_DataSet `json:"dataSets,omitempty"`

	// The duration used to display a comparison chart. A comparison chart
	//  simultaneously shows values from two similar-length time periods
	//  (e.g., week-over-week metrics).
	//  The duration must be positive, and it can only be applied to charts with
	//  data sets of LINE plot type.
	TimeshiftDuration *string `json:"timeshiftDuration,omitempty"`

	// Threshold lines drawn horizontally across the chart.
	Thresholds []Threshold `json:"thresholds,omitempty"`

	// The properties applied to the x-axis.
	XAxis *XyChart_Axis `json:"xAxis,omitempty"`

	// The properties applied to the y-axis.
	YAxis *XyChart_Axis `json:"yAxis,omitempty"`

	/*NOTYET
	// The properties applied to the y2-axis.
	Y2Axis *XyChart_Axis `json:"y2Axis,omitempty"`
	*/

	// Display options for the chart.
	ChartOptions *ChartOptions `json:"chartOptions,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.XyChart.DataSet
type XyChart_DataSet struct {
	// Required. Fields for querying time series data from the
	//  Stackdriver metrics API.
	//
	// +required
	TimeSeriesQuery *TimeSeriesQuery `json:"timeSeriesQuery,omitempty"`

	// How this data should be plotted on the chart.
	PlotType *string `json:"plotType,omitempty"`

	// A template string for naming `TimeSeries` in the resulting data set.
	//  This should be a string with interpolations of the form `${label_name}`,
	//  which will resolve to the label's value.
	LegendTemplate *string `json:"legendTemplate,omitempty"`

	// Optional. The lower bound on data point frequency for this data set,
	//  implemented by specifying the minimum alignment period to use in a time
	//  series query For example, if the data is published once every 10 minutes,
	//  the `min_alignment_period` should be at least 10 minutes. It would not
	//  make sense to fetch and align data at one minute intervals.
	MinAlignmentPeriod *string `json:"minAlignmentPeriod,omitempty"`

	/*NOTYET
	// Optional. The target axis to use for plotting the metric.
	TargetAxis *string `json:"targetAxis,omitempty"`
	*/
}

// +kcc:proto=google.monitoring.dashboard.v1.XyChart.Axis
type XyChart_Axis struct {
	// The label of the axis.
	Label *string `json:"label,omitempty"`

	// The axis scale. By default, a linear scale is used.
	Scale *string `json:"scale,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.ColumnLayout
type ColumnLayout struct {
	// The columns of content to display.
	Columns []ColumnLayout_Column `json:"columns,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.ColumnLayout.Column
type ColumnLayout_Column struct {
	// The relative weight of this column. The column weight is used to adjust
	//  the width of columns on the screen (relative to peers).
	//  Greater the weight, greater the width of the column on the screen.
	//  If omitted, a value of 1 is used while rendering.
	Weight *int64 `json:"weight,omitempty"`

	// The display widgets arranged vertically in this column.
	Widgets []Widget `json:"widgets,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.GridLayout
type GridLayout struct {
	// The number of columns into which the view's width is divided. If omitted
	//  or set to zero, a system default will be used while rendering.
	Columns *int64 `json:"columns,omitempty"`

	// The informational elements that are arranged into the columns row-first.
	Widgets []Widget `json:"widgets,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.MosaicLayout
type MosaicLayout struct {
	// The number of columns in the mosaic grid. The number of columns must be
	//  between 1 and 12, inclusive.
	Columns *int32 `json:"columns,omitempty"`

	// The tiles to display.
	Tiles []MosaicLayout_Tile `json:"tiles,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.MosaicLayout.Tile
type MosaicLayout_Tile struct {
	// The zero-indexed position of the tile in grid blocks relative to the
	//  left edge of the grid. Tiles must be contained within the specified
	//  number of columns. `x_pos` cannot be negative.
	XPos *int32 `json:"xPos,omitempty"`

	// The zero-indexed position of the tile in grid blocks relative to the
	//  top edge of the grid. `y_pos` cannot be negative.
	YPos *int32 `json:"yPos,omitempty"`

	// The width of the tile, measured in grid blocks. Tiles must have a
	//  minimum width of 1.
	Width *int32 `json:"width,omitempty"`

	// The height of the tile, measured in grid blocks. Tiles must have a
	//  minimum height of 1.
	Height *int32 `json:"height,omitempty"`

	// The informational widget contained in the tile. For example an `XyChart`.
	Widget *Widget `json:"widget,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.RowLayout
type RowLayout struct {
	// The rows of content to display.
	Rows []RowLayout_Row `json:"rows,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.RowLayout.Row
type RowLayout_Row struct {
	// The relative weight of this row. The row weight is used to adjust the
	//  height of rows on the screen (relative to peers). Greater the weight,
	//  greater the height of the row on the screen. If omitted, a value
	//  of 1 is used while rendering.
	Weight *int64 `json:"weight,omitempty"`

	// The display widgets arranged horizontally in this row.
	Widgets []Widget `json:"widgets,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.TimeSeriesTable
type TimeSeriesTable struct {
	// Required. The data displayed in this table.
	//
	// +required
	DataSets []TimeSeriesTable_TableDataSet `json:"dataSets,omitempty"`

	/*NOTYET
	// Optional. Store rendering strategy
	MetricVisualization *string `json:"metricVisualization,omitempty"`
	*/

	// Optional. The list of the persistent column settings for the table.
	ColumnSettings []TimeSeriesTable_ColumnSettings `json:"columnSettings,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.TimeSeriesTable.TableDataSet
type TimeSeriesTable_TableDataSet struct {
	// Required. Fields for querying time series data from the
	//  Stackdriver metrics API.
	TimeSeriesQuery *TimeSeriesQuery `json:"timeSeriesQuery,omitempty"`

	// Optional. A template string for naming `TimeSeries` in the resulting data
	//  set. This should be a string with interpolations of the form
	//  `${label_name}`, which will resolve to the label's value i.e.
	//  "${resource.labels.project_id}."
	TableTemplate *string `json:"tableTemplate,omitempty"`

	// Optional. The lower bound on data point frequency for this data set,
	//  implemented by specifying the minimum alignment period to use in a time
	//  series query For example, if the data is published once every 10 minutes,
	//  the `min_alignment_period` should be at least 10 minutes. It would not
	//  make sense to fetch and align data at one minute intervals.
	MinAlignmentPeriod *string `json:"minAlignmentPeriod,omitempty"`

	// Optional. Table display options for configuring how the table is
	//  rendered.
	TableDisplayOptions *TableDisplayOptions `json:"tableDisplayOptions,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.TimeSeriesTable.ColumnSettings
type TimeSeriesTable_ColumnSettings struct {
	// Required. The id of the column.
	//
	// +required
	Column *string `json:"column,omitempty"`

	// Required. Whether the column should be visible on page load.
	//
	// +required
	Visible *bool `json:"visible,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.Text
type Text struct {
	// The text content to be displayed.
	Content *string `json:"content,omitempty"`

	// How the text content is formatted.
	Format *string `json:"format,omitempty"`

	/*NOTYET
	// How the text is styled
	Style *Text_TextStyle `json:"style,omitempty"`
	*/
}

// +kcc:proto=google.monitoring.dashboard.v1.Text.TextStyle
type Text_TextStyle struct {
	// The background color as a hex string. "#RRGGBB" or "#RGB"
	BackgroundColor *string `json:"backgroundColor,omitempty"`

	// The text color as a hex string. "#RRGGBB" or "#RGB"
	TextColor *string `json:"textColor,omitempty"`

	// The horizontal alignment of both the title and content
	HorizontalAlignment *string `json:"horizontalAlignment,omitempty"`

	// The vertical alignment of both the title and content
	VerticalAlignment *string `json:"verticalAlignment,omitempty"`

	// The amount of padding around the widget
	Padding *string `json:"padding,omitempty"`

	// Font sizes for both the title and content. The title will still be larger
	//  relative to the content.
	FontSize *string `json:"fontSize,omitempty"`

	// The pointer location for this widget (also sometimes called a "tail")
	PointerLocation *string `json:"pointerLocation,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.CollapsibleGroup
type CollapsibleGroup struct {
	// The collapsed state of the widget on first page load.
	Collapsed *bool `json:"collapsed,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.ErrorReportingPanel
type ErrorReportingPanel struct {
	// The resource name of the Google Cloud Platform project. Written
	//  as `projects/{projectID}` or `projects/{projectNumber}`, where
	//  `{projectID}` and `{projectNumber}` can be found in the
	//  [Google Cloud console](https://support.google.com/cloud/answer/6158840).
	//
	//  Examples: `projects/my-project-123`, `projects/5551234`.
	ProjectNames []string `json:"projectNames,omitempty"`

	// An identifier of the service, such as the name of the
	//  executable, job, or Google App Engine service name. This field is expected
	//  to have a low number of values that are relatively stable over time, as
	//  opposed to `version`, which can be changed whenever new code is deployed.
	//
	//  Contains the service name for error reports extracted from Google
	//  App Engine logs or `default` if the App Engine default service is used.
	Services []string `json:"services,omitempty"`

	// Represents the source code version that the developer provided,
	//  which could represent a version label or a Git SHA-1 hash, for example.
	//  For App Engine standard environment, the version is set to the version of
	//  the app.
	Versions []string `json:"versions,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.SingleViewGroup
type SingleViewGroup struct {
}

// +kcc:proto=google.monitoring.dashboard.v1.SectionHeader
type SectionHeader struct {
	// The subtitle of the section
	Subtitle *string `json:"subtitle,omitempty"`

	// Whether to insert a divider below the section in the table of contents
	DividerBelow *bool `json:"dividerBelow,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.Aggregation
type Aggregation struct {
	// The `alignment_period` specifies a time interval, in seconds, that is used
	//  to divide the data in all the
	//  [time series][google.monitoring.v3.TimeSeries] into consistent blocks of
	//  time. This will be done before the per-series aligner can be applied to
	//  the data.
	//
	//  The value must be at least 60 seconds. If a per-series aligner other than
	//  `ALIGN_NONE` is specified, this field is required or an error is returned.
	//  If no per-series aligner is specified, or the aligner `ALIGN_NONE` is
	//  specified, then this field is ignored.
	//
	//  The maximum value of the `alignment_period` is 2 years, or 104 weeks.
	AlignmentPeriod *string `json:"alignmentPeriod,omitempty"`

	// An `Aligner` describes how to bring the data points in a single
	//  time series into temporal alignment. Except for `ALIGN_NONE`, all
	//  alignments cause all the data points in an `alignment_period` to be
	//  mathematically grouped together, resulting in a single data point for
	//  each `alignment_period` with end timestamp at the end of the period.
	//
	//  Not all alignment operations may be applied to all time series. The valid
	//  choices depend on the `metric_kind` and `value_type` of the original time
	//  series. Alignment can change the `metric_kind` or the `value_type` of
	//  the time series.
	//
	//  Time series data must be aligned in order to perform cross-time
	//  series reduction. If `cross_series_reducer` is specified, then
	//  `per_series_aligner` must be specified and not equal to `ALIGN_NONE`
	//  and `alignment_period` must be specified; otherwise, an error is
	//  returned.
	PerSeriesAligner *string `json:"perSeriesAligner,omitempty"`

	// The reduction operation to be used to combine time series into a single
	//  time series, where the value of each data point in the resulting series is
	//  a function of all the already aligned values in the input time series.
	//
	//  Not all reducer operations can be applied to all time series. The valid
	//  choices depend on the `metric_kind` and the `value_type` of the original
	//  time series. Reduction can yield a time series with a different
	//  `metric_kind` or `value_type` than the input time series.
	//
	//  Time series data must first be aligned (see `per_series_aligner`) in order
	//  to perform cross-time series reduction. If `cross_series_reducer` is
	//  specified, then `per_series_aligner` must be specified, and must not be
	//  `ALIGN_NONE`. An `alignment_period` must also be specified; otherwise, an
	//  error is returned.
	CrossSeriesReducer *string `json:"crossSeriesReducer,omitempty"`

	// The set of fields to preserve when `cross_series_reducer` is
	//  specified. The `group_by_fields` determine how the time series are
	//  partitioned into subsets prior to applying the aggregation
	//  operation. Each subset contains time series that have the same
	//  value for each of the grouping fields. Each individual time
	//  series is a member of exactly one subset. The
	//  `cross_series_reducer` is applied to each subset of time series.
	//  It is not possible to reduce across different resource types, so
	//  this field implicitly contains `resource.type`.  Fields not
	//  specified in `group_by_fields` are aggregated away.  If
	//  `group_by_fields` is not specified and all the time series have
	//  the same resource type, then the time series are aggregated into
	//  a single output time series. If `cross_series_reducer` is not
	//  defined, this field is ignored.
	GroupByFields []string `json:"groupByFields,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.PickTimeSeriesFilter
type PickTimeSeriesFilter struct {
	// `ranking_method` is applied to each time series independently to produce
	//  the value which will be used to compare the time series to other time
	//  series.
	RankingMethod *string `json:"rankingMethod,omitempty"`

	// How many time series to allow to pass through the filter.
	NumTimeSeries *int32 `json:"numTimeSeries,omitempty"`

	// How to use the ranking to select time series that pass through the filter.
	Direction *string `json:"direction,omitempty"`

	/*NOTYET
	// Select the top N streams/time series within this time interval
	Interval *string `json:"interval,omitempty"`
	*/
}

// +kcc:proto=google.monitoring.dashboard.v1.StatisticalTimeSeriesFilter
type StatisticalTimeSeriesFilter struct {
	// `rankingMethod` is applied to a set of time series, and then the produced
	//  value for each individual time series is used to compare a given time
	//  series to others.
	//  These are methods that cannot be applied stream-by-stream, but rather
	//  require the full context of a request to evaluate time series.
	RankingMethod *string `json:"rankingMethod,omitempty"`

	// How many time series to output.
	NumTimeSeries *int32 `json:"numTimeSeries,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.Widget
type Widget struct {
	// Optional. The title of the widget.
	Title *string `json:"title,omitempty"`

	// A chart of time series data.
	XyChart *XyChart `json:"xyChart,omitempty"`

	// A scorecard summarizing time series data.
	Scorecard *Scorecard `json:"scorecard,omitempty"`

	// A raw string or markdown displaying textual content.
	Text *Text `json:"text,omitempty"`

	// A blank space.
	Blank *Empty `json:"blank,omitempty"`

	/*NOTYET
	// A chart of alert policy data.
	AlertChart *AlertChart `json:"alertChart,omitempty"`

	// A widget that displays time series data in a tabular format.
	TimeSeriesTable *TimeSeriesTable `json:"timeSeriesTable,omitempty"`

	// A widget that groups the other widgets. All widgets that are within
	//  the area spanned by the grouping widget are considered member widgets.
	CollapsibleGroup *CollapsibleGroup `json:"collapsibleGroup,omitempty"`
	*/

	// A widget that shows a stream of logs.
	LogsPanel *LogsPanel `json:"logsPanel,omitempty"`

	/*NOTYET
	// A widget that shows list of incidents.
	IncidentList *IncidentList `json:"incidentList,omitempty"`

	// A widget that displays timeseries data as a pie chart.
	PieChart *PieChart `json:"pieChart,omitempty"`

	// A widget that displays a list of error groups.
	ErrorReportingPanel *ErrorReportingPanel `json:"errorReportingPanel,omitempty"`

	// A widget that defines a section header for easier navigation of the
	//  dashboard.
	SectionHeader *SectionHeader `json:"sectionHeader,omitempty"`

	// A widget that groups the other widgets by using a dropdown menu.
	SingleViewGroup *SingleViewGroup `json:"singleViewGroup,omitempty"`


	// Optional. The widget id. Ids may be made up of alphanumerics, dashes and
	//  underscores. Widget ids are optional.
	Id *string `json:"id,omitempty"`
	*/
}

// +kcc:proto=emptypb.Empty
type Empty struct {
}

// +kcc:proto=google.monitoring.dashboard.v1.Threshold
type Threshold struct {
	// A label for the threshold.
	Label *string `json:"label,omitempty"`

	// The value of the threshold. The value should be defined in the native scale
	//  of the metric.
	// +kubebuilder:validation:Format=double
	Value *float64 `json:"value,omitempty"`

	// The state color for this threshold. Color is not allowed in a XyChart.
	Color *string `json:"color,omitempty"`

	// The direction for the current threshold. Direction is not allowed in a
	//  XyChart.
	Direction *string `json:"direction,omitempty"`

	/*NOTYET
	// The target axis to use for plotting the threshold. Target axis is not
	//  allowed in a Scorecard.
	TargetAxis *string `json:"targetAxis,omitempty"`
	*/
}

// +kcc:proto=google.monitoring.dashboard.v1.TimeSeriesFilter
type TimeSeriesFilter struct {
	// Required. The [monitoring
	//  filter](https://cloud.google.com/monitoring/api/v3/filters) that identifies
	//  the metric types, resources, and projects to query.
	//
	// +required
	Filter *string `json:"filter,omitempty"`

	// By default, the raw time series data is returned.
	//  Use this field to combine multiple time series for different views of the
	//  data.
	Aggregation *Aggregation `json:"aggregation,omitempty"`

	// Apply a second aggregation after `aggregation` is applied.
	SecondaryAggregation *Aggregation `json:"secondaryAggregation,omitempty"`

	// Ranking based time series filter.
	PickTimeSeriesFilter *PickTimeSeriesFilter `json:"pickTimeSeriesFilter,omitempty"`

	/*NOTYET
	// Statistics based time series filter.
	//  Note: This field is deprecated and completely ignored by the API.
	StatisticalTimeSeriesFilter *StatisticalTimeSeriesFilter `json:"statisticalTimeSeriesFilter,omitempty"`
	*/
}

// +kcc:proto=google.monitoring.dashboard.v1.TimeSeriesFilterRatio
type TimeSeriesFilterRatio struct {
	// The numerator of the ratio.
	Numerator *TimeSeriesFilterRatio_RatioPart `json:"numerator,omitempty"`

	// The denominator of the ratio.
	Denominator *TimeSeriesFilterRatio_RatioPart `json:"denominator,omitempty"`

	// Apply a second aggregation after the ratio is computed.
	SecondaryAggregation *Aggregation `json:"secondaryAggregation,omitempty"`

	// Ranking based time series filter.
	PickTimeSeriesFilter *PickTimeSeriesFilter `json:"pickTimeSeriesFilter,omitempty"`

	/*NOTYET
	// Statistics based time series filter.
	//  Note: This field is deprecated and completely ignored by the API.
	StatisticalTimeSeriesFilter *StatisticalTimeSeriesFilter `json:"statisticalTimeSeriesFilter,omitempty"`
	*/
}

// +kcc:proto=google.monitoring.dashboard.v1.TimeSeriesFilterRatio.RatioPart
type TimeSeriesFilterRatio_RatioPart struct {
	// Required. The [monitoring
	//  filter](https://cloud.google.com/monitoring/api/v3/filters) that
	//  identifies the metric types, resources, and projects to query.
	//
	// +required
	Filter *string `json:"filter,omitempty"`

	// By default, the raw time series data is returned.
	//  Use this field to combine multiple time series for different views of the
	//  data.
	Aggregation *Aggregation `json:"aggregation,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.TimeSeriesQuery
type TimeSeriesQuery struct {
	// Filter parameters to fetch time series.
	TimeSeriesFilter *TimeSeriesFilter `json:"timeSeriesFilter,omitempty"`

	// Parameters to fetch a ratio between two time series filters.
	TimeSeriesFilterRatio *TimeSeriesFilterRatio `json:"timeSeriesFilterRatio,omitempty"`

	// A query used to fetch time series with MQL.
	TimeSeriesQueryLanguage *string `json:"timeSeriesQueryLanguage,omitempty"`

	/*NOTYET
	// A query used to fetch time series with PromQL.
	PrometheusQuery *string `json:"prometheusQuery,omitempty"`
	*/

	// The unit of data contained in fetched time series. If non-empty, this
	//  unit will override any unit that accompanies fetched data. The format is
	//  the same as the
	//  [`unit`](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.metricDescriptors)
	//  field in `MetricDescriptor`.
	UnitOverride *string `json:"unitOverride,omitempty"`

	/*
	   NOTYET

	   		   	// Optional. If set, Cloud Monitoring will treat the full query duration as
	   	   	//  the alignment period so that there will be only 1 output value.
	   	   	//
	   	   	//  *Note: This could override the configured alignment period except for
	   	   	//  the cases where a series of data points are expected, like
	   	   	//    - XyChart
	   	   	//    - Scorecard's spark chart
	   	   	OutputFullDuration *bool `json:"outputFullDuration,omitempty"`
	*/
}

// +kcc:proto=google.monitoring.dashboard.v1.IncidentList
type IncidentList struct {
	// Optional. The monitored resource for which incidents are listed.
	//  The resource doesn't need to be fully specified. That is, you can specify
	//  the resource type but not the values of the resource labels.
	//  The resource type and labels are used for filtering.
	MonitoredResources []string `json:"monitoredResources,omitempty"`

	// Optional. A list of alert policy names to filter the incident list by.
	//  Don't include the project ID prefix in the policy name. For
	//  example, use `alertPolicies/utilization`.
	PolicyNames []string `json:"policyNames,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.Dashboard.LabelsEntry
type Dashboard_LabelsEntry struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.TableDisplayOptions
type TableDisplayOptions struct {
	// Optional. This field is unused and has been replaced by
	//  TimeSeriesTable.column_settings
	ShownColumns []string `json:"shownColumns,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.DashboardFilter
type DashboardFilter struct {
	// Required. The key for the label
	//
	// +required
	LabelKey *string `json:"labelKey,omitempty"`

	// The placeholder text that can be referenced in a filter string or MQL
	//  query. If omitted, the dashboard filter will be applied to all relevant
	//  widgets in the dashboard.
	TemplateVariable *string `json:"templateVariable,omitempty"`

	// A variable-length string value.
	StringValue *string `json:"stringValue,omitempty"`

	// The specified filter type
	FilterType *string `json:"filterType,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.PieChart
type PieChart struct {
	// Required. The queries for the chart's data.
	//
	// +required
	DataSets []PieChart_PieChartDataSet `json:"dataSets,omitempty"`

	// Required. Indicates the visualization type for the PieChart.
	//
	// +required
	ChartType *string `json:"chartType,omitempty"`

	// Optional. Indicates whether or not the pie chart should show slices' labels
	ShowLabels *bool `json:"showLabels,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.PieChart.PieChartDataSet
type PieChart_PieChartDataSet struct {
	// Required. The query for the PieChart. See,
	//  `google.monitoring.dashboard.v1.TimeSeriesQuery`.
	//
	// +required
	TimeSeriesQuery *TimeSeriesQuery `json:"timeSeriesQuery,omitempty"`

	// Optional. A template for the name of the slice. This name will be
	//  displayed in the legend and the tooltip of the pie chart. It replaces the
	//  auto-generated names for the slices. For example, if the template is set
	//  to
	//  `${resource.labels.zone}`, the zone's value will be used for the name
	//  instead of the default name.
	SliceNameTemplate *string `json:"sliceNameTemplate,omitempty"`

	// Optional. The lower bound on data point frequency for this data set,
	//  implemented by specifying the minimum alignment period to use in a time
	//  series query. For example, if the data is published once every 10
	//  minutes, the `min_alignment_period` should be at least 10 minutes. It
	//  would not make sense to fetch and align data at one minute intervals.
	MinAlignmentPeriod *string `json:"minAlignmentPeriod,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.Scorecard
type Scorecard struct {
	// Required. Fields for querying time series data from the
	//  Stackdriver metrics API.
	// +required
	TimeSeriesQuery *TimeSeriesQuery `json:"timeSeriesQuery,omitempty"`

	// Will cause the scorecard to show a gauge chart.
	GaugeView *Scorecard_GaugeView `json:"gaugeView,omitempty"`

	// Will cause the scorecard to show a spark chart.
	SparkChartView *Scorecard_SparkChartView `json:"sparkChartView,omitempty"`

	/*NOTYET
	// Will cause the `Scorecard` to show only the value, with no indicator to
	//  its value relative to its thresholds.
	BlankView *Empty `json:"blankView,omitempty"`
	*/

	// The thresholds used to determine the state of the scorecard given the
	//  time series' current value. For an actual value x, the scorecard is in a
	//  danger state if x is less than or equal to a danger threshold that triggers
	//  below, or greater than or equal to a danger threshold that triggers above.
	//  Similarly, if x is above/below a warning threshold that triggers
	//  above/below, then the scorecard is in a warning state - unless x also puts
	//  it in a danger state. (Danger trumps warning.)
	//
	//  As an example, consider a scorecard with the following four thresholds:
	//
	//  ```
	//  {
	//    value: 90,
	//    category: 'DANGER',
	//    trigger: 'ABOVE',
	//  },
	//  {
	//    value: 70,
	//    category: 'WARNING',
	//    trigger: 'ABOVE',
	//  },
	//  {
	//    value: 10,
	//    category: 'DANGER',
	//    trigger: 'BELOW',
	//  },
	//  {
	//    value: 20,
	//    category: 'WARNING',
	//    trigger: 'BELOW',
	//  }
	//  ```
	//
	//  Then: values less than or equal to 10 would put the scorecard in a DANGER
	//  state, values greater than 10 but less than or equal to 20 a WARNING state,
	//  values strictly between 20 and 70 an OK state, values greater than or equal
	//  to 70 but less than 90 a WARNING state, and values greater than or equal to
	//  90 a DANGER state.
	Thresholds []Threshold `json:"thresholds,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.Scorecard.GaugeView
type Scorecard_GaugeView struct {
	// The lower bound for this gauge chart. The value of the chart should
	//  always be greater than or equal to this.
	// +kubebuilder:validation:Format=double
	LowerBound *float64 `json:"lowerBound,omitempty"`

	// The upper bound for this gauge chart. The value of the chart should
	//  always be less than or equal to this.
	// +kubebuilder:validation:Format=double
	UpperBound *float64 `json:"upperBound,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.Scorecard.SparkChartView
type Scorecard_SparkChartView struct {
	// Required. The type of sparkchart to show in this chartView.
	//
	// +required
	SparkChartType *string `json:"sparkChartType,omitempty"`

	// The lower bound on data point frequency in the chart implemented by
	//  specifying the minimum alignment period to use in a time series query.
	//  For example, if the data is published once every 10 minutes it would not
	//  make sense to fetch and align data at one minute intervals. This field is
	//  optional and exists only as a hint.
	MinAlignmentPeriod *string `json:"minAlignmentPeriod,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.LogsPanel
type LogsPanel struct {
	// A filter that chooses which log entries to return.  See [Advanced Logs
	//  Queries](https://cloud.google.com/logging/docs/view/advanced-queries).
	//  Only log entries that match the filter are returned.  An empty filter
	//  matches all log entries.
	Filter *string `json:"filter,omitempty"`

	// The names of logging resources to collect logs for. Currently only projects
	//  are supported. If empty, the widget will default to the host project.
	ResourceNames []v1alpha1.ResourceRef `json:"resourceNames,omitempty"`
}

// +kcc:proto=google.monitoring.dashboard.v1.Dashboard
type MonitoringDashboardSpec struct {
	/* Immutable. The Project that this resource belongs to. */
	ProjectRef refs.ProjectRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The mutable, human-readable name.
	//
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Content is arranged with a basic layout that re-flows a simple list of
	//  informational elements like widgets or tiles.
	GridLayout *GridLayout `json:"gridLayout,omitempty"`

	// The content is arranged as a grid of tiles, with each content widget
	//  occupying one or more grid blocks.
	MosaicLayout *MosaicLayout `json:"mosaicLayout,omitempty"`

	// The content is divided into equally spaced rows and the widgets are
	//  arranged horizontally.
	RowLayout *RowLayout `json:"rowLayout,omitempty"`

	// The content is divided into equally spaced columns and the widgets are
	//  arranged vertically.
	ColumnLayout *ColumnLayout `json:"columnLayout,omitempty"`

	/*NOTYET
	// Filters to reduce the amount of data charted based on the filter criteria.
	DashboardFilters []DashboardFilter `json:"dashboardFilters,omitempty"`

	// Labels applied to the dashboard
	Labels []Dashboard_LabelsEntry `json:"labels,omitempty"`
	*/
}

// +kcc:proto=google.monitoring.dashboard.v1.Dashboard
type MonitoringDashboardStatus struct {
	/* Conditions represent the latest available observations of the
	   MonitoringDashboard's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* \`etag\` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. An \`etag\` is returned in the response to \`GetDashboard\`, and users are expected to put that etag in the request to \`UpdateDashboard\` to ensure that their change will be applied to the same version of the Dashboard configuration. The field should not be passed during dashboard creation. */
	// +optional
	Etag *string `json:"etag,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmonitoringdashboard;gcpmonitoringdashboards
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MonitoringDashboard is the Schema for the monitoring API
// +k8s:openapi-gen=true
type MonitoringDashboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec MonitoringDashboardSpec `json:"spec,omitempty"`

	Status MonitoringDashboardStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MonitoringDashboardList contains a list of MonitoringDashboard
type MonitoringDashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringDashboard `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MonitoringDashboard{}, &MonitoringDashboardList{})
}

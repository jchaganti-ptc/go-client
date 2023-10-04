/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.23307-4b97c8644a61
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSketchSplineDisplayData359 struct for BTSketchSplineDisplayData359
type BTSketchSplineDisplayData359 struct {
	BtType             *string   `json:"btType,omitempty"`
	Points             []float64 `json:"points,omitempty"`
	Closed             *bool     `json:"closed,omitempty"`
	ControlPointCount  *int32    `json:"controlPointCount,omitempty"`
	Degree             *int32    `json:"degree,omitempty"`
	HasHandlesInSketch *bool     `json:"hasHandlesInSketch,omitempty"`
	MaximumParameter   *float64  `json:"maximumParameter,omitempty"`
	MinimumParameter   *float64  `json:"minimumParameter,omitempty"`
	Rational           *bool     `json:"rational,omitempty"`
	Segment            *bool     `json:"segment,omitempty"`
}

// NewBTSketchSplineDisplayData359 instantiates a new BTSketchSplineDisplayData359 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSketchSplineDisplayData359() *BTSketchSplineDisplayData359 {
	this := BTSketchSplineDisplayData359{}
	return &this
}

// NewBTSketchSplineDisplayData359WithDefaults instantiates a new BTSketchSplineDisplayData359 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSketchSplineDisplayData359WithDefaults() *BTSketchSplineDisplayData359 {
	this := BTSketchSplineDisplayData359{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSketchSplineDisplayData359) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchSplineDisplayData359) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSketchSplineDisplayData359) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSketchSplineDisplayData359) SetBtType(v string) {
	o.BtType = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *BTSketchSplineDisplayData359) GetPoints() []float64 {
	if o == nil || o.Points == nil {
		var ret []float64
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchSplineDisplayData359) GetPointsOk() ([]float64, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *BTSketchSplineDisplayData359) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given []float64 and assigns it to the Points field.
func (o *BTSketchSplineDisplayData359) SetPoints(v []float64) {
	o.Points = v
}

// GetClosed returns the Closed field value if set, zero value otherwise.
func (o *BTSketchSplineDisplayData359) GetClosed() bool {
	if o == nil || o.Closed == nil {
		var ret bool
		return ret
	}
	return *o.Closed
}

// GetClosedOk returns a tuple with the Closed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchSplineDisplayData359) GetClosedOk() (*bool, bool) {
	if o == nil || o.Closed == nil {
		return nil, false
	}
	return o.Closed, true
}

// HasClosed returns a boolean if a field has been set.
func (o *BTSketchSplineDisplayData359) HasClosed() bool {
	if o != nil && o.Closed != nil {
		return true
	}

	return false
}

// SetClosed gets a reference to the given bool and assigns it to the Closed field.
func (o *BTSketchSplineDisplayData359) SetClosed(v bool) {
	o.Closed = &v
}

// GetControlPointCount returns the ControlPointCount field value if set, zero value otherwise.
func (o *BTSketchSplineDisplayData359) GetControlPointCount() int32 {
	if o == nil || o.ControlPointCount == nil {
		var ret int32
		return ret
	}
	return *o.ControlPointCount
}

// GetControlPointCountOk returns a tuple with the ControlPointCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchSplineDisplayData359) GetControlPointCountOk() (*int32, bool) {
	if o == nil || o.ControlPointCount == nil {
		return nil, false
	}
	return o.ControlPointCount, true
}

// HasControlPointCount returns a boolean if a field has been set.
func (o *BTSketchSplineDisplayData359) HasControlPointCount() bool {
	if o != nil && o.ControlPointCount != nil {
		return true
	}

	return false
}

// SetControlPointCount gets a reference to the given int32 and assigns it to the ControlPointCount field.
func (o *BTSketchSplineDisplayData359) SetControlPointCount(v int32) {
	o.ControlPointCount = &v
}

// GetDegree returns the Degree field value if set, zero value otherwise.
func (o *BTSketchSplineDisplayData359) GetDegree() int32 {
	if o == nil || o.Degree == nil {
		var ret int32
		return ret
	}
	return *o.Degree
}

// GetDegreeOk returns a tuple with the Degree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchSplineDisplayData359) GetDegreeOk() (*int32, bool) {
	if o == nil || o.Degree == nil {
		return nil, false
	}
	return o.Degree, true
}

// HasDegree returns a boolean if a field has been set.
func (o *BTSketchSplineDisplayData359) HasDegree() bool {
	if o != nil && o.Degree != nil {
		return true
	}

	return false
}

// SetDegree gets a reference to the given int32 and assigns it to the Degree field.
func (o *BTSketchSplineDisplayData359) SetDegree(v int32) {
	o.Degree = &v
}

// GetHasHandlesInSketch returns the HasHandlesInSketch field value if set, zero value otherwise.
func (o *BTSketchSplineDisplayData359) GetHasHandlesInSketch() bool {
	if o == nil || o.HasHandlesInSketch == nil {
		var ret bool
		return ret
	}
	return *o.HasHandlesInSketch
}

// GetHasHandlesInSketchOk returns a tuple with the HasHandlesInSketch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchSplineDisplayData359) GetHasHandlesInSketchOk() (*bool, bool) {
	if o == nil || o.HasHandlesInSketch == nil {
		return nil, false
	}
	return o.HasHandlesInSketch, true
}

// HasHasHandlesInSketch returns a boolean if a field has been set.
func (o *BTSketchSplineDisplayData359) HasHasHandlesInSketch() bool {
	if o != nil && o.HasHandlesInSketch != nil {
		return true
	}

	return false
}

// SetHasHandlesInSketch gets a reference to the given bool and assigns it to the HasHandlesInSketch field.
func (o *BTSketchSplineDisplayData359) SetHasHandlesInSketch(v bool) {
	o.HasHandlesInSketch = &v
}

// GetMaximumParameter returns the MaximumParameter field value if set, zero value otherwise.
func (o *BTSketchSplineDisplayData359) GetMaximumParameter() float64 {
	if o == nil || o.MaximumParameter == nil {
		var ret float64
		return ret
	}
	return *o.MaximumParameter
}

// GetMaximumParameterOk returns a tuple with the MaximumParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchSplineDisplayData359) GetMaximumParameterOk() (*float64, bool) {
	if o == nil || o.MaximumParameter == nil {
		return nil, false
	}
	return o.MaximumParameter, true
}

// HasMaximumParameter returns a boolean if a field has been set.
func (o *BTSketchSplineDisplayData359) HasMaximumParameter() bool {
	if o != nil && o.MaximumParameter != nil {
		return true
	}

	return false
}

// SetMaximumParameter gets a reference to the given float64 and assigns it to the MaximumParameter field.
func (o *BTSketchSplineDisplayData359) SetMaximumParameter(v float64) {
	o.MaximumParameter = &v
}

// GetMinimumParameter returns the MinimumParameter field value if set, zero value otherwise.
func (o *BTSketchSplineDisplayData359) GetMinimumParameter() float64 {
	if o == nil || o.MinimumParameter == nil {
		var ret float64
		return ret
	}
	return *o.MinimumParameter
}

// GetMinimumParameterOk returns a tuple with the MinimumParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchSplineDisplayData359) GetMinimumParameterOk() (*float64, bool) {
	if o == nil || o.MinimumParameter == nil {
		return nil, false
	}
	return o.MinimumParameter, true
}

// HasMinimumParameter returns a boolean if a field has been set.
func (o *BTSketchSplineDisplayData359) HasMinimumParameter() bool {
	if o != nil && o.MinimumParameter != nil {
		return true
	}

	return false
}

// SetMinimumParameter gets a reference to the given float64 and assigns it to the MinimumParameter field.
func (o *BTSketchSplineDisplayData359) SetMinimumParameter(v float64) {
	o.MinimumParameter = &v
}

// GetRational returns the Rational field value if set, zero value otherwise.
func (o *BTSketchSplineDisplayData359) GetRational() bool {
	if o == nil || o.Rational == nil {
		var ret bool
		return ret
	}
	return *o.Rational
}

// GetRationalOk returns a tuple with the Rational field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchSplineDisplayData359) GetRationalOk() (*bool, bool) {
	if o == nil || o.Rational == nil {
		return nil, false
	}
	return o.Rational, true
}

// HasRational returns a boolean if a field has been set.
func (o *BTSketchSplineDisplayData359) HasRational() bool {
	if o != nil && o.Rational != nil {
		return true
	}

	return false
}

// SetRational gets a reference to the given bool and assigns it to the Rational field.
func (o *BTSketchSplineDisplayData359) SetRational(v bool) {
	o.Rational = &v
}

// GetSegment returns the Segment field value if set, zero value otherwise.
func (o *BTSketchSplineDisplayData359) GetSegment() bool {
	if o == nil || o.Segment == nil {
		var ret bool
		return ret
	}
	return *o.Segment
}

// GetSegmentOk returns a tuple with the Segment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchSplineDisplayData359) GetSegmentOk() (*bool, bool) {
	if o == nil || o.Segment == nil {
		return nil, false
	}
	return o.Segment, true
}

// HasSegment returns a boolean if a field has been set.
func (o *BTSketchSplineDisplayData359) HasSegment() bool {
	if o != nil && o.Segment != nil {
		return true
	}

	return false
}

// SetSegment gets a reference to the given bool and assigns it to the Segment field.
func (o *BTSketchSplineDisplayData359) SetSegment(v bool) {
	o.Segment = &v
}

func (o BTSketchSplineDisplayData359) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	if o.Closed != nil {
		toSerialize["closed"] = o.Closed
	}
	if o.ControlPointCount != nil {
		toSerialize["controlPointCount"] = o.ControlPointCount
	}
	if o.Degree != nil {
		toSerialize["degree"] = o.Degree
	}
	if o.HasHandlesInSketch != nil {
		toSerialize["hasHandlesInSketch"] = o.HasHandlesInSketch
	}
	if o.MaximumParameter != nil {
		toSerialize["maximumParameter"] = o.MaximumParameter
	}
	if o.MinimumParameter != nil {
		toSerialize["minimumParameter"] = o.MinimumParameter
	}
	if o.Rational != nil {
		toSerialize["rational"] = o.Rational
	}
	if o.Segment != nil {
		toSerialize["segment"] = o.Segment
	}
	return json.Marshal(toSerialize)
}

type NullableBTSketchSplineDisplayData359 struct {
	value *BTSketchSplineDisplayData359
	isSet bool
}

func (v NullableBTSketchSplineDisplayData359) Get() *BTSketchSplineDisplayData359 {
	return v.value
}

func (v *NullableBTSketchSplineDisplayData359) Set(val *BTSketchSplineDisplayData359) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSketchSplineDisplayData359) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSketchSplineDisplayData359) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSketchSplineDisplayData359(val *BTSketchSplineDisplayData359) *NullableBTSketchSplineDisplayData359 {
	return &NullableBTSketchSplineDisplayData359{value: val, isSet: true}
}

func (v NullableBTSketchSplineDisplayData359) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSketchSplineDisplayData359) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

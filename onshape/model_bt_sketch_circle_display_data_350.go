/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.27960-e195de6ac56c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSketchCircleDisplayData350 struct for BTSketchCircleDisplayData350
type BTSketchCircleDisplayData350 struct {
	BtType *string   `json:"btType,omitempty"`
	Points []float64 `json:"points,omitempty"`
	Radius *float64  `json:"radius,omitempty"`
}

// NewBTSketchCircleDisplayData350 instantiates a new BTSketchCircleDisplayData350 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSketchCircleDisplayData350() *BTSketchCircleDisplayData350 {
	this := BTSketchCircleDisplayData350{}
	return &this
}

// NewBTSketchCircleDisplayData350WithDefaults instantiates a new BTSketchCircleDisplayData350 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSketchCircleDisplayData350WithDefaults() *BTSketchCircleDisplayData350 {
	this := BTSketchCircleDisplayData350{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSketchCircleDisplayData350) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchCircleDisplayData350) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSketchCircleDisplayData350) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSketchCircleDisplayData350) SetBtType(v string) {
	o.BtType = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *BTSketchCircleDisplayData350) GetPoints() []float64 {
	if o == nil || o.Points == nil {
		var ret []float64
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchCircleDisplayData350) GetPointsOk() ([]float64, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *BTSketchCircleDisplayData350) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given []float64 and assigns it to the Points field.
func (o *BTSketchCircleDisplayData350) SetPoints(v []float64) {
	o.Points = v
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *BTSketchCircleDisplayData350) GetRadius() float64 {
	if o == nil || o.Radius == nil {
		var ret float64
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchCircleDisplayData350) GetRadiusOk() (*float64, bool) {
	if o == nil || o.Radius == nil {
		return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *BTSketchCircleDisplayData350) HasRadius() bool {
	if o != nil && o.Radius != nil {
		return true
	}

	return false
}

// SetRadius gets a reference to the given float64 and assigns it to the Radius field.
func (o *BTSketchCircleDisplayData350) SetRadius(v float64) {
	o.Radius = &v
}

func (o BTSketchCircleDisplayData350) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	if o.Radius != nil {
		toSerialize["radius"] = o.Radius
	}
	return json.Marshal(toSerialize)
}

type NullableBTSketchCircleDisplayData350 struct {
	value *BTSketchCircleDisplayData350
	isSet bool
}

func (v NullableBTSketchCircleDisplayData350) Get() *BTSketchCircleDisplayData350 {
	return v.value
}

func (v *NullableBTSketchCircleDisplayData350) Set(val *BTSketchCircleDisplayData350) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSketchCircleDisplayData350) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSketchCircleDisplayData350) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSketchCircleDisplayData350(val *BTSketchCircleDisplayData350) *NullableBTSketchCircleDisplayData350 {
	return &NullableBTSketchCircleDisplayData350{value: val, isSet: true}
}

func (v NullableBTSketchCircleDisplayData350) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSketchCircleDisplayData350) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

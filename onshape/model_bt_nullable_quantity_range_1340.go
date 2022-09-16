/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6489-39ccb1a53c2e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTNullableQuantityRange1340 struct for BTNullableQuantityRange1340
type BTNullableQuantityRange1340 struct {
	BtType           *string            `json:"btType,omitempty"`
	DefaultValue     *float64           `json:"defaultValue,omitempty"`
	Location         *BTLocationInfo226 `json:"location,omitempty"`
	MaxValue         *float64           `json:"maxValue,omitempty"`
	MinValue         *float64           `json:"minValue,omitempty"`
	Units            *string            `json:"units,omitempty"`
	HasDefaultValue_ *bool              `json:"hasDefaultValue,omitempty"`
	HasMaxValue_     *bool              `json:"hasMaxValue,omitempty"`
	HasMinValue_     *bool              `json:"hasMinValue,omitempty"`
}

// NewBTNullableQuantityRange1340 instantiates a new BTNullableQuantityRange1340 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTNullableQuantityRange1340() *BTNullableQuantityRange1340 {
	this := BTNullableQuantityRange1340{}
	return &this
}

// NewBTNullableQuantityRange1340WithDefaults instantiates a new BTNullableQuantityRange1340 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTNullableQuantityRange1340WithDefaults() *BTNullableQuantityRange1340 {
	this := BTNullableQuantityRange1340{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTNullableQuantityRange1340) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNullableQuantityRange1340) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTNullableQuantityRange1340) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTNullableQuantityRange1340) SetBtType(v string) {
	o.BtType = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *BTNullableQuantityRange1340) GetDefaultValue() float64 {
	if o == nil || o.DefaultValue == nil {
		var ret float64
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNullableQuantityRange1340) GetDefaultValueOk() (*float64, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *BTNullableQuantityRange1340) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given float64 and assigns it to the DefaultValue field.
func (o *BTNullableQuantityRange1340) SetDefaultValue(v float64) {
	o.DefaultValue = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *BTNullableQuantityRange1340) GetLocation() BTLocationInfo226 {
	if o == nil || o.Location == nil {
		var ret BTLocationInfo226
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNullableQuantityRange1340) GetLocationOk() (*BTLocationInfo226, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *BTNullableQuantityRange1340) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given BTLocationInfo226 and assigns it to the Location field.
func (o *BTNullableQuantityRange1340) SetLocation(v BTLocationInfo226) {
	o.Location = &v
}

// GetMaxValue returns the MaxValue field value if set, zero value otherwise.
func (o *BTNullableQuantityRange1340) GetMaxValue() float64 {
	if o == nil || o.MaxValue == nil {
		var ret float64
		return ret
	}
	return *o.MaxValue
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNullableQuantityRange1340) GetMaxValueOk() (*float64, bool) {
	if o == nil || o.MaxValue == nil {
		return nil, false
	}
	return o.MaxValue, true
}

// HasMaxValue returns a boolean if a field has been set.
func (o *BTNullableQuantityRange1340) HasMaxValue() bool {
	if o != nil && o.MaxValue != nil {
		return true
	}

	return false
}

// SetMaxValue gets a reference to the given float64 and assigns it to the MaxValue field.
func (o *BTNullableQuantityRange1340) SetMaxValue(v float64) {
	o.MaxValue = &v
}

// GetMinValue returns the MinValue field value if set, zero value otherwise.
func (o *BTNullableQuantityRange1340) GetMinValue() float64 {
	if o == nil || o.MinValue == nil {
		var ret float64
		return ret
	}
	return *o.MinValue
}

// GetMinValueOk returns a tuple with the MinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNullableQuantityRange1340) GetMinValueOk() (*float64, bool) {
	if o == nil || o.MinValue == nil {
		return nil, false
	}
	return o.MinValue, true
}

// HasMinValue returns a boolean if a field has been set.
func (o *BTNullableQuantityRange1340) HasMinValue() bool {
	if o != nil && o.MinValue != nil {
		return true
	}

	return false
}

// SetMinValue gets a reference to the given float64 and assigns it to the MinValue field.
func (o *BTNullableQuantityRange1340) SetMinValue(v float64) {
	o.MinValue = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *BTNullableQuantityRange1340) GetUnits() string {
	if o == nil || o.Units == nil {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNullableQuantityRange1340) GetUnitsOk() (*string, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *BTNullableQuantityRange1340) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *BTNullableQuantityRange1340) SetUnits(v string) {
	o.Units = &v
}

// GetHasDefaultValue_ returns the HasDefaultValue_ field value if set, zero value otherwise.
func (o *BTNullableQuantityRange1340) GetHasDefaultValue_() bool {
	if o == nil || o.HasDefaultValue_ == nil {
		var ret bool
		return ret
	}
	return *o.HasDefaultValue_
}

// GetHasDefaultValue_Ok returns a tuple with the HasDefaultValue_ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNullableQuantityRange1340) GetHasDefaultValue_Ok() (*bool, bool) {
	if o == nil || o.HasDefaultValue_ == nil {
		return nil, false
	}
	return o.HasDefaultValue_, true
}

// HasHasDefaultValue_ returns a boolean if a field has been set.
func (o *BTNullableQuantityRange1340) HasHasDefaultValue_() bool {
	if o != nil && o.HasDefaultValue_ != nil {
		return true
	}

	return false
}

// SetHasDefaultValue_ gets a reference to the given bool and assigns it to the HasDefaultValue_ field.
func (o *BTNullableQuantityRange1340) SetHasDefaultValue_(v bool) {
	o.HasDefaultValue_ = &v
}

// GetHasMaxValue_ returns the HasMaxValue_ field value if set, zero value otherwise.
func (o *BTNullableQuantityRange1340) GetHasMaxValue_() bool {
	if o == nil || o.HasMaxValue_ == nil {
		var ret bool
		return ret
	}
	return *o.HasMaxValue_
}

// GetHasMaxValue_Ok returns a tuple with the HasMaxValue_ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNullableQuantityRange1340) GetHasMaxValue_Ok() (*bool, bool) {
	if o == nil || o.HasMaxValue_ == nil {
		return nil, false
	}
	return o.HasMaxValue_, true
}

// HasHasMaxValue_ returns a boolean if a field has been set.
func (o *BTNullableQuantityRange1340) HasHasMaxValue_() bool {
	if o != nil && o.HasMaxValue_ != nil {
		return true
	}

	return false
}

// SetHasMaxValue_ gets a reference to the given bool and assigns it to the HasMaxValue_ field.
func (o *BTNullableQuantityRange1340) SetHasMaxValue_(v bool) {
	o.HasMaxValue_ = &v
}

// GetHasMinValue_ returns the HasMinValue_ field value if set, zero value otherwise.
func (o *BTNullableQuantityRange1340) GetHasMinValue_() bool {
	if o == nil || o.HasMinValue_ == nil {
		var ret bool
		return ret
	}
	return *o.HasMinValue_
}

// GetHasMinValue_Ok returns a tuple with the HasMinValue_ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNullableQuantityRange1340) GetHasMinValue_Ok() (*bool, bool) {
	if o == nil || o.HasMinValue_ == nil {
		return nil, false
	}
	return o.HasMinValue_, true
}

// HasHasMinValue_ returns a boolean if a field has been set.
func (o *BTNullableQuantityRange1340) HasHasMinValue_() bool {
	if o != nil && o.HasMinValue_ != nil {
		return true
	}

	return false
}

// SetHasMinValue_ gets a reference to the given bool and assigns it to the HasMinValue_ field.
func (o *BTNullableQuantityRange1340) SetHasMinValue_(v bool) {
	o.HasMinValue_ = &v
}

func (o BTNullableQuantityRange1340) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.MaxValue != nil {
		toSerialize["maxValue"] = o.MaxValue
	}
	if o.MinValue != nil {
		toSerialize["minValue"] = o.MinValue
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	if o.HasDefaultValue_ != nil {
		toSerialize["hasDefaultValue"] = o.HasDefaultValue_
	}
	if o.HasMaxValue_ != nil {
		toSerialize["hasMaxValue"] = o.HasMaxValue_
	}
	if o.HasMinValue_ != nil {
		toSerialize["hasMinValue"] = o.HasMinValue_
	}
	return json.Marshal(toSerialize)
}

type NullableBTNullableQuantityRange1340 struct {
	value *BTNullableQuantityRange1340
	isSet bool
}

func (v NullableBTNullableQuantityRange1340) Get() *BTNullableQuantityRange1340 {
	return v.value
}

func (v *NullableBTNullableQuantityRange1340) Set(val *BTNullableQuantityRange1340) {
	v.value = val
	v.isSet = true
}

func (v NullableBTNullableQuantityRange1340) IsSet() bool {
	return v.isSet
}

func (v *NullableBTNullableQuantityRange1340) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTNullableQuantityRange1340(val *BTNullableQuantityRange1340) *NullableBTNullableQuantityRange1340 {
	return &NullableBTNullableQuantityRange1340{value: val, isSet: true}
}

func (v NullableBTNullableQuantityRange1340) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTNullableQuantityRange1340) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

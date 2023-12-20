/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.27881-5dbbe8053cdf
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTApiConfiguration struct for BTApiConfiguration
type BTApiConfiguration struct {
	Current                     *bool              `json:"current,omitempty"`
	Default                     *bool              `json:"default,omitempty"`
	Null                        *bool              `json:"null,omitempty"`
	ParameterMap                *map[string]string `json:"parameterMap,omitempty"`
	StandardContent             *bool              `json:"standardContent,omitempty"`
	StandardContentParametersId *string            `json:"standardContentParametersId,omitempty"`
}

// NewBTApiConfiguration instantiates a new BTApiConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTApiConfiguration() *BTApiConfiguration {
	this := BTApiConfiguration{}
	return &this
}

// NewBTApiConfigurationWithDefaults instantiates a new BTApiConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTApiConfigurationWithDefaults() *BTApiConfiguration {
	this := BTApiConfiguration{}
	return &this
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *BTApiConfiguration) GetCurrent() bool {
	if o == nil || o.Current == nil {
		var ret bool
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiConfiguration) GetCurrentOk() (*bool, bool) {
	if o == nil || o.Current == nil {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *BTApiConfiguration) HasCurrent() bool {
	if o != nil && o.Current != nil {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given bool and assigns it to the Current field.
func (o *BTApiConfiguration) SetCurrent(v bool) {
	o.Current = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *BTApiConfiguration) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiConfiguration) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *BTApiConfiguration) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *BTApiConfiguration) SetDefault(v bool) {
	o.Default = &v
}

// GetNull returns the Null field value if set, zero value otherwise.
func (o *BTApiConfiguration) GetNull() bool {
	if o == nil || o.Null == nil {
		var ret bool
		return ret
	}
	return *o.Null
}

// GetNullOk returns a tuple with the Null field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiConfiguration) GetNullOk() (*bool, bool) {
	if o == nil || o.Null == nil {
		return nil, false
	}
	return o.Null, true
}

// HasNull returns a boolean if a field has been set.
func (o *BTApiConfiguration) HasNull() bool {
	if o != nil && o.Null != nil {
		return true
	}

	return false
}

// SetNull gets a reference to the given bool and assigns it to the Null field.
func (o *BTApiConfiguration) SetNull(v bool) {
	o.Null = &v
}

// GetParameterMap returns the ParameterMap field value if set, zero value otherwise.
func (o *BTApiConfiguration) GetParameterMap() map[string]string {
	if o == nil || o.ParameterMap == nil {
		var ret map[string]string
		return ret
	}
	return *o.ParameterMap
}

// GetParameterMapOk returns a tuple with the ParameterMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiConfiguration) GetParameterMapOk() (*map[string]string, bool) {
	if o == nil || o.ParameterMap == nil {
		return nil, false
	}
	return o.ParameterMap, true
}

// HasParameterMap returns a boolean if a field has been set.
func (o *BTApiConfiguration) HasParameterMap() bool {
	if o != nil && o.ParameterMap != nil {
		return true
	}

	return false
}

// SetParameterMap gets a reference to the given map[string]string and assigns it to the ParameterMap field.
func (o *BTApiConfiguration) SetParameterMap(v map[string]string) {
	o.ParameterMap = &v
}

// GetStandardContent returns the StandardContent field value if set, zero value otherwise.
func (o *BTApiConfiguration) GetStandardContent() bool {
	if o == nil || o.StandardContent == nil {
		var ret bool
		return ret
	}
	return *o.StandardContent
}

// GetStandardContentOk returns a tuple with the StandardContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiConfiguration) GetStandardContentOk() (*bool, bool) {
	if o == nil || o.StandardContent == nil {
		return nil, false
	}
	return o.StandardContent, true
}

// HasStandardContent returns a boolean if a field has been set.
func (o *BTApiConfiguration) HasStandardContent() bool {
	if o != nil && o.StandardContent != nil {
		return true
	}

	return false
}

// SetStandardContent gets a reference to the given bool and assigns it to the StandardContent field.
func (o *BTApiConfiguration) SetStandardContent(v bool) {
	o.StandardContent = &v
}

// GetStandardContentParametersId returns the StandardContentParametersId field value if set, zero value otherwise.
func (o *BTApiConfiguration) GetStandardContentParametersId() string {
	if o == nil || o.StandardContentParametersId == nil {
		var ret string
		return ret
	}
	return *o.StandardContentParametersId
}

// GetStandardContentParametersIdOk returns a tuple with the StandardContentParametersId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiConfiguration) GetStandardContentParametersIdOk() (*string, bool) {
	if o == nil || o.StandardContentParametersId == nil {
		return nil, false
	}
	return o.StandardContentParametersId, true
}

// HasStandardContentParametersId returns a boolean if a field has been set.
func (o *BTApiConfiguration) HasStandardContentParametersId() bool {
	if o != nil && o.StandardContentParametersId != nil {
		return true
	}

	return false
}

// SetStandardContentParametersId gets a reference to the given string and assigns it to the StandardContentParametersId field.
func (o *BTApiConfiguration) SetStandardContentParametersId(v string) {
	o.StandardContentParametersId = &v
}

func (o BTApiConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Current != nil {
		toSerialize["current"] = o.Current
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.Null != nil {
		toSerialize["null"] = o.Null
	}
	if o.ParameterMap != nil {
		toSerialize["parameterMap"] = o.ParameterMap
	}
	if o.StandardContent != nil {
		toSerialize["standardContent"] = o.StandardContent
	}
	if o.StandardContentParametersId != nil {
		toSerialize["standardContentParametersId"] = o.StandardContentParametersId
	}
	return json.Marshal(toSerialize)
}

type NullableBTApiConfiguration struct {
	value *BTApiConfiguration
	isSet bool
}

func (v NullableBTApiConfiguration) Get() *BTApiConfiguration {
	return v.value
}

func (v *NullableBTApiConfiguration) Set(val *BTApiConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableBTApiConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableBTApiConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTApiConfiguration(val *BTApiConfiguration) *NullableBTApiConfiguration {
	return &NullableBTApiConfiguration{value: val, isSet: true}
}

func (v NullableBTApiConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTApiConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

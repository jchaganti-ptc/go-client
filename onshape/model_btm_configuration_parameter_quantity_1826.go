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

// BTMConfigurationParameterQuantity1826 struct for BTMConfigurationParameterQuantity1826
type BTMConfigurationParameterQuantity1826 struct {
	BtType               *string                        `json:"btType,omitempty"`
	GeneratedParameterId *BTTreeNode20                  `json:"generatedParameterId,omitempty"`
	ImportMicroversion   *string                        `json:"importMicroversion,omitempty"`
	NodeId               *string                        `json:"nodeId,omitempty"`
	ParameterId          *string                        `json:"parameterId,omitempty"`
	ParameterName        *string                        `json:"parameterName,omitempty"`
	ParameterType        *GBTConfigurationParameterType `json:"parameterType,omitempty"`
	Valid                *bool                          `json:"valid,omitempty"`
	QuantityType         *GBTQuantityType               `json:"quantityType,omitempty"`
	RangeAndDefault      *BTQuantityRange181            `json:"rangeAndDefault,omitempty"`
}

// NewBTMConfigurationParameterQuantity1826 instantiates a new BTMConfigurationParameterQuantity1826 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMConfigurationParameterQuantity1826() *BTMConfigurationParameterQuantity1826 {
	this := BTMConfigurationParameterQuantity1826{}
	return &this
}

// NewBTMConfigurationParameterQuantity1826WithDefaults instantiates a new BTMConfigurationParameterQuantity1826 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMConfigurationParameterQuantity1826WithDefaults() *BTMConfigurationParameterQuantity1826 {
	this := BTMConfigurationParameterQuantity1826{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMConfigurationParameterQuantity1826) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterQuantity1826) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMConfigurationParameterQuantity1826) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMConfigurationParameterQuantity1826) SetBtType(v string) {
	o.BtType = &v
}

// GetGeneratedParameterId returns the GeneratedParameterId field value if set, zero value otherwise.
func (o *BTMConfigurationParameterQuantity1826) GetGeneratedParameterId() BTTreeNode20 {
	if o == nil || o.GeneratedParameterId == nil {
		var ret BTTreeNode20
		return ret
	}
	return *o.GeneratedParameterId
}

// GetGeneratedParameterIdOk returns a tuple with the GeneratedParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterQuantity1826) GetGeneratedParameterIdOk() (*BTTreeNode20, bool) {
	if o == nil || o.GeneratedParameterId == nil {
		return nil, false
	}
	return o.GeneratedParameterId, true
}

// HasGeneratedParameterId returns a boolean if a field has been set.
func (o *BTMConfigurationParameterQuantity1826) HasGeneratedParameterId() bool {
	if o != nil && o.GeneratedParameterId != nil {
		return true
	}

	return false
}

// SetGeneratedParameterId gets a reference to the given BTTreeNode20 and assigns it to the GeneratedParameterId field.
func (o *BTMConfigurationParameterQuantity1826) SetGeneratedParameterId(v BTTreeNode20) {
	o.GeneratedParameterId = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMConfigurationParameterQuantity1826) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterQuantity1826) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMConfigurationParameterQuantity1826) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMConfigurationParameterQuantity1826) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMConfigurationParameterQuantity1826) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterQuantity1826) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMConfigurationParameterQuantity1826) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMConfigurationParameterQuantity1826) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMConfigurationParameterQuantity1826) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterQuantity1826) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMConfigurationParameterQuantity1826) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMConfigurationParameterQuantity1826) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetParameterName returns the ParameterName field value if set, zero value otherwise.
func (o *BTMConfigurationParameterQuantity1826) GetParameterName() string {
	if o == nil || o.ParameterName == nil {
		var ret string
		return ret
	}
	return *o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterQuantity1826) GetParameterNameOk() (*string, bool) {
	if o == nil || o.ParameterName == nil {
		return nil, false
	}
	return o.ParameterName, true
}

// HasParameterName returns a boolean if a field has been set.
func (o *BTMConfigurationParameterQuantity1826) HasParameterName() bool {
	if o != nil && o.ParameterName != nil {
		return true
	}

	return false
}

// SetParameterName gets a reference to the given string and assigns it to the ParameterName field.
func (o *BTMConfigurationParameterQuantity1826) SetParameterName(v string) {
	o.ParameterName = &v
}

// GetParameterType returns the ParameterType field value if set, zero value otherwise.
func (o *BTMConfigurationParameterQuantity1826) GetParameterType() GBTConfigurationParameterType {
	if o == nil || o.ParameterType == nil {
		var ret GBTConfigurationParameterType
		return ret
	}
	return *o.ParameterType
}

// GetParameterTypeOk returns a tuple with the ParameterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterQuantity1826) GetParameterTypeOk() (*GBTConfigurationParameterType, bool) {
	if o == nil || o.ParameterType == nil {
		return nil, false
	}
	return o.ParameterType, true
}

// HasParameterType returns a boolean if a field has been set.
func (o *BTMConfigurationParameterQuantity1826) HasParameterType() bool {
	if o != nil && o.ParameterType != nil {
		return true
	}

	return false
}

// SetParameterType gets a reference to the given GBTConfigurationParameterType and assigns it to the ParameterType field.
func (o *BTMConfigurationParameterQuantity1826) SetParameterType(v GBTConfigurationParameterType) {
	o.ParameterType = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *BTMConfigurationParameterQuantity1826) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterQuantity1826) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *BTMConfigurationParameterQuantity1826) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *BTMConfigurationParameterQuantity1826) SetValid(v bool) {
	o.Valid = &v
}

// GetQuantityType returns the QuantityType field value if set, zero value otherwise.
func (o *BTMConfigurationParameterQuantity1826) GetQuantityType() GBTQuantityType {
	if o == nil || o.QuantityType == nil {
		var ret GBTQuantityType
		return ret
	}
	return *o.QuantityType
}

// GetQuantityTypeOk returns a tuple with the QuantityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterQuantity1826) GetQuantityTypeOk() (*GBTQuantityType, bool) {
	if o == nil || o.QuantityType == nil {
		return nil, false
	}
	return o.QuantityType, true
}

// HasQuantityType returns a boolean if a field has been set.
func (o *BTMConfigurationParameterQuantity1826) HasQuantityType() bool {
	if o != nil && o.QuantityType != nil {
		return true
	}

	return false
}

// SetQuantityType gets a reference to the given GBTQuantityType and assigns it to the QuantityType field.
func (o *BTMConfigurationParameterQuantity1826) SetQuantityType(v GBTQuantityType) {
	o.QuantityType = &v
}

// GetRangeAndDefault returns the RangeAndDefault field value if set, zero value otherwise.
func (o *BTMConfigurationParameterQuantity1826) GetRangeAndDefault() BTQuantityRange181 {
	if o == nil || o.RangeAndDefault == nil {
		var ret BTQuantityRange181
		return ret
	}
	return *o.RangeAndDefault
}

// GetRangeAndDefaultOk returns a tuple with the RangeAndDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfigurationParameterQuantity1826) GetRangeAndDefaultOk() (*BTQuantityRange181, bool) {
	if o == nil || o.RangeAndDefault == nil {
		return nil, false
	}
	return o.RangeAndDefault, true
}

// HasRangeAndDefault returns a boolean if a field has been set.
func (o *BTMConfigurationParameterQuantity1826) HasRangeAndDefault() bool {
	if o != nil && o.RangeAndDefault != nil {
		return true
	}

	return false
}

// SetRangeAndDefault gets a reference to the given BTQuantityRange181 and assigns it to the RangeAndDefault field.
func (o *BTMConfigurationParameterQuantity1826) SetRangeAndDefault(v BTQuantityRange181) {
	o.RangeAndDefault = &v
}

func (o BTMConfigurationParameterQuantity1826) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.GeneratedParameterId != nil {
		toSerialize["generatedParameterId"] = o.GeneratedParameterId
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.ParameterName != nil {
		toSerialize["parameterName"] = o.ParameterName
	}
	if o.ParameterType != nil {
		toSerialize["parameterType"] = o.ParameterType
	}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}
	if o.QuantityType != nil {
		toSerialize["quantityType"] = o.QuantityType
	}
	if o.RangeAndDefault != nil {
		toSerialize["rangeAndDefault"] = o.RangeAndDefault
	}
	return json.Marshal(toSerialize)
}

type NullableBTMConfigurationParameterQuantity1826 struct {
	value *BTMConfigurationParameterQuantity1826
	isSet bool
}

func (v NullableBTMConfigurationParameterQuantity1826) Get() *BTMConfigurationParameterQuantity1826 {
	return v.value
}

func (v *NullableBTMConfigurationParameterQuantity1826) Set(val *BTMConfigurationParameterQuantity1826) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMConfigurationParameterQuantity1826) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMConfigurationParameterQuantity1826) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMConfigurationParameterQuantity1826(val *BTMConfigurationParameterQuantity1826) *NullableBTMConfigurationParameterQuantity1826 {
	return &NullableBTMConfigurationParameterQuantity1826{value: val, isSet: true}
}

func (v NullableBTMConfigurationParameterQuantity1826) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMConfigurationParameterQuantity1826) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

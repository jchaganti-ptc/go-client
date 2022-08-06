/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5928-bd774e9c9f97
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMParameterNullableQuantity807 struct for BTMParameterNullableQuantity807
type BTMParameterNullableQuantity807 struct {
	ImportMicroversion *string  `json:"importMicroversion,omitempty"`
	NodeId             *string  `json:"nodeId,omitempty"`
	ParameterId        *string  `json:"parameterId,omitempty"`
	BtType             *string  `json:"btType,omitempty"`
	Expression         *string  `json:"expression,omitempty"`
	IsInteger          *bool    `json:"isInteger,omitempty"`
	Units              *string  `json:"units,omitempty"`
	Value              *float64 `json:"value,omitempty"`
	IsNull             *bool    `json:"isNull,omitempty"`
	NullValue          *string  `json:"nullValue,omitempty"`
}

// NewBTMParameterNullableQuantity807 instantiates a new BTMParameterNullableQuantity807 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterNullableQuantity807() *BTMParameterNullableQuantity807 {
	this := BTMParameterNullableQuantity807{}
	return &this
}

// NewBTMParameterNullableQuantity807WithDefaults instantiates a new BTMParameterNullableQuantity807 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterNullableQuantity807WithDefaults() *BTMParameterNullableQuantity807 {
	this := BTMParameterNullableQuantity807{}
	return &this
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMParameterNullableQuantity807) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterNullableQuantity807) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMParameterNullableQuantity807) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMParameterNullableQuantity807) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMParameterNullableQuantity807) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterNullableQuantity807) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMParameterNullableQuantity807) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMParameterNullableQuantity807) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMParameterNullableQuantity807) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterNullableQuantity807) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMParameterNullableQuantity807) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMParameterNullableQuantity807) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterNullableQuantity807) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterNullableQuantity807) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMParameterNullableQuantity807) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMParameterNullableQuantity807) SetBtType(v string) {
	o.BtType = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *BTMParameterNullableQuantity807) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterNullableQuantity807) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *BTMParameterNullableQuantity807) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *BTMParameterNullableQuantity807) SetExpression(v string) {
	o.Expression = &v
}

// GetIsInteger returns the IsInteger field value if set, zero value otherwise.
func (o *BTMParameterNullableQuantity807) GetIsInteger() bool {
	if o == nil || o.IsInteger == nil {
		var ret bool
		return ret
	}
	return *o.IsInteger
}

// GetIsIntegerOk returns a tuple with the IsInteger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterNullableQuantity807) GetIsIntegerOk() (*bool, bool) {
	if o == nil || o.IsInteger == nil {
		return nil, false
	}
	return o.IsInteger, true
}

// HasIsInteger returns a boolean if a field has been set.
func (o *BTMParameterNullableQuantity807) HasIsInteger() bool {
	if o != nil && o.IsInteger != nil {
		return true
	}

	return false
}

// SetIsInteger gets a reference to the given bool and assigns it to the IsInteger field.
func (o *BTMParameterNullableQuantity807) SetIsInteger(v bool) {
	o.IsInteger = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *BTMParameterNullableQuantity807) GetUnits() string {
	if o == nil || o.Units == nil {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterNullableQuantity807) GetUnitsOk() (*string, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *BTMParameterNullableQuantity807) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *BTMParameterNullableQuantity807) SetUnits(v string) {
	o.Units = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTMParameterNullableQuantity807) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterNullableQuantity807) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTMParameterNullableQuantity807) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *BTMParameterNullableQuantity807) SetValue(v float64) {
	o.Value = &v
}

// GetIsNull returns the IsNull field value if set, zero value otherwise.
func (o *BTMParameterNullableQuantity807) GetIsNull() bool {
	if o == nil || o.IsNull == nil {
		var ret bool
		return ret
	}
	return *o.IsNull
}

// GetIsNullOk returns a tuple with the IsNull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterNullableQuantity807) GetIsNullOk() (*bool, bool) {
	if o == nil || o.IsNull == nil {
		return nil, false
	}
	return o.IsNull, true
}

// HasIsNull returns a boolean if a field has been set.
func (o *BTMParameterNullableQuantity807) HasIsNull() bool {
	if o != nil && o.IsNull != nil {
		return true
	}

	return false
}

// SetIsNull gets a reference to the given bool and assigns it to the IsNull field.
func (o *BTMParameterNullableQuantity807) SetIsNull(v bool) {
	o.IsNull = &v
}

// GetNullValue returns the NullValue field value if set, zero value otherwise.
func (o *BTMParameterNullableQuantity807) GetNullValue() string {
	if o == nil || o.NullValue == nil {
		var ret string
		return ret
	}
	return *o.NullValue
}

// GetNullValueOk returns a tuple with the NullValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterNullableQuantity807) GetNullValueOk() (*string, bool) {
	if o == nil || o.NullValue == nil {
		return nil, false
	}
	return o.NullValue, true
}

// HasNullValue returns a boolean if a field has been set.
func (o *BTMParameterNullableQuantity807) HasNullValue() bool {
	if o != nil && o.NullValue != nil {
		return true
	}

	return false
}

// SetNullValue gets a reference to the given string and assigns it to the NullValue field.
func (o *BTMParameterNullableQuantity807) SetNullValue(v string) {
	o.NullValue = &v
}

func (o BTMParameterNullableQuantity807) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	if o.IsInteger != nil {
		toSerialize["isInteger"] = o.IsInteger
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.IsNull != nil {
		toSerialize["isNull"] = o.IsNull
	}
	if o.NullValue != nil {
		toSerialize["nullValue"] = o.NullValue
	}
	return json.Marshal(toSerialize)
}

type NullableBTMParameterNullableQuantity807 struct {
	value *BTMParameterNullableQuantity807
	isSet bool
}

func (v NullableBTMParameterNullableQuantity807) Get() *BTMParameterNullableQuantity807 {
	return v.value
}

func (v *NullableBTMParameterNullableQuantity807) Set(val *BTMParameterNullableQuantity807) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterNullableQuantity807) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterNullableQuantity807) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterNullableQuantity807(val *BTMParameterNullableQuantity807) *NullableBTMParameterNullableQuantity807 {
	return &NullableBTMParameterNullableQuantity807{value: val, isSet: true}
}

func (v NullableBTMParameterNullableQuantity807) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterNullableQuantity807) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

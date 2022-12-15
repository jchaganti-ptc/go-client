/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.157.8712-62a9cfa549d9
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTTableColumnInfo1222 - struct for BTTableColumnInfo1222
type BTTableColumnInfo1222 struct {
	implBTTableColumnInfo1222 interface{}
}

// BTNamedPositionValuesColumnInfo816AsBTTableColumnInfo1222 is a convenience function that returns BTNamedPositionValuesColumnInfo816 wrapped in BTTableColumnInfo1222
func (o *BTNamedPositionValuesColumnInfo816) AsBTTableColumnInfo1222() *BTTableColumnInfo1222 {
	return &BTTableColumnInfo1222{o}
}

// BTConfiguredParameterColumnInfo2900AsBTTableColumnInfo1222 is a convenience function that returns BTConfiguredParameterColumnInfo2900 wrapped in BTTableColumnInfo1222
func (o *BTConfiguredParameterColumnInfo2900) AsBTTableColumnInfo1222() *BTTableColumnInfo1222 {
	return &BTTableColumnInfo1222{o}
}

// BTConfiguredSuppressionColumnInfo2498AsBTTableColumnInfo1222 is a convenience function that returns BTConfiguredSuppressionColumnInfo2498 wrapped in BTTableColumnInfo1222
func (o *BTConfiguredSuppressionColumnInfo2498) AsBTTableColumnInfo1222() *BTTableColumnInfo1222 {
	return &BTTableColumnInfo1222{o}
}

// BTPropertyTableColumnInfo2161AsBTTableColumnInfo1222 is a convenience function that returns BTPropertyTableColumnInfo2161 wrapped in BTTableColumnInfo1222
func (o *BTPropertyTableColumnInfo2161) AsBTTableColumnInfo1222() *BTTableColumnInfo1222 {
	return &BTTableColumnInfo1222{o}
}

// BTConfiguredValuesColumnInfo1025AsBTTableColumnInfo1222 is a convenience function that returns BTConfiguredValuesColumnInfo1025 wrapped in BTTableColumnInfo1222
func (o *BTConfiguredValuesColumnInfo1025) AsBTTableColumnInfo1222() *BTTableColumnInfo1222 {
	return &BTTableColumnInfo1222{o}
}

// BTConfiguredFeatureColumnInfo1014AsBTTableColumnInfo1222 is a convenience function that returns BTConfiguredFeatureColumnInfo1014 wrapped in BTTableColumnInfo1222
func (o *BTConfiguredFeatureColumnInfo1014) AsBTTableColumnInfo1222() *BTTableColumnInfo1222 {
	return &BTTableColumnInfo1222{o}
}

// BTConfiguredDimensionColumnInfo2168AsBTTableColumnInfo1222 is a convenience function that returns BTConfiguredDimensionColumnInfo2168 wrapped in BTTableColumnInfo1222
func (o *BTConfiguredDimensionColumnInfo2168) AsBTTableColumnInfo1222() *BTTableColumnInfo1222 {
	return &BTTableColumnInfo1222{o}
}

// BTSimulationTableColumnInfo1785AsBTTableColumnInfo1222 is a convenience function that returns BTSimulationTableColumnInfo1785 wrapped in BTTableColumnInfo1222
func (o *BTSimulationTableColumnInfo1785) AsBTTableColumnInfo1222() *BTTableColumnInfo1222 {
	return &BTTableColumnInfo1222{o}
}

// BTFSTableColumnInfo623AsBTTableColumnInfo1222 is a convenience function that returns BTFSTableColumnInfo623 wrapped in BTTableColumnInfo1222
func (o *BTFSTableColumnInfo623) AsBTTableColumnInfo1222() *BTTableColumnInfo1222 {
	return &BTTableColumnInfo1222{o}
}

// NewBTTableColumnInfo1222 instantiates a new BTTableColumnInfo1222 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableColumnInfo1222() *BTTableColumnInfo1222 {
	this := BTTableColumnInfo1222{Newbase_BTTableColumnInfo1222()}
	return &this
}

// NewBTTableColumnInfo1222WithDefaults instantiates a new BTTableColumnInfo1222 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableColumnInfo1222WithDefaults() *BTTableColumnInfo1222 {
	this := BTTableColumnInfo1222{Newbase_BTTableColumnInfo1222WithDefaults()}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTTableColumnInfo1222) GetId() string {
	type getResult interface {
		GetId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetId()
	} else {
		var de string
		return de
	}
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableColumnInfo1222) GetIdOk() (*string, bool) {
	type getResult interface {
		GetIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIdOk()
	} else {
		return nil, false
	}
}

// HasId returns a boolean if a field has been set.
func (o *BTTableColumnInfo1222) HasId() bool {
	type getResult interface {
		HasId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasId()
	} else {
		return false
	}
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTTableColumnInfo1222) SetId(v string) {
	type getResult interface {
		SetId(v string)
	}

	o.GetActualInstance().(getResult).SetId(v)
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTTableColumnInfo1222) GetNodeId() string {
	type getResult interface {
		GetNodeId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeId()
	} else {
		var de string
		return de
	}
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableColumnInfo1222) GetNodeIdOk() (*string, bool) {
	type getResult interface {
		GetNodeIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeIdOk()
	} else {
		return nil, false
	}
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTTableColumnInfo1222) HasNodeId() bool {
	type getResult interface {
		HasNodeId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasNodeId()
	} else {
		return false
	}
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTTableColumnInfo1222) SetNodeId(v string) {
	type getResult interface {
		SetNodeId(v string)
	}

	o.GetActualInstance().(getResult).SetNodeId(v)
}

// GetSpecification returns the Specification field value if set, zero value otherwise.
func (o *BTTableColumnInfo1222) GetSpecification() BTTableColumnSpec1967 {
	type getResult interface {
		GetSpecification() BTTableColumnSpec1967
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSpecification()
	} else {
		var de BTTableColumnSpec1967
		return de
	}
}

// GetSpecificationOk returns a tuple with the Specification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableColumnInfo1222) GetSpecificationOk() (*BTTableColumnSpec1967, bool) {
	type getResult interface {
		GetSpecificationOk() (*BTTableColumnSpec1967, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetSpecificationOk()
	} else {
		return nil, false
	}
}

// HasSpecification returns a boolean if a field has been set.
func (o *BTTableColumnInfo1222) HasSpecification() bool {
	type getResult interface {
		HasSpecification() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasSpecification()
	} else {
		return false
	}
}

// SetSpecification gets a reference to the given BTTableColumnSpec1967 and assigns it to the Specification field.
func (o *BTTableColumnInfo1222) SetSpecification(v BTTableColumnSpec1967) {
	type getResult interface {
		SetSpecification(v BTTableColumnSpec1967)
	}

	o.GetActualInstance().(getResult).SetSpecification(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTTableColumnInfo1222) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTConfiguredDimensionColumnInfo-2168'
	if jsonDict["btType"] == "BTConfiguredDimensionColumnInfo-2168" {
		// try to unmarshal JSON data into BTConfiguredDimensionColumnInfo2168
		var qr *BTConfiguredDimensionColumnInfo2168
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableColumnInfo1222 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableColumnInfo1222 = nil
			return fmt.Errorf("Failed to unmarshal BTTableColumnInfo1222 as BTConfiguredDimensionColumnInfo2168: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTConfiguredFeatureColumnInfo-1014'
	if jsonDict["btType"] == "BTConfiguredFeatureColumnInfo-1014" {
		// try to unmarshal JSON data into BTConfiguredFeatureColumnInfo1014
		var qr *BTConfiguredFeatureColumnInfo1014
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableColumnInfo1222 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableColumnInfo1222 = nil
			return fmt.Errorf("Failed to unmarshal BTTableColumnInfo1222 as BTConfiguredFeatureColumnInfo1014: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTConfiguredParameterColumnInfo-2900'
	if jsonDict["btType"] == "BTConfiguredParameterColumnInfo-2900" {
		// try to unmarshal JSON data into BTConfiguredParameterColumnInfo2900
		var qr *BTConfiguredParameterColumnInfo2900
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableColumnInfo1222 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableColumnInfo1222 = nil
			return fmt.Errorf("Failed to unmarshal BTTableColumnInfo1222 as BTConfiguredParameterColumnInfo2900: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTConfiguredSuppressionColumnInfo-2498'
	if jsonDict["btType"] == "BTConfiguredSuppressionColumnInfo-2498" {
		// try to unmarshal JSON data into BTConfiguredSuppressionColumnInfo2498
		var qr *BTConfiguredSuppressionColumnInfo2498
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableColumnInfo1222 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableColumnInfo1222 = nil
			return fmt.Errorf("Failed to unmarshal BTTableColumnInfo1222 as BTConfiguredSuppressionColumnInfo2498: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTConfiguredValuesColumnInfo-1025'
	if jsonDict["btType"] == "BTConfiguredValuesColumnInfo-1025" {
		// try to unmarshal JSON data into BTConfiguredValuesColumnInfo1025
		var qr *BTConfiguredValuesColumnInfo1025
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableColumnInfo1222 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableColumnInfo1222 = nil
			return fmt.Errorf("Failed to unmarshal BTTableColumnInfo1222 as BTConfiguredValuesColumnInfo1025: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTFSTableColumnInfo-623'
	if jsonDict["btType"] == "BTFSTableColumnInfo-623" {
		// try to unmarshal JSON data into BTFSTableColumnInfo623
		var qr *BTFSTableColumnInfo623
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableColumnInfo1222 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableColumnInfo1222 = nil
			return fmt.Errorf("Failed to unmarshal BTTableColumnInfo1222 as BTFSTableColumnInfo623: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTNamedPositionValuesColumnInfo-816'
	if jsonDict["btType"] == "BTNamedPositionValuesColumnInfo-816" {
		// try to unmarshal JSON data into BTNamedPositionValuesColumnInfo816
		var qr *BTNamedPositionValuesColumnInfo816
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableColumnInfo1222 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableColumnInfo1222 = nil
			return fmt.Errorf("Failed to unmarshal BTTableColumnInfo1222 as BTNamedPositionValuesColumnInfo816: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTPropertyTableColumnInfo-2161'
	if jsonDict["btType"] == "BTPropertyTableColumnInfo-2161" {
		// try to unmarshal JSON data into BTPropertyTableColumnInfo2161
		var qr *BTPropertyTableColumnInfo2161
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableColumnInfo1222 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableColumnInfo1222 = nil
			return fmt.Errorf("Failed to unmarshal BTTableColumnInfo1222 as BTPropertyTableColumnInfo2161: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTSimulationTableColumnInfo-1785'
	if jsonDict["btType"] == "BTSimulationTableColumnInfo-1785" {
		// try to unmarshal JSON data into BTSimulationTableColumnInfo1785
		var qr *BTSimulationTableColumnInfo1785
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableColumnInfo1222 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableColumnInfo1222 = nil
			return fmt.Errorf("Failed to unmarshal BTTableColumnInfo1222 as BTSimulationTableColumnInfo1785: %s", err.Error())
		}
	}

	var qtx *base_BTTableColumnInfo1222
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTTableColumnInfo1222 = qtx
		return nil // data stored in dst.base_BTTableColumnInfo1222, return on the first match
	} else {
		dst.implBTTableColumnInfo1222 = nil
		return fmt.Errorf("Failed to unmarshal BTTableColumnInfo1222 as base_BTTableColumnInfo1222: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTTableColumnInfo1222) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTTableColumnInfo1222) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTTableColumnInfo1222
}

type NullableBTTableColumnInfo1222 struct {
	value *BTTableColumnInfo1222
	isSet bool
}

func (v NullableBTTableColumnInfo1222) Get() *BTTableColumnInfo1222 {
	return v.value
}

func (v *NullableBTTableColumnInfo1222) Set(val *BTTableColumnInfo1222) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableColumnInfo1222) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableColumnInfo1222) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableColumnInfo1222(val *BTTableColumnInfo1222) *NullableBTTableColumnInfo1222 {
	return &NullableBTTableColumnInfo1222{value: val, isSet: true}
}

func (v NullableBTTableColumnInfo1222) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableColumnInfo1222) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTTableColumnInfo1222 struct {
	Id            *string                `json:"id,omitempty"`
	NodeId        *string                `json:"nodeId,omitempty"`
	Specification *BTTableColumnSpec1967 `json:"specification,omitempty"`
}

// Newbase_BTTableColumnInfo1222 instantiates a new base_BTTableColumnInfo1222 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTTableColumnInfo1222() *base_BTTableColumnInfo1222 {
	this := base_BTTableColumnInfo1222{}
	return &this
}

// Newbase_BTTableColumnInfo1222WithDefaults instantiates a new base_BTTableColumnInfo1222 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTTableColumnInfo1222WithDefaults() *base_BTTableColumnInfo1222 {
	this := base_BTTableColumnInfo1222{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *base_BTTableColumnInfo1222) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableColumnInfo1222) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *base_BTTableColumnInfo1222) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *base_BTTableColumnInfo1222) SetId(v string) {
	o.Id = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *base_BTTableColumnInfo1222) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableColumnInfo1222) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *base_BTTableColumnInfo1222) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *base_BTTableColumnInfo1222) SetNodeId(v string) {
	o.NodeId = &v
}

// GetSpecification returns the Specification field value if set, zero value otherwise.
func (o *base_BTTableColumnInfo1222) GetSpecification() BTTableColumnSpec1967 {
	if o == nil || o.Specification == nil {
		var ret BTTableColumnSpec1967
		return ret
	}
	return *o.Specification
}

// GetSpecificationOk returns a tuple with the Specification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableColumnInfo1222) GetSpecificationOk() (*BTTableColumnSpec1967, bool) {
	if o == nil || o.Specification == nil {
		return nil, false
	}
	return o.Specification, true
}

// HasSpecification returns a boolean if a field has been set.
func (o *base_BTTableColumnInfo1222) HasSpecification() bool {
	if o != nil && o.Specification != nil {
		return true
	}

	return false
}

// SetSpecification gets a reference to the given BTTableColumnSpec1967 and assigns it to the Specification field.
func (o *base_BTTableColumnInfo1222) SetSpecification(v BTTableColumnSpec1967) {
	o.Specification = &v
}

func (o base_BTTableColumnInfo1222) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Specification != nil {
		toSerialize["specification"] = o.Specification
	}
	return json.Marshal(toSerialize)
}

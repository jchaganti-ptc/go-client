/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6524-ea93b01144ea
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTBillOfMaterialsTableRowMetadata1300 struct for BTBillOfMaterialsTableRowMetadata1300
type BTBillOfMaterialsTableRowMetadata1300 struct {
	CrossHighlightDataIfAny *BTTableAssemblyCrossHighlightData2675 `json:"crossHighlightDataIfAny,omitempty"`
	BtType                  *string                                `json:"btType,omitempty"`
	CrossHighlightData      *BTTableAssemblyCrossHighlightData2675 `json:"crossHighlightData,omitempty"`
}

// NewBTBillOfMaterialsTableRowMetadata1300 instantiates a new BTBillOfMaterialsTableRowMetadata1300 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBillOfMaterialsTableRowMetadata1300() *BTBillOfMaterialsTableRowMetadata1300 {
	this := BTBillOfMaterialsTableRowMetadata1300{}
	return &this
}

// NewBTBillOfMaterialsTableRowMetadata1300WithDefaults instantiates a new BTBillOfMaterialsTableRowMetadata1300 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBillOfMaterialsTableRowMetadata1300WithDefaults() *BTBillOfMaterialsTableRowMetadata1300 {
	this := BTBillOfMaterialsTableRowMetadata1300{}
	return &this
}

// GetCrossHighlightDataIfAny returns the CrossHighlightDataIfAny field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRowMetadata1300) GetCrossHighlightDataIfAny() BTTableAssemblyCrossHighlightData2675 {
	if o == nil || o.CrossHighlightDataIfAny == nil {
		var ret BTTableAssemblyCrossHighlightData2675
		return ret
	}
	return *o.CrossHighlightDataIfAny
}

// GetCrossHighlightDataIfAnyOk returns a tuple with the CrossHighlightDataIfAny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRowMetadata1300) GetCrossHighlightDataIfAnyOk() (*BTTableAssemblyCrossHighlightData2675, bool) {
	if o == nil || o.CrossHighlightDataIfAny == nil {
		return nil, false
	}
	return o.CrossHighlightDataIfAny, true
}

// HasCrossHighlightDataIfAny returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRowMetadata1300) HasCrossHighlightDataIfAny() bool {
	if o != nil && o.CrossHighlightDataIfAny != nil {
		return true
	}

	return false
}

// SetCrossHighlightDataIfAny gets a reference to the given BTTableAssemblyCrossHighlightData2675 and assigns it to the CrossHighlightDataIfAny field.
func (o *BTBillOfMaterialsTableRowMetadata1300) SetCrossHighlightDataIfAny(v BTTableAssemblyCrossHighlightData2675) {
	o.CrossHighlightDataIfAny = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRowMetadata1300) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRowMetadata1300) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRowMetadata1300) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTBillOfMaterialsTableRowMetadata1300) SetBtType(v string) {
	o.BtType = &v
}

// GetCrossHighlightData returns the CrossHighlightData field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRowMetadata1300) GetCrossHighlightData() BTTableAssemblyCrossHighlightData2675 {
	if o == nil || o.CrossHighlightData == nil {
		var ret BTTableAssemblyCrossHighlightData2675
		return ret
	}
	return *o.CrossHighlightData
}

// GetCrossHighlightDataOk returns a tuple with the CrossHighlightData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRowMetadata1300) GetCrossHighlightDataOk() (*BTTableAssemblyCrossHighlightData2675, bool) {
	if o == nil || o.CrossHighlightData == nil {
		return nil, false
	}
	return o.CrossHighlightData, true
}

// HasCrossHighlightData returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRowMetadata1300) HasCrossHighlightData() bool {
	if o != nil && o.CrossHighlightData != nil {
		return true
	}

	return false
}

// SetCrossHighlightData gets a reference to the given BTTableAssemblyCrossHighlightData2675 and assigns it to the CrossHighlightData field.
func (o *BTBillOfMaterialsTableRowMetadata1300) SetCrossHighlightData(v BTTableAssemblyCrossHighlightData2675) {
	o.CrossHighlightData = &v
}

func (o BTBillOfMaterialsTableRowMetadata1300) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CrossHighlightDataIfAny != nil {
		toSerialize["crossHighlightDataIfAny"] = o.CrossHighlightDataIfAny
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.CrossHighlightData != nil {
		toSerialize["crossHighlightData"] = o.CrossHighlightData
	}
	return json.Marshal(toSerialize)
}

type NullableBTBillOfMaterialsTableRowMetadata1300 struct {
	value *BTBillOfMaterialsTableRowMetadata1300
	isSet bool
}

func (v NullableBTBillOfMaterialsTableRowMetadata1300) Get() *BTBillOfMaterialsTableRowMetadata1300 {
	return v.value
}

func (v *NullableBTBillOfMaterialsTableRowMetadata1300) Set(val *BTBillOfMaterialsTableRowMetadata1300) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBillOfMaterialsTableRowMetadata1300) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBillOfMaterialsTableRowMetadata1300) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBillOfMaterialsTableRowMetadata1300(val *BTBillOfMaterialsTableRowMetadata1300) *NullableBTBillOfMaterialsTableRowMetadata1300 {
	return &NullableBTBillOfMaterialsTableRowMetadata1300{value: val, isSet: true}
}

func (v NullableBTBillOfMaterialsTableRowMetadata1300) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBillOfMaterialsTableRowMetadata1300) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5901-446ed8116a32
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTComputedPartPropertyConfig struct for BTComputedPartPropertyConfig
type BTComputedPartPropertyConfig struct {
	ComputedPartPropertySpecFunction   *string `json:"computedPartPropertySpecFunction,omitempty"`
	ComputedPartPropertySpecNamespace  *string `json:"computedPartPropertySpecNamespace,omitempty"`
	ComputedPropertyFunctionReturnType *int32  `json:"computedPropertyFunctionReturnType,omitempty"`
	PropertyFunctionDocumentId         *string `json:"propertyFunctionDocumentId,omitempty"`
}

// NewBTComputedPartPropertyConfig instantiates a new BTComputedPartPropertyConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTComputedPartPropertyConfig() *BTComputedPartPropertyConfig {
	this := BTComputedPartPropertyConfig{}
	return &this
}

// NewBTComputedPartPropertyConfigWithDefaults instantiates a new BTComputedPartPropertyConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTComputedPartPropertyConfigWithDefaults() *BTComputedPartPropertyConfig {
	this := BTComputedPartPropertyConfig{}
	return &this
}

// GetComputedPartPropertySpecFunction returns the ComputedPartPropertySpecFunction field value if set, zero value otherwise.
func (o *BTComputedPartPropertyConfig) GetComputedPartPropertySpecFunction() string {
	if o == nil || o.ComputedPartPropertySpecFunction == nil {
		var ret string
		return ret
	}
	return *o.ComputedPartPropertySpecFunction
}

// GetComputedPartPropertySpecFunctionOk returns a tuple with the ComputedPartPropertySpecFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTComputedPartPropertyConfig) GetComputedPartPropertySpecFunctionOk() (*string, bool) {
	if o == nil || o.ComputedPartPropertySpecFunction == nil {
		return nil, false
	}
	return o.ComputedPartPropertySpecFunction, true
}

// HasComputedPartPropertySpecFunction returns a boolean if a field has been set.
func (o *BTComputedPartPropertyConfig) HasComputedPartPropertySpecFunction() bool {
	if o != nil && o.ComputedPartPropertySpecFunction != nil {
		return true
	}

	return false
}

// SetComputedPartPropertySpecFunction gets a reference to the given string and assigns it to the ComputedPartPropertySpecFunction field.
func (o *BTComputedPartPropertyConfig) SetComputedPartPropertySpecFunction(v string) {
	o.ComputedPartPropertySpecFunction = &v
}

// GetComputedPartPropertySpecNamespace returns the ComputedPartPropertySpecNamespace field value if set, zero value otherwise.
func (o *BTComputedPartPropertyConfig) GetComputedPartPropertySpecNamespace() string {
	if o == nil || o.ComputedPartPropertySpecNamespace == nil {
		var ret string
		return ret
	}
	return *o.ComputedPartPropertySpecNamespace
}

// GetComputedPartPropertySpecNamespaceOk returns a tuple with the ComputedPartPropertySpecNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTComputedPartPropertyConfig) GetComputedPartPropertySpecNamespaceOk() (*string, bool) {
	if o == nil || o.ComputedPartPropertySpecNamespace == nil {
		return nil, false
	}
	return o.ComputedPartPropertySpecNamespace, true
}

// HasComputedPartPropertySpecNamespace returns a boolean if a field has been set.
func (o *BTComputedPartPropertyConfig) HasComputedPartPropertySpecNamespace() bool {
	if o != nil && o.ComputedPartPropertySpecNamespace != nil {
		return true
	}

	return false
}

// SetComputedPartPropertySpecNamespace gets a reference to the given string and assigns it to the ComputedPartPropertySpecNamespace field.
func (o *BTComputedPartPropertyConfig) SetComputedPartPropertySpecNamespace(v string) {
	o.ComputedPartPropertySpecNamespace = &v
}

// GetComputedPropertyFunctionReturnType returns the ComputedPropertyFunctionReturnType field value if set, zero value otherwise.
func (o *BTComputedPartPropertyConfig) GetComputedPropertyFunctionReturnType() int32 {
	if o == nil || o.ComputedPropertyFunctionReturnType == nil {
		var ret int32
		return ret
	}
	return *o.ComputedPropertyFunctionReturnType
}

// GetComputedPropertyFunctionReturnTypeOk returns a tuple with the ComputedPropertyFunctionReturnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTComputedPartPropertyConfig) GetComputedPropertyFunctionReturnTypeOk() (*int32, bool) {
	if o == nil || o.ComputedPropertyFunctionReturnType == nil {
		return nil, false
	}
	return o.ComputedPropertyFunctionReturnType, true
}

// HasComputedPropertyFunctionReturnType returns a boolean if a field has been set.
func (o *BTComputedPartPropertyConfig) HasComputedPropertyFunctionReturnType() bool {
	if o != nil && o.ComputedPropertyFunctionReturnType != nil {
		return true
	}

	return false
}

// SetComputedPropertyFunctionReturnType gets a reference to the given int32 and assigns it to the ComputedPropertyFunctionReturnType field.
func (o *BTComputedPartPropertyConfig) SetComputedPropertyFunctionReturnType(v int32) {
	o.ComputedPropertyFunctionReturnType = &v
}

// GetPropertyFunctionDocumentId returns the PropertyFunctionDocumentId field value if set, zero value otherwise.
func (o *BTComputedPartPropertyConfig) GetPropertyFunctionDocumentId() string {
	if o == nil || o.PropertyFunctionDocumentId == nil {
		var ret string
		return ret
	}
	return *o.PropertyFunctionDocumentId
}

// GetPropertyFunctionDocumentIdOk returns a tuple with the PropertyFunctionDocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTComputedPartPropertyConfig) GetPropertyFunctionDocumentIdOk() (*string, bool) {
	if o == nil || o.PropertyFunctionDocumentId == nil {
		return nil, false
	}
	return o.PropertyFunctionDocumentId, true
}

// HasPropertyFunctionDocumentId returns a boolean if a field has been set.
func (o *BTComputedPartPropertyConfig) HasPropertyFunctionDocumentId() bool {
	if o != nil && o.PropertyFunctionDocumentId != nil {
		return true
	}

	return false
}

// SetPropertyFunctionDocumentId gets a reference to the given string and assigns it to the PropertyFunctionDocumentId field.
func (o *BTComputedPartPropertyConfig) SetPropertyFunctionDocumentId(v string) {
	o.PropertyFunctionDocumentId = &v
}

func (o BTComputedPartPropertyConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ComputedPartPropertySpecFunction != nil {
		toSerialize["computedPartPropertySpecFunction"] = o.ComputedPartPropertySpecFunction
	}
	if o.ComputedPartPropertySpecNamespace != nil {
		toSerialize["computedPartPropertySpecNamespace"] = o.ComputedPartPropertySpecNamespace
	}
	if o.ComputedPropertyFunctionReturnType != nil {
		toSerialize["computedPropertyFunctionReturnType"] = o.ComputedPropertyFunctionReturnType
	}
	if o.PropertyFunctionDocumentId != nil {
		toSerialize["propertyFunctionDocumentId"] = o.PropertyFunctionDocumentId
	}
	return json.Marshal(toSerialize)
}

type NullableBTComputedPartPropertyConfig struct {
	value *BTComputedPartPropertyConfig
	isSet bool
}

func (v NullableBTComputedPartPropertyConfig) Get() *BTComputedPartPropertyConfig {
	return v.value
}

func (v *NullableBTComputedPartPropertyConfig) Set(val *BTComputedPartPropertyConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBTComputedPartPropertyConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBTComputedPartPropertyConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTComputedPartPropertyConfig(val *BTComputedPartPropertyConfig) *NullableBTComputedPartPropertyConfig {
	return &NullableBTComputedPartPropertyConfig{value: val, isSet: true}
}

func (v NullableBTComputedPartPropertyConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTComputedPartPropertyConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

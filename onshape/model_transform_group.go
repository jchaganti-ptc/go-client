/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.157.8827-f82e65cdc920
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// TransformGroup struct for TransformGroup
type TransformGroup struct {
	Instances []BTAssemblyInstanceDefinitionParams `json:"instances,omitempty"`
	Transform []float64                            `json:"transform,omitempty"`
}

// NewTransformGroup instantiates a new TransformGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransformGroup() *TransformGroup {
	this := TransformGroup{}
	return &this
}

// NewTransformGroupWithDefaults instantiates a new TransformGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformGroupWithDefaults() *TransformGroup {
	this := TransformGroup{}
	return &this
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *TransformGroup) GetInstances() []BTAssemblyInstanceDefinitionParams {
	if o == nil || o.Instances == nil {
		var ret []BTAssemblyInstanceDefinitionParams
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformGroup) GetInstancesOk() ([]BTAssemblyInstanceDefinitionParams, bool) {
	if o == nil || o.Instances == nil {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *TransformGroup) HasInstances() bool {
	if o != nil && o.Instances != nil {
		return true
	}

	return false
}

// SetInstances gets a reference to the given []BTAssemblyInstanceDefinitionParams and assigns it to the Instances field.
func (o *TransformGroup) SetInstances(v []BTAssemblyInstanceDefinitionParams) {
	o.Instances = v
}

// GetTransform returns the Transform field value if set, zero value otherwise.
func (o *TransformGroup) GetTransform() []float64 {
	if o == nil || o.Transform == nil {
		var ret []float64
		return ret
	}
	return o.Transform
}

// GetTransformOk returns a tuple with the Transform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformGroup) GetTransformOk() ([]float64, bool) {
	if o == nil || o.Transform == nil {
		return nil, false
	}
	return o.Transform, true
}

// HasTransform returns a boolean if a field has been set.
func (o *TransformGroup) HasTransform() bool {
	if o != nil && o.Transform != nil {
		return true
	}

	return false
}

// SetTransform gets a reference to the given []float64 and assigns it to the Transform field.
func (o *TransformGroup) SetTransform(v []float64) {
	o.Transform = v
}

func (o TransformGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Instances != nil {
		toSerialize["instances"] = o.Instances
	}
	if o.Transform != nil {
		toSerialize["transform"] = o.Transform
	}
	return json.Marshal(toSerialize)
}

type NullableTransformGroup struct {
	value *TransformGroup
	isSet bool
}

func (v NullableTransformGroup) Get() *TransformGroup {
	return v.value
}

func (v *NullableTransformGroup) Set(val *TransformGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformGroup(val *TransformGroup) *NullableTransformGroup {
	return &NullableTransformGroup{value: val, isSet: true}
}

func (v NullableTransformGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

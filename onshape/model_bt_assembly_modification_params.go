/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.164.16419-6916b772b99f
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAssemblyModificationParams struct for BTAssemblyModificationParams
type BTAssemblyModificationParams struct {
	DeleteInstances      []string                              `json:"deleteInstances,omitempty"`
	EditDescription      *string                               `json:"editDescription,omitempty"`
	SuppressInstances    []string                              `json:"suppressInstances,omitempty"`
	TransformDefinitions []BTAssemblyTransformDefinitionParams `json:"transformDefinitions,omitempty"`
	UnsuppressInstances  []string                              `json:"unsuppressInstances,omitempty"`
}

// NewBTAssemblyModificationParams instantiates a new BTAssemblyModificationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyModificationParams() *BTAssemblyModificationParams {
	this := BTAssemblyModificationParams{}
	return &this
}

// NewBTAssemblyModificationParamsWithDefaults instantiates a new BTAssemblyModificationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyModificationParamsWithDefaults() *BTAssemblyModificationParams {
	this := BTAssemblyModificationParams{}
	return &this
}

// GetDeleteInstances returns the DeleteInstances field value if set, zero value otherwise.
func (o *BTAssemblyModificationParams) GetDeleteInstances() []string {
	if o == nil || o.DeleteInstances == nil {
		var ret []string
		return ret
	}
	return o.DeleteInstances
}

// GetDeleteInstancesOk returns a tuple with the DeleteInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyModificationParams) GetDeleteInstancesOk() ([]string, bool) {
	if o == nil || o.DeleteInstances == nil {
		return nil, false
	}
	return o.DeleteInstances, true
}

// HasDeleteInstances returns a boolean if a field has been set.
func (o *BTAssemblyModificationParams) HasDeleteInstances() bool {
	if o != nil && o.DeleteInstances != nil {
		return true
	}

	return false
}

// SetDeleteInstances gets a reference to the given []string and assigns it to the DeleteInstances field.
func (o *BTAssemblyModificationParams) SetDeleteInstances(v []string) {
	o.DeleteInstances = v
}

// GetEditDescription returns the EditDescription field value if set, zero value otherwise.
func (o *BTAssemblyModificationParams) GetEditDescription() string {
	if o == nil || o.EditDescription == nil {
		var ret string
		return ret
	}
	return *o.EditDescription
}

// GetEditDescriptionOk returns a tuple with the EditDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyModificationParams) GetEditDescriptionOk() (*string, bool) {
	if o == nil || o.EditDescription == nil {
		return nil, false
	}
	return o.EditDescription, true
}

// HasEditDescription returns a boolean if a field has been set.
func (o *BTAssemblyModificationParams) HasEditDescription() bool {
	if o != nil && o.EditDescription != nil {
		return true
	}

	return false
}

// SetEditDescription gets a reference to the given string and assigns it to the EditDescription field.
func (o *BTAssemblyModificationParams) SetEditDescription(v string) {
	o.EditDescription = &v
}

// GetSuppressInstances returns the SuppressInstances field value if set, zero value otherwise.
func (o *BTAssemblyModificationParams) GetSuppressInstances() []string {
	if o == nil || o.SuppressInstances == nil {
		var ret []string
		return ret
	}
	return o.SuppressInstances
}

// GetSuppressInstancesOk returns a tuple with the SuppressInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyModificationParams) GetSuppressInstancesOk() ([]string, bool) {
	if o == nil || o.SuppressInstances == nil {
		return nil, false
	}
	return o.SuppressInstances, true
}

// HasSuppressInstances returns a boolean if a field has been set.
func (o *BTAssemblyModificationParams) HasSuppressInstances() bool {
	if o != nil && o.SuppressInstances != nil {
		return true
	}

	return false
}

// SetSuppressInstances gets a reference to the given []string and assigns it to the SuppressInstances field.
func (o *BTAssemblyModificationParams) SetSuppressInstances(v []string) {
	o.SuppressInstances = v
}

// GetTransformDefinitions returns the TransformDefinitions field value if set, zero value otherwise.
func (o *BTAssemblyModificationParams) GetTransformDefinitions() []BTAssemblyTransformDefinitionParams {
	if o == nil || o.TransformDefinitions == nil {
		var ret []BTAssemblyTransformDefinitionParams
		return ret
	}
	return o.TransformDefinitions
}

// GetTransformDefinitionsOk returns a tuple with the TransformDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyModificationParams) GetTransformDefinitionsOk() ([]BTAssemblyTransformDefinitionParams, bool) {
	if o == nil || o.TransformDefinitions == nil {
		return nil, false
	}
	return o.TransformDefinitions, true
}

// HasTransformDefinitions returns a boolean if a field has been set.
func (o *BTAssemblyModificationParams) HasTransformDefinitions() bool {
	if o != nil && o.TransformDefinitions != nil {
		return true
	}

	return false
}

// SetTransformDefinitions gets a reference to the given []BTAssemblyTransformDefinitionParams and assigns it to the TransformDefinitions field.
func (o *BTAssemblyModificationParams) SetTransformDefinitions(v []BTAssemblyTransformDefinitionParams) {
	o.TransformDefinitions = v
}

// GetUnsuppressInstances returns the UnsuppressInstances field value if set, zero value otherwise.
func (o *BTAssemblyModificationParams) GetUnsuppressInstances() []string {
	if o == nil || o.UnsuppressInstances == nil {
		var ret []string
		return ret
	}
	return o.UnsuppressInstances
}

// GetUnsuppressInstancesOk returns a tuple with the UnsuppressInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyModificationParams) GetUnsuppressInstancesOk() ([]string, bool) {
	if o == nil || o.UnsuppressInstances == nil {
		return nil, false
	}
	return o.UnsuppressInstances, true
}

// HasUnsuppressInstances returns a boolean if a field has been set.
func (o *BTAssemblyModificationParams) HasUnsuppressInstances() bool {
	if o != nil && o.UnsuppressInstances != nil {
		return true
	}

	return false
}

// SetUnsuppressInstances gets a reference to the given []string and assigns it to the UnsuppressInstances field.
func (o *BTAssemblyModificationParams) SetUnsuppressInstances(v []string) {
	o.UnsuppressInstances = v
}

func (o BTAssemblyModificationParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteInstances != nil {
		toSerialize["deleteInstances"] = o.DeleteInstances
	}
	if o.EditDescription != nil {
		toSerialize["editDescription"] = o.EditDescription
	}
	if o.SuppressInstances != nil {
		toSerialize["suppressInstances"] = o.SuppressInstances
	}
	if o.TransformDefinitions != nil {
		toSerialize["transformDefinitions"] = o.TransformDefinitions
	}
	if o.UnsuppressInstances != nil {
		toSerialize["unsuppressInstances"] = o.UnsuppressInstances
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyModificationParams struct {
	value *BTAssemblyModificationParams
	isSet bool
}

func (v NullableBTAssemblyModificationParams) Get() *BTAssemblyModificationParams {
	return v.value
}

func (v *NullableBTAssemblyModificationParams) Set(val *BTAssemblyModificationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyModificationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyModificationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyModificationParams(val *BTAssemblyModificationParams) *NullableBTAssemblyModificationParams {
	return &NullableBTAssemblyModificationParams{value: val, isSet: true}
}

func (v NullableBTAssemblyModificationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyModificationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

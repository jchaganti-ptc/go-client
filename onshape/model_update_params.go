/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6621-03f431879e4b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// UpdateParams struct for UpdateParams
type UpdateParams struct {
	FromReference *BTUniqueDocumentItemParams `json:"fromReference,omitempty"`
	IdsToUpdate   []string                    `json:"idsToUpdate,omitempty"`
	ToReference   *BTUniqueDocumentItemParams `json:"toReference,omitempty"`
}

// NewUpdateParams instantiates a new UpdateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateParams() *UpdateParams {
	this := UpdateParams{}
	return &this
}

// NewUpdateParamsWithDefaults instantiates a new UpdateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateParamsWithDefaults() *UpdateParams {
	this := UpdateParams{}
	return &this
}

// GetFromReference returns the FromReference field value if set, zero value otherwise.
func (o *UpdateParams) GetFromReference() BTUniqueDocumentItemParams {
	if o == nil || o.FromReference == nil {
		var ret BTUniqueDocumentItemParams
		return ret
	}
	return *o.FromReference
}

// GetFromReferenceOk returns a tuple with the FromReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateParams) GetFromReferenceOk() (*BTUniqueDocumentItemParams, bool) {
	if o == nil || o.FromReference == nil {
		return nil, false
	}
	return o.FromReference, true
}

// HasFromReference returns a boolean if a field has been set.
func (o *UpdateParams) HasFromReference() bool {
	if o != nil && o.FromReference != nil {
		return true
	}

	return false
}

// SetFromReference gets a reference to the given BTUniqueDocumentItemParams and assigns it to the FromReference field.
func (o *UpdateParams) SetFromReference(v BTUniqueDocumentItemParams) {
	o.FromReference = &v
}

// GetIdsToUpdate returns the IdsToUpdate field value if set, zero value otherwise.
func (o *UpdateParams) GetIdsToUpdate() []string {
	if o == nil || o.IdsToUpdate == nil {
		var ret []string
		return ret
	}
	return o.IdsToUpdate
}

// GetIdsToUpdateOk returns a tuple with the IdsToUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateParams) GetIdsToUpdateOk() ([]string, bool) {
	if o == nil || o.IdsToUpdate == nil {
		return nil, false
	}
	return o.IdsToUpdate, true
}

// HasIdsToUpdate returns a boolean if a field has been set.
func (o *UpdateParams) HasIdsToUpdate() bool {
	if o != nil && o.IdsToUpdate != nil {
		return true
	}

	return false
}

// SetIdsToUpdate gets a reference to the given []string and assigns it to the IdsToUpdate field.
func (o *UpdateParams) SetIdsToUpdate(v []string) {
	o.IdsToUpdate = v
}

// GetToReference returns the ToReference field value if set, zero value otherwise.
func (o *UpdateParams) GetToReference() BTUniqueDocumentItemParams {
	if o == nil || o.ToReference == nil {
		var ret BTUniqueDocumentItemParams
		return ret
	}
	return *o.ToReference
}

// GetToReferenceOk returns a tuple with the ToReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateParams) GetToReferenceOk() (*BTUniqueDocumentItemParams, bool) {
	if o == nil || o.ToReference == nil {
		return nil, false
	}
	return o.ToReference, true
}

// HasToReference returns a boolean if a field has been set.
func (o *UpdateParams) HasToReference() bool {
	if o != nil && o.ToReference != nil {
		return true
	}

	return false
}

// SetToReference gets a reference to the given BTUniqueDocumentItemParams and assigns it to the ToReference field.
func (o *UpdateParams) SetToReference(v BTUniqueDocumentItemParams) {
	o.ToReference = &v
}

func (o UpdateParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FromReference != nil {
		toSerialize["fromReference"] = o.FromReference
	}
	if o.IdsToUpdate != nil {
		toSerialize["idsToUpdate"] = o.IdsToUpdate
	}
	if o.ToReference != nil {
		toSerialize["toReference"] = o.ToReference
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateParams struct {
	value *UpdateParams
	isSet bool
}

func (v NullableUpdateParams) Get() *UpdateParams {
	return v.value
}

func (v *NullableUpdateParams) Set(val *UpdateParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateParams(val *UpdateParams) *NullableUpdateParams {
	return &NullableUpdateParams{value: val, isSet: true}
}

func (v NullableUpdateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

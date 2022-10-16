/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.6926-35d1d6168263
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMaterialParams struct for BTMaterialParams
type BTMaterialParams struct {
	DisplayName      *string                         `json:"displayName,omitempty"`
	Id               *string                         `json:"id,omitempty"`
	LibraryName      *string                         `json:"libraryName,omitempty"`
	LibraryReference *BTExternalElementReferenceInfo `json:"libraryReference,omitempty"`
	Properties       []BTMaterialPropertyParams      `json:"properties,omitempty"`
}

// NewBTMaterialParams instantiates a new BTMaterialParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMaterialParams() *BTMaterialParams {
	this := BTMaterialParams{}
	return &this
}

// NewBTMaterialParamsWithDefaults instantiates a new BTMaterialParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMaterialParamsWithDefaults() *BTMaterialParams {
	this := BTMaterialParams{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *BTMaterialParams) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialParams) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *BTMaterialParams) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *BTMaterialParams) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTMaterialParams) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialParams) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTMaterialParams) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTMaterialParams) SetId(v string) {
	o.Id = &v
}

// GetLibraryName returns the LibraryName field value if set, zero value otherwise.
func (o *BTMaterialParams) GetLibraryName() string {
	if o == nil || o.LibraryName == nil {
		var ret string
		return ret
	}
	return *o.LibraryName
}

// GetLibraryNameOk returns a tuple with the LibraryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialParams) GetLibraryNameOk() (*string, bool) {
	if o == nil || o.LibraryName == nil {
		return nil, false
	}
	return o.LibraryName, true
}

// HasLibraryName returns a boolean if a field has been set.
func (o *BTMaterialParams) HasLibraryName() bool {
	if o != nil && o.LibraryName != nil {
		return true
	}

	return false
}

// SetLibraryName gets a reference to the given string and assigns it to the LibraryName field.
func (o *BTMaterialParams) SetLibraryName(v string) {
	o.LibraryName = &v
}

// GetLibraryReference returns the LibraryReference field value if set, zero value otherwise.
func (o *BTMaterialParams) GetLibraryReference() BTExternalElementReferenceInfo {
	if o == nil || o.LibraryReference == nil {
		var ret BTExternalElementReferenceInfo
		return ret
	}
	return *o.LibraryReference
}

// GetLibraryReferenceOk returns a tuple with the LibraryReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialParams) GetLibraryReferenceOk() (*BTExternalElementReferenceInfo, bool) {
	if o == nil || o.LibraryReference == nil {
		return nil, false
	}
	return o.LibraryReference, true
}

// HasLibraryReference returns a boolean if a field has been set.
func (o *BTMaterialParams) HasLibraryReference() bool {
	if o != nil && o.LibraryReference != nil {
		return true
	}

	return false
}

// SetLibraryReference gets a reference to the given BTExternalElementReferenceInfo and assigns it to the LibraryReference field.
func (o *BTMaterialParams) SetLibraryReference(v BTExternalElementReferenceInfo) {
	o.LibraryReference = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *BTMaterialParams) GetProperties() []BTMaterialPropertyParams {
	if o == nil || o.Properties == nil {
		var ret []BTMaterialPropertyParams
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMaterialParams) GetPropertiesOk() ([]BTMaterialPropertyParams, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *BTMaterialParams) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []BTMaterialPropertyParams and assigns it to the Properties field.
func (o *BTMaterialParams) SetProperties(v []BTMaterialPropertyParams) {
	o.Properties = v
}

func (o BTMaterialParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LibraryName != nil {
		toSerialize["libraryName"] = o.LibraryName
	}
	if o.LibraryReference != nil {
		toSerialize["libraryReference"] = o.LibraryReference
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableBTMaterialParams struct {
	value *BTMaterialParams
	isSet bool
}

func (v NullableBTMaterialParams) Get() *BTMaterialParams {
	return v.value
}

func (v *NullableBTMaterialParams) Set(val *BTMaterialParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMaterialParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMaterialParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMaterialParams(val *BTMaterialParams) *NullableBTMaterialParams {
	return &NullableBTMaterialParams{value: val, isSet: true}
}

func (v NullableBTMaterialParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMaterialParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.156.7192-0ed4c121c7d8
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTBillOfMaterialsElementInfo struct for BTBillOfMaterialsElementInfo
type BTBillOfMaterialsElementInfo struct {
	Configuration   *string                           `json:"configuration,omitempty"`
	HeaderIdToValue map[string]map[string]interface{} `json:"headerIdToValue,omitempty"`
	Href            *string                           `json:"href,omitempty"`
	Id              *string                           `json:"id,omitempty"`
	Name            *string                           `json:"name,omitempty"`
	ViewRef         *string                           `json:"viewRef,omitempty"`
}

// NewBTBillOfMaterialsElementInfo instantiates a new BTBillOfMaterialsElementInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBillOfMaterialsElementInfo() *BTBillOfMaterialsElementInfo {
	this := BTBillOfMaterialsElementInfo{}
	return &this
}

// NewBTBillOfMaterialsElementInfoWithDefaults instantiates a new BTBillOfMaterialsElementInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBillOfMaterialsElementInfoWithDefaults() *BTBillOfMaterialsElementInfo {
	this := BTBillOfMaterialsElementInfo{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *BTBillOfMaterialsElementInfo) GetConfiguration() string {
	if o == nil || o.Configuration == nil {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsElementInfo) GetConfigurationOk() (*string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *BTBillOfMaterialsElementInfo) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *BTBillOfMaterialsElementInfo) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetHeaderIdToValue returns the HeaderIdToValue field value if set, zero value otherwise.
func (o *BTBillOfMaterialsElementInfo) GetHeaderIdToValue() map[string]map[string]interface{} {
	if o == nil || o.HeaderIdToValue == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.HeaderIdToValue
}

// GetHeaderIdToValueOk returns a tuple with the HeaderIdToValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsElementInfo) GetHeaderIdToValueOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.HeaderIdToValue == nil {
		return nil, false
	}
	return o.HeaderIdToValue, true
}

// HasHeaderIdToValue returns a boolean if a field has been set.
func (o *BTBillOfMaterialsElementInfo) HasHeaderIdToValue() bool {
	if o != nil && o.HeaderIdToValue != nil {
		return true
	}

	return false
}

// SetHeaderIdToValue gets a reference to the given map[string]map[string]interface{} and assigns it to the HeaderIdToValue field.
func (o *BTBillOfMaterialsElementInfo) SetHeaderIdToValue(v map[string]map[string]interface{}) {
	o.HeaderIdToValue = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTBillOfMaterialsElementInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsElementInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTBillOfMaterialsElementInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTBillOfMaterialsElementInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTBillOfMaterialsElementInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsElementInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsElementInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTBillOfMaterialsElementInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTBillOfMaterialsElementInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsElementInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTBillOfMaterialsElementInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTBillOfMaterialsElementInfo) SetName(v string) {
	o.Name = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTBillOfMaterialsElementInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsElementInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTBillOfMaterialsElementInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTBillOfMaterialsElementInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTBillOfMaterialsElementInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.HeaderIdToValue != nil {
		toSerialize["headerIdToValue"] = o.HeaderIdToValue
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTBillOfMaterialsElementInfo struct {
	value *BTBillOfMaterialsElementInfo
	isSet bool
}

func (v NullableBTBillOfMaterialsElementInfo) Get() *BTBillOfMaterialsElementInfo {
	return v.value
}

func (v *NullableBTBillOfMaterialsElementInfo) Set(val *BTBillOfMaterialsElementInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBillOfMaterialsElementInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBillOfMaterialsElementInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBillOfMaterialsElementInfo(val *BTBillOfMaterialsElementInfo) *NullableBTBillOfMaterialsElementInfo {
	return &NullableBTBillOfMaterialsElementInfo{value: val, isSet: true}
}

func (v NullableBTBillOfMaterialsElementInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBillOfMaterialsElementInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

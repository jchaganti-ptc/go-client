/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.168.21206-13add30fbba2
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTBillOfMaterialsObjectWithPropertiesInfo struct for BTBillOfMaterialsObjectWithPropertiesInfo
type BTBillOfMaterialsObjectWithPropertiesInfo struct {
	HeaderIdToValue map[string]map[string]interface{} `json:"headerIdToValue,omitempty"`
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id *string `json:"id,omitempty"`
	// Name of the resource.
	Name *string `json:"name,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef *string `json:"viewRef,omitempty"`
}

// NewBTBillOfMaterialsObjectWithPropertiesInfo instantiates a new BTBillOfMaterialsObjectWithPropertiesInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBillOfMaterialsObjectWithPropertiesInfo() *BTBillOfMaterialsObjectWithPropertiesInfo {
	this := BTBillOfMaterialsObjectWithPropertiesInfo{}
	return &this
}

// NewBTBillOfMaterialsObjectWithPropertiesInfoWithDefaults instantiates a new BTBillOfMaterialsObjectWithPropertiesInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBillOfMaterialsObjectWithPropertiesInfoWithDefaults() *BTBillOfMaterialsObjectWithPropertiesInfo {
	this := BTBillOfMaterialsObjectWithPropertiesInfo{}
	return &this
}

// GetHeaderIdToValue returns the HeaderIdToValue field value if set, zero value otherwise.
func (o *BTBillOfMaterialsObjectWithPropertiesInfo) GetHeaderIdToValue() map[string]map[string]interface{} {
	if o == nil || o.HeaderIdToValue == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.HeaderIdToValue
}

// GetHeaderIdToValueOk returns a tuple with the HeaderIdToValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsObjectWithPropertiesInfo) GetHeaderIdToValueOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.HeaderIdToValue == nil {
		return nil, false
	}
	return o.HeaderIdToValue, true
}

// HasHeaderIdToValue returns a boolean if a field has been set.
func (o *BTBillOfMaterialsObjectWithPropertiesInfo) HasHeaderIdToValue() bool {
	if o != nil && o.HeaderIdToValue != nil {
		return true
	}

	return false
}

// SetHeaderIdToValue gets a reference to the given map[string]map[string]interface{} and assigns it to the HeaderIdToValue field.
func (o *BTBillOfMaterialsObjectWithPropertiesInfo) SetHeaderIdToValue(v map[string]map[string]interface{}) {
	o.HeaderIdToValue = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTBillOfMaterialsObjectWithPropertiesInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsObjectWithPropertiesInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTBillOfMaterialsObjectWithPropertiesInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTBillOfMaterialsObjectWithPropertiesInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTBillOfMaterialsObjectWithPropertiesInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsObjectWithPropertiesInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTBillOfMaterialsObjectWithPropertiesInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTBillOfMaterialsObjectWithPropertiesInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTBillOfMaterialsObjectWithPropertiesInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsObjectWithPropertiesInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTBillOfMaterialsObjectWithPropertiesInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTBillOfMaterialsObjectWithPropertiesInfo) SetName(v string) {
	o.Name = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTBillOfMaterialsObjectWithPropertiesInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsObjectWithPropertiesInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTBillOfMaterialsObjectWithPropertiesInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTBillOfMaterialsObjectWithPropertiesInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTBillOfMaterialsObjectWithPropertiesInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableBTBillOfMaterialsObjectWithPropertiesInfo struct {
	value *BTBillOfMaterialsObjectWithPropertiesInfo
	isSet bool
}

func (v NullableBTBillOfMaterialsObjectWithPropertiesInfo) Get() *BTBillOfMaterialsObjectWithPropertiesInfo {
	return v.value
}

func (v *NullableBTBillOfMaterialsObjectWithPropertiesInfo) Set(val *BTBillOfMaterialsObjectWithPropertiesInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBillOfMaterialsObjectWithPropertiesInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBillOfMaterialsObjectWithPropertiesInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBillOfMaterialsObjectWithPropertiesInfo(val *BTBillOfMaterialsObjectWithPropertiesInfo) *NullableBTBillOfMaterialsObjectWithPropertiesInfo {
	return &NullableBTBillOfMaterialsObjectWithPropertiesInfo{value: val, isSet: true}
}

func (v NullableBTBillOfMaterialsObjectWithPropertiesInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBillOfMaterialsObjectWithPropertiesInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

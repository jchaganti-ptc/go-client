/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.27881-5dbbe8053cdf
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTUserBasicSummaryInfo struct for BTUserBasicSummaryInfo
type BTUserBasicSummaryInfo struct {
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id *string `json:"id,omitempty"`
	// Name of the resource.
	Name *string `json:"name,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef  *string `json:"viewRef,omitempty"`
	Image    *string `json:"image,omitempty"`
	State    *int32  `json:"state,omitempty"`
	JsonType string  `json:"jsonType"`
}

// NewBTUserBasicSummaryInfo instantiates a new BTUserBasicSummaryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUserBasicSummaryInfo(jsonType string) *BTUserBasicSummaryInfo {
	this := BTUserBasicSummaryInfo{}
	this.JsonType = jsonType
	return &this
}

// NewBTUserBasicSummaryInfoWithDefaults instantiates a new BTUserBasicSummaryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUserBasicSummaryInfoWithDefaults() *BTUserBasicSummaryInfo {
	this := BTUserBasicSummaryInfo{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTUserBasicSummaryInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserBasicSummaryInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTUserBasicSummaryInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTUserBasicSummaryInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTUserBasicSummaryInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserBasicSummaryInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTUserBasicSummaryInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTUserBasicSummaryInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTUserBasicSummaryInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserBasicSummaryInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTUserBasicSummaryInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTUserBasicSummaryInfo) SetName(v string) {
	o.Name = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTUserBasicSummaryInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserBasicSummaryInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTUserBasicSummaryInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTUserBasicSummaryInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *BTUserBasicSummaryInfo) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserBasicSummaryInfo) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *BTUserBasicSummaryInfo) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *BTUserBasicSummaryInfo) SetImage(v string) {
	o.Image = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BTUserBasicSummaryInfo) GetState() int32 {
	if o == nil || o.State == nil {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserBasicSummaryInfo) GetStateOk() (*int32, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BTUserBasicSummaryInfo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *BTUserBasicSummaryInfo) SetState(v int32) {
	o.State = &v
}

// GetJsonType returns the JsonType field value
func (o *BTUserBasicSummaryInfo) GetJsonType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JsonType
}

// GetJsonTypeOk returns a tuple with the JsonType field value
// and a boolean to check if the value has been set.
func (o *BTUserBasicSummaryInfo) GetJsonTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JsonType, true
}

// SetJsonType sets field value
func (o *BTUserBasicSummaryInfo) SetJsonType(v string) {
	o.JsonType = v
}

func (o BTUserBasicSummaryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["jsonType"] = o.JsonType
	}
	return json.Marshal(toSerialize)
}

type NullableBTUserBasicSummaryInfo struct {
	value *BTUserBasicSummaryInfo
	isSet bool
}

func (v NullableBTUserBasicSummaryInfo) Get() *BTUserBasicSummaryInfo {
	return v.value
}

func (v *NullableBTUserBasicSummaryInfo) Set(val *BTUserBasicSummaryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUserBasicSummaryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUserBasicSummaryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUserBasicSummaryInfo(val *BTUserBasicSummaryInfo) *NullableBTUserBasicSummaryInfo {
	return &NullableBTUserBasicSummaryInfo{value: val, isSet: true}
}

func (v NullableBTUserBasicSummaryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUserBasicSummaryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

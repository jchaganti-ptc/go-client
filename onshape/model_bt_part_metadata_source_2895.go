/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.169.22266-e2d421ffb3ea
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPartMetadataSource2895 struct for BTPartMetadataSource2895
type BTPartMetadataSource2895 struct {
	BtType     *string                `json:"btType,omitempty"`
	SourceId   *string                `json:"sourceId,omitempty"`
	SourceType *GBTMetadataSourceType `json:"sourceType,omitempty"`
}

// NewBTPartMetadataSource2895 instantiates a new BTPartMetadataSource2895 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPartMetadataSource2895() *BTPartMetadataSource2895 {
	this := BTPartMetadataSource2895{}
	return &this
}

// NewBTPartMetadataSource2895WithDefaults instantiates a new BTPartMetadataSource2895 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPartMetadataSource2895WithDefaults() *BTPartMetadataSource2895 {
	this := BTPartMetadataSource2895{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPartMetadataSource2895) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataSource2895) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPartMetadataSource2895) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPartMetadataSource2895) SetBtType(v string) {
	o.BtType = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *BTPartMetadataSource2895) GetSourceId() string {
	if o == nil || o.SourceId == nil {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataSource2895) GetSourceIdOk() (*string, bool) {
	if o == nil || o.SourceId == nil {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *BTPartMetadataSource2895) HasSourceId() bool {
	if o != nil && o.SourceId != nil {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *BTPartMetadataSource2895) SetSourceId(v string) {
	o.SourceId = &v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *BTPartMetadataSource2895) GetSourceType() GBTMetadataSourceType {
	if o == nil || o.SourceType == nil {
		var ret GBTMetadataSourceType
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartMetadataSource2895) GetSourceTypeOk() (*GBTMetadataSourceType, bool) {
	if o == nil || o.SourceType == nil {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *BTPartMetadataSource2895) HasSourceType() bool {
	if o != nil && o.SourceType != nil {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given GBTMetadataSourceType and assigns it to the SourceType field.
func (o *BTPartMetadataSource2895) SetSourceType(v GBTMetadataSourceType) {
	o.SourceType = &v
}

func (o BTPartMetadataSource2895) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.SourceId != nil {
		toSerialize["sourceId"] = o.SourceId
	}
	if o.SourceType != nil {
		toSerialize["sourceType"] = o.SourceType
	}
	return json.Marshal(toSerialize)
}

type NullableBTPartMetadataSource2895 struct {
	value *BTPartMetadataSource2895
	isSet bool
}

func (v NullableBTPartMetadataSource2895) Get() *BTPartMetadataSource2895 {
	return v.value
}

func (v *NullableBTPartMetadataSource2895) Set(val *BTPartMetadataSource2895) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPartMetadataSource2895) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPartMetadataSource2895) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPartMetadataSource2895(val *BTPartMetadataSource2895) *NullableBTPartMetadataSource2895 {
	return &NullableBTPartMetadataSource2895{value: val, isSet: true}
}

func (v NullableBTPartMetadataSource2895) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPartMetadataSource2895) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

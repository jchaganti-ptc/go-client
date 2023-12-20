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

// ExternalDocumentation struct for ExternalDocumentation
type ExternalDocumentation struct {
	Description *string                           `json:"description,omitempty"`
	Extensions  map[string]map[string]interface{} `json:"extensions,omitempty"`
	Url         *string                           `json:"url,omitempty"`
}

// NewExternalDocumentation instantiates a new ExternalDocumentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalDocumentation() *ExternalDocumentation {
	this := ExternalDocumentation{}
	return &this
}

// NewExternalDocumentationWithDefaults instantiates a new ExternalDocumentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalDocumentationWithDefaults() *ExternalDocumentation {
	this := ExternalDocumentation{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExternalDocumentation) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalDocumentation) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExternalDocumentation) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExternalDocumentation) SetDescription(v string) {
	o.Description = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *ExternalDocumentation) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalDocumentation) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *ExternalDocumentation) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *ExternalDocumentation) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ExternalDocumentation) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalDocumentation) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ExternalDocumentation) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ExternalDocumentation) SetUrl(v string) {
	o.Url = &v
}

func (o ExternalDocumentation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableExternalDocumentation struct {
	value *ExternalDocumentation
	isSet bool
}

func (v NullableExternalDocumentation) Get() *ExternalDocumentation {
	return v.value
}

func (v *NullableExternalDocumentation) Set(val *ExternalDocumentation) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalDocumentation) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalDocumentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalDocumentation(val *ExternalDocumentation) *NullableExternalDocumentation {
	return &NullableExternalDocumentation{value: val, isSet: true}
}

func (v NullableExternalDocumentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalDocumentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

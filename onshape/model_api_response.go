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

// ApiResponse struct for ApiResponse
type ApiResponse struct {
	Content     *ApiResponseContent               `json:"content,omitempty"`
	Description *string                           `json:"description,omitempty"`
	Extensions  map[string]map[string]interface{} `json:"extensions,omitempty"`
	Getref      *string                           `json:"get$ref,omitempty"`
	Headers     *map[string]Header                `json:"headers,omitempty"`
	Links       *map[string]Link                  `json:"links,omitempty"`
}

// NewApiResponse instantiates a new ApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiResponse() *ApiResponse {
	this := ApiResponse{}
	return &this
}

// NewApiResponseWithDefaults instantiates a new ApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiResponseWithDefaults() *ApiResponse {
	this := ApiResponse{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ApiResponse) GetContent() ApiResponseContent {
	if o == nil || o.Content == nil {
		var ret ApiResponseContent
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse) GetContentOk() (*ApiResponseContent, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ApiResponse) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given ApiResponseContent and assigns it to the Content field.
func (o *ApiResponse) SetContent(v ApiResponseContent) {
	o.Content = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiResponse) SetDescription(v string) {
	o.Description = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *ApiResponse) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *ApiResponse) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *ApiResponse) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetGetref returns the Getref field value if set, zero value otherwise.
func (o *ApiResponse) GetGetref() string {
	if o == nil || o.Getref == nil {
		var ret string
		return ret
	}
	return *o.Getref
}

// GetGetrefOk returns a tuple with the Getref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse) GetGetrefOk() (*string, bool) {
	if o == nil || o.Getref == nil {
		return nil, false
	}
	return o.Getref, true
}

// HasGetref returns a boolean if a field has been set.
func (o *ApiResponse) HasGetref() bool {
	if o != nil && o.Getref != nil {
		return true
	}

	return false
}

// SetGetref gets a reference to the given string and assigns it to the Getref field.
func (o *ApiResponse) SetGetref(v string) {
	o.Getref = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *ApiResponse) GetHeaders() map[string]Header {
	if o == nil || o.Headers == nil {
		var ret map[string]Header
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse) GetHeadersOk() (*map[string]Header, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *ApiResponse) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]Header and assigns it to the Headers field.
func (o *ApiResponse) SetHeaders(v map[string]Header) {
	o.Headers = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApiResponse) GetLinks() map[string]Link {
	if o == nil || o.Links == nil {
		var ret map[string]Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse) GetLinksOk() (*map[string]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApiResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]Link and assigns it to the Links field.
func (o *ApiResponse) SetLinks(v map[string]Link) {
	o.Links = &v
}

func (o ApiResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Getref != nil {
		toSerialize["get$ref"] = o.Getref
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableApiResponse struct {
	value *ApiResponse
	isSet bool
}

func (v NullableApiResponse) Get() *ApiResponse {
	return v.value
}

func (v *NullableApiResponse) Set(val *ApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiResponse(val *ApiResponse) *NullableApiResponse {
	return &NullableApiResponse{value: val, isSet: true}
}

func (v NullableApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

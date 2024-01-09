/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.28658-06d4d4923fc7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// Link struct for Link
type Link struct {
	Description  *string                           `json:"description,omitempty"`
	Extensions   map[string]map[string]interface{} `json:"extensions,omitempty"`
	Getref       *string                           `json:"get$ref,omitempty"`
	Headers      *map[string]Header                `json:"headers,omitempty"`
	OperationId  *string                           `json:"operationId,omitempty"`
	OperationRef *string                           `json:"operationRef,omitempty"`
	Parameters   *map[string]string                `json:"parameters,omitempty"`
	RequestBody  *map[string]interface{}           `json:"requestBody,omitempty"`
	Server       *Server                           `json:"server,omitempty"`
}

// NewLink instantiates a new Link object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLink() *Link {
	this := Link{}
	return &this
}

// NewLinkWithDefaults instantiates a new Link object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkWithDefaults() *Link {
	this := Link{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Link) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Link) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Link) SetDescription(v string) {
	o.Description = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Link) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Link) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *Link) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetGetref returns the Getref field value if set, zero value otherwise.
func (o *Link) GetGetref() string {
	if o == nil || o.Getref == nil {
		var ret string
		return ret
	}
	return *o.Getref
}

// GetGetrefOk returns a tuple with the Getref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetGetrefOk() (*string, bool) {
	if o == nil || o.Getref == nil {
		return nil, false
	}
	return o.Getref, true
}

// HasGetref returns a boolean if a field has been set.
func (o *Link) HasGetref() bool {
	if o != nil && o.Getref != nil {
		return true
	}

	return false
}

// SetGetref gets a reference to the given string and assigns it to the Getref field.
func (o *Link) SetGetref(v string) {
	o.Getref = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *Link) GetHeaders() map[string]Header {
	if o == nil || o.Headers == nil {
		var ret map[string]Header
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetHeadersOk() (*map[string]Header, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *Link) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]Header and assigns it to the Headers field.
func (o *Link) SetHeaders(v map[string]Header) {
	o.Headers = &v
}

// GetOperationId returns the OperationId field value if set, zero value otherwise.
func (o *Link) GetOperationId() string {
	if o == nil || o.OperationId == nil {
		var ret string
		return ret
	}
	return *o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetOperationIdOk() (*string, bool) {
	if o == nil || o.OperationId == nil {
		return nil, false
	}
	return o.OperationId, true
}

// HasOperationId returns a boolean if a field has been set.
func (o *Link) HasOperationId() bool {
	if o != nil && o.OperationId != nil {
		return true
	}

	return false
}

// SetOperationId gets a reference to the given string and assigns it to the OperationId field.
func (o *Link) SetOperationId(v string) {
	o.OperationId = &v
}

// GetOperationRef returns the OperationRef field value if set, zero value otherwise.
func (o *Link) GetOperationRef() string {
	if o == nil || o.OperationRef == nil {
		var ret string
		return ret
	}
	return *o.OperationRef
}

// GetOperationRefOk returns a tuple with the OperationRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetOperationRefOk() (*string, bool) {
	if o == nil || o.OperationRef == nil {
		return nil, false
	}
	return o.OperationRef, true
}

// HasOperationRef returns a boolean if a field has been set.
func (o *Link) HasOperationRef() bool {
	if o != nil && o.OperationRef != nil {
		return true
	}

	return false
}

// SetOperationRef gets a reference to the given string and assigns it to the OperationRef field.
func (o *Link) SetOperationRef(v string) {
	o.OperationRef = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *Link) GetParameters() map[string]string {
	if o == nil || o.Parameters == nil {
		var ret map[string]string
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetParametersOk() (*map[string]string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *Link) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *Link) SetParameters(v map[string]string) {
	o.Parameters = &v
}

// GetRequestBody returns the RequestBody field value if set, zero value otherwise.
func (o *Link) GetRequestBody() map[string]interface{} {
	if o == nil || o.RequestBody == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.RequestBody
}

// GetRequestBodyOk returns a tuple with the RequestBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetRequestBodyOk() (*map[string]interface{}, bool) {
	if o == nil || o.RequestBody == nil {
		return nil, false
	}
	return o.RequestBody, true
}

// HasRequestBody returns a boolean if a field has been set.
func (o *Link) HasRequestBody() bool {
	if o != nil && o.RequestBody != nil {
		return true
	}

	return false
}

// SetRequestBody gets a reference to the given map[string]interface{} and assigns it to the RequestBody field.
func (o *Link) SetRequestBody(v map[string]interface{}) {
	o.RequestBody = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *Link) GetServer() Server {
	if o == nil || o.Server == nil {
		var ret Server
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetServerOk() (*Server, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *Link) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given Server and assigns it to the Server field.
func (o *Link) SetServer(v Server) {
	o.Server = &v
}

func (o Link) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.OperationId != nil {
		toSerialize["operationId"] = o.OperationId
	}
	if o.OperationRef != nil {
		toSerialize["operationRef"] = o.OperationRef
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.RequestBody != nil {
		toSerialize["requestBody"] = o.RequestBody
	}
	if o.Server != nil {
		toSerialize["server"] = o.Server
	}
	return json.Marshal(toSerialize)
}

type NullableLink struct {
	value *Link
	isSet bool
}

func (v NullableLink) Get() *Link {
	return v.value
}

func (v *NullableLink) Set(val *Link) {
	v.value = val
	v.isSet = true
}

func (v NullableLink) IsSet() bool {
	return v.isSet
}

func (v *NullableLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLink(val *Link) *NullableLink {
	return &NullableLink{value: val, isSet: true}
}

func (v NullableLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

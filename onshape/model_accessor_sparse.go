/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.173.26882-0482adeaa8aa
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// AccessorSparse struct for AccessorSparse
type AccessorSparse struct {
	Count      *int32                            `json:"count,omitempty"`
	Extensions map[string]map[string]interface{} `json:"extensions,omitempty"`
	Extras     *map[string]interface{}           `json:"extras,omitempty"`
	Indices    *AccessorSparseIndices            `json:"indices,omitempty"`
	Values     *AccessorSparseValues             `json:"values,omitempty"`
}

// NewAccessorSparse instantiates a new AccessorSparse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessorSparse() *AccessorSparse {
	this := AccessorSparse{}
	return &this
}

// NewAccessorSparseWithDefaults instantiates a new AccessorSparse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessorSparseWithDefaults() *AccessorSparse {
	this := AccessorSparse{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *AccessorSparse) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessorSparse) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *AccessorSparse) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *AccessorSparse) SetCount(v int32) {
	o.Count = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *AccessorSparse) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessorSparse) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *AccessorSparse) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *AccessorSparse) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *AccessorSparse) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessorSparse) GetExtrasOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *AccessorSparse) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *AccessorSparse) SetExtras(v map[string]interface{}) {
	o.Extras = &v
}

// GetIndices returns the Indices field value if set, zero value otherwise.
func (o *AccessorSparse) GetIndices() AccessorSparseIndices {
	if o == nil || o.Indices == nil {
		var ret AccessorSparseIndices
		return ret
	}
	return *o.Indices
}

// GetIndicesOk returns a tuple with the Indices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessorSparse) GetIndicesOk() (*AccessorSparseIndices, bool) {
	if o == nil || o.Indices == nil {
		return nil, false
	}
	return o.Indices, true
}

// HasIndices returns a boolean if a field has been set.
func (o *AccessorSparse) HasIndices() bool {
	if o != nil && o.Indices != nil {
		return true
	}

	return false
}

// SetIndices gets a reference to the given AccessorSparseIndices and assigns it to the Indices field.
func (o *AccessorSparse) SetIndices(v AccessorSparseIndices) {
	o.Indices = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *AccessorSparse) GetValues() AccessorSparseValues {
	if o == nil || o.Values == nil {
		var ret AccessorSparseValues
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessorSparse) GetValuesOk() (*AccessorSparseValues, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *AccessorSparse) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given AccessorSparseValues and assigns it to the Values field.
func (o *AccessorSparse) SetValues(v AccessorSparseValues) {
	o.Values = &v
}

func (o AccessorSparse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.Indices != nil {
		toSerialize["indices"] = o.Indices
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableAccessorSparse struct {
	value *AccessorSparse
	isSet bool
}

func (v NullableAccessorSparse) Get() *AccessorSparse {
	return v.value
}

func (v *NullableAccessorSparse) Set(val *AccessorSparse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessorSparse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessorSparse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessorSparse(val *AccessorSparse) *NullableAccessorSparse {
	return &NullableAccessorSparse{value: val, isSet: true}
}

func (v NullableAccessorSparse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessorSparse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

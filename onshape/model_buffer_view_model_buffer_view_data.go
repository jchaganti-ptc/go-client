/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.173.27313-c3b730633f3c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BufferViewModelBufferViewData struct for BufferViewModelBufferViewData
type BufferViewModelBufferViewData struct {
	Short    *int32   `json:"short,omitempty"`
	Char     *string  `json:"char,omitempty"`
	Int      *int32   `json:"int,omitempty"`
	Long     *int64   `json:"long,omitempty"`
	Float    *float32 `json:"float,omitempty"`
	Double   *float64 `json:"double,omitempty"`
	Direct   *bool    `json:"direct,omitempty"`
	ReadOnly *bool    `json:"readOnly,omitempty"`
}

// NewBufferViewModelBufferViewData instantiates a new BufferViewModelBufferViewData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBufferViewModelBufferViewData() *BufferViewModelBufferViewData {
	this := BufferViewModelBufferViewData{}
	return &this
}

// NewBufferViewModelBufferViewDataWithDefaults instantiates a new BufferViewModelBufferViewData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBufferViewModelBufferViewDataWithDefaults() *BufferViewModelBufferViewData {
	this := BufferViewModelBufferViewData{}
	return &this
}

// GetShort returns the Short field value if set, zero value otherwise.
func (o *BufferViewModelBufferViewData) GetShort() int32 {
	if o == nil || o.Short == nil {
		var ret int32
		return ret
	}
	return *o.Short
}

// GetShortOk returns a tuple with the Short field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferViewModelBufferViewData) GetShortOk() (*int32, bool) {
	if o == nil || o.Short == nil {
		return nil, false
	}
	return o.Short, true
}

// HasShort returns a boolean if a field has been set.
func (o *BufferViewModelBufferViewData) HasShort() bool {
	if o != nil && o.Short != nil {
		return true
	}

	return false
}

// SetShort gets a reference to the given int32 and assigns it to the Short field.
func (o *BufferViewModelBufferViewData) SetShort(v int32) {
	o.Short = &v
}

// GetChar returns the Char field value if set, zero value otherwise.
func (o *BufferViewModelBufferViewData) GetChar() string {
	if o == nil || o.Char == nil {
		var ret string
		return ret
	}
	return *o.Char
}

// GetCharOk returns a tuple with the Char field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferViewModelBufferViewData) GetCharOk() (*string, bool) {
	if o == nil || o.Char == nil {
		return nil, false
	}
	return o.Char, true
}

// HasChar returns a boolean if a field has been set.
func (o *BufferViewModelBufferViewData) HasChar() bool {
	if o != nil && o.Char != nil {
		return true
	}

	return false
}

// SetChar gets a reference to the given string and assigns it to the Char field.
func (o *BufferViewModelBufferViewData) SetChar(v string) {
	o.Char = &v
}

// GetInt returns the Int field value if set, zero value otherwise.
func (o *BufferViewModelBufferViewData) GetInt() int32 {
	if o == nil || o.Int == nil {
		var ret int32
		return ret
	}
	return *o.Int
}

// GetIntOk returns a tuple with the Int field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferViewModelBufferViewData) GetIntOk() (*int32, bool) {
	if o == nil || o.Int == nil {
		return nil, false
	}
	return o.Int, true
}

// HasInt returns a boolean if a field has been set.
func (o *BufferViewModelBufferViewData) HasInt() bool {
	if o != nil && o.Int != nil {
		return true
	}

	return false
}

// SetInt gets a reference to the given int32 and assigns it to the Int field.
func (o *BufferViewModelBufferViewData) SetInt(v int32) {
	o.Int = &v
}

// GetLong returns the Long field value if set, zero value otherwise.
func (o *BufferViewModelBufferViewData) GetLong() int64 {
	if o == nil || o.Long == nil {
		var ret int64
		return ret
	}
	return *o.Long
}

// GetLongOk returns a tuple with the Long field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferViewModelBufferViewData) GetLongOk() (*int64, bool) {
	if o == nil || o.Long == nil {
		return nil, false
	}
	return o.Long, true
}

// HasLong returns a boolean if a field has been set.
func (o *BufferViewModelBufferViewData) HasLong() bool {
	if o != nil && o.Long != nil {
		return true
	}

	return false
}

// SetLong gets a reference to the given int64 and assigns it to the Long field.
func (o *BufferViewModelBufferViewData) SetLong(v int64) {
	o.Long = &v
}

// GetFloat returns the Float field value if set, zero value otherwise.
func (o *BufferViewModelBufferViewData) GetFloat() float32 {
	if o == nil || o.Float == nil {
		var ret float32
		return ret
	}
	return *o.Float
}

// GetFloatOk returns a tuple with the Float field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferViewModelBufferViewData) GetFloatOk() (*float32, bool) {
	if o == nil || o.Float == nil {
		return nil, false
	}
	return o.Float, true
}

// HasFloat returns a boolean if a field has been set.
func (o *BufferViewModelBufferViewData) HasFloat() bool {
	if o != nil && o.Float != nil {
		return true
	}

	return false
}

// SetFloat gets a reference to the given float32 and assigns it to the Float field.
func (o *BufferViewModelBufferViewData) SetFloat(v float32) {
	o.Float = &v
}

// GetDouble returns the Double field value if set, zero value otherwise.
func (o *BufferViewModelBufferViewData) GetDouble() float64 {
	if o == nil || o.Double == nil {
		var ret float64
		return ret
	}
	return *o.Double
}

// GetDoubleOk returns a tuple with the Double field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferViewModelBufferViewData) GetDoubleOk() (*float64, bool) {
	if o == nil || o.Double == nil {
		return nil, false
	}
	return o.Double, true
}

// HasDouble returns a boolean if a field has been set.
func (o *BufferViewModelBufferViewData) HasDouble() bool {
	if o != nil && o.Double != nil {
		return true
	}

	return false
}

// SetDouble gets a reference to the given float64 and assigns it to the Double field.
func (o *BufferViewModelBufferViewData) SetDouble(v float64) {
	o.Double = &v
}

// GetDirect returns the Direct field value if set, zero value otherwise.
func (o *BufferViewModelBufferViewData) GetDirect() bool {
	if o == nil || o.Direct == nil {
		var ret bool
		return ret
	}
	return *o.Direct
}

// GetDirectOk returns a tuple with the Direct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferViewModelBufferViewData) GetDirectOk() (*bool, bool) {
	if o == nil || o.Direct == nil {
		return nil, false
	}
	return o.Direct, true
}

// HasDirect returns a boolean if a field has been set.
func (o *BufferViewModelBufferViewData) HasDirect() bool {
	if o != nil && o.Direct != nil {
		return true
	}

	return false
}

// SetDirect gets a reference to the given bool and assigns it to the Direct field.
func (o *BufferViewModelBufferViewData) SetDirect(v bool) {
	o.Direct = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *BufferViewModelBufferViewData) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferViewModelBufferViewData) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *BufferViewModelBufferViewData) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *BufferViewModelBufferViewData) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

func (o BufferViewModelBufferViewData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Short != nil {
		toSerialize["short"] = o.Short
	}
	if o.Char != nil {
		toSerialize["char"] = o.Char
	}
	if o.Int != nil {
		toSerialize["int"] = o.Int
	}
	if o.Long != nil {
		toSerialize["long"] = o.Long
	}
	if o.Float != nil {
		toSerialize["float"] = o.Float
	}
	if o.Double != nil {
		toSerialize["double"] = o.Double
	}
	if o.Direct != nil {
		toSerialize["direct"] = o.Direct
	}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}
	return json.Marshal(toSerialize)
}

type NullableBufferViewModelBufferViewData struct {
	value *BufferViewModelBufferViewData
	isSet bool
}

func (v NullableBufferViewModelBufferViewData) Get() *BufferViewModelBufferViewData {
	return v.value
}

func (v *NullableBufferViewModelBufferViewData) Set(val *BufferViewModelBufferViewData) {
	v.value = val
	v.isSet = true
}

func (v NullableBufferViewModelBufferViewData) IsSet() bool {
	return v.isSet
}

func (v *NullableBufferViewModelBufferViewData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBufferViewModelBufferViewData(val *BufferViewModelBufferViewData) *NullableBufferViewModelBufferViewData {
	return &NullableBufferViewModelBufferViewData{value: val, isSet: true}
}

func (v NullableBufferViewModelBufferViewData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBufferViewModelBufferViewData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

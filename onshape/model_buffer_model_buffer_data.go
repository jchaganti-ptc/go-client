/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.25652-944cf1af37c9
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BufferModelBufferData struct for BufferModelBufferData
type BufferModelBufferData struct {
	Short    *int32   `json:"short,omitempty"`
	Char     *string  `json:"char,omitempty"`
	Int      *int32   `json:"int,omitempty"`
	Long     *int64   `json:"long,omitempty"`
	Float    *float32 `json:"float,omitempty"`
	Double   *float64 `json:"double,omitempty"`
	Direct   *bool    `json:"direct,omitempty"`
	ReadOnly *bool    `json:"readOnly,omitempty"`
}

// NewBufferModelBufferData instantiates a new BufferModelBufferData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBufferModelBufferData() *BufferModelBufferData {
	this := BufferModelBufferData{}
	return &this
}

// NewBufferModelBufferDataWithDefaults instantiates a new BufferModelBufferData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBufferModelBufferDataWithDefaults() *BufferModelBufferData {
	this := BufferModelBufferData{}
	return &this
}

// GetShort returns the Short field value if set, zero value otherwise.
func (o *BufferModelBufferData) GetShort() int32 {
	if o == nil || o.Short == nil {
		var ret int32
		return ret
	}
	return *o.Short
}

// GetShortOk returns a tuple with the Short field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferModelBufferData) GetShortOk() (*int32, bool) {
	if o == nil || o.Short == nil {
		return nil, false
	}
	return o.Short, true
}

// HasShort returns a boolean if a field has been set.
func (o *BufferModelBufferData) HasShort() bool {
	if o != nil && o.Short != nil {
		return true
	}

	return false
}

// SetShort gets a reference to the given int32 and assigns it to the Short field.
func (o *BufferModelBufferData) SetShort(v int32) {
	o.Short = &v
}

// GetChar returns the Char field value if set, zero value otherwise.
func (o *BufferModelBufferData) GetChar() string {
	if o == nil || o.Char == nil {
		var ret string
		return ret
	}
	return *o.Char
}

// GetCharOk returns a tuple with the Char field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferModelBufferData) GetCharOk() (*string, bool) {
	if o == nil || o.Char == nil {
		return nil, false
	}
	return o.Char, true
}

// HasChar returns a boolean if a field has been set.
func (o *BufferModelBufferData) HasChar() bool {
	if o != nil && o.Char != nil {
		return true
	}

	return false
}

// SetChar gets a reference to the given string and assigns it to the Char field.
func (o *BufferModelBufferData) SetChar(v string) {
	o.Char = &v
}

// GetInt returns the Int field value if set, zero value otherwise.
func (o *BufferModelBufferData) GetInt() int32 {
	if o == nil || o.Int == nil {
		var ret int32
		return ret
	}
	return *o.Int
}

// GetIntOk returns a tuple with the Int field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferModelBufferData) GetIntOk() (*int32, bool) {
	if o == nil || o.Int == nil {
		return nil, false
	}
	return o.Int, true
}

// HasInt returns a boolean if a field has been set.
func (o *BufferModelBufferData) HasInt() bool {
	if o != nil && o.Int != nil {
		return true
	}

	return false
}

// SetInt gets a reference to the given int32 and assigns it to the Int field.
func (o *BufferModelBufferData) SetInt(v int32) {
	o.Int = &v
}

// GetLong returns the Long field value if set, zero value otherwise.
func (o *BufferModelBufferData) GetLong() int64 {
	if o == nil || o.Long == nil {
		var ret int64
		return ret
	}
	return *o.Long
}

// GetLongOk returns a tuple with the Long field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferModelBufferData) GetLongOk() (*int64, bool) {
	if o == nil || o.Long == nil {
		return nil, false
	}
	return o.Long, true
}

// HasLong returns a boolean if a field has been set.
func (o *BufferModelBufferData) HasLong() bool {
	if o != nil && o.Long != nil {
		return true
	}

	return false
}

// SetLong gets a reference to the given int64 and assigns it to the Long field.
func (o *BufferModelBufferData) SetLong(v int64) {
	o.Long = &v
}

// GetFloat returns the Float field value if set, zero value otherwise.
func (o *BufferModelBufferData) GetFloat() float32 {
	if o == nil || o.Float == nil {
		var ret float32
		return ret
	}
	return *o.Float
}

// GetFloatOk returns a tuple with the Float field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferModelBufferData) GetFloatOk() (*float32, bool) {
	if o == nil || o.Float == nil {
		return nil, false
	}
	return o.Float, true
}

// HasFloat returns a boolean if a field has been set.
func (o *BufferModelBufferData) HasFloat() bool {
	if o != nil && o.Float != nil {
		return true
	}

	return false
}

// SetFloat gets a reference to the given float32 and assigns it to the Float field.
func (o *BufferModelBufferData) SetFloat(v float32) {
	o.Float = &v
}

// GetDouble returns the Double field value if set, zero value otherwise.
func (o *BufferModelBufferData) GetDouble() float64 {
	if o == nil || o.Double == nil {
		var ret float64
		return ret
	}
	return *o.Double
}

// GetDoubleOk returns a tuple with the Double field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferModelBufferData) GetDoubleOk() (*float64, bool) {
	if o == nil || o.Double == nil {
		return nil, false
	}
	return o.Double, true
}

// HasDouble returns a boolean if a field has been set.
func (o *BufferModelBufferData) HasDouble() bool {
	if o != nil && o.Double != nil {
		return true
	}

	return false
}

// SetDouble gets a reference to the given float64 and assigns it to the Double field.
func (o *BufferModelBufferData) SetDouble(v float64) {
	o.Double = &v
}

// GetDirect returns the Direct field value if set, zero value otherwise.
func (o *BufferModelBufferData) GetDirect() bool {
	if o == nil || o.Direct == nil {
		var ret bool
		return ret
	}
	return *o.Direct
}

// GetDirectOk returns a tuple with the Direct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferModelBufferData) GetDirectOk() (*bool, bool) {
	if o == nil || o.Direct == nil {
		return nil, false
	}
	return o.Direct, true
}

// HasDirect returns a boolean if a field has been set.
func (o *BufferModelBufferData) HasDirect() bool {
	if o != nil && o.Direct != nil {
		return true
	}

	return false
}

// SetDirect gets a reference to the given bool and assigns it to the Direct field.
func (o *BufferModelBufferData) SetDirect(v bool) {
	o.Direct = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *BufferModelBufferData) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BufferModelBufferData) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *BufferModelBufferData) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *BufferModelBufferData) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

func (o BufferModelBufferData) MarshalJSON() ([]byte, error) {
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

type NullableBufferModelBufferData struct {
	value *BufferModelBufferData
	isSet bool
}

func (v NullableBufferModelBufferData) Get() *BufferModelBufferData {
	return v.value
}

func (v *NullableBufferModelBufferData) Set(val *BufferModelBufferData) {
	v.value = val
	v.isSet = true
}

func (v NullableBufferModelBufferData) IsSet() bool {
	return v.isSet
}

func (v *NullableBufferModelBufferData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBufferModelBufferData(val *BufferModelBufferData) *NullableBufferModelBufferData {
	return &NullableBufferModelBufferData{value: val, isSet: true}
}

func (v NullableBufferModelBufferData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBufferModelBufferData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

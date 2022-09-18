/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6510-4a0bae47e25c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAppearanceOverride2517 struct for BTAppearanceOverride2517
type BTAppearanceOverride2517 struct {
	Appearance             *BTGraphicsAppearance1152 `json:"appearance,omitempty"`
	AppearanceReset        *bool                     `json:"appearanceReset,omitempty"`
	CopyWithoutEntities    *BTAppearanceOverride2517 `json:"copyWithoutEntities,omitempty"`
	EntityDeterministicIds []string                  `json:"entityDeterministicIds,omitempty"`
	IsDeletion             *bool                     `json:"isDeletion,omitempty"`
}

// NewBTAppearanceOverride2517 instantiates a new BTAppearanceOverride2517 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppearanceOverride2517() *BTAppearanceOverride2517 {
	this := BTAppearanceOverride2517{}
	return &this
}

// NewBTAppearanceOverride2517WithDefaults instantiates a new BTAppearanceOverride2517 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppearanceOverride2517WithDefaults() *BTAppearanceOverride2517 {
	this := BTAppearanceOverride2517{}
	return &this
}

// GetAppearance returns the Appearance field value if set, zero value otherwise.
func (o *BTAppearanceOverride2517) GetAppearance() BTGraphicsAppearance1152 {
	if o == nil || o.Appearance == nil {
		var ret BTGraphicsAppearance1152
		return ret
	}
	return *o.Appearance
}

// GetAppearanceOk returns a tuple with the Appearance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppearanceOverride2517) GetAppearanceOk() (*BTGraphicsAppearance1152, bool) {
	if o == nil || o.Appearance == nil {
		return nil, false
	}
	return o.Appearance, true
}

// HasAppearance returns a boolean if a field has been set.
func (o *BTAppearanceOverride2517) HasAppearance() bool {
	if o != nil && o.Appearance != nil {
		return true
	}

	return false
}

// SetAppearance gets a reference to the given BTGraphicsAppearance1152 and assigns it to the Appearance field.
func (o *BTAppearanceOverride2517) SetAppearance(v BTGraphicsAppearance1152) {
	o.Appearance = &v
}

// GetAppearanceReset returns the AppearanceReset field value if set, zero value otherwise.
func (o *BTAppearanceOverride2517) GetAppearanceReset() bool {
	if o == nil || o.AppearanceReset == nil {
		var ret bool
		return ret
	}
	return *o.AppearanceReset
}

// GetAppearanceResetOk returns a tuple with the AppearanceReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppearanceOverride2517) GetAppearanceResetOk() (*bool, bool) {
	if o == nil || o.AppearanceReset == nil {
		return nil, false
	}
	return o.AppearanceReset, true
}

// HasAppearanceReset returns a boolean if a field has been set.
func (o *BTAppearanceOverride2517) HasAppearanceReset() bool {
	if o != nil && o.AppearanceReset != nil {
		return true
	}

	return false
}

// SetAppearanceReset gets a reference to the given bool and assigns it to the AppearanceReset field.
func (o *BTAppearanceOverride2517) SetAppearanceReset(v bool) {
	o.AppearanceReset = &v
}

// GetCopyWithoutEntities returns the CopyWithoutEntities field value if set, zero value otherwise.
func (o *BTAppearanceOverride2517) GetCopyWithoutEntities() BTAppearanceOverride2517 {
	if o == nil || o.CopyWithoutEntities == nil {
		var ret BTAppearanceOverride2517
		return ret
	}
	return *o.CopyWithoutEntities
}

// GetCopyWithoutEntitiesOk returns a tuple with the CopyWithoutEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppearanceOverride2517) GetCopyWithoutEntitiesOk() (*BTAppearanceOverride2517, bool) {
	if o == nil || o.CopyWithoutEntities == nil {
		return nil, false
	}
	return o.CopyWithoutEntities, true
}

// HasCopyWithoutEntities returns a boolean if a field has been set.
func (o *BTAppearanceOverride2517) HasCopyWithoutEntities() bool {
	if o != nil && o.CopyWithoutEntities != nil {
		return true
	}

	return false
}

// SetCopyWithoutEntities gets a reference to the given BTAppearanceOverride2517 and assigns it to the CopyWithoutEntities field.
func (o *BTAppearanceOverride2517) SetCopyWithoutEntities(v BTAppearanceOverride2517) {
	o.CopyWithoutEntities = &v
}

// GetEntityDeterministicIds returns the EntityDeterministicIds field value if set, zero value otherwise.
func (o *BTAppearanceOverride2517) GetEntityDeterministicIds() []string {
	if o == nil || o.EntityDeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.EntityDeterministicIds
}

// GetEntityDeterministicIdsOk returns a tuple with the EntityDeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppearanceOverride2517) GetEntityDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.EntityDeterministicIds == nil {
		return nil, false
	}
	return o.EntityDeterministicIds, true
}

// HasEntityDeterministicIds returns a boolean if a field has been set.
func (o *BTAppearanceOverride2517) HasEntityDeterministicIds() bool {
	if o != nil && o.EntityDeterministicIds != nil {
		return true
	}

	return false
}

// SetEntityDeterministicIds gets a reference to the given []string and assigns it to the EntityDeterministicIds field.
func (o *BTAppearanceOverride2517) SetEntityDeterministicIds(v []string) {
	o.EntityDeterministicIds = v
}

// GetIsDeletion returns the IsDeletion field value if set, zero value otherwise.
func (o *BTAppearanceOverride2517) GetIsDeletion() bool {
	if o == nil || o.IsDeletion == nil {
		var ret bool
		return ret
	}
	return *o.IsDeletion
}

// GetIsDeletionOk returns a tuple with the IsDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppearanceOverride2517) GetIsDeletionOk() (*bool, bool) {
	if o == nil || o.IsDeletion == nil {
		return nil, false
	}
	return o.IsDeletion, true
}

// HasIsDeletion returns a boolean if a field has been set.
func (o *BTAppearanceOverride2517) HasIsDeletion() bool {
	if o != nil && o.IsDeletion != nil {
		return true
	}

	return false
}

// SetIsDeletion gets a reference to the given bool and assigns it to the IsDeletion field.
func (o *BTAppearanceOverride2517) SetIsDeletion(v bool) {
	o.IsDeletion = &v
}

func (o BTAppearanceOverride2517) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Appearance != nil {
		toSerialize["appearance"] = o.Appearance
	}
	if o.AppearanceReset != nil {
		toSerialize["appearanceReset"] = o.AppearanceReset
	}
	if o.CopyWithoutEntities != nil {
		toSerialize["copyWithoutEntities"] = o.CopyWithoutEntities
	}
	if o.EntityDeterministicIds != nil {
		toSerialize["entityDeterministicIds"] = o.EntityDeterministicIds
	}
	if o.IsDeletion != nil {
		toSerialize["isDeletion"] = o.IsDeletion
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppearanceOverride2517 struct {
	value *BTAppearanceOverride2517
	isSet bool
}

func (v NullableBTAppearanceOverride2517) Get() *BTAppearanceOverride2517 {
	return v.value
}

func (v *NullableBTAppearanceOverride2517) Set(val *BTAppearanceOverride2517) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppearanceOverride2517) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppearanceOverride2517) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppearanceOverride2517(val *BTAppearanceOverride2517) *NullableBTAppearanceOverride2517 {
	return &NullableBTAppearanceOverride2517{value: val, isSet: true}
}

func (v NullableBTAppearanceOverride2517) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppearanceOverride2517) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

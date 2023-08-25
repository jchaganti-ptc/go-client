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

// BTCommonUnitInfo struct for BTCommonUnitInfo
type BTCommonUnitInfo struct {
	Abbreviation     *string  `json:"abbreviation,omitempty"`
	Unit             *string  `json:"unit,omitempty"`
	UnitName         *string  `json:"unitName,omitempty"`
	UnitType         *string  `json:"unitType,omitempty"`
	ValueInBaseUnits *float64 `json:"valueInBaseUnits,omitempty"`
}

// NewBTCommonUnitInfo instantiates a new BTCommonUnitInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCommonUnitInfo() *BTCommonUnitInfo {
	this := BTCommonUnitInfo{}
	return &this
}

// NewBTCommonUnitInfoWithDefaults instantiates a new BTCommonUnitInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCommonUnitInfoWithDefaults() *BTCommonUnitInfo {
	this := BTCommonUnitInfo{}
	return &this
}

// GetAbbreviation returns the Abbreviation field value if set, zero value otherwise.
func (o *BTCommonUnitInfo) GetAbbreviation() string {
	if o == nil || o.Abbreviation == nil {
		var ret string
		return ret
	}
	return *o.Abbreviation
}

// GetAbbreviationOk returns a tuple with the Abbreviation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommonUnitInfo) GetAbbreviationOk() (*string, bool) {
	if o == nil || o.Abbreviation == nil {
		return nil, false
	}
	return o.Abbreviation, true
}

// HasAbbreviation returns a boolean if a field has been set.
func (o *BTCommonUnitInfo) HasAbbreviation() bool {
	if o != nil && o.Abbreviation != nil {
		return true
	}

	return false
}

// SetAbbreviation gets a reference to the given string and assigns it to the Abbreviation field.
func (o *BTCommonUnitInfo) SetAbbreviation(v string) {
	o.Abbreviation = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *BTCommonUnitInfo) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommonUnitInfo) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *BTCommonUnitInfo) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *BTCommonUnitInfo) SetUnit(v string) {
	o.Unit = &v
}

// GetUnitName returns the UnitName field value if set, zero value otherwise.
func (o *BTCommonUnitInfo) GetUnitName() string {
	if o == nil || o.UnitName == nil {
		var ret string
		return ret
	}
	return *o.UnitName
}

// GetUnitNameOk returns a tuple with the UnitName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommonUnitInfo) GetUnitNameOk() (*string, bool) {
	if o == nil || o.UnitName == nil {
		return nil, false
	}
	return o.UnitName, true
}

// HasUnitName returns a boolean if a field has been set.
func (o *BTCommonUnitInfo) HasUnitName() bool {
	if o != nil && o.UnitName != nil {
		return true
	}

	return false
}

// SetUnitName gets a reference to the given string and assigns it to the UnitName field.
func (o *BTCommonUnitInfo) SetUnitName(v string) {
	o.UnitName = &v
}

// GetUnitType returns the UnitType field value if set, zero value otherwise.
func (o *BTCommonUnitInfo) GetUnitType() string {
	if o == nil || o.UnitType == nil {
		var ret string
		return ret
	}
	return *o.UnitType
}

// GetUnitTypeOk returns a tuple with the UnitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommonUnitInfo) GetUnitTypeOk() (*string, bool) {
	if o == nil || o.UnitType == nil {
		return nil, false
	}
	return o.UnitType, true
}

// HasUnitType returns a boolean if a field has been set.
func (o *BTCommonUnitInfo) HasUnitType() bool {
	if o != nil && o.UnitType != nil {
		return true
	}

	return false
}

// SetUnitType gets a reference to the given string and assigns it to the UnitType field.
func (o *BTCommonUnitInfo) SetUnitType(v string) {
	o.UnitType = &v
}

// GetValueInBaseUnits returns the ValueInBaseUnits field value if set, zero value otherwise.
func (o *BTCommonUnitInfo) GetValueInBaseUnits() float64 {
	if o == nil || o.ValueInBaseUnits == nil {
		var ret float64
		return ret
	}
	return *o.ValueInBaseUnits
}

// GetValueInBaseUnitsOk returns a tuple with the ValueInBaseUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCommonUnitInfo) GetValueInBaseUnitsOk() (*float64, bool) {
	if o == nil || o.ValueInBaseUnits == nil {
		return nil, false
	}
	return o.ValueInBaseUnits, true
}

// HasValueInBaseUnits returns a boolean if a field has been set.
func (o *BTCommonUnitInfo) HasValueInBaseUnits() bool {
	if o != nil && o.ValueInBaseUnits != nil {
		return true
	}

	return false
}

// SetValueInBaseUnits gets a reference to the given float64 and assigns it to the ValueInBaseUnits field.
func (o *BTCommonUnitInfo) SetValueInBaseUnits(v float64) {
	o.ValueInBaseUnits = &v
}

func (o BTCommonUnitInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Abbreviation != nil {
		toSerialize["abbreviation"] = o.Abbreviation
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.UnitName != nil {
		toSerialize["unitName"] = o.UnitName
	}
	if o.UnitType != nil {
		toSerialize["unitType"] = o.UnitType
	}
	if o.ValueInBaseUnits != nil {
		toSerialize["valueInBaseUnits"] = o.ValueInBaseUnits
	}
	return json.Marshal(toSerialize)
}

type NullableBTCommonUnitInfo struct {
	value *BTCommonUnitInfo
	isSet bool
}

func (v NullableBTCommonUnitInfo) Get() *BTCommonUnitInfo {
	return v.value
}

func (v *NullableBTCommonUnitInfo) Set(val *BTCommonUnitInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCommonUnitInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCommonUnitInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCommonUnitInfo(val *BTCommonUnitInfo) *NullableBTCommonUnitInfo {
	return &NullableBTCommonUnitInfo{value: val, isSet: true}
}

func (v NullableBTCommonUnitInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCommonUnitInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

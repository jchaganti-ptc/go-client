/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.157.8827-f82e65cdc920
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTElementDisplayData326 - struct for BTElementDisplayData326
type BTElementDisplayData326 struct {
	implBTElementDisplayData326 interface{}
}

// BTPartStudioDisplayDataBase2751AsBTElementDisplayData326 is a convenience function that returns BTPartStudioDisplayDataBase2751 wrapped in BTElementDisplayData326
func (o *BTPartStudioDisplayDataBase2751) AsBTElementDisplayData326() *BTElementDisplayData326 {
	return &BTElementDisplayData326{o}
}

// BTAssemblyReferencesDisplayData1562AsBTElementDisplayData326 is a convenience function that returns BTAssemblyReferencesDisplayData1562 wrapped in BTElementDisplayData326
func (o *BTAssemblyReferencesDisplayData1562) AsBTElementDisplayData326() *BTElementDisplayData326 {
	return &BTElementDisplayData326{o}
}

// BTRootAssemblyDisplayData96AsBTElementDisplayData326 is a convenience function that returns BTRootAssemblyDisplayData96 wrapped in BTElementDisplayData326
func (o *BTRootAssemblyDisplayData96) AsBTElementDisplayData326() *BTElementDisplayData326 {
	return &BTElementDisplayData326{o}
}

// NewBTElementDisplayData326 instantiates a new BTElementDisplayData326 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTElementDisplayData326() *BTElementDisplayData326 {
	this := BTElementDisplayData326{Newbase_BTElementDisplayData326()}
	return &this
}

// NewBTElementDisplayData326WithDefaults instantiates a new BTElementDisplayData326 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTElementDisplayData326WithDefaults() *BTElementDisplayData326 {
	this := BTElementDisplayData326{Newbase_BTElementDisplayData326WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTElementDisplayData326) GetBtType() string {
	type getResult interface {
		GetBtType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtType()
	} else {
		var de string
		return de
	}
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementDisplayData326) GetBtTypeOk() (*string, bool) {
	type getResult interface {
		GetBtTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtTypeOk()
	} else {
		return nil, false
	}
}

// HasBtType returns a boolean if a field has been set.
func (o *BTElementDisplayData326) HasBtType() bool {
	type getResult interface {
		HasBtType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasBtType()
	} else {
		return false
	}
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTElementDisplayData326) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTElementDisplayData326) GetElementId() string {
	type getResult interface {
		GetElementId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetElementId()
	} else {
		var de string
		return de
	}
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementDisplayData326) GetElementIdOk() (*string, bool) {
	type getResult interface {
		GetElementIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetElementIdOk()
	} else {
		return nil, false
	}
}

// HasElementId returns a boolean if a field has been set.
func (o *BTElementDisplayData326) HasElementId() bool {
	type getResult interface {
		HasElementId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasElementId()
	} else {
		return false
	}
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTElementDisplayData326) SetElementId(v string) {
	type getResult interface {
		SetElementId(v string)
	}

	o.GetActualInstance().(getResult).SetElementId(v)
}

// GetFromFullElementId returns the FromFullElementId field value if set, zero value otherwise.
func (o *BTElementDisplayData326) GetFromFullElementId() BTFullElementId756 {
	type getResult interface {
		GetFromFullElementId() BTFullElementId756
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFromFullElementId()
	} else {
		var de BTFullElementId756
		return de
	}
}

// GetFromFullElementIdOk returns a tuple with the FromFullElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementDisplayData326) GetFromFullElementIdOk() (*BTFullElementId756, bool) {
	type getResult interface {
		GetFromFullElementIdOk() (*BTFullElementId756, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFromFullElementIdOk()
	} else {
		return nil, false
	}
}

// HasFromFullElementId returns a boolean if a field has been set.
func (o *BTElementDisplayData326) HasFromFullElementId() bool {
	type getResult interface {
		HasFromFullElementId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasFromFullElementId()
	} else {
		return false
	}
}

// SetFromFullElementId gets a reference to the given BTFullElementId756 and assigns it to the FromFullElementId field.
func (o *BTElementDisplayData326) SetFromFullElementId(v BTFullElementId756) {
	type getResult interface {
		SetFromFullElementId(v BTFullElementId756)
	}

	o.GetActualInstance().(getResult).SetFromFullElementId(v)
}

// GetFullElementId returns the FullElementId field value if set, zero value otherwise.
func (o *BTElementDisplayData326) GetFullElementId() BTFullElementId756 {
	type getResult interface {
		GetFullElementId() BTFullElementId756
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFullElementId()
	} else {
		var de BTFullElementId756
		return de
	}
}

// GetFullElementIdOk returns a tuple with the FullElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementDisplayData326) GetFullElementIdOk() (*BTFullElementId756, bool) {
	type getResult interface {
		GetFullElementIdOk() (*BTFullElementId756, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFullElementIdOk()
	} else {
		return nil, false
	}
}

// HasFullElementId returns a boolean if a field has been set.
func (o *BTElementDisplayData326) HasFullElementId() bool {
	type getResult interface {
		HasFullElementId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasFullElementId()
	} else {
		return false
	}
}

// SetFullElementId gets a reference to the given BTFullElementId756 and assigns it to the FullElementId field.
func (o *BTElementDisplayData326) SetFullElementId(v BTFullElementId756) {
	type getResult interface {
		SetFullElementId(v BTFullElementId756)
	}

	o.GetActualInstance().(getResult).SetFullElementId(v)
}

// GetIncremental returns the Incremental field value if set, zero value otherwise.
func (o *BTElementDisplayData326) GetIncremental() bool {
	type getResult interface {
		GetIncremental() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIncremental()
	} else {
		var de bool
		return de
	}
}

// GetIncrementalOk returns a tuple with the Incremental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementDisplayData326) GetIncrementalOk() (*bool, bool) {
	type getResult interface {
		GetIncrementalOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIncrementalOk()
	} else {
		return nil, false
	}
}

// HasIncremental returns a boolean if a field has been set.
func (o *BTElementDisplayData326) HasIncremental() bool {
	type getResult interface {
		HasIncremental() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasIncremental()
	} else {
		return false
	}
}

// SetIncremental gets a reference to the given bool and assigns it to the Incremental field.
func (o *BTElementDisplayData326) SetIncremental(v bool) {
	type getResult interface {
		SetIncremental(v bool)
	}

	o.GetActualInstance().(getResult).SetIncremental(v)
}

// GetInstanceCount returns the InstanceCount field value if set, zero value otherwise.
func (o *BTElementDisplayData326) GetInstanceCount() int32 {
	type getResult interface {
		GetInstanceCount() int32
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetInstanceCount()
	} else {
		var de int32
		return de
	}
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementDisplayData326) GetInstanceCountOk() (*int32, bool) {
	type getResult interface {
		GetInstanceCountOk() (*int32, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetInstanceCountOk()
	} else {
		return nil, false
	}
}

// HasInstanceCount returns a boolean if a field has been set.
func (o *BTElementDisplayData326) HasInstanceCount() bool {
	type getResult interface {
		HasInstanceCount() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasInstanceCount()
	} else {
		return false
	}
}

// SetInstanceCount gets a reference to the given int32 and assigns it to the InstanceCount field.
func (o *BTElementDisplayData326) SetInstanceCount(v int32) {
	type getResult interface {
		SetInstanceCount(v int32)
	}

	o.GetActualInstance().(getResult).SetInstanceCount(v)
}

// GetKeepFromMicroversion returns the KeepFromMicroversion field value if set, zero value otherwise.
func (o *BTElementDisplayData326) GetKeepFromMicroversion() bool {
	type getResult interface {
		GetKeepFromMicroversion() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetKeepFromMicroversion()
	} else {
		var de bool
		return de
	}
}

// GetKeepFromMicroversionOk returns a tuple with the KeepFromMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementDisplayData326) GetKeepFromMicroversionOk() (*bool, bool) {
	type getResult interface {
		GetKeepFromMicroversionOk() (*bool, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetKeepFromMicroversionOk()
	} else {
		return nil, false
	}
}

// HasKeepFromMicroversion returns a boolean if a field has been set.
func (o *BTElementDisplayData326) HasKeepFromMicroversion() bool {
	type getResult interface {
		HasKeepFromMicroversion() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasKeepFromMicroversion()
	} else {
		return false
	}
}

// SetKeepFromMicroversion gets a reference to the given bool and assigns it to the KeepFromMicroversion field.
func (o *BTElementDisplayData326) SetKeepFromMicroversion(v bool) {
	type getResult interface {
		SetKeepFromMicroversion(v bool)
	}

	o.GetActualInstance().(getResult).SetKeepFromMicroversion(v)
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTElementDisplayData326) GetMicroversionId() BTMicroversionId366 {
	type getResult interface {
		GetMicroversionId() BTMicroversionId366
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMicroversionId()
	} else {
		var de BTMicroversionId366
		return de
	}
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementDisplayData326) GetMicroversionIdOk() (*BTMicroversionId366, bool) {
	type getResult interface {
		GetMicroversionIdOk() (*BTMicroversionId366, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMicroversionIdOk()
	} else {
		return nil, false
	}
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTElementDisplayData326) HasMicroversionId() bool {
	type getResult interface {
		HasMicroversionId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasMicroversionId()
	} else {
		return false
	}
}

// SetMicroversionId gets a reference to the given BTMicroversionId366 and assigns it to the MicroversionId field.
func (o *BTElementDisplayData326) SetMicroversionId(v BTMicroversionId366) {
	type getResult interface {
		SetMicroversionId(v BTMicroversionId366)
	}

	o.GetActualInstance().(getResult).SetMicroversionId(v)
}

// GetMicroversionIdAndConfigurationInterval returns the MicroversionIdAndConfigurationInterval field value if set, zero value otherwise.
func (o *BTElementDisplayData326) GetMicroversionIdAndConfigurationInterval() BTMicroversionIdAndConfigurationInterval2364 {
	type getResult interface {
		GetMicroversionIdAndConfigurationInterval() BTMicroversionIdAndConfigurationInterval2364
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMicroversionIdAndConfigurationInterval()
	} else {
		var de BTMicroversionIdAndConfigurationInterval2364
		return de
	}
}

// GetMicroversionIdAndConfigurationIntervalOk returns a tuple with the MicroversionIdAndConfigurationInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementDisplayData326) GetMicroversionIdAndConfigurationIntervalOk() (*BTMicroversionIdAndConfigurationInterval2364, bool) {
	type getResult interface {
		GetMicroversionIdAndConfigurationIntervalOk() (*BTMicroversionIdAndConfigurationInterval2364, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMicroversionIdAndConfigurationIntervalOk()
	} else {
		return nil, false
	}
}

// HasMicroversionIdAndConfigurationInterval returns a boolean if a field has been set.
func (o *BTElementDisplayData326) HasMicroversionIdAndConfigurationInterval() bool {
	type getResult interface {
		HasMicroversionIdAndConfigurationInterval() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasMicroversionIdAndConfigurationInterval()
	} else {
		return false
	}
}

// SetMicroversionIdAndConfigurationInterval gets a reference to the given BTMicroversionIdAndConfigurationInterval2364 and assigns it to the MicroversionIdAndConfigurationInterval field.
func (o *BTElementDisplayData326) SetMicroversionIdAndConfigurationInterval(v BTMicroversionIdAndConfigurationInterval2364) {
	type getResult interface {
		SetMicroversionIdAndConfigurationInterval(v BTMicroversionIdAndConfigurationInterval2364)
	}

	o.GetActualInstance().(getResult).SetMicroversionIdAndConfigurationInterval(v)
}

// GetMicroversionInterval returns the MicroversionInterval field value if set, zero value otherwise.
func (o *BTElementDisplayData326) GetMicroversionInterval() BTMicroversionIdInterval367 {
	type getResult interface {
		GetMicroversionInterval() BTMicroversionIdInterval367
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMicroversionInterval()
	} else {
		var de BTMicroversionIdInterval367
		return de
	}
}

// GetMicroversionIntervalOk returns a tuple with the MicroversionInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementDisplayData326) GetMicroversionIntervalOk() (*BTMicroversionIdInterval367, bool) {
	type getResult interface {
		GetMicroversionIntervalOk() (*BTMicroversionIdInterval367, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMicroversionIntervalOk()
	} else {
		return nil, false
	}
}

// HasMicroversionInterval returns a boolean if a field has been set.
func (o *BTElementDisplayData326) HasMicroversionInterval() bool {
	type getResult interface {
		HasMicroversionInterval() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasMicroversionInterval()
	} else {
		return false
	}
}

// SetMicroversionInterval gets a reference to the given BTMicroversionIdInterval367 and assigns it to the MicroversionInterval field.
func (o *BTElementDisplayData326) SetMicroversionInterval(v BTMicroversionIdInterval367) {
	type getResult interface {
		SetMicroversionInterval(v BTMicroversionIdInterval367)
	}

	o.GetActualInstance().(getResult).SetMicroversionInterval(v)
}

// GetVersionForRasterization returns the VersionForRasterization field value if set, zero value otherwise.
func (o *BTElementDisplayData326) GetVersionForRasterization() BTElementDisplayData326 {
	type getResult interface {
		GetVersionForRasterization() BTElementDisplayData326
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetVersionForRasterization()
	} else {
		var de BTElementDisplayData326
		return de
	}
}

// GetVersionForRasterizationOk returns a tuple with the VersionForRasterization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTElementDisplayData326) GetVersionForRasterizationOk() (*BTElementDisplayData326, bool) {
	type getResult interface {
		GetVersionForRasterizationOk() (*BTElementDisplayData326, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetVersionForRasterizationOk()
	} else {
		return nil, false
	}
}

// HasVersionForRasterization returns a boolean if a field has been set.
func (o *BTElementDisplayData326) HasVersionForRasterization() bool {
	type getResult interface {
		HasVersionForRasterization() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasVersionForRasterization()
	} else {
		return false
	}
}

// SetVersionForRasterization gets a reference to the given BTElementDisplayData326 and assigns it to the VersionForRasterization field.
func (o *BTElementDisplayData326) SetVersionForRasterization(v BTElementDisplayData326) {
	type getResult interface {
		SetVersionForRasterization(v BTElementDisplayData326)
	}

	o.GetActualInstance().(getResult).SetVersionForRasterization(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTElementDisplayData326) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTAssemblyReferencesDisplayData-1562'
	if jsonDict["btType"] == "BTAssemblyReferencesDisplayData-1562" {
		// try to unmarshal JSON data into BTAssemblyReferencesDisplayData1562
		var qr *BTAssemblyReferencesDisplayData1562
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTElementDisplayData326 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTElementDisplayData326 = nil
			return fmt.Errorf("Failed to unmarshal BTElementDisplayData326 as BTAssemblyReferencesDisplayData1562: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTPartStudioDisplayDataBase-2751'
	if jsonDict["btType"] == "BTPartStudioDisplayDataBase-2751" {
		// try to unmarshal JSON data into BTPartStudioDisplayDataBase2751
		var qr *BTPartStudioDisplayDataBase2751
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTElementDisplayData326 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTElementDisplayData326 = nil
			return fmt.Errorf("Failed to unmarshal BTElementDisplayData326 as BTPartStudioDisplayDataBase2751: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTRootAssemblyDisplayData-96'
	if jsonDict["btType"] == "BTRootAssemblyDisplayData-96" {
		// try to unmarshal JSON data into BTRootAssemblyDisplayData96
		var qr *BTRootAssemblyDisplayData96
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTElementDisplayData326 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTElementDisplayData326 = nil
			return fmt.Errorf("Failed to unmarshal BTElementDisplayData326 as BTRootAssemblyDisplayData96: %s", err.Error())
		}
	}

	var qtx *base_BTElementDisplayData326
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTElementDisplayData326 = qtx
		return nil // data stored in dst.base_BTElementDisplayData326, return on the first match
	} else {
		dst.implBTElementDisplayData326 = nil
		return fmt.Errorf("Failed to unmarshal BTElementDisplayData326 as base_BTElementDisplayData326: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTElementDisplayData326) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTElementDisplayData326) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTElementDisplayData326
}

type NullableBTElementDisplayData326 struct {
	value *BTElementDisplayData326
	isSet bool
}

func (v NullableBTElementDisplayData326) Get() *BTElementDisplayData326 {
	return v.value
}

func (v *NullableBTElementDisplayData326) Set(val *BTElementDisplayData326) {
	v.value = val
	v.isSet = true
}

func (v NullableBTElementDisplayData326) IsSet() bool {
	return v.isSet
}

func (v *NullableBTElementDisplayData326) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTElementDisplayData326(val *BTElementDisplayData326) *NullableBTElementDisplayData326 {
	return &NullableBTElementDisplayData326{value: val, isSet: true}
}

func (v NullableBTElementDisplayData326) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTElementDisplayData326) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTElementDisplayData326 struct {
	BtType                                 *string                                       `json:"btType,omitempty"`
	ElementId                              *string                                       `json:"elementId,omitempty"`
	FromFullElementId                      *BTFullElementId756                           `json:"fromFullElementId,omitempty"`
	FullElementId                          *BTFullElementId756                           `json:"fullElementId,omitempty"`
	Incremental                            *bool                                         `json:"incremental,omitempty"`
	InstanceCount                          *int32                                        `json:"instanceCount,omitempty"`
	KeepFromMicroversion                   *bool                                         `json:"keepFromMicroversion,omitempty"`
	MicroversionId                         *BTMicroversionId366                          `json:"microversionId,omitempty"`
	MicroversionIdAndConfigurationInterval *BTMicroversionIdAndConfigurationInterval2364 `json:"microversionIdAndConfigurationInterval,omitempty"`
	MicroversionInterval                   *BTMicroversionIdInterval367                  `json:"microversionInterval,omitempty"`
	VersionForRasterization                *BTElementDisplayData326                      `json:"versionForRasterization,omitempty"`
}

// Newbase_BTElementDisplayData326 instantiates a new base_BTElementDisplayData326 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTElementDisplayData326() *base_BTElementDisplayData326 {
	this := base_BTElementDisplayData326{}
	return &this
}

// Newbase_BTElementDisplayData326WithDefaults instantiates a new base_BTElementDisplayData326 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTElementDisplayData326WithDefaults() *base_BTElementDisplayData326 {
	this := base_BTElementDisplayData326{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTElementDisplayData326) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTElementDisplayData326) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTElementDisplayData326) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTElementDisplayData326) SetBtType(v string) {
	o.BtType = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *base_BTElementDisplayData326) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTElementDisplayData326) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *base_BTElementDisplayData326) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *base_BTElementDisplayData326) SetElementId(v string) {
	o.ElementId = &v
}

// GetFromFullElementId returns the FromFullElementId field value if set, zero value otherwise.
func (o *base_BTElementDisplayData326) GetFromFullElementId() BTFullElementId756 {
	if o == nil || o.FromFullElementId == nil {
		var ret BTFullElementId756
		return ret
	}
	return *o.FromFullElementId
}

// GetFromFullElementIdOk returns a tuple with the FromFullElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTElementDisplayData326) GetFromFullElementIdOk() (*BTFullElementId756, bool) {
	if o == nil || o.FromFullElementId == nil {
		return nil, false
	}
	return o.FromFullElementId, true
}

// HasFromFullElementId returns a boolean if a field has been set.
func (o *base_BTElementDisplayData326) HasFromFullElementId() bool {
	if o != nil && o.FromFullElementId != nil {
		return true
	}

	return false
}

// SetFromFullElementId gets a reference to the given BTFullElementId756 and assigns it to the FromFullElementId field.
func (o *base_BTElementDisplayData326) SetFromFullElementId(v BTFullElementId756) {
	o.FromFullElementId = &v
}

// GetFullElementId returns the FullElementId field value if set, zero value otherwise.
func (o *base_BTElementDisplayData326) GetFullElementId() BTFullElementId756 {
	if o == nil || o.FullElementId == nil {
		var ret BTFullElementId756
		return ret
	}
	return *o.FullElementId
}

// GetFullElementIdOk returns a tuple with the FullElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTElementDisplayData326) GetFullElementIdOk() (*BTFullElementId756, bool) {
	if o == nil || o.FullElementId == nil {
		return nil, false
	}
	return o.FullElementId, true
}

// HasFullElementId returns a boolean if a field has been set.
func (o *base_BTElementDisplayData326) HasFullElementId() bool {
	if o != nil && o.FullElementId != nil {
		return true
	}

	return false
}

// SetFullElementId gets a reference to the given BTFullElementId756 and assigns it to the FullElementId field.
func (o *base_BTElementDisplayData326) SetFullElementId(v BTFullElementId756) {
	o.FullElementId = &v
}

// GetIncremental returns the Incremental field value if set, zero value otherwise.
func (o *base_BTElementDisplayData326) GetIncremental() bool {
	if o == nil || o.Incremental == nil {
		var ret bool
		return ret
	}
	return *o.Incremental
}

// GetIncrementalOk returns a tuple with the Incremental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTElementDisplayData326) GetIncrementalOk() (*bool, bool) {
	if o == nil || o.Incremental == nil {
		return nil, false
	}
	return o.Incremental, true
}

// HasIncremental returns a boolean if a field has been set.
func (o *base_BTElementDisplayData326) HasIncremental() bool {
	if o != nil && o.Incremental != nil {
		return true
	}

	return false
}

// SetIncremental gets a reference to the given bool and assigns it to the Incremental field.
func (o *base_BTElementDisplayData326) SetIncremental(v bool) {
	o.Incremental = &v
}

// GetInstanceCount returns the InstanceCount field value if set, zero value otherwise.
func (o *base_BTElementDisplayData326) GetInstanceCount() int32 {
	if o == nil || o.InstanceCount == nil {
		var ret int32
		return ret
	}
	return *o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTElementDisplayData326) GetInstanceCountOk() (*int32, bool) {
	if o == nil || o.InstanceCount == nil {
		return nil, false
	}
	return o.InstanceCount, true
}

// HasInstanceCount returns a boolean if a field has been set.
func (o *base_BTElementDisplayData326) HasInstanceCount() bool {
	if o != nil && o.InstanceCount != nil {
		return true
	}

	return false
}

// SetInstanceCount gets a reference to the given int32 and assigns it to the InstanceCount field.
func (o *base_BTElementDisplayData326) SetInstanceCount(v int32) {
	o.InstanceCount = &v
}

// GetKeepFromMicroversion returns the KeepFromMicroversion field value if set, zero value otherwise.
func (o *base_BTElementDisplayData326) GetKeepFromMicroversion() bool {
	if o == nil || o.KeepFromMicroversion == nil {
		var ret bool
		return ret
	}
	return *o.KeepFromMicroversion
}

// GetKeepFromMicroversionOk returns a tuple with the KeepFromMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTElementDisplayData326) GetKeepFromMicroversionOk() (*bool, bool) {
	if o == nil || o.KeepFromMicroversion == nil {
		return nil, false
	}
	return o.KeepFromMicroversion, true
}

// HasKeepFromMicroversion returns a boolean if a field has been set.
func (o *base_BTElementDisplayData326) HasKeepFromMicroversion() bool {
	if o != nil && o.KeepFromMicroversion != nil {
		return true
	}

	return false
}

// SetKeepFromMicroversion gets a reference to the given bool and assigns it to the KeepFromMicroversion field.
func (o *base_BTElementDisplayData326) SetKeepFromMicroversion(v bool) {
	o.KeepFromMicroversion = &v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *base_BTElementDisplayData326) GetMicroversionId() BTMicroversionId366 {
	if o == nil || o.MicroversionId == nil {
		var ret BTMicroversionId366
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTElementDisplayData326) GetMicroversionIdOk() (*BTMicroversionId366, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *base_BTElementDisplayData326) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given BTMicroversionId366 and assigns it to the MicroversionId field.
func (o *base_BTElementDisplayData326) SetMicroversionId(v BTMicroversionId366) {
	o.MicroversionId = &v
}

// GetMicroversionIdAndConfigurationInterval returns the MicroversionIdAndConfigurationInterval field value if set, zero value otherwise.
func (o *base_BTElementDisplayData326) GetMicroversionIdAndConfigurationInterval() BTMicroversionIdAndConfigurationInterval2364 {
	if o == nil || o.MicroversionIdAndConfigurationInterval == nil {
		var ret BTMicroversionIdAndConfigurationInterval2364
		return ret
	}
	return *o.MicroversionIdAndConfigurationInterval
}

// GetMicroversionIdAndConfigurationIntervalOk returns a tuple with the MicroversionIdAndConfigurationInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTElementDisplayData326) GetMicroversionIdAndConfigurationIntervalOk() (*BTMicroversionIdAndConfigurationInterval2364, bool) {
	if o == nil || o.MicroversionIdAndConfigurationInterval == nil {
		return nil, false
	}
	return o.MicroversionIdAndConfigurationInterval, true
}

// HasMicroversionIdAndConfigurationInterval returns a boolean if a field has been set.
func (o *base_BTElementDisplayData326) HasMicroversionIdAndConfigurationInterval() bool {
	if o != nil && o.MicroversionIdAndConfigurationInterval != nil {
		return true
	}

	return false
}

// SetMicroversionIdAndConfigurationInterval gets a reference to the given BTMicroversionIdAndConfigurationInterval2364 and assigns it to the MicroversionIdAndConfigurationInterval field.
func (o *base_BTElementDisplayData326) SetMicroversionIdAndConfigurationInterval(v BTMicroversionIdAndConfigurationInterval2364) {
	o.MicroversionIdAndConfigurationInterval = &v
}

// GetMicroversionInterval returns the MicroversionInterval field value if set, zero value otherwise.
func (o *base_BTElementDisplayData326) GetMicroversionInterval() BTMicroversionIdInterval367 {
	if o == nil || o.MicroversionInterval == nil {
		var ret BTMicroversionIdInterval367
		return ret
	}
	return *o.MicroversionInterval
}

// GetMicroversionIntervalOk returns a tuple with the MicroversionInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTElementDisplayData326) GetMicroversionIntervalOk() (*BTMicroversionIdInterval367, bool) {
	if o == nil || o.MicroversionInterval == nil {
		return nil, false
	}
	return o.MicroversionInterval, true
}

// HasMicroversionInterval returns a boolean if a field has been set.
func (o *base_BTElementDisplayData326) HasMicroversionInterval() bool {
	if o != nil && o.MicroversionInterval != nil {
		return true
	}

	return false
}

// SetMicroversionInterval gets a reference to the given BTMicroversionIdInterval367 and assigns it to the MicroversionInterval field.
func (o *base_BTElementDisplayData326) SetMicroversionInterval(v BTMicroversionIdInterval367) {
	o.MicroversionInterval = &v
}

// GetVersionForRasterization returns the VersionForRasterization field value if set, zero value otherwise.
func (o *base_BTElementDisplayData326) GetVersionForRasterization() BTElementDisplayData326 {
	if o == nil || o.VersionForRasterization == nil {
		var ret BTElementDisplayData326
		return ret
	}
	return *o.VersionForRasterization
}

// GetVersionForRasterizationOk returns a tuple with the VersionForRasterization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTElementDisplayData326) GetVersionForRasterizationOk() (*BTElementDisplayData326, bool) {
	if o == nil || o.VersionForRasterization == nil {
		return nil, false
	}
	return o.VersionForRasterization, true
}

// HasVersionForRasterization returns a boolean if a field has been set.
func (o *base_BTElementDisplayData326) HasVersionForRasterization() bool {
	if o != nil && o.VersionForRasterization != nil {
		return true
	}

	return false
}

// SetVersionForRasterization gets a reference to the given BTElementDisplayData326 and assigns it to the VersionForRasterization field.
func (o *base_BTElementDisplayData326) SetVersionForRasterization(v BTElementDisplayData326) {
	o.VersionForRasterization = &v
}

func (o base_BTElementDisplayData326) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.FromFullElementId != nil {
		toSerialize["fromFullElementId"] = o.FromFullElementId
	}
	if o.FullElementId != nil {
		toSerialize["fullElementId"] = o.FullElementId
	}
	if o.Incremental != nil {
		toSerialize["incremental"] = o.Incremental
	}
	if o.InstanceCount != nil {
		toSerialize["instanceCount"] = o.InstanceCount
	}
	if o.KeepFromMicroversion != nil {
		toSerialize["keepFromMicroversion"] = o.KeepFromMicroversion
	}
	if o.MicroversionId != nil {
		toSerialize["microversionId"] = o.MicroversionId
	}
	if o.MicroversionIdAndConfigurationInterval != nil {
		toSerialize["microversionIdAndConfigurationInterval"] = o.MicroversionIdAndConfigurationInterval
	}
	if o.MicroversionInterval != nil {
		toSerialize["microversionInterval"] = o.MicroversionInterval
	}
	if o.VersionForRasterization != nil {
		toSerialize["versionForRasterization"] = o.VersionForRasterization
	}
	return json.Marshal(toSerialize)
}

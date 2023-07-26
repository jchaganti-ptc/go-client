/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.167.19458-7ff87863110f
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPartStudioDisplayDataReference4360 struct for BTPartStudioDisplayDataReference4360
type BTPartStudioDisplayDataReference4360 struct {
	BtType                                 *string                                       `json:"btType,omitempty"`
	ElementId                              *string                                       `json:"elementId,omitempty"`
	FromFullElementId                      *BTFullElementId756                           `json:"fromFullElementId,omitempty"`
	FullElementId                          *BTFullElementId756                           `json:"fullElementId,omitempty"`
	Incremental                            *bool                                         `json:"incremental,omitempty"`
	InstanceCount                          *int32                                        `json:"instanceCount,omitempty"`
	IsNoop                                 *bool                                         `json:"isNoop,omitempty"`
	KeepFromMicroversion                   *bool                                         `json:"keepFromMicroversion,omitempty"`
	MicroversionId                         *BTMicroversionId366                          `json:"microversionId,omitempty"`
	MicroversionIdAndConfigurationInterval *BTMicroversionIdAndConfigurationInterval2364 `json:"microversionIdAndConfigurationInterval,omitempty"`
	MicroversionInterval                   *BTMicroversionIdInterval367                  `json:"microversionInterval,omitempty"`
	NumberOfSketchEntities                 *int32                                        `json:"numberOfSketchEntities,omitempty"`
	VersionForRasterization                *BTElementDisplayData326                      `json:"versionForRasterization,omitempty"`
	CacheDataPath                          *BTCacheDataPath191                           `json:"cacheDataPath,omitempty"`
	TessellationSettings                   *map[string][]int32                           `json:"tessellationSettings,omitempty"`
}

// NewBTPartStudioDisplayDataReference4360 instantiates a new BTPartStudioDisplayDataReference4360 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPartStudioDisplayDataReference4360() *BTPartStudioDisplayDataReference4360 {
	this := BTPartStudioDisplayDataReference4360{}
	return &this
}

// NewBTPartStudioDisplayDataReference4360WithDefaults instantiates a new BTPartStudioDisplayDataReference4360 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPartStudioDisplayDataReference4360WithDefaults() *BTPartStudioDisplayDataReference4360 {
	this := BTPartStudioDisplayDataReference4360{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPartStudioDisplayDataReference4360) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartStudioDisplayDataReference4360) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPartStudioDisplayDataReference4360) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPartStudioDisplayDataReference4360) SetBtType(v string) {
	o.BtType = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTPartStudioDisplayDataReference4360) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartStudioDisplayDataReference4360) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTPartStudioDisplayDataReference4360) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTPartStudioDisplayDataReference4360) SetElementId(v string) {
	o.ElementId = &v
}

// GetFromFullElementId returns the FromFullElementId field value if set, zero value otherwise.
func (o *BTPartStudioDisplayDataReference4360) GetFromFullElementId() BTFullElementId756 {
	if o == nil || o.FromFullElementId == nil {
		var ret BTFullElementId756
		return ret
	}
	return *o.FromFullElementId
}

// GetFromFullElementIdOk returns a tuple with the FromFullElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartStudioDisplayDataReference4360) GetFromFullElementIdOk() (*BTFullElementId756, bool) {
	if o == nil || o.FromFullElementId == nil {
		return nil, false
	}
	return o.FromFullElementId, true
}

// HasFromFullElementId returns a boolean if a field has been set.
func (o *BTPartStudioDisplayDataReference4360) HasFromFullElementId() bool {
	if o != nil && o.FromFullElementId != nil {
		return true
	}

	return false
}

// SetFromFullElementId gets a reference to the given BTFullElementId756 and assigns it to the FromFullElementId field.
func (o *BTPartStudioDisplayDataReference4360) SetFromFullElementId(v BTFullElementId756) {
	o.FromFullElementId = &v
}

// GetFullElementId returns the FullElementId field value if set, zero value otherwise.
func (o *BTPartStudioDisplayDataReference4360) GetFullElementId() BTFullElementId756 {
	if o == nil || o.FullElementId == nil {
		var ret BTFullElementId756
		return ret
	}
	return *o.FullElementId
}

// GetFullElementIdOk returns a tuple with the FullElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartStudioDisplayDataReference4360) GetFullElementIdOk() (*BTFullElementId756, bool) {
	if o == nil || o.FullElementId == nil {
		return nil, false
	}
	return o.FullElementId, true
}

// HasFullElementId returns a boolean if a field has been set.
func (o *BTPartStudioDisplayDataReference4360) HasFullElementId() bool {
	if o != nil && o.FullElementId != nil {
		return true
	}

	return false
}

// SetFullElementId gets a reference to the given BTFullElementId756 and assigns it to the FullElementId field.
func (o *BTPartStudioDisplayDataReference4360) SetFullElementId(v BTFullElementId756) {
	o.FullElementId = &v
}

// GetIncremental returns the Incremental field value if set, zero value otherwise.
func (o *BTPartStudioDisplayDataReference4360) GetIncremental() bool {
	if o == nil || o.Incremental == nil {
		var ret bool
		return ret
	}
	return *o.Incremental
}

// GetIncrementalOk returns a tuple with the Incremental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartStudioDisplayDataReference4360) GetIncrementalOk() (*bool, bool) {
	if o == nil || o.Incremental == nil {
		return nil, false
	}
	return o.Incremental, true
}

// HasIncremental returns a boolean if a field has been set.
func (o *BTPartStudioDisplayDataReference4360) HasIncremental() bool {
	if o != nil && o.Incremental != nil {
		return true
	}

	return false
}

// SetIncremental gets a reference to the given bool and assigns it to the Incremental field.
func (o *BTPartStudioDisplayDataReference4360) SetIncremental(v bool) {
	o.Incremental = &v
}

// GetInstanceCount returns the InstanceCount field value if set, zero value otherwise.
func (o *BTPartStudioDisplayDataReference4360) GetInstanceCount() int32 {
	if o == nil || o.InstanceCount == nil {
		var ret int32
		return ret
	}
	return *o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartStudioDisplayDataReference4360) GetInstanceCountOk() (*int32, bool) {
	if o == nil || o.InstanceCount == nil {
		return nil, false
	}
	return o.InstanceCount, true
}

// HasInstanceCount returns a boolean if a field has been set.
func (o *BTPartStudioDisplayDataReference4360) HasInstanceCount() bool {
	if o != nil && o.InstanceCount != nil {
		return true
	}

	return false
}

// SetInstanceCount gets a reference to the given int32 and assigns it to the InstanceCount field.
func (o *BTPartStudioDisplayDataReference4360) SetInstanceCount(v int32) {
	o.InstanceCount = &v
}

// GetIsNoop returns the IsNoop field value if set, zero value otherwise.
func (o *BTPartStudioDisplayDataReference4360) GetIsNoop() bool {
	if o == nil || o.IsNoop == nil {
		var ret bool
		return ret
	}
	return *o.IsNoop
}

// GetIsNoopOk returns a tuple with the IsNoop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartStudioDisplayDataReference4360) GetIsNoopOk() (*bool, bool) {
	if o == nil || o.IsNoop == nil {
		return nil, false
	}
	return o.IsNoop, true
}

// HasIsNoop returns a boolean if a field has been set.
func (o *BTPartStudioDisplayDataReference4360) HasIsNoop() bool {
	if o != nil && o.IsNoop != nil {
		return true
	}

	return false
}

// SetIsNoop gets a reference to the given bool and assigns it to the IsNoop field.
func (o *BTPartStudioDisplayDataReference4360) SetIsNoop(v bool) {
	o.IsNoop = &v
}

// GetKeepFromMicroversion returns the KeepFromMicroversion field value if set, zero value otherwise.
func (o *BTPartStudioDisplayDataReference4360) GetKeepFromMicroversion() bool {
	if o == nil || o.KeepFromMicroversion == nil {
		var ret bool
		return ret
	}
	return *o.KeepFromMicroversion
}

// GetKeepFromMicroversionOk returns a tuple with the KeepFromMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartStudioDisplayDataReference4360) GetKeepFromMicroversionOk() (*bool, bool) {
	if o == nil || o.KeepFromMicroversion == nil {
		return nil, false
	}
	return o.KeepFromMicroversion, true
}

// HasKeepFromMicroversion returns a boolean if a field has been set.
func (o *BTPartStudioDisplayDataReference4360) HasKeepFromMicroversion() bool {
	if o != nil && o.KeepFromMicroversion != nil {
		return true
	}

	return false
}

// SetKeepFromMicroversion gets a reference to the given bool and assigns it to the KeepFromMicroversion field.
func (o *BTPartStudioDisplayDataReference4360) SetKeepFromMicroversion(v bool) {
	o.KeepFromMicroversion = &v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTPartStudioDisplayDataReference4360) GetMicroversionId() BTMicroversionId366 {
	if o == nil || o.MicroversionId == nil {
		var ret BTMicroversionId366
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartStudioDisplayDataReference4360) GetMicroversionIdOk() (*BTMicroversionId366, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTPartStudioDisplayDataReference4360) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given BTMicroversionId366 and assigns it to the MicroversionId field.
func (o *BTPartStudioDisplayDataReference4360) SetMicroversionId(v BTMicroversionId366) {
	o.MicroversionId = &v
}

// GetMicroversionIdAndConfigurationInterval returns the MicroversionIdAndConfigurationInterval field value if set, zero value otherwise.
func (o *BTPartStudioDisplayDataReference4360) GetMicroversionIdAndConfigurationInterval() BTMicroversionIdAndConfigurationInterval2364 {
	if o == nil || o.MicroversionIdAndConfigurationInterval == nil {
		var ret BTMicroversionIdAndConfigurationInterval2364
		return ret
	}
	return *o.MicroversionIdAndConfigurationInterval
}

// GetMicroversionIdAndConfigurationIntervalOk returns a tuple with the MicroversionIdAndConfigurationInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartStudioDisplayDataReference4360) GetMicroversionIdAndConfigurationIntervalOk() (*BTMicroversionIdAndConfigurationInterval2364, bool) {
	if o == nil || o.MicroversionIdAndConfigurationInterval == nil {
		return nil, false
	}
	return o.MicroversionIdAndConfigurationInterval, true
}

// HasMicroversionIdAndConfigurationInterval returns a boolean if a field has been set.
func (o *BTPartStudioDisplayDataReference4360) HasMicroversionIdAndConfigurationInterval() bool {
	if o != nil && o.MicroversionIdAndConfigurationInterval != nil {
		return true
	}

	return false
}

// SetMicroversionIdAndConfigurationInterval gets a reference to the given BTMicroversionIdAndConfigurationInterval2364 and assigns it to the MicroversionIdAndConfigurationInterval field.
func (o *BTPartStudioDisplayDataReference4360) SetMicroversionIdAndConfigurationInterval(v BTMicroversionIdAndConfigurationInterval2364) {
	o.MicroversionIdAndConfigurationInterval = &v
}

// GetMicroversionInterval returns the MicroversionInterval field value if set, zero value otherwise.
func (o *BTPartStudioDisplayDataReference4360) GetMicroversionInterval() BTMicroversionIdInterval367 {
	if o == nil || o.MicroversionInterval == nil {
		var ret BTMicroversionIdInterval367
		return ret
	}
	return *o.MicroversionInterval
}

// GetMicroversionIntervalOk returns a tuple with the MicroversionInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartStudioDisplayDataReference4360) GetMicroversionIntervalOk() (*BTMicroversionIdInterval367, bool) {
	if o == nil || o.MicroversionInterval == nil {
		return nil, false
	}
	return o.MicroversionInterval, true
}

// HasMicroversionInterval returns a boolean if a field has been set.
func (o *BTPartStudioDisplayDataReference4360) HasMicroversionInterval() bool {
	if o != nil && o.MicroversionInterval != nil {
		return true
	}

	return false
}

// SetMicroversionInterval gets a reference to the given BTMicroversionIdInterval367 and assigns it to the MicroversionInterval field.
func (o *BTPartStudioDisplayDataReference4360) SetMicroversionInterval(v BTMicroversionIdInterval367) {
	o.MicroversionInterval = &v
}

// GetNumberOfSketchEntities returns the NumberOfSketchEntities field value if set, zero value otherwise.
func (o *BTPartStudioDisplayDataReference4360) GetNumberOfSketchEntities() int32 {
	if o == nil || o.NumberOfSketchEntities == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfSketchEntities
}

// GetNumberOfSketchEntitiesOk returns a tuple with the NumberOfSketchEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartStudioDisplayDataReference4360) GetNumberOfSketchEntitiesOk() (*int32, bool) {
	if o == nil || o.NumberOfSketchEntities == nil {
		return nil, false
	}
	return o.NumberOfSketchEntities, true
}

// HasNumberOfSketchEntities returns a boolean if a field has been set.
func (o *BTPartStudioDisplayDataReference4360) HasNumberOfSketchEntities() bool {
	if o != nil && o.NumberOfSketchEntities != nil {
		return true
	}

	return false
}

// SetNumberOfSketchEntities gets a reference to the given int32 and assigns it to the NumberOfSketchEntities field.
func (o *BTPartStudioDisplayDataReference4360) SetNumberOfSketchEntities(v int32) {
	o.NumberOfSketchEntities = &v
}

// GetVersionForRasterization returns the VersionForRasterization field value if set, zero value otherwise.
func (o *BTPartStudioDisplayDataReference4360) GetVersionForRasterization() BTElementDisplayData326 {
	if o == nil || o.VersionForRasterization == nil {
		var ret BTElementDisplayData326
		return ret
	}
	return *o.VersionForRasterization
}

// GetVersionForRasterizationOk returns a tuple with the VersionForRasterization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartStudioDisplayDataReference4360) GetVersionForRasterizationOk() (*BTElementDisplayData326, bool) {
	if o == nil || o.VersionForRasterization == nil {
		return nil, false
	}
	return o.VersionForRasterization, true
}

// HasVersionForRasterization returns a boolean if a field has been set.
func (o *BTPartStudioDisplayDataReference4360) HasVersionForRasterization() bool {
	if o != nil && o.VersionForRasterization != nil {
		return true
	}

	return false
}

// SetVersionForRasterization gets a reference to the given BTElementDisplayData326 and assigns it to the VersionForRasterization field.
func (o *BTPartStudioDisplayDataReference4360) SetVersionForRasterization(v BTElementDisplayData326) {
	o.VersionForRasterization = &v
}

// GetCacheDataPath returns the CacheDataPath field value if set, zero value otherwise.
func (o *BTPartStudioDisplayDataReference4360) GetCacheDataPath() BTCacheDataPath191 {
	if o == nil || o.CacheDataPath == nil {
		var ret BTCacheDataPath191
		return ret
	}
	return *o.CacheDataPath
}

// GetCacheDataPathOk returns a tuple with the CacheDataPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartStudioDisplayDataReference4360) GetCacheDataPathOk() (*BTCacheDataPath191, bool) {
	if o == nil || o.CacheDataPath == nil {
		return nil, false
	}
	return o.CacheDataPath, true
}

// HasCacheDataPath returns a boolean if a field has been set.
func (o *BTPartStudioDisplayDataReference4360) HasCacheDataPath() bool {
	if o != nil && o.CacheDataPath != nil {
		return true
	}

	return false
}

// SetCacheDataPath gets a reference to the given BTCacheDataPath191 and assigns it to the CacheDataPath field.
func (o *BTPartStudioDisplayDataReference4360) SetCacheDataPath(v BTCacheDataPath191) {
	o.CacheDataPath = &v
}

// GetTessellationSettings returns the TessellationSettings field value if set, zero value otherwise.
func (o *BTPartStudioDisplayDataReference4360) GetTessellationSettings() map[string][]int32 {
	if o == nil || o.TessellationSettings == nil {
		var ret map[string][]int32
		return ret
	}
	return *o.TessellationSettings
}

// GetTessellationSettingsOk returns a tuple with the TessellationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartStudioDisplayDataReference4360) GetTessellationSettingsOk() (*map[string][]int32, bool) {
	if o == nil || o.TessellationSettings == nil {
		return nil, false
	}
	return o.TessellationSettings, true
}

// HasTessellationSettings returns a boolean if a field has been set.
func (o *BTPartStudioDisplayDataReference4360) HasTessellationSettings() bool {
	if o != nil && o.TessellationSettings != nil {
		return true
	}

	return false
}

// SetTessellationSettings gets a reference to the given map[string][]int32 and assigns it to the TessellationSettings field.
func (o *BTPartStudioDisplayDataReference4360) SetTessellationSettings(v map[string][]int32) {
	o.TessellationSettings = &v
}

func (o BTPartStudioDisplayDataReference4360) MarshalJSON() ([]byte, error) {
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
	if o.IsNoop != nil {
		toSerialize["isNoop"] = o.IsNoop
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
	if o.NumberOfSketchEntities != nil {
		toSerialize["numberOfSketchEntities"] = o.NumberOfSketchEntities
	}
	if o.VersionForRasterization != nil {
		toSerialize["versionForRasterization"] = o.VersionForRasterization
	}
	if o.CacheDataPath != nil {
		toSerialize["cacheDataPath"] = o.CacheDataPath
	}
	if o.TessellationSettings != nil {
		toSerialize["tessellationSettings"] = o.TessellationSettings
	}
	return json.Marshal(toSerialize)
}

type NullableBTPartStudioDisplayDataReference4360 struct {
	value *BTPartStudioDisplayDataReference4360
	isSet bool
}

func (v NullableBTPartStudioDisplayDataReference4360) Get() *BTPartStudioDisplayDataReference4360 {
	return v.value
}

func (v *NullableBTPartStudioDisplayDataReference4360) Set(val *BTPartStudioDisplayDataReference4360) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPartStudioDisplayDataReference4360) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPartStudioDisplayDataReference4360) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPartStudioDisplayDataReference4360(val *BTPartStudioDisplayDataReference4360) *NullableBTPartStudioDisplayDataReference4360 {
	return &NullableBTPartStudioDisplayDataReference4360{value: val, isSet: true}
}

func (v NullableBTPartStudioDisplayDataReference4360) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPartStudioDisplayDataReference4360) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

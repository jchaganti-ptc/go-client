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

// BTViewManipulationMouseKeyMappingInfo struct for BTViewManipulationMouseKeyMappingInfo
type BTViewManipulationMouseKeyMappingInfo struct {
	AxisRotate3DMapping []BTKeyMouseValuesInfo `json:"axisRotate3DMapping,omitempty"`
	Pan2DMapping        []BTKeyMouseValuesInfo `json:"pan2DMapping,omitempty"`
	Pan3DMapping        []BTKeyMouseValuesInfo `json:"pan3DMapping,omitempty"`
	Rotate3DMapping     []BTKeyMouseValuesInfo `json:"rotate3DMapping,omitempty"`
	Zoom2DMapping       []BTKeyMouseValuesInfo `json:"zoom2DMapping,omitempty"`
	Zoom3DMapping       []BTKeyMouseValuesInfo `json:"zoom3DMapping,omitempty"`
}

// NewBTViewManipulationMouseKeyMappingInfo instantiates a new BTViewManipulationMouseKeyMappingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTViewManipulationMouseKeyMappingInfo() *BTViewManipulationMouseKeyMappingInfo {
	this := BTViewManipulationMouseKeyMappingInfo{}
	return &this
}

// NewBTViewManipulationMouseKeyMappingInfoWithDefaults instantiates a new BTViewManipulationMouseKeyMappingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTViewManipulationMouseKeyMappingInfoWithDefaults() *BTViewManipulationMouseKeyMappingInfo {
	this := BTViewManipulationMouseKeyMappingInfo{}
	return &this
}

// GetAxisRotate3DMapping returns the AxisRotate3DMapping field value if set, zero value otherwise.
func (o *BTViewManipulationMouseKeyMappingInfo) GetAxisRotate3DMapping() []BTKeyMouseValuesInfo {
	if o == nil || o.AxisRotate3DMapping == nil {
		var ret []BTKeyMouseValuesInfo
		return ret
	}
	return o.AxisRotate3DMapping
}

// GetAxisRotate3DMappingOk returns a tuple with the AxisRotate3DMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTViewManipulationMouseKeyMappingInfo) GetAxisRotate3DMappingOk() ([]BTKeyMouseValuesInfo, bool) {
	if o == nil || o.AxisRotate3DMapping == nil {
		return nil, false
	}
	return o.AxisRotate3DMapping, true
}

// HasAxisRotate3DMapping returns a boolean if a field has been set.
func (o *BTViewManipulationMouseKeyMappingInfo) HasAxisRotate3DMapping() bool {
	if o != nil && o.AxisRotate3DMapping != nil {
		return true
	}

	return false
}

// SetAxisRotate3DMapping gets a reference to the given []BTKeyMouseValuesInfo and assigns it to the AxisRotate3DMapping field.
func (o *BTViewManipulationMouseKeyMappingInfo) SetAxisRotate3DMapping(v []BTKeyMouseValuesInfo) {
	o.AxisRotate3DMapping = v
}

// GetPan2DMapping returns the Pan2DMapping field value if set, zero value otherwise.
func (o *BTViewManipulationMouseKeyMappingInfo) GetPan2DMapping() []BTKeyMouseValuesInfo {
	if o == nil || o.Pan2DMapping == nil {
		var ret []BTKeyMouseValuesInfo
		return ret
	}
	return o.Pan2DMapping
}

// GetPan2DMappingOk returns a tuple with the Pan2DMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTViewManipulationMouseKeyMappingInfo) GetPan2DMappingOk() ([]BTKeyMouseValuesInfo, bool) {
	if o == nil || o.Pan2DMapping == nil {
		return nil, false
	}
	return o.Pan2DMapping, true
}

// HasPan2DMapping returns a boolean if a field has been set.
func (o *BTViewManipulationMouseKeyMappingInfo) HasPan2DMapping() bool {
	if o != nil && o.Pan2DMapping != nil {
		return true
	}

	return false
}

// SetPan2DMapping gets a reference to the given []BTKeyMouseValuesInfo and assigns it to the Pan2DMapping field.
func (o *BTViewManipulationMouseKeyMappingInfo) SetPan2DMapping(v []BTKeyMouseValuesInfo) {
	o.Pan2DMapping = v
}

// GetPan3DMapping returns the Pan3DMapping field value if set, zero value otherwise.
func (o *BTViewManipulationMouseKeyMappingInfo) GetPan3DMapping() []BTKeyMouseValuesInfo {
	if o == nil || o.Pan3DMapping == nil {
		var ret []BTKeyMouseValuesInfo
		return ret
	}
	return o.Pan3DMapping
}

// GetPan3DMappingOk returns a tuple with the Pan3DMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTViewManipulationMouseKeyMappingInfo) GetPan3DMappingOk() ([]BTKeyMouseValuesInfo, bool) {
	if o == nil || o.Pan3DMapping == nil {
		return nil, false
	}
	return o.Pan3DMapping, true
}

// HasPan3DMapping returns a boolean if a field has been set.
func (o *BTViewManipulationMouseKeyMappingInfo) HasPan3DMapping() bool {
	if o != nil && o.Pan3DMapping != nil {
		return true
	}

	return false
}

// SetPan3DMapping gets a reference to the given []BTKeyMouseValuesInfo and assigns it to the Pan3DMapping field.
func (o *BTViewManipulationMouseKeyMappingInfo) SetPan3DMapping(v []BTKeyMouseValuesInfo) {
	o.Pan3DMapping = v
}

// GetRotate3DMapping returns the Rotate3DMapping field value if set, zero value otherwise.
func (o *BTViewManipulationMouseKeyMappingInfo) GetRotate3DMapping() []BTKeyMouseValuesInfo {
	if o == nil || o.Rotate3DMapping == nil {
		var ret []BTKeyMouseValuesInfo
		return ret
	}
	return o.Rotate3DMapping
}

// GetRotate3DMappingOk returns a tuple with the Rotate3DMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTViewManipulationMouseKeyMappingInfo) GetRotate3DMappingOk() ([]BTKeyMouseValuesInfo, bool) {
	if o == nil || o.Rotate3DMapping == nil {
		return nil, false
	}
	return o.Rotate3DMapping, true
}

// HasRotate3DMapping returns a boolean if a field has been set.
func (o *BTViewManipulationMouseKeyMappingInfo) HasRotate3DMapping() bool {
	if o != nil && o.Rotate3DMapping != nil {
		return true
	}

	return false
}

// SetRotate3DMapping gets a reference to the given []BTKeyMouseValuesInfo and assigns it to the Rotate3DMapping field.
func (o *BTViewManipulationMouseKeyMappingInfo) SetRotate3DMapping(v []BTKeyMouseValuesInfo) {
	o.Rotate3DMapping = v
}

// GetZoom2DMapping returns the Zoom2DMapping field value if set, zero value otherwise.
func (o *BTViewManipulationMouseKeyMappingInfo) GetZoom2DMapping() []BTKeyMouseValuesInfo {
	if o == nil || o.Zoom2DMapping == nil {
		var ret []BTKeyMouseValuesInfo
		return ret
	}
	return o.Zoom2DMapping
}

// GetZoom2DMappingOk returns a tuple with the Zoom2DMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTViewManipulationMouseKeyMappingInfo) GetZoom2DMappingOk() ([]BTKeyMouseValuesInfo, bool) {
	if o == nil || o.Zoom2DMapping == nil {
		return nil, false
	}
	return o.Zoom2DMapping, true
}

// HasZoom2DMapping returns a boolean if a field has been set.
func (o *BTViewManipulationMouseKeyMappingInfo) HasZoom2DMapping() bool {
	if o != nil && o.Zoom2DMapping != nil {
		return true
	}

	return false
}

// SetZoom2DMapping gets a reference to the given []BTKeyMouseValuesInfo and assigns it to the Zoom2DMapping field.
func (o *BTViewManipulationMouseKeyMappingInfo) SetZoom2DMapping(v []BTKeyMouseValuesInfo) {
	o.Zoom2DMapping = v
}

// GetZoom3DMapping returns the Zoom3DMapping field value if set, zero value otherwise.
func (o *BTViewManipulationMouseKeyMappingInfo) GetZoom3DMapping() []BTKeyMouseValuesInfo {
	if o == nil || o.Zoom3DMapping == nil {
		var ret []BTKeyMouseValuesInfo
		return ret
	}
	return o.Zoom3DMapping
}

// GetZoom3DMappingOk returns a tuple with the Zoom3DMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTViewManipulationMouseKeyMappingInfo) GetZoom3DMappingOk() ([]BTKeyMouseValuesInfo, bool) {
	if o == nil || o.Zoom3DMapping == nil {
		return nil, false
	}
	return o.Zoom3DMapping, true
}

// HasZoom3DMapping returns a boolean if a field has been set.
func (o *BTViewManipulationMouseKeyMappingInfo) HasZoom3DMapping() bool {
	if o != nil && o.Zoom3DMapping != nil {
		return true
	}

	return false
}

// SetZoom3DMapping gets a reference to the given []BTKeyMouseValuesInfo and assigns it to the Zoom3DMapping field.
func (o *BTViewManipulationMouseKeyMappingInfo) SetZoom3DMapping(v []BTKeyMouseValuesInfo) {
	o.Zoom3DMapping = v
}

func (o BTViewManipulationMouseKeyMappingInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AxisRotate3DMapping != nil {
		toSerialize["axisRotate3DMapping"] = o.AxisRotate3DMapping
	}
	if o.Pan2DMapping != nil {
		toSerialize["pan2DMapping"] = o.Pan2DMapping
	}
	if o.Pan3DMapping != nil {
		toSerialize["pan3DMapping"] = o.Pan3DMapping
	}
	if o.Rotate3DMapping != nil {
		toSerialize["rotate3DMapping"] = o.Rotate3DMapping
	}
	if o.Zoom2DMapping != nil {
		toSerialize["zoom2DMapping"] = o.Zoom2DMapping
	}
	if o.Zoom3DMapping != nil {
		toSerialize["zoom3DMapping"] = o.Zoom3DMapping
	}
	return json.Marshal(toSerialize)
}

type NullableBTViewManipulationMouseKeyMappingInfo struct {
	value *BTViewManipulationMouseKeyMappingInfo
	isSet bool
}

func (v NullableBTViewManipulationMouseKeyMappingInfo) Get() *BTViewManipulationMouseKeyMappingInfo {
	return v.value
}

func (v *NullableBTViewManipulationMouseKeyMappingInfo) Set(val *BTViewManipulationMouseKeyMappingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTViewManipulationMouseKeyMappingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTViewManipulationMouseKeyMappingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTViewManipulationMouseKeyMappingInfo(val *BTViewManipulationMouseKeyMappingInfo) *NullableBTViewManipulationMouseKeyMappingInfo {
	return &NullableBTViewManipulationMouseKeyMappingInfo{value: val, isSet: true}
}

func (v NullableBTViewManipulationMouseKeyMappingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTViewManipulationMouseKeyMappingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

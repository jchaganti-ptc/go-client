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
	"fmt"
)

// BTMetadataObjectType the model 'BTMetadataObjectType'
type BTMetadataObjectType string

// List of BTMetadataObjectType
const (
	BTMetadataObjectTypeGlobal             BTMetadataObjectType = "GLOBAL"
	BTMetadataObjectTypeDocument           BTMetadataObjectType = "DOCUMENT"
	BTMetadataObjectTypePart               BTMetadataObjectType = "PART"
	BTMetadataObjectTypeAssembly           BTMetadataObjectType = "ASSEMBLY"
	BTMetadataObjectTypeDrawing            BTMetadataObjectType = "DRAWING"
	BTMetadataObjectTypePartStudio         BTMetadataObjectType = "PART_STUDIO"
	BTMetadataObjectTypeBlobElement        BTMetadataObjectType = "BLOB_ELEMENT"
	BTMetadataObjectTypeAppElement         BTMetadataObjectType = "APP_ELEMENT"
	BTMetadataObjectTypeVersion            BTMetadataObjectType = "VERSION"
	BTMetadataObjectTypeWorkspace          BTMetadataObjectType = "WORKSPACE"
	BTMetadataObjectTypeProject            BTMetadataObjectType = "PROJECT"
	BTMetadataObjectTypeItem               BTMetadataObjectType = "ITEM"
	BTMetadataObjectTypeFeatureStudio      BTMetadataObjectType = "FEATURE_STUDIO"
	BTMetadataObjectTypeChangeRequest      BTMetadataObjectType = "CHANGE_REQUEST"
	BTMetadataObjectTypeTask               BTMetadataObjectType = "TASK"
	BTMetadataObjectTypeChangeOrder        BTMetadataObjectType = "CHANGE_ORDER"
	BTMetadataObjectTypeChangeTask         BTMetadataObjectType = "CHANGE_TASK"
	BTMetadataObjectTypeVariableStudio     BTMetadataObjectType = "VARIABLE_STUDIO"
	BTMetadataObjectTypeDrawingAnnotations BTMetadataObjectType = "DRAWING_ANNOTATIONS"
)

// All allowed values of BTMetadataObjectType enum
var AllowedBTMetadataObjectTypeEnumValues = []BTMetadataObjectType{
	"GLOBAL",
	"DOCUMENT",
	"PART",
	"ASSEMBLY",
	"DRAWING",
	"PART_STUDIO",
	"BLOB_ELEMENT",
	"APP_ELEMENT",
	"VERSION",
	"WORKSPACE",
	"PROJECT",
	"ITEM",
	"FEATURE_STUDIO",
	"CHANGE_REQUEST",
	"TASK",
	"CHANGE_ORDER",
	"CHANGE_TASK",
	"VARIABLE_STUDIO",
	"DRAWING_ANNOTATIONS",
}

func (v *BTMetadataObjectType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BTMetadataObjectType(value)
	for _, existing := range AllowedBTMetadataObjectTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BTMetadataObjectType", value)
}

// NewBTMetadataObjectTypeFromValue returns a pointer to a valid BTMetadataObjectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBTMetadataObjectTypeFromValue(v string) (*BTMetadataObjectType, error) {
	ev := BTMetadataObjectType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BTMetadataObjectType: valid values are %v", v, AllowedBTMetadataObjectTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BTMetadataObjectType) IsValid() bool {
	for _, existing := range AllowedBTMetadataObjectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BTMetadataObjectType value
func (v BTMetadataObjectType) Ptr() *BTMetadataObjectType {
	return &v
}

type NullableBTMetadataObjectType struct {
	value *BTMetadataObjectType
	isSet bool
}

func (v NullableBTMetadataObjectType) Get() *BTMetadataObjectType {
	return v.value
}

func (v *NullableBTMetadataObjectType) Set(val *BTMetadataObjectType) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMetadataObjectType) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMetadataObjectType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMetadataObjectType(val *BTMetadataObjectType) *NullableBTMetadataObjectType {
	return &NullableBTMetadataObjectType{value: val, isSet: true}
}

func (v NullableBTMetadataObjectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMetadataObjectType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

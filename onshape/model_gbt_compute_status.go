/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.165.17666-197c9d1638c5
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// GBTComputeStatus the model 'GBTComputeStatus'
type GBTComputeStatus string

// List of GBTComputeStatus
const (
	GBTComputeStatusComputed     GBTComputeStatus = "COMPUTED"
	GBTComputeStatusStale        GBTComputeStatus = "STALE"
	GBTComputeStatusComputing    GBTComputeStatus = "COMPUTING"
	GBTComputeStatusError        GBTComputeStatus = "ERROR"
	GBTComputeStatusUnderdefined GBTComputeStatus = "UNDERDEFINED"
	GBTComputeStatusPreparing    GBTComputeStatus = "PREPARING"
	GBTComputeStatusUnknown      GBTComputeStatus = "UNKNOWN"
)

// All allowed values of GBTComputeStatus enum
var AllowedGBTComputeStatusEnumValues = []GBTComputeStatus{
	"COMPUTED",
	"STALE",
	"COMPUTING",
	"ERROR",
	"UNDERDEFINED",
	"PREPARING",
	"UNKNOWN",
}

func (v *GBTComputeStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GBTComputeStatus(value)
	for _, existing := range AllowedGBTComputeStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GBTComputeStatus", value)
}

// NewGBTComputeStatusFromValue returns a pointer to a valid GBTComputeStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGBTComputeStatusFromValue(v string) (*GBTComputeStatus, error) {
	ev := GBTComputeStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GBTComputeStatus: valid values are %v", v, AllowedGBTComputeStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GBTComputeStatus) IsValid() bool {
	for _, existing := range AllowedGBTComputeStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GBTComputeStatus value
func (v GBTComputeStatus) Ptr() *GBTComputeStatus {
	return &v
}

type NullableGBTComputeStatus struct {
	value *GBTComputeStatus
	isSet bool
}

func (v NullableGBTComputeStatus) Get() *GBTComputeStatus {
	return v.value
}

func (v *NullableGBTComputeStatus) Set(val *GBTComputeStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGBTComputeStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGBTComputeStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGBTComputeStatus(val *GBTComputeStatus) *NullableGBTComputeStatus {
	return &NullableGBTComputeStatus{value: val, isSet: true}
}

func (v NullableGBTComputeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGBTComputeStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
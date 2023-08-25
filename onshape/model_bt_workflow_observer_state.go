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
	"fmt"
)

// BTWorkflowObserverState the model 'BTWorkflowObserverState'
type BTWorkflowObserverState string

// List of BTWorkflowObserverState
const (
	BTWorkflowObserverStateNone       BTWorkflowObserverState = "NONE"
	BTWorkflowObserverStateApproved   BTWorkflowObserverState = "APPROVED"
	BTWorkflowObserverStateRejected   BTWorkflowObserverState = "REJECTED"
	BTWorkflowObserverStateOverridden BTWorkflowObserverState = "OVERRIDDEN"
)

// All allowed values of BTWorkflowObserverState enum
var AllowedBTWorkflowObserverStateEnumValues = []BTWorkflowObserverState{
	"NONE",
	"APPROVED",
	"REJECTED",
	"OVERRIDDEN",
}

func (v *BTWorkflowObserverState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BTWorkflowObserverState(value)
	for _, existing := range AllowedBTWorkflowObserverStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BTWorkflowObserverState", value)
}

// NewBTWorkflowObserverStateFromValue returns a pointer to a valid BTWorkflowObserverState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBTWorkflowObserverStateFromValue(v string) (*BTWorkflowObserverState, error) {
	ev := BTWorkflowObserverState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BTWorkflowObserverState: valid values are %v", v, AllowedBTWorkflowObserverStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BTWorkflowObserverState) IsValid() bool {
	for _, existing := range AllowedBTWorkflowObserverStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BTWorkflowObserverState value
func (v BTWorkflowObserverState) Ptr() *BTWorkflowObserverState {
	return &v
}

type NullableBTWorkflowObserverState struct {
	value *BTWorkflowObserverState
	isSet bool
}

func (v NullableBTWorkflowObserverState) Get() *BTWorkflowObserverState {
	return v.value
}

func (v *NullableBTWorkflowObserverState) Set(val *BTWorkflowObserverState) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWorkflowObserverState) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWorkflowObserverState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWorkflowObserverState(val *BTWorkflowObserverState) *NullableBTWorkflowObserverState {
	return &NullableBTWorkflowObserverState{value: val, isSet: true}
}

func (v NullableBTWorkflowObserverState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWorkflowObserverState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

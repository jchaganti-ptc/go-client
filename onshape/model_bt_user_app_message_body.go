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

// BTUserAppMessageBody struct for BTUserAppMessageBody
type BTUserAppMessageBody struct {
	AppElementSessionId *string                    `json:"appElementSessionId,omitempty"`
	ClientId            *string                    `json:"clientId,omitempty"`
	Data                *string                    `json:"data,omitempty"`
	Event               *string                    `json:"event,omitempty"`
	IdentityId          *string                    `json:"identityId,omitempty"`
	MessageId           *string                    `json:"messageId,omitempty"`
	SettingType         *BTApplicationSettingsType `json:"settingType,omitempty"`
	Timestamp           *JSONTime                  `json:"timestamp,omitempty"`
	WebhookId           *string                    `json:"webhookId,omitempty"`
}

// NewBTUserAppMessageBody instantiates a new BTUserAppMessageBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUserAppMessageBody() *BTUserAppMessageBody {
	this := BTUserAppMessageBody{}
	return &this
}

// NewBTUserAppMessageBodyWithDefaults instantiates a new BTUserAppMessageBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUserAppMessageBodyWithDefaults() *BTUserAppMessageBody {
	this := BTUserAppMessageBody{}
	return &this
}

// GetAppElementSessionId returns the AppElementSessionId field value if set, zero value otherwise.
func (o *BTUserAppMessageBody) GetAppElementSessionId() string {
	if o == nil || o.AppElementSessionId == nil {
		var ret string
		return ret
	}
	return *o.AppElementSessionId
}

// GetAppElementSessionIdOk returns a tuple with the AppElementSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAppMessageBody) GetAppElementSessionIdOk() (*string, bool) {
	if o == nil || o.AppElementSessionId == nil {
		return nil, false
	}
	return o.AppElementSessionId, true
}

// HasAppElementSessionId returns a boolean if a field has been set.
func (o *BTUserAppMessageBody) HasAppElementSessionId() bool {
	if o != nil && o.AppElementSessionId != nil {
		return true
	}

	return false
}

// SetAppElementSessionId gets a reference to the given string and assigns it to the AppElementSessionId field.
func (o *BTUserAppMessageBody) SetAppElementSessionId(v string) {
	o.AppElementSessionId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *BTUserAppMessageBody) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAppMessageBody) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *BTUserAppMessageBody) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *BTUserAppMessageBody) SetClientId(v string) {
	o.ClientId = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BTUserAppMessageBody) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAppMessageBody) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BTUserAppMessageBody) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *BTUserAppMessageBody) SetData(v string) {
	o.Data = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *BTUserAppMessageBody) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAppMessageBody) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *BTUserAppMessageBody) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *BTUserAppMessageBody) SetEvent(v string) {
	o.Event = &v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *BTUserAppMessageBody) GetIdentityId() string {
	if o == nil || o.IdentityId == nil {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAppMessageBody) GetIdentityIdOk() (*string, bool) {
	if o == nil || o.IdentityId == nil {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *BTUserAppMessageBody) HasIdentityId() bool {
	if o != nil && o.IdentityId != nil {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *BTUserAppMessageBody) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *BTUserAppMessageBody) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAppMessageBody) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *BTUserAppMessageBody) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *BTUserAppMessageBody) SetMessageId(v string) {
	o.MessageId = &v
}

// GetSettingType returns the SettingType field value if set, zero value otherwise.
func (o *BTUserAppMessageBody) GetSettingType() BTApplicationSettingsType {
	if o == nil || o.SettingType == nil {
		var ret BTApplicationSettingsType
		return ret
	}
	return *o.SettingType
}

// GetSettingTypeOk returns a tuple with the SettingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAppMessageBody) GetSettingTypeOk() (*BTApplicationSettingsType, bool) {
	if o == nil || o.SettingType == nil {
		return nil, false
	}
	return o.SettingType, true
}

// HasSettingType returns a boolean if a field has been set.
func (o *BTUserAppMessageBody) HasSettingType() bool {
	if o != nil && o.SettingType != nil {
		return true
	}

	return false
}

// SetSettingType gets a reference to the given BTApplicationSettingsType and assigns it to the SettingType field.
func (o *BTUserAppMessageBody) SetSettingType(v BTApplicationSettingsType) {
	o.SettingType = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *BTUserAppMessageBody) GetTimestamp() JSONTime {
	if o == nil || o.Timestamp == nil {
		var ret JSONTime
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAppMessageBody) GetTimestampOk() (*JSONTime, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *BTUserAppMessageBody) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given JSONTime and assigns it to the Timestamp field.
func (o *BTUserAppMessageBody) SetTimestamp(v JSONTime) {
	o.Timestamp = &v
}

// GetWebhookId returns the WebhookId field value if set, zero value otherwise.
func (o *BTUserAppMessageBody) GetWebhookId() string {
	if o == nil || o.WebhookId == nil {
		var ret string
		return ret
	}
	return *o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAppMessageBody) GetWebhookIdOk() (*string, bool) {
	if o == nil || o.WebhookId == nil {
		return nil, false
	}
	return o.WebhookId, true
}

// HasWebhookId returns a boolean if a field has been set.
func (o *BTUserAppMessageBody) HasWebhookId() bool {
	if o != nil && o.WebhookId != nil {
		return true
	}

	return false
}

// SetWebhookId gets a reference to the given string and assigns it to the WebhookId field.
func (o *BTUserAppMessageBody) SetWebhookId(v string) {
	o.WebhookId = &v
}

func (o BTUserAppMessageBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppElementSessionId != nil {
		toSerialize["appElementSessionId"] = o.AppElementSessionId
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.IdentityId != nil {
		toSerialize["identityId"] = o.IdentityId
	}
	if o.MessageId != nil {
		toSerialize["messageId"] = o.MessageId
	}
	if o.SettingType != nil {
		toSerialize["settingType"] = o.SettingType
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.WebhookId != nil {
		toSerialize["webhookId"] = o.WebhookId
	}
	return json.Marshal(toSerialize)
}

type NullableBTUserAppMessageBody struct {
	value *BTUserAppMessageBody
	isSet bool
}

func (v NullableBTUserAppMessageBody) Get() *BTUserAppMessageBody {
	return v.value
}

func (v *NullableBTUserAppMessageBody) Set(val *BTUserAppMessageBody) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUserAppMessageBody) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUserAppMessageBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUserAppMessageBody(val *BTUserAppMessageBody) *NullableBTUserAppMessageBody {
	return &NullableBTUserAppMessageBody{value: val, isSet: true}
}

func (v NullableBTUserAppMessageBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUserAppMessageBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

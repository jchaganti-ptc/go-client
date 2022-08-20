/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6126-5c3a878ad24b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// Coupon struct for Coupon
type Coupon struct {
	AmountOff        *int64             `json:"amountOff,omitempty"`
	Created          *int64             `json:"created,omitempty"`
	Currency         *string            `json:"currency,omitempty"`
	Duration         *string            `json:"duration,omitempty"`
	DurationInMonths *int32             `json:"durationInMonths,omitempty"`
	Id               *string            `json:"id,omitempty"`
	Livemode         *bool              `json:"livemode,omitempty"`
	MaxRedemptions   *int64             `json:"maxRedemptions,omitempty"`
	Metadata         *map[string]string `json:"metadata,omitempty"`
	Object           *string            `json:"object,omitempty"`
	PercentOff       *int32             `json:"percentOff,omitempty"`
	RedeemBy         *int64             `json:"redeemBy,omitempty"`
	TimesRedeemed    *int32             `json:"timesRedeemed,omitempty"`
	Valid            *bool              `json:"valid,omitempty"`
}

// NewCoupon instantiates a new Coupon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoupon() *Coupon {
	this := Coupon{}
	return &this
}

// NewCouponWithDefaults instantiates a new Coupon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponWithDefaults() *Coupon {
	this := Coupon{}
	return &this
}

// GetAmountOff returns the AmountOff field value if set, zero value otherwise.
func (o *Coupon) GetAmountOff() int64 {
	if o == nil || o.AmountOff == nil {
		var ret int64
		return ret
	}
	return *o.AmountOff
}

// GetAmountOffOk returns a tuple with the AmountOff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetAmountOffOk() (*int64, bool) {
	if o == nil || o.AmountOff == nil {
		return nil, false
	}
	return o.AmountOff, true
}

// HasAmountOff returns a boolean if a field has been set.
func (o *Coupon) HasAmountOff() bool {
	if o != nil && o.AmountOff != nil {
		return true
	}

	return false
}

// SetAmountOff gets a reference to the given int64 and assigns it to the AmountOff field.
func (o *Coupon) SetAmountOff(v int64) {
	o.AmountOff = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Coupon) GetCreated() int64 {
	if o == nil || o.Created == nil {
		var ret int64
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetCreatedOk() (*int64, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Coupon) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int64 and assigns it to the Created field.
func (o *Coupon) SetCreated(v int64) {
	o.Created = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Coupon) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Coupon) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Coupon) SetCurrency(v string) {
	o.Currency = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *Coupon) GetDuration() string {
	if o == nil || o.Duration == nil {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetDurationOk() (*string, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *Coupon) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *Coupon) SetDuration(v string) {
	o.Duration = &v
}

// GetDurationInMonths returns the DurationInMonths field value if set, zero value otherwise.
func (o *Coupon) GetDurationInMonths() int32 {
	if o == nil || o.DurationInMonths == nil {
		var ret int32
		return ret
	}
	return *o.DurationInMonths
}

// GetDurationInMonthsOk returns a tuple with the DurationInMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetDurationInMonthsOk() (*int32, bool) {
	if o == nil || o.DurationInMonths == nil {
		return nil, false
	}
	return o.DurationInMonths, true
}

// HasDurationInMonths returns a boolean if a field has been set.
func (o *Coupon) HasDurationInMonths() bool {
	if o != nil && o.DurationInMonths != nil {
		return true
	}

	return false
}

// SetDurationInMonths gets a reference to the given int32 and assigns it to the DurationInMonths field.
func (o *Coupon) SetDurationInMonths(v int32) {
	o.DurationInMonths = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Coupon) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Coupon) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Coupon) SetId(v string) {
	o.Id = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *Coupon) GetLivemode() bool {
	if o == nil || o.Livemode == nil {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetLivemodeOk() (*bool, bool) {
	if o == nil || o.Livemode == nil {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *Coupon) HasLivemode() bool {
	if o != nil && o.Livemode != nil {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *Coupon) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetMaxRedemptions returns the MaxRedemptions field value if set, zero value otherwise.
func (o *Coupon) GetMaxRedemptions() int64 {
	if o == nil || o.MaxRedemptions == nil {
		var ret int64
		return ret
	}
	return *o.MaxRedemptions
}

// GetMaxRedemptionsOk returns a tuple with the MaxRedemptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetMaxRedemptionsOk() (*int64, bool) {
	if o == nil || o.MaxRedemptions == nil {
		return nil, false
	}
	return o.MaxRedemptions, true
}

// HasMaxRedemptions returns a boolean if a field has been set.
func (o *Coupon) HasMaxRedemptions() bool {
	if o != nil && o.MaxRedemptions != nil {
		return true
	}

	return false
}

// SetMaxRedemptions gets a reference to the given int64 and assigns it to the MaxRedemptions field.
func (o *Coupon) SetMaxRedemptions(v int64) {
	o.MaxRedemptions = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Coupon) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Coupon) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *Coupon) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *Coupon) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *Coupon) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *Coupon) SetObject(v string) {
	o.Object = &v
}

// GetPercentOff returns the PercentOff field value if set, zero value otherwise.
func (o *Coupon) GetPercentOff() int32 {
	if o == nil || o.PercentOff == nil {
		var ret int32
		return ret
	}
	return *o.PercentOff
}

// GetPercentOffOk returns a tuple with the PercentOff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetPercentOffOk() (*int32, bool) {
	if o == nil || o.PercentOff == nil {
		return nil, false
	}
	return o.PercentOff, true
}

// HasPercentOff returns a boolean if a field has been set.
func (o *Coupon) HasPercentOff() bool {
	if o != nil && o.PercentOff != nil {
		return true
	}

	return false
}

// SetPercentOff gets a reference to the given int32 and assigns it to the PercentOff field.
func (o *Coupon) SetPercentOff(v int32) {
	o.PercentOff = &v
}

// GetRedeemBy returns the RedeemBy field value if set, zero value otherwise.
func (o *Coupon) GetRedeemBy() int64 {
	if o == nil || o.RedeemBy == nil {
		var ret int64
		return ret
	}
	return *o.RedeemBy
}

// GetRedeemByOk returns a tuple with the RedeemBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetRedeemByOk() (*int64, bool) {
	if o == nil || o.RedeemBy == nil {
		return nil, false
	}
	return o.RedeemBy, true
}

// HasRedeemBy returns a boolean if a field has been set.
func (o *Coupon) HasRedeemBy() bool {
	if o != nil && o.RedeemBy != nil {
		return true
	}

	return false
}

// SetRedeemBy gets a reference to the given int64 and assigns it to the RedeemBy field.
func (o *Coupon) SetRedeemBy(v int64) {
	o.RedeemBy = &v
}

// GetTimesRedeemed returns the TimesRedeemed field value if set, zero value otherwise.
func (o *Coupon) GetTimesRedeemed() int32 {
	if o == nil || o.TimesRedeemed == nil {
		var ret int32
		return ret
	}
	return *o.TimesRedeemed
}

// GetTimesRedeemedOk returns a tuple with the TimesRedeemed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetTimesRedeemedOk() (*int32, bool) {
	if o == nil || o.TimesRedeemed == nil {
		return nil, false
	}
	return o.TimesRedeemed, true
}

// HasTimesRedeemed returns a boolean if a field has been set.
func (o *Coupon) HasTimesRedeemed() bool {
	if o != nil && o.TimesRedeemed != nil {
		return true
	}

	return false
}

// SetTimesRedeemed gets a reference to the given int32 and assigns it to the TimesRedeemed field.
func (o *Coupon) SetTimesRedeemed(v int32) {
	o.TimesRedeemed = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *Coupon) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *Coupon) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *Coupon) SetValid(v bool) {
	o.Valid = &v
}

func (o Coupon) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AmountOff != nil {
		toSerialize["amountOff"] = o.AmountOff
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.DurationInMonths != nil {
		toSerialize["durationInMonths"] = o.DurationInMonths
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Livemode != nil {
		toSerialize["livemode"] = o.Livemode
	}
	if o.MaxRedemptions != nil {
		toSerialize["maxRedemptions"] = o.MaxRedemptions
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.PercentOff != nil {
		toSerialize["percentOff"] = o.PercentOff
	}
	if o.RedeemBy != nil {
		toSerialize["redeemBy"] = o.RedeemBy
	}
	if o.TimesRedeemed != nil {
		toSerialize["timesRedeemed"] = o.TimesRedeemed
	}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}
	return json.Marshal(toSerialize)
}

type NullableCoupon struct {
	value *Coupon
	isSet bool
}

func (v NullableCoupon) Get() *Coupon {
	return v.value
}

func (v *NullableCoupon) Set(val *Coupon) {
	v.value = val
	v.isSet = true
}

func (v NullableCoupon) IsSet() bool {
	return v.isSet
}

func (v *NullableCoupon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoupon(val *Coupon) *NullableCoupon {
	return &NullableCoupon{value: val, isSet: true}
}

func (v NullableCoupon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoupon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

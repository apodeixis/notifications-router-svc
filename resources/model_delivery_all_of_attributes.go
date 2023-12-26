/*
Apodeixis notifications-router-svc

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the DeliveryAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryAllOfAttributes{}

// DeliveryAllOfAttributes struct for DeliveryAllOfAttributes
type DeliveryAllOfAttributes struct {
	Destination     string `json:"destination"`
	DestinationType string `json:"destination_type"`
	Status          string `json:"status"`
	SentAt          *int64 `json:"sent_at,omitempty"`
}

type _DeliveryAllOfAttributes DeliveryAllOfAttributes

// NewDeliveryAllOfAttributes instantiates a new DeliveryAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryAllOfAttributes(destination string, destinationType string, status string) *DeliveryAllOfAttributes {
	this := DeliveryAllOfAttributes{}
	this.Destination = destination
	this.DestinationType = destinationType
	this.Status = status
	return &this
}

// NewDeliveryAllOfAttributesWithDefaults instantiates a new DeliveryAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryAllOfAttributesWithDefaults() *DeliveryAllOfAttributes {
	this := DeliveryAllOfAttributes{}
	return &this
}

// GetDestination returns the Destination field value
func (o *DeliveryAllOfAttributes) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *DeliveryAllOfAttributes) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *DeliveryAllOfAttributes) SetDestination(v string) {
	o.Destination = v
}

// GetDestinationType returns the DestinationType field value
func (o *DeliveryAllOfAttributes) GetDestinationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationType
}

// GetDestinationTypeOk returns a tuple with the DestinationType field value
// and a boolean to check if the value has been set.
func (o *DeliveryAllOfAttributes) GetDestinationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationType, true
}

// SetDestinationType sets field value
func (o *DeliveryAllOfAttributes) SetDestinationType(v string) {
	o.DestinationType = v
}

// GetStatus returns the Status field value
func (o *DeliveryAllOfAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DeliveryAllOfAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DeliveryAllOfAttributes) SetStatus(v string) {
	o.Status = v
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *DeliveryAllOfAttributes) GetSentAt() int64 {
	if o == nil || IsNil(o.SentAt) {
		var ret int64
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryAllOfAttributes) GetSentAtOk() (*int64, bool) {
	if o == nil || IsNil(o.SentAt) {
		return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *DeliveryAllOfAttributes) HasSentAt() bool {
	if o != nil && !IsNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given int64 and assigns it to the SentAt field.
func (o *DeliveryAllOfAttributes) SetSentAt(v int64) {
	o.SentAt = &v
}

func (o DeliveryAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destination"] = o.Destination
	toSerialize["destination_type"] = o.DestinationType
	toSerialize["status"] = o.Status
	if !IsNil(o.SentAt) {
		toSerialize["sent_at"] = o.SentAt
	}
	return toSerialize, nil
}

func (o *DeliveryAllOfAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destination",
		"destination_type",
		"status",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDeliveryAllOfAttributes := _DeliveryAllOfAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeliveryAllOfAttributes)

	if err != nil {
		return err
	}

	*o = DeliveryAllOfAttributes(varDeliveryAllOfAttributes)

	return err
}

type NullableDeliveryAllOfAttributes struct {
	value *DeliveryAllOfAttributes
	isSet bool
}

func (v NullableDeliveryAllOfAttributes) Get() *DeliveryAllOfAttributes {
	return v.value
}

func (v *NullableDeliveryAllOfAttributes) Set(val *DeliveryAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryAllOfAttributes(val *DeliveryAllOfAttributes) *NullableDeliveryAllOfAttributes {
	return &NullableDeliveryAllOfAttributes{value: val, isSet: true}
}

func (v NullableDeliveryAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
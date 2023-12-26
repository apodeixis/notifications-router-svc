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

// checks if the NotificationAllOfRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationAllOfRelationships{}

// NotificationAllOfRelationships struct for NotificationAllOfRelationships
type NotificationAllOfRelationships struct {
	Deliveries NotificationAllOfRelationshipsDeliveries `json:"deliveries"`
}

type _NotificationAllOfRelationships NotificationAllOfRelationships

// NewNotificationAllOfRelationships instantiates a new NotificationAllOfRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationAllOfRelationships(deliveries NotificationAllOfRelationshipsDeliveries) *NotificationAllOfRelationships {
	this := NotificationAllOfRelationships{}
	this.Deliveries = deliveries
	return &this
}

// NewNotificationAllOfRelationshipsWithDefaults instantiates a new NotificationAllOfRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationAllOfRelationshipsWithDefaults() *NotificationAllOfRelationships {
	this := NotificationAllOfRelationships{}
	return &this
}

// GetDeliveries returns the Deliveries field value
func (o *NotificationAllOfRelationships) GetDeliveries() NotificationAllOfRelationshipsDeliveries {
	if o == nil {
		var ret NotificationAllOfRelationshipsDeliveries
		return ret
	}

	return o.Deliveries
}

// GetDeliveriesOk returns a tuple with the Deliveries field value
// and a boolean to check if the value has been set.
func (o *NotificationAllOfRelationships) GetDeliveriesOk() (*NotificationAllOfRelationshipsDeliveries, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deliveries, true
}

// SetDeliveries sets field value
func (o *NotificationAllOfRelationships) SetDeliveries(v NotificationAllOfRelationshipsDeliveries) {
	o.Deliveries = v
}

func (o NotificationAllOfRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationAllOfRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deliveries"] = o.Deliveries
	return toSerialize, nil
}

func (o *NotificationAllOfRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deliveries",
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

	varNotificationAllOfRelationships := _NotificationAllOfRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotificationAllOfRelationships)

	if err != nil {
		return err
	}

	*o = NotificationAllOfRelationships(varNotificationAllOfRelationships)

	return err
}

type NullableNotificationAllOfRelationships struct {
	value *NotificationAllOfRelationships
	isSet bool
}

func (v NullableNotificationAllOfRelationships) Get() *NotificationAllOfRelationships {
	return v.value
}

func (v *NullableNotificationAllOfRelationships) Set(val *NotificationAllOfRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationAllOfRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationAllOfRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationAllOfRelationships(val *NotificationAllOfRelationships) *NullableNotificationAllOfRelationships {
	return &NullableNotificationAllOfRelationships{value: val, isSet: true}
}

func (v NullableNotificationAllOfRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationAllOfRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
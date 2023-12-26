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

// checks if the NotificationAllOfRelationshipsDeliveries type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationAllOfRelationshipsDeliveries{}

// NotificationAllOfRelationshipsDeliveries struct for NotificationAllOfRelationshipsDeliveries
type NotificationAllOfRelationshipsDeliveries struct {
	Data []DeliveryKey `json:"data"`
}

type _NotificationAllOfRelationshipsDeliveries NotificationAllOfRelationshipsDeliveries

// NewNotificationAllOfRelationshipsDeliveries instantiates a new NotificationAllOfRelationshipsDeliveries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationAllOfRelationshipsDeliveries(data []DeliveryKey) *NotificationAllOfRelationshipsDeliveries {
	this := NotificationAllOfRelationshipsDeliveries{}
	this.Data = data
	return &this
}

// NewNotificationAllOfRelationshipsDeliveriesWithDefaults instantiates a new NotificationAllOfRelationshipsDeliveries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationAllOfRelationshipsDeliveriesWithDefaults() *NotificationAllOfRelationshipsDeliveries {
	this := NotificationAllOfRelationshipsDeliveries{}
	return &this
}

// GetData returns the Data field value
func (o *NotificationAllOfRelationshipsDeliveries) GetData() []DeliveryKey {
	if o == nil {
		var ret []DeliveryKey
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *NotificationAllOfRelationshipsDeliveries) GetDataOk() ([]DeliveryKey, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *NotificationAllOfRelationshipsDeliveries) SetData(v []DeliveryKey) {
	o.Data = v
}

func (o NotificationAllOfRelationshipsDeliveries) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationAllOfRelationshipsDeliveries) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *NotificationAllOfRelationshipsDeliveries) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varNotificationAllOfRelationshipsDeliveries := _NotificationAllOfRelationshipsDeliveries{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotificationAllOfRelationshipsDeliveries)

	if err != nil {
		return err
	}

	*o = NotificationAllOfRelationshipsDeliveries(varNotificationAllOfRelationshipsDeliveries)

	return err
}

type NullableNotificationAllOfRelationshipsDeliveries struct {
	value *NotificationAllOfRelationshipsDeliveries
	isSet bool
}

func (v NullableNotificationAllOfRelationshipsDeliveries) Get() *NotificationAllOfRelationshipsDeliveries {
	return v.value
}

func (v *NullableNotificationAllOfRelationshipsDeliveries) Set(val *NotificationAllOfRelationshipsDeliveries) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationAllOfRelationshipsDeliveries) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationAllOfRelationshipsDeliveries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationAllOfRelationshipsDeliveries(val *NotificationAllOfRelationshipsDeliveries) *NullableNotificationAllOfRelationshipsDeliveries {
	return &NullableNotificationAllOfRelationshipsDeliveries{value: val, isSet: true}
}

func (v NullableNotificationAllOfRelationshipsDeliveries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationAllOfRelationshipsDeliveries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

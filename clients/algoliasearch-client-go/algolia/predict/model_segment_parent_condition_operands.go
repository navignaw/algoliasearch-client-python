// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package predict

import (
	"encoding/json"
	"fmt"
)

// SegmentParentConditionOperands - struct for SegmentParentConditionOperands
type SegmentParentConditionOperands struct {
	SegmentChildConditions    *SegmentChildConditions
	SegmentOperandAffinity    *SegmentOperandAffinity
	SegmentOperandFunnelStage *SegmentOperandFunnelStage
	SegmentOperandOrderValue  *SegmentOperandOrderValue
	SegmentOperandProperty    *SegmentOperandProperty
}

// SegmentChildConditionsAsSegmentParentConditionOperands is a convenience function that returns SegmentChildConditions wrapped in SegmentParentConditionOperands
func SegmentChildConditionsAsSegmentParentConditionOperands(v *SegmentChildConditions) SegmentParentConditionOperands {
	return SegmentParentConditionOperands{
		SegmentChildConditions: v,
	}
}

// SegmentOperandAffinityAsSegmentParentConditionOperands is a convenience function that returns SegmentOperandAffinity wrapped in SegmentParentConditionOperands
func SegmentOperandAffinityAsSegmentParentConditionOperands(v *SegmentOperandAffinity) SegmentParentConditionOperands {
	return SegmentParentConditionOperands{
		SegmentOperandAffinity: v,
	}
}

// SegmentOperandFunnelStageAsSegmentParentConditionOperands is a convenience function that returns SegmentOperandFunnelStage wrapped in SegmentParentConditionOperands
func SegmentOperandFunnelStageAsSegmentParentConditionOperands(v *SegmentOperandFunnelStage) SegmentParentConditionOperands {
	return SegmentParentConditionOperands{
		SegmentOperandFunnelStage: v,
	}
}

// SegmentOperandOrderValueAsSegmentParentConditionOperands is a convenience function that returns SegmentOperandOrderValue wrapped in SegmentParentConditionOperands
func SegmentOperandOrderValueAsSegmentParentConditionOperands(v *SegmentOperandOrderValue) SegmentParentConditionOperands {
	return SegmentParentConditionOperands{
		SegmentOperandOrderValue: v,
	}
}

// SegmentOperandPropertyAsSegmentParentConditionOperands is a convenience function that returns SegmentOperandProperty wrapped in SegmentParentConditionOperands
func SegmentOperandPropertyAsSegmentParentConditionOperands(v *SegmentOperandProperty) SegmentParentConditionOperands {
	return SegmentParentConditionOperands{
		SegmentOperandProperty: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SegmentParentConditionOperands) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SegmentChildConditions
	err = newStrictDecoder(data).Decode(&dst.SegmentChildConditions)
	if err == nil {
		jsonSegmentChildConditions, _ := json.Marshal(dst.SegmentChildConditions)
		if string(jsonSegmentChildConditions) == "{}" { // empty struct
			dst.SegmentChildConditions = nil
		} else {
			match++
		}
	} else {
		dst.SegmentChildConditions = nil
	}

	// try to unmarshal data into SegmentOperandAffinity
	err = newStrictDecoder(data).Decode(&dst.SegmentOperandAffinity)
	if err == nil {
		jsonSegmentOperandAffinity, _ := json.Marshal(dst.SegmentOperandAffinity)
		if string(jsonSegmentOperandAffinity) == "{}" { // empty struct
			dst.SegmentOperandAffinity = nil
		} else {
			match++
		}
	} else {
		dst.SegmentOperandAffinity = nil
	}

	// try to unmarshal data into SegmentOperandFunnelStage
	err = newStrictDecoder(data).Decode(&dst.SegmentOperandFunnelStage)
	if err == nil {
		jsonSegmentOperandFunnelStage, _ := json.Marshal(dst.SegmentOperandFunnelStage)
		if string(jsonSegmentOperandFunnelStage) == "{}" { // empty struct
			dst.SegmentOperandFunnelStage = nil
		} else {
			match++
		}
	} else {
		dst.SegmentOperandFunnelStage = nil
	}

	// try to unmarshal data into SegmentOperandOrderValue
	err = newStrictDecoder(data).Decode(&dst.SegmentOperandOrderValue)
	if err == nil {
		jsonSegmentOperandOrderValue, _ := json.Marshal(dst.SegmentOperandOrderValue)
		if string(jsonSegmentOperandOrderValue) == "{}" { // empty struct
			dst.SegmentOperandOrderValue = nil
		} else {
			match++
		}
	} else {
		dst.SegmentOperandOrderValue = nil
	}

	// try to unmarshal data into SegmentOperandProperty
	err = newStrictDecoder(data).Decode(&dst.SegmentOperandProperty)
	if err == nil {
		jsonSegmentOperandProperty, _ := json.Marshal(dst.SegmentOperandProperty)
		if string(jsonSegmentOperandProperty) == "{}" { // empty struct
			dst.SegmentOperandProperty = nil
		} else {
			match++
		}
	} else {
		dst.SegmentOperandProperty = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SegmentChildConditions = nil
		dst.SegmentOperandAffinity = nil
		dst.SegmentOperandFunnelStage = nil
		dst.SegmentOperandOrderValue = nil
		dst.SegmentOperandProperty = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(SegmentParentConditionOperands)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(SegmentParentConditionOperands)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SegmentParentConditionOperands) MarshalJSON() ([]byte, error) {
	if src.SegmentChildConditions != nil {
		return json.Marshal(&src.SegmentChildConditions)
	}

	if src.SegmentOperandAffinity != nil {
		return json.Marshal(&src.SegmentOperandAffinity)
	}

	if src.SegmentOperandFunnelStage != nil {
		return json.Marshal(&src.SegmentOperandFunnelStage)
	}

	if src.SegmentOperandOrderValue != nil {
		return json.Marshal(&src.SegmentOperandOrderValue)
	}

	if src.SegmentOperandProperty != nil {
		return json.Marshal(&src.SegmentOperandProperty)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SegmentParentConditionOperands) GetActualInstance() any {
	if obj == nil {
		return nil
	}
	if obj.SegmentChildConditions != nil {
		return obj.SegmentChildConditions
	}

	if obj.SegmentOperandAffinity != nil {
		return obj.SegmentOperandAffinity
	}

	if obj.SegmentOperandFunnelStage != nil {
		return obj.SegmentOperandFunnelStage
	}

	if obj.SegmentOperandOrderValue != nil {
		return obj.SegmentOperandOrderValue
	}

	if obj.SegmentOperandProperty != nil {
		return obj.SegmentOperandProperty
	}

	// all schemas are nil
	return nil
}

type NullableSegmentParentConditionOperands struct {
	value *SegmentParentConditionOperands
	isSet bool
}

func (v NullableSegmentParentConditionOperands) Get() *SegmentParentConditionOperands {
	return v.value
}

func (v *NullableSegmentParentConditionOperands) Set(val *SegmentParentConditionOperands) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentParentConditionOperands) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentParentConditionOperands) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentParentConditionOperands(val *SegmentParentConditionOperands) *NullableSegmentParentConditionOperands {
	return &NullableSegmentParentConditionOperands{value: val, isSet: true}
}

func (v NullableSegmentParentConditionOperands) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentParentConditionOperands) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HubbleStatus Status of the Hubble server
// swagger:model HubbleStatus
// +k8s:deepcopy-gen=true
type HubbleStatus struct {

	// metrics
	Metrics *HubbleStatusMetrics `json:"metrics,omitempty"`

	// Human readable status/error/warning message
	Msg string `json:"msg,omitempty"`

	// observer
	Observer *HubbleStatusObserver `json:"observer,omitempty"`

	// State the component is in
	// Enum: [Ok Warning Failure Disabled]
	State string `json:"state,omitempty"`
}

// Validate validates this hubble status
func (m *HubbleStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObserver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HubbleStatus) validateMetrics(formats strfmt.Registry) error {

	if swag.IsZero(m.Metrics) { // not required
		return nil
	}

	if m.Metrics != nil {
		if err := m.Metrics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics")
			}
			return err
		}
	}

	return nil
}

func (m *HubbleStatus) validateObserver(formats strfmt.Registry) error {

	if swag.IsZero(m.Observer) { // not required
		return nil
	}

	if m.Observer != nil {
		if err := m.Observer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("observer")
			}
			return err
		}
	}

	return nil
}

var hubbleStatusTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Ok","Warning","Failure","Disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hubbleStatusTypeStatePropEnum = append(hubbleStatusTypeStatePropEnum, v)
	}
}

const (

	// HubbleStatusStateOk captures enum value "Ok"
	HubbleStatusStateOk string = "Ok"

	// HubbleStatusStateWarning captures enum value "Warning"
	HubbleStatusStateWarning string = "Warning"

	// HubbleStatusStateFailure captures enum value "Failure"
	HubbleStatusStateFailure string = "Failure"

	// HubbleStatusStateDisabled captures enum value "Disabled"
	HubbleStatusStateDisabled string = "Disabled"
)

// prop value enum
func (m *HubbleStatus) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hubbleStatusTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HubbleStatus) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HubbleStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HubbleStatus) UnmarshalBinary(b []byte) error {
	var res HubbleStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HubbleStatusMetrics Status of the Hubble metrics server
// swagger:model HubbleStatusMetrics
type HubbleStatusMetrics struct {

	// State of the Hubble metrics
	// Enum: [Ok Warning Failure Disabled]
	State string `json:"state,omitempty"`
}

// Validate validates this hubble status metrics
func (m *HubbleStatusMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var hubbleStatusMetricsTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Ok","Warning","Failure","Disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hubbleStatusMetricsTypeStatePropEnum = append(hubbleStatusMetricsTypeStatePropEnum, v)
	}
}

const (

	// HubbleStatusMetricsStateOk captures enum value "Ok"
	HubbleStatusMetricsStateOk string = "Ok"

	// HubbleStatusMetricsStateWarning captures enum value "Warning"
	HubbleStatusMetricsStateWarning string = "Warning"

	// HubbleStatusMetricsStateFailure captures enum value "Failure"
	HubbleStatusMetricsStateFailure string = "Failure"

	// HubbleStatusMetricsStateDisabled captures enum value "Disabled"
	HubbleStatusMetricsStateDisabled string = "Disabled"
)

// prop value enum
func (m *HubbleStatusMetrics) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hubbleStatusMetricsTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HubbleStatusMetrics) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("metrics"+"."+"state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HubbleStatusMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HubbleStatusMetrics) UnmarshalBinary(b []byte) error {
	var res HubbleStatusMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HubbleStatusObserver Status of the Hubble observer
// swagger:model HubbleStatusObserver
// +k8s:deepcopy-gen=true
type HubbleStatusObserver struct {

	// Current number of flows this Hubble observer stores
	CurrentFlows int64 `json:"current-flows,omitempty"`

	// Maximum number of flows this Hubble observer is able to store
	MaxFlows int64 `json:"max-flows,omitempty"`

	// Total number of flows this Hubble observer has seen
	SeenFlows int64 `json:"seen-flows,omitempty"`

	// Uptime of this Hubble observer instance
	// Format: duration
	Uptime strfmt.Duration `json:"uptime,omitempty"`
}

// Validate validates this hubble status observer
func (m *HubbleStatusObserver) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUptime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HubbleStatusObserver) validateUptime(formats strfmt.Registry) error {

	if swag.IsZero(m.Uptime) { // not required
		return nil
	}

	if err := validate.FormatOf("observer"+"."+"uptime", "body", "duration", m.Uptime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HubbleStatusObserver) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HubbleStatusObserver) UnmarshalBinary(b []byte) error {
	var res HubbleStatusObserver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

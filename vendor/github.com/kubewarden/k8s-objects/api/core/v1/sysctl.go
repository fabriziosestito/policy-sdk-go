// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// Sysctl Sysctl defines a kernel parameter to be set
//
// swagger:model Sysctl
type Sysctl struct {

	// Name of a property to set
	// Required: true
	Name *string `json:"name"`

	// Value of a property to set
	// Required: true
	Value *string `json:"value"`
}
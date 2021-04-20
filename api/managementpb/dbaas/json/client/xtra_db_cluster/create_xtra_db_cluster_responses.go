// Code generated by go-swagger; DO NOT EDIT.

package xtra_db_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateXtraDBClusterReader is a Reader for the CreateXtraDBCluster structure.
type CreateXtraDBClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateXtraDBClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateXtraDBClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateXtraDBClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateXtraDBClusterOK creates a CreateXtraDBClusterOK with default headers values
func NewCreateXtraDBClusterOK() *CreateXtraDBClusterOK {
	return &CreateXtraDBClusterOK{}
}

/*CreateXtraDBClusterOK handles this case with default header values.

A successful response.
*/
type CreateXtraDBClusterOK struct {
	Payload interface{}
}

func (o *CreateXtraDBClusterOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/XtraDBCluster/Create][%d] createXtraDbClusterOk  %+v", 200, o.Payload)
}

func (o *CreateXtraDBClusterOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateXtraDBClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateXtraDBClusterDefault creates a CreateXtraDBClusterDefault with default headers values
func NewCreateXtraDBClusterDefault(code int) *CreateXtraDBClusterDefault {
	return &CreateXtraDBClusterDefault{
		_statusCode: code,
	}
}

/*CreateXtraDBClusterDefault handles this case with default header values.

An unexpected error response.
*/
type CreateXtraDBClusterDefault struct {
	_statusCode int

	Payload *CreateXtraDBClusterDefaultBody
}

// Code gets the status code for the create xtra DB cluster default response
func (o *CreateXtraDBClusterDefault) Code() int {
	return o._statusCode
}

func (o *CreateXtraDBClusterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/XtraDBCluster/Create][%d] CreateXtraDBCluster default  %+v", o._statusCode, o.Payload)
}

func (o *CreateXtraDBClusterDefault) GetPayload() *CreateXtraDBClusterDefaultBody {
	return o.Payload
}

func (o *CreateXtraDBClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateXtraDBClusterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateXtraDBClusterBody create xtra DB cluster body
swagger:model CreateXtraDBClusterBody
*/
type CreateXtraDBClusterBody struct {

	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// XtraDB cluster name.
	// a DNS-1035 label must consist of lower case alphanumeric characters or '-',
	// start with an alphabetic character, and end with an alphanumeric character
	// (e.g. 'my-name',  or 'abc-123', regex used for validation is '[a-z]([-a-z0-9]*[a-z0-9])?')
	Name string `json:"name,omitempty"`

	// params
	Params *CreateXtraDBClusterParamsBodyParams `json:"params,omitempty"`
}

// Validate validates this create xtra DB cluster body
func (o *CreateXtraDBClusterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateXtraDBClusterBody) validateParams(formats strfmt.Registry) error {

	if swag.IsZero(o.Params) { // not required
		return nil
	}

	if o.Params != nil {
		if err := o.Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateXtraDBClusterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateXtraDBClusterBody) UnmarshalBinary(b []byte) error {
	var res CreateXtraDBClusterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateXtraDBClusterDefaultBody create xtra DB cluster default body
swagger:model CreateXtraDBClusterDefaultBody
*/
type CreateXtraDBClusterDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this create xtra DB cluster default body
func (o *CreateXtraDBClusterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateXtraDBClusterDefaultBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CreateXtraDBCluster default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateXtraDBClusterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateXtraDBClusterDefaultBody) UnmarshalBinary(b []byte) error {
	var res CreateXtraDBClusterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateXtraDBClusterParamsBodyParams XtraDBClusterParams represents XtraDB cluster parameters that can be updated.
swagger:model CreateXtraDBClusterParamsBodyParams
*/
type CreateXtraDBClusterParamsBodyParams struct {

	// Cluster size.
	ClusterSize int32 `json:"cluster_size,omitempty"`

	// haproxy
	Haproxy *CreateXtraDBClusterParamsBodyParamsHaproxy `json:"haproxy,omitempty"`

	// proxysql
	Proxysql *CreateXtraDBClusterParamsBodyParamsProxysql `json:"proxysql,omitempty"`

	// pxc
	Pxc *CreateXtraDBClusterParamsBodyParamsPxc `json:"pxc,omitempty"`
}

// Validate validates this create xtra DB cluster params body params
func (o *CreateXtraDBClusterParamsBodyParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHaproxy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProxysql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePxc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateXtraDBClusterParamsBodyParams) validateHaproxy(formats strfmt.Registry) error {

	if swag.IsZero(o.Haproxy) { // not required
		return nil
	}

	if o.Haproxy != nil {
		if err := o.Haproxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "haproxy")
			}
			return err
		}
	}

	return nil
}

func (o *CreateXtraDBClusterParamsBodyParams) validateProxysql(formats strfmt.Registry) error {

	if swag.IsZero(o.Proxysql) { // not required
		return nil
	}

	if o.Proxysql != nil {
		if err := o.Proxysql.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "proxysql")
			}
			return err
		}
	}

	return nil
}

func (o *CreateXtraDBClusterParamsBodyParams) validatePxc(formats strfmt.Registry) error {

	if swag.IsZero(o.Pxc) { // not required
		return nil
	}

	if o.Pxc != nil {
		if err := o.Pxc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "pxc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateXtraDBClusterParamsBodyParams) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateXtraDBClusterParamsBodyParams) UnmarshalBinary(b []byte) error {
	var res CreateXtraDBClusterParamsBodyParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateXtraDBClusterParamsBodyParamsHaproxy HAProxy container parameters.
// NOTE: HAProxy does not need disk size as ProxySQL does because the container does not require it.
swagger:model CreateXtraDBClusterParamsBodyParamsHaproxy
*/
type CreateXtraDBClusterParamsBodyParamsHaproxy struct {

	// Docker image used for HAProxy.
	Image string `json:"image,omitempty"`

	// compute resources
	ComputeResources *CreateXtraDBClusterParamsBodyParamsHaproxyComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this create xtra DB cluster params body params haproxy
func (o *CreateXtraDBClusterParamsBodyParamsHaproxy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateXtraDBClusterParamsBodyParamsHaproxy) validateComputeResources(formats strfmt.Registry) error {

	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "haproxy" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateXtraDBClusterParamsBodyParamsHaproxy) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateXtraDBClusterParamsBodyParamsHaproxy) UnmarshalBinary(b []byte) error {
	var res CreateXtraDBClusterParamsBodyParamsHaproxy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateXtraDBClusterParamsBodyParamsHaproxyComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model CreateXtraDBClusterParamsBodyParamsHaproxyComputeResources
*/
type CreateXtraDBClusterParamsBodyParamsHaproxyComputeResources struct {

	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this create xtra DB cluster params body params haproxy compute resources
func (o *CreateXtraDBClusterParamsBodyParamsHaproxyComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateXtraDBClusterParamsBodyParamsHaproxyComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateXtraDBClusterParamsBodyParamsHaproxyComputeResources) UnmarshalBinary(b []byte) error {
	var res CreateXtraDBClusterParamsBodyParamsHaproxyComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateXtraDBClusterParamsBodyParamsProxysql ProxySQL container parameters.
// TODO Do not use inner messages in all public APIs (for consistency).
swagger:model CreateXtraDBClusterParamsBodyParamsProxysql
*/
type CreateXtraDBClusterParamsBodyParamsProxysql struct {

	// Docker image used for ProxySQL.
	Image string `json:"image,omitempty"`

	// Disk size in bytes.
	DiskSize string `json:"disk_size,omitempty"`

	// compute resources
	ComputeResources *CreateXtraDBClusterParamsBodyParamsProxysqlComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this create xtra DB cluster params body params proxysql
func (o *CreateXtraDBClusterParamsBodyParamsProxysql) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateXtraDBClusterParamsBodyParamsProxysql) validateComputeResources(formats strfmt.Registry) error {

	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "proxysql" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateXtraDBClusterParamsBodyParamsProxysql) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateXtraDBClusterParamsBodyParamsProxysql) UnmarshalBinary(b []byte) error {
	var res CreateXtraDBClusterParamsBodyParamsProxysql
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateXtraDBClusterParamsBodyParamsProxysqlComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model CreateXtraDBClusterParamsBodyParamsProxysqlComputeResources
*/
type CreateXtraDBClusterParamsBodyParamsProxysqlComputeResources struct {

	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this create xtra DB cluster params body params proxysql compute resources
func (o *CreateXtraDBClusterParamsBodyParamsProxysqlComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateXtraDBClusterParamsBodyParamsProxysqlComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateXtraDBClusterParamsBodyParamsProxysqlComputeResources) UnmarshalBinary(b []byte) error {
	var res CreateXtraDBClusterParamsBodyParamsProxysqlComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateXtraDBClusterParamsBodyParamsPxc PXC container parameters.
// TODO Do not use inner messages in all public APIs (for consistency).
swagger:model CreateXtraDBClusterParamsBodyParamsPxc
*/
type CreateXtraDBClusterParamsBodyParamsPxc struct {

	// Docker image used for PXC.
	Image string `json:"image,omitempty"`

	// Disk size in bytes.
	DiskSize string `json:"disk_size,omitempty"`

	// compute resources
	ComputeResources *CreateXtraDBClusterParamsBodyParamsPxcComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this create xtra DB cluster params body params pxc
func (o *CreateXtraDBClusterParamsBodyParamsPxc) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateXtraDBClusterParamsBodyParamsPxc) validateComputeResources(formats strfmt.Registry) error {

	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "params" + "." + "pxc" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateXtraDBClusterParamsBodyParamsPxc) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateXtraDBClusterParamsBodyParamsPxc) UnmarshalBinary(b []byte) error {
	var res CreateXtraDBClusterParamsBodyParamsPxc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateXtraDBClusterParamsBodyParamsPxcComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model CreateXtraDBClusterParamsBodyParamsPxcComputeResources
*/
type CreateXtraDBClusterParamsBodyParamsPxcComputeResources struct {

	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this create xtra DB cluster params body params pxc compute resources
func (o *CreateXtraDBClusterParamsBodyParamsPxcComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateXtraDBClusterParamsBodyParamsPxcComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateXtraDBClusterParamsBodyParamsPxcComputeResources) UnmarshalBinary(b []byte) error {
	var res CreateXtraDBClusterParamsBodyParamsPxcComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DetailsItems0 details items0
swagger:model DetailsItems0
*/
type DetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this details items0
func (o *DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

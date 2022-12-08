/*
Copyright 2021 The Dapr Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
   http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sdk

import (
	msgpack "github.com/wapc/tinygo-msgpack"
)

// OperationType describes a CRUD operation performed against a state store.
type OperationType string

// Upsert is an update or create operation.
const Upsert OperationType = "upsert"

// Delete is a delete operation.
const Delete OperationType = "delete"

// TransactionalRequest describes a set of stateful operations for a given actor that are performed in a transactional manner.
//type TransactionalRequest struct {
//	Operations []TransactionalOperation `json:"operations"`
//	ActorType  string
//	ActorID    string
//}
//
//// TransactionalOperation is the request object for a state operation participating in a transaction.
//type TransactionalOperation struct {
//	Operation OperationType `json:"operation"`
//	Request   Request       `json:"request"`
//}
//
//// Request defines a key/value pair for an upsert/delete operation.
//type Request struct {
//	Key   string `json:"key"`
//	Value string `json:"value"`
//}

// Description of Uppercase service
type GetActorStateRequest struct {
	ActorID   string `json:"ActorID"`
	ActorType string `json:"ActorType"`
	KeyName   string `json:"KeyName"`
}

// Encode serializes a GetActorStateRequest using msgpack
func (o *GetActorStateRequest) Encode() ([]byte, error) {
	sizer := msgpack.NewSizer()
	err := o.encode(&sizer)
	if err != nil {
		return nil, err
	}
	size := sizer.Len()
	buf := make([]byte, size)
	encoder := msgpack.NewEncoder(buf)
	err = o.encode(&encoder)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (o *GetActorStateRequest) encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("ActorID")
	encoder.WriteString(o.ActorID)
	encoder.WriteString("ActorType")
	encoder.WriteString(o.ActorType)
	encoder.WriteString("KeyName")
	encoder.WriteString(o.KeyName)

	return encoder.Err()
}

// DecodeGetActorStateRequest deserializes a GetActorStateRequest using msgpack
func DecodeGetActorStateRequest(d *msgpack.Decoder) (GetActorStateRequest, error) {
	var val GetActorStateRequest
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "ActorID":
			val.ActorID, err = d.ReadString()
		case "ActorType":
			val.ActorType, err = d.ReadString()
		case "KeyName":
			val.KeyName, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

type GetActorStateResponse struct {
	Data []byte `json:"Data"`
}

// Encode serializes a GetActorStateResponse using msgpack
func (o *GetActorStateResponse) Encode() ([]byte, error) {
	sizer := msgpack.NewSizer()
	err := o.encode(&sizer)
	if err != nil {
		return nil, err
	}
	size := sizer.Len()
	buf := make([]byte, size)
	encoder := msgpack.NewEncoder(buf)
	err = o.encode(&encoder)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (o *GetActorStateResponse) encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(1)
	encoder.WriteString("Data")
	encoder.WriteByteArray(o.Data)

	return encoder.Err()
}

// DecodeGetActorStateResponse deserializes a GetActorStateResponse using msgpack
func DecodeGetActorStateResponse(d *msgpack.Decoder) (GetActorStateResponse, error) {
	var val GetActorStateResponse
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "Data":
			val.Data, err = d.ReadNillableByteArray()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

type InvokeActorRequest struct {
	ActorID     string `json:"ActorID"`
	ActorType   string `json:"ActorType"`
	ContentType string `json:"ContentType"`
	Data        []byte `json:"Data"`
	Method      string `json:"Method"`
}

// Encode serializes a InvokeActorRequest using msgpack
func (o *InvokeActorRequest) Encode() ([]byte, error) {
	sizer := msgpack.NewSizer()
	err := o.encode(&sizer)
	if err != nil {
		return nil, err
	}
	size := sizer.Len()
	buf := make([]byte, size)
	encoder := msgpack.NewEncoder(buf)
	err = o.encode(&encoder)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (o *InvokeActorRequest) encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(5)
	encoder.WriteString("ActorID")
	encoder.WriteString(o.ActorID)
	encoder.WriteString("ActorType")
	encoder.WriteString(o.ActorType)
	encoder.WriteString("ContentType")
	encoder.WriteString(o.ContentType)
	encoder.WriteString("Data")
	encoder.WriteByteArray(o.Data)
	encoder.WriteString("Method")
	encoder.WriteString(o.Method)

	return encoder.Err()
}

// DecodeInvokeActorRequest deserializes a InvokeActorRequest using msgpack
func DecodeInvokeActorRequest(d *msgpack.Decoder) (InvokeActorRequest, error) {
	var val InvokeActorRequest
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "ActorID":
			val.ActorID, err = d.ReadString()
		case "ActorType":
			val.ActorType, err = d.ReadString()
		case "ContentType":
			val.ContentType, err = d.ReadString()
		case "Data":
			val.Data, err = d.ReadNillableByteArray()
		case "Method":
			val.Method, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

type InvokeActorResponse struct {
	ContentType string `json:"ContentType"`
	Data        []byte `json:"Data"`
	StatusCode  int32  `json:"StatusCode"`
}

// Encode serializes a InvokeActorResponse using msgpack
func (o *InvokeActorResponse) Encode() ([]byte, error) {
	sizer := msgpack.NewSizer()
	err := o.encode(&sizer)
	if err != nil {
		return nil, err
	}
	size := sizer.Len()
	buf := make([]byte, size)
	encoder := msgpack.NewEncoder(buf)
	err = o.encode(&encoder)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (o *InvokeActorResponse) encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("ContentType")
	encoder.WriteString(o.ContentType)
	encoder.WriteString("Data")
	encoder.WriteByteArray(o.Data)
	encoder.WriteString("StatusCode")
	encoder.WriteInt32(o.StatusCode)

	return encoder.Err()
}

// DecodeInvokeActorResponse deserializes a InvokeActorResponse using msgpack
func DecodeInvokeActorResponse(d *msgpack.Decoder) (InvokeActorResponse, error) {
	var val InvokeActorResponse
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "ContentType":
			val.ContentType, err = d.ReadString()
		case "Data":
			val.Data, err = d.ReadNillableByteArray()
		case "StatusCode":
			val.StatusCode, err = d.ReadInt32()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

type OperationRequest struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

// Encode serializes a OperationRequest using msgpack
func (o *OperationRequest) Encode() ([]byte, error) {
	sizer := msgpack.NewSizer()
	err := o.encode(&sizer)
	if err != nil {
		return nil, err
	}
	size := sizer.Len()
	buf := make([]byte, size)
	encoder := msgpack.NewEncoder(buf)
	err = o.encode(&encoder)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
func (o *OperationRequest) encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("Key")
	encoder.WriteString(o.Key)
	encoder.WriteString("Value")
	encoder.WriteString(o.Value)

	return encoder.Err()
}

// DecodeOperationRequest deserializes a OperationRequest using msgpack
func DecodeOperationRequest(d *msgpack.Decoder) (OperationRequest, error) {
	var val OperationRequest
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "Key":
			val.Key, err = d.ReadString()
		case "Value":
			val.Value, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

type TransactionalOperation struct {
	Operation OperationType    `json:"Operation"`
	Request   OperationRequest `json:"Request"`
}

// Encode serializes a TransactionalOperation using msgpack
func (o *TransactionalOperation) Encode() ([]byte, error) {
	sizer := msgpack.NewSizer()
	err := o.encode(&sizer)
	if err != nil {
		return nil, err
	}
	size := sizer.Len()
	buf := make([]byte, size)
	encoder := msgpack.NewEncoder(buf)
	err = o.encode(&encoder)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
func (o *TransactionalOperation) encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("Operation")
	encoder.WriteString(string(o.Operation))
	encoder.WriteString("Request")
	o.Request.encode(encoder)

	return encoder.Err()
}

// DecodeTransactionalOperation deserializes a TransactionalOperation using msgpack
func DecodeTransactionalOperation(d *msgpack.Decoder) (TransactionalOperation, error) {
	var val TransactionalOperation
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "Operation":
			var v string
			v, err = d.ReadString()
			val.Operation = OperationType(v)
		case "Request":
			val.Request, err = DecodeOperationRequest(d)
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

type TransactionalOperationList []TransactionalOperation

// Encode serializes a TransactionalOperationList using msgpack
func (o *TransactionalOperationList) Encode() ([]byte, error) {
	sizer := msgpack.NewSizer()
	err := o.encode(&sizer)
	if err != nil {
		return nil, err
	}
	size := sizer.Len()
	buf := make([]byte, size)
	encoder := msgpack.NewEncoder(buf)
	err = o.encode(&encoder)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
func (o *TransactionalOperationList) encode(encoder msgpack.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.encode(encoder)
	}

	return encoder.Err()
}

// DecodeTransactionalOperationList deserializes a TransactionalOperationList using msgpack
func DecodeTransactionalOperationList(d *msgpack.Decoder) (TransactionalOperationList, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]TransactionalOperation, 0), err
	}
	size, err := d.ReadArraySize()
	if err != nil {
		return make([]TransactionalOperation, 0), err
	}
	val := make([]TransactionalOperation, 0, size)
	for i := uint32(0); i < size; i++ {
		item, err := DecodeTransactionalOperation(d)
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

type TransactionalRequest struct {
	ActorID    string                     `json:"ActorID"`
	ActorType  string                     `json:"ActorType"`
	Operations TransactionalOperationList `json:"Operations"`
}

// Encode serializes a TransactionalRequest using msgpack
func (o *TransactionalRequest) Encode() ([]byte, error) {
	sizer := msgpack.NewSizer()
	err := o.encode(&sizer)
	size := sizer.Len()
	buf := make([]byte, size)
	encoder := msgpack.NewEncoder(buf)
	err = o.encode(&encoder)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
func (o *TransactionalRequest) encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("ActorID")
	encoder.WriteString(o.ActorID)
	encoder.WriteString("ActorType")
	encoder.WriteString(o.ActorType)
	encoder.WriteString("Operations")
	o.Operations.encode(encoder)

	return encoder.Err()
}

// DecodeTransactionalRequest deserializes a TransactionalRequest using msgpack
func DecodeTransactionalRequest(d *msgpack.Decoder) (TransactionalRequest, error) {
	var val TransactionalRequest
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "ActorID":
			val.ActorID, err = d.ReadString()
		case "ActorType":
			val.ActorType, err = d.ReadString()
		case "Operations":
			val.Operations, err = DecodeTransactionalOperationList(d)
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

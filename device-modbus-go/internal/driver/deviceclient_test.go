// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"bytes"
	"testing"
	"time"

	sdkModel "github.com/edgexfoundry/device-sdk-go/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
)

func init() {
	driver = new(Driver)
	driver.Logger = logger.NewClient("test", false, "./device-Modbus.log", "DEBUG")
}

func TestTransformDataBytesToResult_INT16(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Int16,
		Attributes: map[string]string{
			PRIMARY_TABLE:    INPUT_REGISTERS,
			STARTING_ADDRESS: "10",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{255, 231} // => big-endian [231,255] => -25
	expected := int16(-25)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)

	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Int16Value()
	if err != nil || expected != result {
		t.Fatalf("Unexpected result. Error: %v", err)
	}
}

func TestTransformDataBytesToResult_INT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Int32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    INPUT_REGISTERS,
			STARTING_ADDRESS: "10",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{0, 0, 1, 11} // big-endian [11,1,0,0] => 11+2^8=267
	expected := int32(267)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)
	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Int32Value()
	if err != nil || expected != result {
		t.Fatalf("Unexpected result. Error: %v", err)
	}
}

func TestTransformDataBytesToResult_INT64(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Int64,
		Attributes: map[string]string{
			PRIMARY_TABLE:    INPUT_REGISTERS,
			STARTING_ADDRESS: "10",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{0, 1, 0, 0, 0, 2, 1, 1} // big-endian [1,1,2,0,0,0,1,0] => 1+2^8+2^17+2^48=281474976841985
	expected := int64(281474976841985)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)
	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Int64Value()
	if err != nil || expected != result {
		t.Fatalf("Unexpected result. Error: %v", err)
	}
}

func TestTransformDataBytesToResult_UINT16(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Uint16,
		Attributes: map[string]string{
			PRIMARY_TABLE:    INPUT_REGISTERS,
			STARTING_ADDRESS: "10",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{0, 11} // => big-endian [11,0] => 11
	expected := uint16(11)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)

	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Uint16Value()
	if err != nil || expected != result {
		t.Fatalf("Unexpected result. Error: %v", err)
	}
}

func TestTransformDataBytesToResult_UINT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Uint32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{0, 0, 1, 11} // big-endian [11,1,0,0] => 11+2^8=267
	expected := uint32(267)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)
	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Uint32Value()
	if err != nil || expected != result {
		t.Fatalf("Unexpected result. Error: %v", err)
	}
}

func TestTransformDataBytesToResult_UINT64(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Uint64,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{0, 1, 0, 0, 0, 2, 1, 1} // big-endian [1,1,2,0,0,0,1,0] => 1+2^8+2^17+2^48=281474976841985

	expected := uint64(281474976841985)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)
	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Uint64Value()
	if err != nil || expected != result {
		t.Fatalf("Unexpected result. Error: %v", err)
	}
}

func TestTransformDataBytesToResult_FLOAT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Float32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{63, 143, 92, 41} // big-endian [41,92,143,63] => 1.12
	expected := float32(1.12)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)
	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Float32Value()
	if err != nil || expected != result {
		t.Fatalf("Unexpected result. Error: %v", err)
	}
}

func TestTransformDataBytesToResult_FLOAT64(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Float64,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{63, 241, 235, 133, 30, 184, 81, 236} // => 1.12
	expected := float64(1.12)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)
	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Float64Value()
	if err != nil || expected != result {
		t.Fatalf("Unexpected result. Error: %v", err)
	}
}

func TestTransformDataBytesToResult_BOOL(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Bool,
		Attributes: map[string]string{
			PRIMARY_TABLE:    DISCRETES_INPUT,
			STARTING_ADDRESS: "10",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{1} // => 00000001
	expected := true

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)
	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.BoolValue()
	if err != nil || expected != result {
		t.Fatalf("Unexpected result. Error: %v", err)
	}
}

func TestTransformDataBytesToResult_RawType_INT16_ValueType_FLOAT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Float32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    INPUT_REGISTERS,
			STARTING_ADDRESS: "10",
			RAW_TYPE:         INT16,
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{255, 231} // => big-endian [231,255] => -25
	expected := float32(-25)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)

	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Float32Value()
	if err != nil {
		t.Fatalf("Unexpected result. Error: %v", err)
	} else if expected != result {
		t.Fatalf("Unexpected result. expected result %v should equal to %v", expected, result)
	}
}

func TestTransformDataBytesToResult_RawType_UINT16_ValueType_FLOAT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Float32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    INPUT_REGISTERS,
			STARTING_ADDRESS: "10",
			RAW_TYPE:         UINT16,
		},
	}

	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{0, 11} // => big-endian [11,0] => 11
	expected := float32(11)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)

	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Float32Value()
	if err != nil {
		t.Fatalf("Unexpected result. Error: %v", err)
	} else if expected != result {
		t.Fatalf("Unexpected result. expected result %v should equal to %v", expected, result)
	}
}

func TestTransformDataBytesToResult_RawType_INT16_ValueType_FLOAT64(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Float64,
		Attributes: map[string]string{
			PRIMARY_TABLE:    INPUT_REGISTERS,
			STARTING_ADDRESS: "10",
			RAW_TYPE:         INT16,
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{255, 231} // => big-endian [231,255] => -25
	expected := float64(-25)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)

	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Float64Value()
	if err != nil {
		t.Fatalf("Unexpected result. Error: %v", err)
	} else if expected != result {
		t.Fatalf("Unexpected result. expected result %v should equal to %v", expected, result)
	}
}

func TestTransformDataBytesToResult_RawType_UINT16_ValueType_FLOAT64(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Float64,
		Attributes: map[string]string{
			PRIMARY_TABLE:    INPUT_REGISTERS,
			STARTING_ADDRESS: "10",
			RAW_TYPE:         UINT16,
		},
	}

	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{0, 11} // => big-endian [11,0] => 11
	expected := float64(11)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)

	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Float64Value()
	if err != nil {
		t.Fatalf("Unexpected result. Error: %v", err)
	} else if expected != result {
		t.Fatalf("Unexpected result. expected result %v should equal to %v", expected, result)
	}
}

func TestTransformCommandValueToDataBytes_INT16(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Int16,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewInt16Value(req.DeviceResourceName, resTime, -25)
	expected := []byte{255, 231}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil || !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	}
}

func TestTransformCommandValueToDataBytes_INT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Int32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewInt32Value(req.DeviceResourceName, resTime, 267)
	expected := []byte{0, 0, 1, 11}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil || !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	}
}

func TestTransformCommandValueToDataBytes_INT64(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Int64,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewInt64Value(req.DeviceResourceName, resTime, 281474976841985)
	expected := []byte{0, 1, 0, 0, 0, 2, 1, 1}
	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil || !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	}
}

func TestTransformCommandValueToDataBytes_UINT16(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Uint16,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewUint16Value(req.DeviceResourceName, resTime, 11)
	expected := []byte{0, 11}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil || !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	}
}

func TestTransformCommandValueToDataBytes_UINT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Uint32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewUint32Value(req.DeviceResourceName, resTime, 267)
	expected := []byte{0, 0, 1, 11}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil || !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	}
}

func TestTransformCommandValueToDataBytes_UINT64(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Uint64,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewUint64Value(req.DeviceResourceName, resTime, 281474976841985)
	expected := []byte{0, 1, 0, 0, 0, 2, 1, 1}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil || !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	}
}

func TestTransformCommandValueToDataBytes_FLOAT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Float32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewFloat32Value(req.DeviceResourceName, resTime, 1.12)
	expected := []byte{63, 143, 92, 41}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil || !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	}
}

func TestTransformCommandValueToDataBytes_FLOAT64(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Float64,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewFloat64Value(req.DeviceResourceName, resTime, 1.12)
	expected := []byte{63, 241, 235, 133, 30, 184, 81, 236}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil || !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	}
}

func TestTransformCommandValueToDataBytes_BOOL(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Bool,
		Attributes: map[string]string{
			PRIMARY_TABLE:    COILS,
			STARTING_ADDRESS: "10",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewBoolValue(req.DeviceResourceName, resTime, true)
	expected := []byte{1}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil || !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	}
}

func TestTransformCommandValueToDataBytes_ValueType_FLOAT32_RawType_INT16(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Float32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
			RAW_TYPE:         INT16,
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewFloat32Value(req.DeviceResourceName, resTime, -52.0)
	expected := []byte{255, 204}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	} else if !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Unexpected result. expected result %v should equal to %v", expected, dataBytes)
	}
}

func TestTransformCommandValueToDataBytes_ValueType_FLOAT32_RawType_UINT16(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Float32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
			RAW_TYPE:         UINT16,
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewFloat32Value(req.DeviceResourceName, resTime, 112.1)
	expected := []byte{0, 112}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	} else if !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Unexpected result. expected result %v should equal to %v", expected, dataBytes)
	}
}

func TestTransformCommandValueToDataBytes_ValueType_FLOAT64_RawType_INT16(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Float64,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
			RAW_TYPE:         INT16,
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewFloat64Value(req.DeviceResourceName, resTime, -52.0)
	expected := []byte{255, 204}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	} else if !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Unexpected result. expected result %v should equal to %v", expected, dataBytes)
	}
}

func TestTransformCommandValueToDataBytes_ValueType_FLOAT64_RawType_UINT16(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Float64,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
			RAW_TYPE:         UINT16,
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewFloat64Value(req.DeviceResourceName, resTime, 112.1)
	expected := []byte{0, 112}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	} else if !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Unexpected result. expected result %v should equal to %v", expected, dataBytes)
	}
}

// Test swap operation for read command
func TestTransformDataBytesToResultWithByteSwap_INT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Int32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    INPUT_REGISTERS,
			STARTING_ADDRESS: "10",
			IS_BYTE_SWAP:     "true",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{0, 0, 11, 1} // bytes swap & big-endian => [11,1,0,0] => 11+2^8=267
	expected := int32(267)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)
	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Int32Value()
	if err != nil {
		t.Fatalf("Fail to test the TransformDataBytesToResult function. Error: %v", err)
	}
	if expected != result {
		t.Fatalf("Unexpected result, the result %d should be equal to the expected value %d", result, expected)
	}
}

func TestTransformDataBytesToResultWithWordSwap_INT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Int32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    INPUT_REGISTERS,
			STARTING_ADDRESS: "10",
			IS_WORD_SWAP:     "true",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{1, 11, 0, 0} // words swap & big-endian => [11,1,0,0] => 11+2^8=267
	expected := int32(267)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)
	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Int32Value()
	if err != nil {
		t.Fatalf("Fail to test the TransformDataBytesToResult function. Error: %v", err)
	}
	if expected != result {
		t.Fatalf("Unexpected result, the result %d should be equal to the expected value %d", result, expected)
	}
}

func TestTransformDataBytesToResultWithByteAndWordSwap_INT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Int32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    INPUT_REGISTERS,
			STARTING_ADDRESS: "10",
			IS_BYTE_SWAP:     "true",
			IS_WORD_SWAP:     "true",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{11, 1, 0, 0} // bytes and words swap & big-endian => [11,1,0,0] => 11+2^8=267
	expected := int32(267)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)
	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Int32Value()
	if err != nil {
		t.Fatalf("Fail to test the TransformDataBytesToResult function. Error: %v", err)
	}
	if expected != result {
		t.Fatalf("Unexpected result, the result %d should be equal to the expected value %d", result, expected)
	}
}

func TestTransformDataBytesToResultWithByteSwap_UINT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Uint32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
			IS_BYTE_SWAP:     "true",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{0, 0, 11, 1} // bytes swap & big-endian => [11,1,0,0] => 11+2^8=267
	expected := uint32(267)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)
	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Uint32Value()
	if err != nil {
		t.Fatalf("Fail to test the TransformDataBytesToResult function. Error: %v", err)
	}
	if expected != result {
		t.Fatalf("Unexpected result, the result %d should be equal to the expected value %d", result, expected)
	}
}

func TestTransformDataBytesToResultWithWordSwap_UINT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Uint32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
			IS_WORD_SWAP:     "true",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{1, 11, 0, 0} // words swap & big-endian => [11,1,0,0] => 11+2^8=267
	expected := uint32(267)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)
	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Uint32Value()
	if err != nil {
		t.Fatalf("Fail to test the TransformDataBytesToResult function. Error: %v", err)
	}
	if expected != result {
		t.Fatalf("Unexpected result, the result %d should be equal to the expected value %d", result, expected)
	}
}

func TestTransformDataBytesToResultWithByteAndWordSwap_UINT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Uint32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
			IS_BYTE_SWAP:     "true",
			IS_WORD_SWAP:     "true",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{11, 1, 0, 0} // bytes and words swap & big-endian => [11,1,0,0] => 11+2^8=267
	expected := uint32(267)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)
	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Uint32Value()
	if err != nil {
		t.Fatalf("Fail to test the TransformDataBytesToResult function. Error: %v", err)
	}
	if expected != result {
		t.Fatalf("Unexpected result, the result %d should be equal to the expected value %d", result, expected)
	}
}

func TestTransformDataBytesToResultWithByteSwap_FLOAT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Float32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
			IS_BYTE_SWAP:     "true",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{143, 63, 41, 92} // bytes swap & big-endian => [41,92,143,63] => 1.12
	expected := float32(1.12)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)
	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Float32Value()
	if err != nil {
		t.Fatalf("Fail to test the TransformDataBytesToResult function. Error: %v", err)
	}
	if expected != result {
		t.Fatalf("Unexpected result, the result %f should be equal to the expected value %f", result, expected)
	}
}

func TestTransformDataBytesToResultWithWordSwap_FLOAT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Float32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
			IS_WORD_SWAP:     "true",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{92, 41, 63, 143} // words swap & big-endian => [41,92,143,63] => 1.12
	expected := float32(1.12)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)
	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Float32Value()
	if err != nil {
		t.Fatalf("Fail to test the TransformDataBytesToResult function. Error: %v", err)
	}
	if expected != result {
		t.Fatalf("Unexpected result, the result %f should be equal to the expected value %f", result, expected)
	}
}

func TestTransformDataBytesToResultWithByteAndWordSwap_FLOAT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Float32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
			IS_BYTE_SWAP:     "true",
			IS_WORD_SWAP:     "true",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	dataBytes := []byte{41, 92, 143, 63} // bytes and words swap & big-endian => [41,92,143,63] => 1.12
	expected := float32(1.12)

	commandValue, err := TransformDataBytesToResult(&req, dataBytes, commandInfo)
	if err != nil {
		t.Fatalf("Fail to tramsform data bytes to result. Error: %v", err)
	}
	result, err := commandValue.Float32Value()
	if err != nil {
		t.Fatalf("Fail to test the TransformDataBytesToResult function. Error: %v", err)
	}
	if expected != result {
		t.Fatalf("Unexpected result, the result %f should be equal to the expected value %f", result, expected)
	}
}

// Test swap operation for write command
func TestTransformCommandValueToDataBytesWithByteSwap_INT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Int32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
			IS_BYTE_SWAP:     "true",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewInt32Value(req.DeviceResourceName, resTime, 267)
	expected := []byte{0, 0, 11, 1}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	}
	if !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Unexpected result, the result %v should be equal to the expected value %v", dataBytes, expected)
	}
}

func TestTransformCommandValueToDataBytesWithWordSwap_INT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Int32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
			IS_WORD_SWAP:     "true",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewInt32Value(req.DeviceResourceName, resTime, 267)
	expected := []byte{1, 11, 0, 0}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	}
	if !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Unexpected result, the result %v should be equal to the expected value %v", dataBytes, expected)
	}
}

func TestTransformCommandValueToDataBytesWithByteAndWordSwap_INT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Int32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
			IS_BYTE_SWAP:     "true",
			IS_WORD_SWAP:     "true",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewInt32Value(req.DeviceResourceName, resTime, 267)
	expected := []byte{11, 1, 0, 0}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	}
	if !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Unexpected result, the result %v should be equal to the expected value %v", dataBytes, expected)
	}
}

func TestTransformCommandValueToDataBytesWithByteSwap_UINT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Uint32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
			IS_BYTE_SWAP:     "true",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewUint32Value(req.DeviceResourceName, resTime, 267)
	expected := []byte{0, 0, 11, 1}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	}
	if !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Unexpected result, the result %v should be equal to the expected value %v", dataBytes, expected)
	}
}

func TestTransformCommandValueToDataBytesWithWordSwap_UINT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Uint32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
			IS_WORD_SWAP:     "true",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewUint32Value(req.DeviceResourceName, resTime, 267)
	expected := []byte{1, 11, 0, 0}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	}
	if !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Unexpected result, the result %v should be equal to the expected value %v", dataBytes, expected)
	}
}

func TestTransformCommandValueToDataBytesWithByteAndWordSwap_UINT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Uint32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
			IS_BYTE_SWAP:     "true",
			IS_WORD_SWAP:     "true",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewUint32Value(req.DeviceResourceName, resTime, 267)
	expected := []byte{11, 1, 0, 0}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	}
	if !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Unexpected result, the result %v should be equal to the expected value %v", dataBytes, expected)
	}
}

func TestTransformCommandValueToDataBytesWithByteSwap_FLOAT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Float32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
			IS_BYTE_SWAP:     "true",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewFloat32Value(req.DeviceResourceName, resTime, 1.12)
	expected := []byte{143, 63, 41, 92}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	}
	if !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Unexpected result, the result %v should be equal to the expected value %v", dataBytes, expected)
	}
}

func TestTransformCommandValueToDataBytesWithWordSwap_FLOAT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Float32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
			IS_WORD_SWAP:     "true",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewFloat32Value(req.DeviceResourceName, resTime, 1.12)
	expected := []byte{92, 41, 63, 143}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	}
	if !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Unexpected result, the result %v should be equal to the expected value %v", dataBytes, expected)
	}
}

func TestTransformCommandValueToDataBytesWithByteAndWordSwap_FLOAT32(t *testing.T) {
	req := sdkModel.CommandRequest{
		DeviceResourceName: "light",
		Type:               sdkModel.Float32,
		Attributes: map[string]string{
			PRIMARY_TABLE:    HOLDING_REGISTERS,
			STARTING_ADDRESS: "10",
			IS_BYTE_SWAP:     "true",
			IS_WORD_SWAP:     "true",
		},
	}
	commandInfo, err := createCommandInfo(&req)
	if err != nil {
		t.Fatalf("Fail to createcommandInfo. Error: %v", err)
	}
	resTime := time.Now().UnixNano()
	val, _ := sdkModel.NewFloat32Value(req.DeviceResourceName, resTime, 1.12)
	expected := []byte{41, 92, 143, 63}

	dataBytes, err := TransformCommandValueToDataBytes(commandInfo, val)

	if err != nil {
		t.Fatalf("Fail to tramsform command value to data bytes. Error: %v", err)
	}
	if !bytes.Equal(dataBytes, expected) {
		t.Fatalf("Unexpected result, the result %v should be equal to the expected value %v", dataBytes, expected)
	}
}

package posapi

/*
#cgo LDFLAGS: -ldl
#include <stdlib.h>
#include <stdio.h>
#include <dlfcn.h>
#include "posapi.h"

typedef char* (*get_information_t)();
typedef char* (*check_api_t)();
typedef char* (*call_function_t)(const char*, const char*);
typedef char* (*put_t)(const char*);
typedef char* (*return_bill_t)(const char*);
typedef char* (*send_data_t)();

char* get_information(void* handle) {
    get_information_t get_information = (get_information_t)dlsym(handle, "getInformation");
    if (get_information == NULL) {
        return "";
    }
    return get_information();
}

char* check_api(void* handle) {
    check_api_t check_api = (check_api_t)dlsym(handle, "checkApi");
    if (check_api == NULL) {
        return "";
    }
    return check_api();
}

char* call_function(void* handle, const char* functionName, const char* params) {
    call_function_t call_function = (call_function_t)dlsym(handle, "callFunction");
    if (call_function == NULL) {
        return "";
    }
    return call_function(functionName, params);
}

char* putF(void* handle, char* data) {
    put_t putF = (put_t)dlsym(handle, "put");
    if (putF == NULL) {
        return "";
    }
    char* result = putF(data);
    return result;
}

char* return_bill(void* handle, const char* data) {
    return_bill_t return_bill = (return_bill_t)dlsym(handle, "returnBill");
    if (return_bill == NULL) {
        return "";
    }
    return return_bill(data);
}

char* send_data(void* handle) {
    send_data_t send_data = (send_data_t)dlsym(handle, "sendData");
    if (send_data == NULL) {
        return "";
    }
    return send_data();
}
*/
import "C"

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type PosAPI struct {
	handle unsafe.Pointer
}

func NewPosAPI(libPath string) (*PosAPI, error) {
	C.dlopen(C.CString(libPath), C.RTLD_NOW)
	handle := C.dlopen(C.CString(libPath), C.RTLD_NOW)
	if handle == nil {
		return nil, fmt.Errorf("Failed to load shared library: %s", libPath)
	}

	return &PosAPI{handle}, nil
}

func (api *PosAPI) Close() {
	C.dlclose(api.handle)
}

func (api *PosAPI) GetInformation() (InformationOutput, error) {

	// Call the getInformation function.
	var resultBytes []byte
	resultPtr := C.get_information(api.handle)
	if resultPtr != nil {
		resultBytes = C.GoBytes(unsafe.Pointer(resultPtr), C.int(len(C.GoString(resultPtr))))
	}
	//
	// Parse the JSON result.
	var response InformationOutput
	err := json.Unmarshal(resultBytes, &response)
	if err != nil {
		return InformationOutput{}, err
	}

	return response, nil
}
func (api *PosAPI) CheckApi() (CheckOutput, error) {
	var resultBytes []byte
	resultPtr := C.check_api(api.handle)
	if resultPtr != nil {
		resultBytes = C.GoBytes(unsafe.Pointer(resultPtr), C.int(len(C.GoString(resultPtr))))
	}
	// Parse the JSON result.
	var response CheckOutput
	err := json.Unmarshal(resultBytes, &response)
	if err != nil {
		return CheckOutput{}, err
	}
	return response, nil
}

func (api *PosAPI) CallFunction(functionName string, params string) (string, error) {
	cFn := C.CString(functionName)
	defer C.free(unsafe.Pointer(cFn))
	cParams := C.CString(params)
	defer C.free(unsafe.Pointer(cParams))
	var resultBytes []byte
	resultPtr := C.call_function(api.handle, cFn, cParams)
	if resultPtr != nil {
		resultBytes = C.GoBytes(unsafe.Pointer(resultPtr), C.int(len(C.GoString(resultPtr))))
	}
	return string(resultBytes), nil
}

func (api *PosAPI) Put(input PutInput) (PutOutput, error) {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return PutOutput{}, err
	}
	cData := C.CString(string(jsonBytes))
	defer C.free(unsafe.Pointer(cData))

	var resultBytes []byte
	resultPtr := C.putF(api.handle, cData)
	if resultPtr != nil {
		resultBytes = C.GoBytes(unsafe.Pointer(resultPtr), C.int(len(C.GoString(resultPtr))))
	}
	// Parse the JSON result.
	var response PutOutput
	err = json.Unmarshal(resultBytes, &response)
	if err != nil {
		return PutOutput{}, err
	}
	return response, nil

}
func (api *PosAPI) PutBatch(input PutInputBatch) (PutOutput, error) {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return PutOutput{}, err
	}
	cData := C.CString(string(jsonBytes))
	defer C.free(unsafe.Pointer(cData))

	var resultBytes []byte
	resultPtr := C.putF(api.handle, cData)
	if resultPtr != nil {
		resultBytes = C.GoBytes(unsafe.Pointer(resultPtr), C.int(len(C.GoString(resultPtr))))
	}
	// Parse the JSON result.
	var response PutOutput
	err = json.Unmarshal(resultBytes, &response)
	if err != nil {
		return PutOutput{}, err
	}
	return response, nil

}

func (api *PosAPI) ReturnBill(input BillInput) (BillOutput, error) {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return BillOutput{}, err
	}
	cData := C.CString(string(jsonBytes))
	defer C.free(unsafe.Pointer(cData))
	var resultBytes []byte
	resultPtr := C.return_bill(api.handle, cData)
	if resultPtr != nil {
		resultBytes = C.GoBytes(unsafe.Pointer(resultPtr), C.int(len(C.GoString(resultPtr))))
	}

	// Parse the JSON result.
	var response BillOutput
	err = json.Unmarshal(resultBytes, &response)
	if err != nil {
		return BillOutput{}, err
	}
	return response, nil
}

func (api *PosAPI) SendData() (DataOutput, error) {
	var resultBytes []byte
	resultPtr := C.send_data(api.handle)
	if resultPtr != nil {
		resultBytes = C.GoBytes(unsafe.Pointer(resultPtr), C.int(len(C.GoString(resultPtr))))
	}
	// Parse the JSON result.
	var response DataOutput
	err := json.Unmarshal(resultBytes, &response)
	if err != nil {
		return DataOutput{}, err
	}
	return response, nil
}

// Code generated by "enumer -type=AllocationStatus -trimprefix=AllocationStatus"; DO NOT EDIT.

package core

import (
	"fmt"
)

const _AllocationStatusName = "AllocationUndefinedGrantedExhaustedNamespaceQuotaExceeded"

var _AllocationStatusIndex = [...]uint8{0, 19, 26, 35, 57}

func (i AllocationStatus) String() string {
	if i < 0 || i >= AllocationStatus(len(_AllocationStatusIndex)-1) {
		return fmt.Sprintf("AllocationStatus(%d)", i)
	}
	return _AllocationStatusName[_AllocationStatusIndex[i]:_AllocationStatusIndex[i+1]]
}

var _AllocationStatusValues = []AllocationStatus{0, 1, 2, 3}

var _AllocationStatusNameToValueMap = map[string]AllocationStatus{
	_AllocationStatusName[0:19]:  0,
	_AllocationStatusName[19:26]: 1,
	_AllocationStatusName[26:35]: 2,
	_AllocationStatusName[35:57]: 3,
}

// AllocationStatusString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func AllocationStatusString(s string) (AllocationStatus, error) {
	if val, ok := _AllocationStatusNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to AllocationStatus values", s)
}

// AllocationStatusValues returns all values of the enum
func AllocationStatusValues() []AllocationStatus {
	return _AllocationStatusValues
}

// IsAAllocationStatus returns "true" if the value is listed in the enum definition. "false" otherwise
func (i AllocationStatus) IsAAllocationStatus() bool {
	for _, v := range _AllocationStatusValues {
		if i == v {
			return true
		}
	}
	return false
}

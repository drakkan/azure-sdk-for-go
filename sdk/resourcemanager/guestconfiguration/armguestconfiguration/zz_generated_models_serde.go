//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armguestconfiguration

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AssignmentProperties.
func (a AssignmentProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "assignmentHash", a.AssignmentHash)
	populate(objectMap, "complianceStatus", a.ComplianceStatus)
	populate(objectMap, "context", a.Context)
	populate(objectMap, "guestConfiguration", a.GuestConfiguration)
	populateTimeRFC3339(objectMap, "lastComplianceStatusChecked", a.LastComplianceStatusChecked)
	populate(objectMap, "latestAssignmentReport", a.LatestAssignmentReport)
	populate(objectMap, "latestReportId", a.LatestReportID)
	populate(objectMap, "parameterHash", a.ParameterHash)
	populate(objectMap, "provisioningState", a.ProvisioningState)
	populate(objectMap, "resourceType", a.ResourceType)
	populate(objectMap, "targetResourceId", a.TargetResourceID)
	populate(objectMap, "vmssVMList", a.VmssVMList)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AssignmentProperties.
func (a *AssignmentProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "assignmentHash":
			err = unpopulate(val, "AssignmentHash", &a.AssignmentHash)
			delete(rawMsg, key)
		case "complianceStatus":
			err = unpopulate(val, "ComplianceStatus", &a.ComplianceStatus)
			delete(rawMsg, key)
		case "context":
			err = unpopulate(val, "Context", &a.Context)
			delete(rawMsg, key)
		case "guestConfiguration":
			err = unpopulate(val, "GuestConfiguration", &a.GuestConfiguration)
			delete(rawMsg, key)
		case "lastComplianceStatusChecked":
			err = unpopulateTimeRFC3339(val, "LastComplianceStatusChecked", &a.LastComplianceStatusChecked)
			delete(rawMsg, key)
		case "latestAssignmentReport":
			err = unpopulate(val, "LatestAssignmentReport", &a.LatestAssignmentReport)
			delete(rawMsg, key)
		case "latestReportId":
			err = unpopulate(val, "LatestReportID", &a.LatestReportID)
			delete(rawMsg, key)
		case "parameterHash":
			err = unpopulate(val, "ParameterHash", &a.ParameterHash)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &a.ProvisioningState)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, "ResourceType", &a.ResourceType)
			delete(rawMsg, key)
		case "targetResourceId":
			err = unpopulate(val, "TargetResourceID", &a.TargetResourceID)
			delete(rawMsg, key)
		case "vmssVMList":
			err = unpopulate(val, "VmssVMList", &a.VmssVMList)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AssignmentReportDetails.
func (a *AssignmentReportDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "complianceStatus":
			err = unpopulate(val, "ComplianceStatus", &a.ComplianceStatus)
			delete(rawMsg, key)
		case "endTime":
			err = unpopulateTimeRFC3339(val, "EndTime", &a.EndTime)
			delete(rawMsg, key)
		case "jobId":
			err = unpopulate(val, "JobID", &a.JobID)
			delete(rawMsg, key)
		case "operationType":
			err = unpopulate(val, "OperationType", &a.OperationType)
			delete(rawMsg, key)
		case "resources":
			err = unpopulate(val, "Resources", &a.Resources)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, "StartTime", &a.StartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AssignmentReportProperties.
func (a *AssignmentReportProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "assignment":
			err = unpopulate(val, "Assignment", &a.Assignment)
			delete(rawMsg, key)
		case "complianceStatus":
			err = unpopulate(val, "ComplianceStatus", &a.ComplianceStatus)
			delete(rawMsg, key)
		case "details":
			err = unpopulate(val, "Details", &a.Details)
			delete(rawMsg, key)
		case "endTime":
			err = unpopulateTimeRFC3339(val, "EndTime", &a.EndTime)
			delete(rawMsg, key)
		case "reportId":
			err = unpopulate(val, "ReportID", &a.ReportID)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, "StartTime", &a.StartTime)
			delete(rawMsg, key)
		case "vm":
			err = unpopulate(val, "VM", &a.VM)
			delete(rawMsg, key)
		case "vmssResourceId":
			err = unpopulate(val, "VmssResourceID", &a.VmssResourceID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AssignmentReportResource.
func (a AssignmentReportResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "complianceStatus", a.ComplianceStatus)
	populate(objectMap, "properties", &a.Properties)
	populate(objectMap, "reasons", a.Reasons)
	populate(objectMap, "resourceId", a.ResourceID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CommonAssignmentReport.
func (c CommonAssignmentReport) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "assignment", c.Assignment)
	populate(objectMap, "complianceStatus", c.ComplianceStatus)
	populateTimeRFC3339(objectMap, "endTime", c.EndTime)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "operationType", c.OperationType)
	populate(objectMap, "reportId", c.ReportID)
	populate(objectMap, "resources", c.Resources)
	populateTimeRFC3339(objectMap, "startTime", c.StartTime)
	populate(objectMap, "vm", c.VM)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CommonAssignmentReport.
func (c *CommonAssignmentReport) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "assignment":
			err = unpopulate(val, "Assignment", &c.Assignment)
			delete(rawMsg, key)
		case "complianceStatus":
			err = unpopulate(val, "ComplianceStatus", &c.ComplianceStatus)
			delete(rawMsg, key)
		case "endTime":
			err = unpopulateTimeRFC3339(val, "EndTime", &c.EndTime)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &c.ID)
			delete(rawMsg, key)
		case "operationType":
			err = unpopulate(val, "OperationType", &c.OperationType)
			delete(rawMsg, key)
		case "reportId":
			err = unpopulate(val, "ReportID", &c.ReportID)
			delete(rawMsg, key)
		case "resources":
			err = unpopulate(val, "Resources", &c.Resources)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, "StartTime", &c.StartTime)
			delete(rawMsg, key)
		case "vm":
			err = unpopulate(val, "VM", &c.VM)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Navigation.
func (n Navigation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "assignmentSource", n.AssignmentSource)
	populate(objectMap, "assignmentType", n.AssignmentType)
	populate(objectMap, "configurationParameter", n.ConfigurationParameter)
	populate(objectMap, "configurationProtectedParameter", n.ConfigurationProtectedParameter)
	populate(objectMap, "configurationSetting", n.ConfigurationSetting)
	populate(objectMap, "contentHash", n.ContentHash)
	populate(objectMap, "contentType", n.ContentType)
	populate(objectMap, "contentUri", n.ContentURI)
	populate(objectMap, "kind", n.Kind)
	populate(objectMap, "name", n.Name)
	populate(objectMap, "version", n.Version)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, "LastModifiedByType", &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VMSSVMInfo.
func (v VMSSVMInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "complianceStatus", v.ComplianceStatus)
	populateTimeRFC3339(objectMap, "lastComplianceChecked", v.LastComplianceChecked)
	populate(objectMap, "latestReportId", v.LatestReportID)
	populate(objectMap, "vmId", v.VMID)
	populate(objectMap, "vmResourceId", v.VMResourceID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VMSSVMInfo.
func (v *VMSSVMInfo) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "complianceStatus":
			err = unpopulate(val, "ComplianceStatus", &v.ComplianceStatus)
			delete(rawMsg, key)
		case "lastComplianceChecked":
			err = unpopulateTimeRFC3339(val, "LastComplianceChecked", &v.LastComplianceChecked)
			delete(rawMsg, key)
		case "latestReportId":
			err = unpopulate(val, "LatestReportID", &v.LatestReportID)
			delete(rawMsg, key)
		case "vmId":
			err = unpopulate(val, "VMID", &v.VMID)
			delete(rawMsg, key)
		case "vmResourceId":
			err = unpopulate(val, "VMResourceID", &v.VMResourceID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}

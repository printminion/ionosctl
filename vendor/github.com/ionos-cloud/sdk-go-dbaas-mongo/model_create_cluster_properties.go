/*
 * IONOS DBaaS MongoDB REST API
 *
 * With IONOS Cloud Database as a Service, you have the ability to quickly set up and manage a MongoDB database. You can also delete clusters, manage backups and users via the API.  MongoDB is an open source, cross-platform, document-oriented database program. Classified as a NoSQL database program, it uses JSON-like documents with optional schemas.  The MongoDB API allows you to create additional database clusters or modify existing ones. Both tools, the Data Center Designer (DCD) and the API use the same concepts consistently and are well suited for smooth and intuitive use.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// CreateClusterProperties The properties with all data needed to create a new MongoDB cluster.
type CreateClusterProperties struct {
	// The unique ID of the template, which specifies the number of cores, storage size, and memory. You cannot downgrade to a smaller template or minor edition (e.g. from business to playground). To get a list of all templates to confirm the changes use the /templates endpoint.
	TemplateID *string `json:"templateID"`
	// The MongoDB version of your cluster.
	MongoDBVersion *string `json:"mongoDBVersion,omitempty"`
	// The total number of instances in the cluster (one primary and n-1 secondaries).
	Instances   *int32        `json:"instances"`
	Connections *[]Connection `json:"connections"`
	// The physical location where the cluster will be created. This is the location where all your instances will be located. This property is immutable.
	Location *string `json:"location"`
	// The name of your cluster.
	DisplayName       *string            `json:"displayName"`
	MaintenanceWindow *MaintenanceWindow `json:"maintenanceWindow,omitempty"`
}

// NewCreateClusterProperties instantiates a new CreateClusterProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClusterProperties(templateID string, instances int32, connections []Connection, location string, displayName string) *CreateClusterProperties {
	this := CreateClusterProperties{}

	this.TemplateID = &templateID
	this.Instances = &instances
	this.Connections = &connections
	this.Location = &location
	this.DisplayName = &displayName

	return &this
}

// NewCreateClusterPropertiesWithDefaults instantiates a new CreateClusterProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClusterPropertiesWithDefaults() *CreateClusterProperties {
	this := CreateClusterProperties{}
	return &this
}

// GetTemplateID returns the TemplateID field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateClusterProperties) GetTemplateID() *string {
	if o == nil {
		return nil
	}

	return o.TemplateID

}

// GetTemplateIDOk returns a tuple with the TemplateID field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClusterProperties) GetTemplateIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.TemplateID, true
}

// SetTemplateID sets field value
func (o *CreateClusterProperties) SetTemplateID(v string) {

	o.TemplateID = &v

}

// HasTemplateID returns a boolean if a field has been set.
func (o *CreateClusterProperties) HasTemplateID() bool {
	if o != nil && o.TemplateID != nil {
		return true
	}

	return false
}

// GetMongoDBVersion returns the MongoDBVersion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateClusterProperties) GetMongoDBVersion() *string {
	if o == nil {
		return nil
	}

	return o.MongoDBVersion

}

// GetMongoDBVersionOk returns a tuple with the MongoDBVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClusterProperties) GetMongoDBVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.MongoDBVersion, true
}

// SetMongoDBVersion sets field value
func (o *CreateClusterProperties) SetMongoDBVersion(v string) {

	o.MongoDBVersion = &v

}

// HasMongoDBVersion returns a boolean if a field has been set.
func (o *CreateClusterProperties) HasMongoDBVersion() bool {
	if o != nil && o.MongoDBVersion != nil {
		return true
	}

	return false
}

// GetInstances returns the Instances field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *CreateClusterProperties) GetInstances() *int32 {
	if o == nil {
		return nil
	}

	return o.Instances

}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClusterProperties) GetInstancesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Instances, true
}

// SetInstances sets field value
func (o *CreateClusterProperties) SetInstances(v int32) {

	o.Instances = &v

}

// HasInstances returns a boolean if a field has been set.
func (o *CreateClusterProperties) HasInstances() bool {
	if o != nil && o.Instances != nil {
		return true
	}

	return false
}

// GetConnections returns the Connections field value
// If the value is explicit nil, the zero value for []Connection will be returned
func (o *CreateClusterProperties) GetConnections() *[]Connection {
	if o == nil {
		return nil
	}

	return o.Connections

}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClusterProperties) GetConnectionsOk() (*[]Connection, bool) {
	if o == nil {
		return nil, false
	}

	return o.Connections, true
}

// SetConnections sets field value
func (o *CreateClusterProperties) SetConnections(v []Connection) {

	o.Connections = &v

}

// HasConnections returns a boolean if a field has been set.
func (o *CreateClusterProperties) HasConnections() bool {
	if o != nil && o.Connections != nil {
		return true
	}

	return false
}

// GetLocation returns the Location field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateClusterProperties) GetLocation() *string {
	if o == nil {
		return nil
	}

	return o.Location

}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClusterProperties) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Location, true
}

// SetLocation sets field value
func (o *CreateClusterProperties) SetLocation(v string) {

	o.Location = &v

}

// HasLocation returns a boolean if a field has been set.
func (o *CreateClusterProperties) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// GetDisplayName returns the DisplayName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateClusterProperties) GetDisplayName() *string {
	if o == nil {
		return nil
	}

	return o.DisplayName

}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClusterProperties) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.DisplayName, true
}

// SetDisplayName sets field value
func (o *CreateClusterProperties) SetDisplayName(v string) {

	o.DisplayName = &v

}

// HasDisplayName returns a boolean if a field has been set.
func (o *CreateClusterProperties) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// GetMaintenanceWindow returns the MaintenanceWindow field value
// If the value is explicit nil, the zero value for MaintenanceWindow will be returned
func (o *CreateClusterProperties) GetMaintenanceWindow() *MaintenanceWindow {
	if o == nil {
		return nil
	}

	return o.MaintenanceWindow

}

// GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateClusterProperties) GetMaintenanceWindowOk() (*MaintenanceWindow, bool) {
	if o == nil {
		return nil, false
	}

	return o.MaintenanceWindow, true
}

// SetMaintenanceWindow sets field value
func (o *CreateClusterProperties) SetMaintenanceWindow(v MaintenanceWindow) {

	o.MaintenanceWindow = &v

}

// HasMaintenanceWindow returns a boolean if a field has been set.
func (o *CreateClusterProperties) HasMaintenanceWindow() bool {
	if o != nil && o.MaintenanceWindow != nil {
		return true
	}

	return false
}

func (o CreateClusterProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TemplateID != nil {
		toSerialize["templateID"] = o.TemplateID
	}

	if o.MongoDBVersion != nil {
		toSerialize["mongoDBVersion"] = o.MongoDBVersion
	}

	if o.Instances != nil {
		toSerialize["instances"] = o.Instances
	}

	if o.Connections != nil {
		toSerialize["connections"] = o.Connections
	}

	if o.Location != nil {
		toSerialize["location"] = o.Location
	}

	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}

	if o.MaintenanceWindow != nil {
		toSerialize["maintenanceWindow"] = o.MaintenanceWindow
	}

	return json.Marshal(toSerialize)
}

type NullableCreateClusterProperties struct {
	value *CreateClusterProperties
	isSet bool
}

func (v NullableCreateClusterProperties) Get() *CreateClusterProperties {
	return v.value
}

func (v *NullableCreateClusterProperties) Set(val *CreateClusterProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClusterProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClusterProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClusterProperties(val *CreateClusterProperties) *NullableCreateClusterProperties {
	return &NullableCreateClusterProperties{value: val, isSet: true}
}

func (v NullableCreateClusterProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClusterProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
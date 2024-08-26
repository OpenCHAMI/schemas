package cloudinit

import "encoding/json"

// V1 represents the structure of the "v1" JSON object
type InstanceV1 struct {
	// The region in which the datacenter is located
	Region string `json:"region"`

	// The datacenter the instance is in
	DataCenter string `json:"data_center"`

	// The name of the system the instance belongs to which may span datacenters and regions.
	// Most often, a system is in a single datacenter and region.
	SystemName string `json:"system_name"`

	// The  failure domain of the instance.  Each system is composed of one or more failure domains
	// Common failure domains are "cooling group, "rack", "row", "datacenter", "region", "cloud", etc.
	FailureDomain string `json:"failure_domain"`

	// The cloud provider the instance is in (e.g. AWS, GCP, Azure, LANL, CSCS, etc.)
	// Many existing cloudinit scripts use this field to determine the cloud provider
	CloudName string `json:"cloud_name"`

	// The instance ID of the instance which is unique, but not stable.
	// Rebooting the instance should preserve the ID.  Changing the OS or other parameters may change the ID.
	InstanceID string `json:"instance_id"`

	// Many existing cloudinit scripts use the distro information to determine what to do next in init scripts
	// HPC Users are more likely to parse the RootImageId to determine what to do next
	RootImageId            string `json:"root_image_id"`
	RootImageDistro        string `json:"distro"`
	RootImageDistroRelease string `json:"distro_release"`
	RootImageDistroVersion string `json:"distro_version"`

	// Machine is a text description like SKU that describes what kind of machine is targeted.
	// Some common uses for this field include "x86_64 dual socket HBM2", "POWER9", "ARM64", etc.
	Machine string `json:"machine"`

	// Location is a text field for identifying physical location of the instance within the datacenter.
	// For CSM systems, this may be an xname.  For other systems, this may be a combinantion of row, rack, and slot numbers.
	Location string `json:"location"`

	// In OpenCHAMI, each group a node is a part of may provide additional metadata that can be used to configure the node.
	// These are serialized into an array of JSON objects.  The key is the group name and the value is the metadata.
	GroupsMetaData GroupsMetaData `json:"groups_metadata"`
}

// GroupsMetaData represents the structure of each group’s metadata
type GroupsMetaData map[string]json.RawMessage

// Config represents the top-level structure of the JSON object
type Config struct {
	V1 InstanceV1 `json:"v1"`
}

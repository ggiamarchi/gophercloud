package bootfromvolume

import (
	"errors"
	"strconv"

	"github.com/ggiamarchi/gophercloud"
	"github.com/ggiamarchi/gophercloud/openstack/compute/v2/servers"

	"github.com/racker/perigee"
)

// SourceType represents the type of medium being used to create the volume.
type SourceType string

const (
	Volume   SourceType = "volume"
	Snapshot SourceType = "snapshot"
	Image    SourceType = "image"
)

// BlockDevice is a structure with options for booting a server instance
// from a volume. The volume may be created from an image, snapshot, or another
// volume.
type BlockDevice struct {
	// BootIndex [optional] is the boot index. It defaults to 0.
	BootIndex int `json:"boot_index"`

	// DeleteOnTermination [optional] specifies whether or not to delete the attached volume
	// when the server is deleted. Defaults to `false`.
	DeleteOnTermination bool `json:"delete_on_termination"`

	// DestinationType [optional] is the type that gets created. Possible values are "volume"
	// and "local".
	DestinationType string `json:"destination_type"`

	// SourceType [required] must be one of: "volume", "snapshot", "image".
	SourceType SourceType `json:"source_type"`

	// UUID [required] is the unique identifier for the volume, snapshot, or image (see above)
	UUID string `json:"uuid"`

	// VolumeSize [optional] is the size of the volume to create (in gigabytes).
	VolumeSize int `json:"volume_size"`
}

// CreateOptsExt is a structure that extends the server `CreateOpts` structure
// by allowing for a block device mapping.
type CreateOptsExt struct {
	servers.CreateOptsBuilder
	BlockDevice []BlockDevice `json:"block_device_mapping_v2,omitempty"`
}

// ToServerCreateMap adds the block device mapping option to the base server
// creation options.
func (opts CreateOptsExt) ToServerCreateMap() (map[string]interface{}, error) {
	base, err := opts.CreateOptsBuilder.ToServerCreateMap()
	if err != nil {
		return nil, err
	}

	if len(opts.BlockDevice) == 0 {
		return nil, errors.New("Required fields UUID and SourceType not set.")
	}

	serverMap := base["server"].(map[string]interface{})

	blockDevice := make([]map[string]interface{}, len(opts.BlockDevice))

	for i, bd := range opts.BlockDevice {
		if string(bd.SourceType) == "" {
			return nil, errors.New("SourceType must be one of: volume, image, snapshot.")
		}

		blockDevice[i] = make(map[string]interface{})

		blockDevice[i]["source_type"] = bd.SourceType
		blockDevice[i]["boot_index"] = strconv.Itoa(bd.BootIndex)
		blockDevice[i]["delete_on_termination"] = strconv.FormatBool(bd.DeleteOnTermination)
		blockDevice[i]["volume_size"] = strconv.Itoa(bd.VolumeSize)
		if bd.UUID != "" {
			blockDevice[i]["uuid"] = bd.UUID
		}
		if bd.DestinationType != "" {
			blockDevice[i]["destination_type"] = bd.DestinationType
		}

	}
	serverMap["block_device_mapping_v2"] = blockDevice

	return base, nil
}

// Create requests the creation of a server from the given block device mapping.
func Create(client *gophercloud.ServiceClient, opts servers.CreateOptsBuilder) servers.CreateResult {
	var res servers.CreateResult

	reqBody, err := opts.ToServerCreateMap()
	if err != nil {
		res.Err = err
		return res
	}

	_, res.Err = perigee.Request("POST", createURL(client), perigee.Options{
		MoreHeaders: client.AuthenticatedHeaders(),
		ReqBody:     reqBody,
		Results:     &res.Body,
		OkCodes:     []int{200, 202},
	})
	return res
}

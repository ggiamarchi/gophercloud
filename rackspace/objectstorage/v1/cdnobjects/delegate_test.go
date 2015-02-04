package cdnobjects

import (
	"testing"

	os "github.com/ggiamarchi/gophercloud/openstack/objectstorage/v1/objects"
	th "github.com/ggiamarchi/gophercloud/testhelper"
	fake "github.com/ggiamarchi/gophercloud/testhelper/client"
)

func TestDeleteCDNObject(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	os.HandleDeleteObjectSuccessfully(t)

	res := Delete(fake.ServiceClient(), "testContainer", "testObject", nil)
	th.AssertNoErr(t, res.Err)

}

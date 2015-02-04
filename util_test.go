package gophercloud

import (
	"testing"

	th "github.com/ggiamarchi/gophercloud/testhelper"
)

func TestWaitFor(t *testing.T) {
	err := WaitFor(5, func() (bool, error) {
		return true, nil
	})
	th.CheckNoErr(t, err)
}

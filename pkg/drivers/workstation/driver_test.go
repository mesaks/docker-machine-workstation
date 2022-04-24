package workstation

import (
	"testing"

	"github.com/docker/machine/libmachine/drivers"
	"github.com/stretchr/testify/assert"
)

func TestSetConfigFromFlags(t *testing.T) {
	driver := NewDriver("default", "path")

	checkFlags := &drivers.CheckDriverOptions{
		FlagsValues: map[string]interface{}{},
		CreateFlags: driver.GetCreateFlags(),
	}

	err := driver.SetConfigFromFlags(checkFlags)

	assert.NoError(t, err)
	assert.Empty(t, checkFlags.InvalidFlags)
}

func TestParseShareFolderWindows(t *testing.T) {
	ShareFolder := `C:\Projects\`

	ShareName, GuestFolder, GuestCompatLink := parseShareFolderWindows(ShareFolder)

	assert.Equal(t, "Projects", ShareName)
	assert.Equal(t, "/Projects", GuestFolder)
	assert.Equal(t, "/c/", GuestCompatLink)

	ShareFolder = `D:\Work\Projects\`
	ShareName, GuestFolder, GuestCompatLink = parseShareFolderWindows(ShareFolder)

	assert.Equal(t, "Projects", ShareName)
	assert.Equal(t, "/Projects", GuestFolder)
	assert.Equal(t, "/d/Work", GuestCompatLink)
}

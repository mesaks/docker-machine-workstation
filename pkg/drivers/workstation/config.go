package workstation

import (
	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/mcnflag"
)

const (
	B2DUser        = "docker"
	B2DPass        = "tcuser"
	isoFilename    = "boot2docker.iso"
	isoConfigDrive = "configdrive.iso"

	defaultSSHUser  = B2DUser
	defaultSSHPass  = B2DPass
	defaultDiskSize = 20000
	defaultCpus     = 1
	defaultMemory   = 1024
)

// Driver for VMware Workstation
type Driver struct {
	*drivers.BaseDriver
	Memory         int
	DiskSize       int
	CPU            int
	ISO            string
	Boot2DockerURL string
	CPUS           int

	SSHPassword    string
	ConfigDriveISO string
	ConfigDriveURL string

	NoShare         bool
	ShareName       string
	ShareFolder     string
	GuestFolder     string
	GuestCompatLink string
}

// GetCreateFlags registers the flags this driver adds to
// "docker hosts create"
func (d *Driver) GetCreateFlags() []mcnflag.Flag {
	return []mcnflag.Flag{
		mcnflag.StringFlag{
			EnvVar: "WORKSTATION_BOOT2DOCKER_URL",
			Name:   "workstation-boot2docker-url",
			Usage:  "VMWare Workstation URL for boot2docker image",
			Value:  "",
		},
		mcnflag.StringFlag{
			EnvVar: "WORKSTATION_CONFIGDRIVE_URL",
			Name:   "workstation-configdrive-url",
			Usage:  "VMWare Workstation URL for cloud-init configdrive",
			Value:  "",
		},
		mcnflag.IntFlag{
			EnvVar: "WORKSTATION_CPU_COUNT",
			Name:   "workstation-cpu-count",
			Usage:  "number of CPUs for the machine (-1 to use the number of CPUs available)",
			Value:  defaultCpus,
		},
		mcnflag.IntFlag{
			EnvVar: "WORKSTATION_MEMORY_SIZE",
			Name:   "workstation-memory-size",
			Usage:  "VMWare Workstation size of memory for host VM (in MB)",
			Value:  defaultMemory,
		},
		mcnflag.IntFlag{
			EnvVar: "WORKSTATION_DISK_SIZE",
			Name:   "workstation-disk-size",
			Usage:  "VMWare Workstation size of disk for host VM (in MB)",
			Value:  defaultDiskSize,
		},
		mcnflag.StringFlag{
			EnvVar: "WORKSTATION_SSH_USER",
			Name:   "workstation-ssh-user",
			Usage:  "SSH user",
			Value:  defaultSSHUser,
		},
		mcnflag.StringFlag{
			EnvVar: "WORKSTATION_SSH_PASSWORD",
			Name:   "workstation-ssh-password",
			Usage:  "SSH password",
			Value:  defaultSSHPass,
		},
		mcnflag.BoolFlag{
			EnvVar: "WORKSTATION_NO_SHARE",
			Name:   "workstation-no-share",
			Usage:  "Disable the mount of your home directory",
		},
		mcnflag.StringFlag{
			EnvVar: "WORKSTATION_SHARE_FOLDER",
			Name:   "workstation-share-folder",
			Usage:  "Mount the specified directory instead of the default home location. Format: name:dir",
		},
	}
}

package libvirt

import (
	"github.com/code-ready/crc/pkg/crc/constants"
	"path/filepath"

	"github.com/code-ready/crc/pkg/crc/machine/config"

	"github.com/code-ready/machine/libmachine/drivers"
)

type libvirtDriver struct {
	*drivers.BaseDriver

	// Driver specific configuration
	Memory      int
	CPU         int
	Network     string
	DiskPath    string
	CacheMode   string
	IOMode      string
	DiskPathUrl string
	SSHKeyPath  string
}

func CreateHost(machineConfig config.MachineConfig) *libvirtDriver {
	return &libvirtDriver{
		BaseDriver: &drivers.BaseDriver{
			MachineName: machineConfig.Name,
			StorePath:   constants.MachineBaseDir,
			SSHUser:     constants.DefaultSSHUser,
			BundleName:  machineConfig.BundleName,
		},
		Memory:    machineConfig.Memory,
		CPU:       machineConfig.CPUs,
		Network:   DefaultNetwork,
		CacheMode: DefaultCacheMode,
		IOMode:    DefaultIOMode,
		// This force to add entry of DiskPath under crc machine config.json
		DiskPath:    filepath.Join(constants.MachineBaseDir, "machines", machineConfig.Name, constants.DefaultName),
		DiskPathUrl: machineConfig.DiskPathURL,
		SSHKeyPath:  machineConfig.SSHKeyPath,
	}
}

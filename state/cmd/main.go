package main

import (
	"crypto/rand"
	"flag"
	"os"
	"path/filepath"

	"github.com/edgelesssys/constellation/coordinator/config"
	"github.com/edgelesssys/constellation/internal/utils"
	"github.com/edgelesssys/constellation/state/keyservice"
	"github.com/edgelesssys/constellation/state/mapper"
)

const (
	keyPath = "/run/cryptsetup-keys.d"
	keyFile = "state.key"
)

var csp = flag.String("csp", "", "Cloud Service Provider the image is running on")

func main() {
	flag.Parse()
	diskPath, err := mapper.GetDiskPath(*csp)
	if err != nil {
		utils.KernelPanic(err)
	}

	mapper, err := mapper.New(diskPath)
	if err != nil {
		utils.KernelPanic(err)
	}
	defer mapper.Close()

	if mapper.IsLUKSDevice() {
		uuid := mapper.DiskUUID()
		_, err = keyservice.WaitForDecryptionKey(*csp, uuid)
	} else {
		err = formatDisk(mapper)
	}

	if err != nil {
		utils.KernelPanic(err)
	}
}

func formatDisk(mapper *mapper.Mapper) error {
	// generate and save temporary passphrase
	if err := os.MkdirAll(keyPath, os.ModePerm); err != nil {
		utils.KernelPanic(err)
	}
	passphrase := make([]byte, config.RNGLengthDefault)
	if _, err := rand.Read(passphrase); err != nil {
		utils.KernelPanic(err)
	}
	if err := os.WriteFile(filepath.Join(keyPath, keyFile), passphrase, 0o400); err != nil {
		utils.KernelPanic(err)
	}

	if err := mapper.FormatDisk(string(passphrase)); err != nil {
		utils.KernelPanic(err)
	}
	if err := mapper.MapDisk("state", string(passphrase)); err != nil {
		utils.KernelPanic(err)
	}

	return nil
}

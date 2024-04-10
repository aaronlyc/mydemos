package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Step 1: Create a physical volume
	pvName := "/dev/sdb"
	if err := createPhysicalVolume(pvName); err != nil {
		fmt.Printf("Error creating physical volume: %v\n", err)
		return
	}
	fmt.Printf("Physical volume %s created successfully.\n", pvName)

	// Step 2: Create a volume group
	vgName := "my_vg"
	if err := createVolumeGroup(vgName, pvName); err != nil {
		fmt.Printf("Error creating volume group: %v\n", err)
		return
	}
	fmt.Printf("Volume group %s created successfully.\n", vgName)

	// Step 3: Create a logical volume
	lvName := "my_lv1"
	lvSize := "10G"
	if err := createLogicalVolume(lvName, lvSize, vgName); err != nil {
		fmt.Printf("Error creating logical volume: %v\n", err)
		return
	}
	fmt.Printf("Logical volume %s created successfully.\n", lvName)

	// Step 4: Format the logical volume
	if err := formatLogicalVolume(lvName, vgName); err != nil {
		fmt.Printf("Error formatting logical volume: %v\n", err)
		return
	}
	fmt.Printf("Logical volume %s formatted successfully.\n", lvName)

	// Step 5: Mount the logical volume
	mountPath := "/mnt/my_lv1"
	if err := mountLogicalVolume(lvName, vgName, mountPath); err != nil {
		fmt.Printf("Error mounting logical volume: %v\n", err)
		return
	}
	fmt.Printf("Logical volume %s mounted at %s successfully.\n", lvName, mountPath)
}

func createPhysicalVolume(pvName string) error {
	if pvName == "" {
		return errors.New("pvName is empty")
	}
	return commandWithOutput("pvcreate", pvName)
}

func createVolumeGroup(vgName, pvName string) error {
	if pvName == "" || vgName == "" {
		return errors.New("pvName or vgName is empty")
	}
	return commandWithOutput("vgcreate", vgName, pvName)
}

func createLogicalVolume(lvName, lvSize, vgName string) error {
	if lvName == "" || lvSize == "" || vgName == "" {
		return errors.New("pvName or lvSize or vgName is empty")
	}
	return commandWithOutput("lvcreate", "-L", lvSize, "-n", lvName, vgName)
}

func formatLogicalVolume(lvName, vgName string) error {
	if lvName == "" {
		return errors.New("lvName is empty")
	}
	return commandWithOutput("mkfs.ext4", "/dev/"+vgName+"/"+lvName)
}

func mountLogicalVolume(lvName, vgName, mountPath string) error {
	if lvName == "" || mountPath == "" {
		return errors.New("lvName or mountPath is empty")
	}

	// Create mount directory if not exists
	if _, err := os.Stat(mountPath); os.IsNotExist(err) {
		if err := os.MkdirAll(mountPath, 0755); err != nil {
			return err
		}
	}

	return commandWithOutput("mount", "/dev/"+vgName+"/"+lvName, mountPath)
}

func commandWithOutput(name string, arg ...string) error {
	// 创建一个执行命令的cmd对象
	cmd := exec.Command(name, arg...)

	fmt.Printf("exec: [%s]\n", cmd.String())

	output, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(output))
	}
	return nil
}

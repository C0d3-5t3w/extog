package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if os.Geteuid() != 0 {
		fmt.Println("This program must be run as root.")
		return
	}

	fmt.Print("Do you want to remove /boot/firmware/config.txt? (yes/no): ")
	var removeChoice string
	fmt.Scanln(&removeChoice)

	removeChoice = strings.ToLower(removeChoice)
	if removeChoice == "yes" {
		_, err := os.Stat("/boot/firmware/config.txtbk")
		if os.IsNotExist(err) {
			cpCmd := exec.Command("cp", "/boot/firmware/config.txt", "/boot/firmware/config.txtbk")
			err = cpCmd.Run()
			if err != nil {
				fmt.Println("Error creating backup:", err)
				return
			}
			fmt.Println("Backup created successfully.")
		}

		rmCmd := exec.Command("rm", "/boot/firmware/config.txt")
		err = rmCmd.Run()
		if err != nil {
			fmt.Println("Error removing config.txt:", err)
			return
		}
		fmt.Println("config.txt removed successfully.")
	} else {
		fmt.Println("Exiting program because config.txt was not removed.")
		return
	}

	var choice string
	for {
		fmt.Print("Enter 'on'(yes external) or 'off'(no external): ")
		fmt.Scanln(&choice)

		choice = strings.ToLower(choice)
		if choice == "on" || choice == "off" {
			break
		} else {
			fmt.Println("Invalid input, restoring backup...")
			restoreCmd := exec.Command("cp", "/boot/firmware/config.txtbk", "/boot/firmware/config.txt")
			err := restoreCmd.Run()
			if err != nil {
				fmt.Println("Error restoring backup:", err)
				return
			}
			fmt.Println("Backup restored successfully. Exiting program.")
			return
		}
	}

	var fileName string
	if choice == "on" {
		fileName = "on.txt"
	} else if choice == "off" {
		fileName = "off.txt"
	}

	cmd := exec.Command("cp", fileName, "-r", "/boot/firmware/config.txt")

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error, you might not have a config, please rerun:", err)
		fmt.Println("If you are having trouble try running, 'sudo cp off.txt -r /boot/firmware/config.txt' yourself.")
		fmt.Println("If you are still having trouble, please contact 5T3W.")
	} else {
		fmt.Println("File copied successfully.")
		rebootCmd := exec.Command("reboot")
		err := rebootCmd.Run()
		if err != nil {
			fmt.Println("Error rebooting the system:", err)
			fmt.Println("Please reboot the system manually.")
		} else {
			fmt.Println("System rebooting...")
		}
	}
}

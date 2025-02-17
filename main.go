package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if os.Geteuid() != 0 {
		fmt.Println("This program must be run as root.")
		return
	}

	fmt.Print("Do you want to remove /boot/firmware/config.txt? (yes/no): ")
	var removeChoice string
	fmt.Scanln(&removeChoice)

	if removeChoice == "yes" {
		rmCmd := exec.Command("rm", "/boot/firmware/config.txt")

		err := rmCmd.Run()
		if err != nil {
			fmt.Println("Error removing config.txt:", err)
			return
		} else {
			fmt.Println("config.txt removed successfully.")
		}
	} else {
		fmt.Println("Exiting program as config.txt was not removed.")
		return
	}

	var choice string
	for {
		fmt.Print("Enter 'on'(yes external) or 'off'(no external): ")
		fmt.Scanln(&choice)

		if choice == "on" || choice == "off" {
			break
		} else {
			fmt.Println("Invalid choice. Please enter 'on' or 'off'.")
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

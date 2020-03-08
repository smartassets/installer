package ui

import (
	"fmt"
	"os"
)

func ReportErr(message string){
	fmt.Printf("[ERROR] %s\n", message)
	os.Exit(1)
}

func ReportInfo(message string, args... string){
	fmt.Printf("[INFO] %s\n", fmt.Sprintf(message, args))
}

func ReportSimple(message string, args... string){
	fmt.Printf("%s\n", fmt.Sprintf(message, args))
}

func ReportWarning(message string, args... string){
	fmt.Printf("[WARNING] %s\n", fmt.Sprintf(message, args))
}
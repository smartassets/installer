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

func ReportInfoWithoutArgs(message string){
	fmt.Printf("[INFO] %s\n", message)
}

func ReportSimple(message string, args... string){
	fmt.Printf("%s\n", fmt.Sprintf(message, args))
}

func ReportSimpleWithoutArgs(message string){
	fmt.Printf("%s\n", message)
}

func ReportWarning(message string, args... string){
	fmt.Printf("[WARNING] %s\n", fmt.Sprintf(message, args))
}
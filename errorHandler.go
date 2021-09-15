package butil

import (
	"fmt"
	"log"
)

// check error with log.Fatal
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// just print the error
func CheckErrP(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// print error then return
func CheckErrPR(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

package butil

import (
	"fmt"
	"log"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckErrP(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func CheckErrPR(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

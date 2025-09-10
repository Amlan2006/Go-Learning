package main

import "fmt"

func conditions(x int) {
	if x < 18 {
		fmt.Println("You are a minor")
	}
	if x >= 18 && x < 60 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a senior citizen")
	}
}

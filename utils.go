package main

import (
	"strings"
)

func check(data Question) (int, string) {

	var total int
	var operation_type string

	math := map[string][]string{
		"addition":       {"add", "addition", "plus", "sum"},
		"subtraction":    {"subtract", "subtraction", "minus"},
		"multiplication": {"multiply", "multiplication", "times"},
	}

	for _, val := range math["addition"] {
		if strings.Contains(data.OperationType, val) {
			total = data.X + data.Y
			operation_type = "addition"
		}
	}

	for _, val := range math["multiplication"] {
		if strings.Contains(data.OperationType, val) {
			total = data.X * data.Y
			operation_type = "multiplication"

		}
	}

	for _, val := range math["subtraction"] {
		if strings.Contains(data.OperationType, val) {
			total = data.X - data.Y
			operation_type = "subtraction"

		}
	}

	return total, operation_type
}

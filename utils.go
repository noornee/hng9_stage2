package main

import (
	"strconv"
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

	// checks if the values of X and Y is nil
	// if the OperationType contains the values of X and Y,
	// loop through it and reassign them to X and Y
	if data.X == 0 && data.Y == 0 {
		arr := strings.Split(data.OperationType, " ")
		var num []int

		for _, word := range arr {
			val, err := strconv.Atoi(word)
			if err != nil {
				continue
			}
			num = append(num, val)
		}

		if len(num) == 2 {
			data.X = num[0]
			data.Y = num[1]
		}

	}

	for _, val := range math["addition"] {
		if strings.Contains(strings.ToLower(data.OperationType), val) {
			total = data.X + data.Y
			operation_type = "addition"
		}
	}

	for _, val := range math["multiplication"] {
		if strings.Contains(strings.ToLower(data.OperationType), val) {
			total = data.X * data.Y
			operation_type = "multiplication"

		}
	}

	for _, val := range math["subtraction"] {
		if strings.Contains(strings.ToLower(data.OperationType), val) {
			total = data.X - data.Y
			operation_type = "subtraction"

		}
	}

	return total, operation_type
}

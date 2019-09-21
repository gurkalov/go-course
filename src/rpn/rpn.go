package rpn

import (
	"course/calc"
	"course/calc/extend"
	"course/stack"
	"strconv"
	"strings"
)

func getPriority(str string) uint {
	_, err := strconv.ParseFloat(str, 32)
	if err == nil {
		return 0
	}

	mapPriority := map[string]uint{
		"(": 1,
		")": 1,
		"+": 2,
		"-": 2,
		"*": 3,
		"/": 3,
		"^": 4,
	}

	if v, ok := mapPriority[str]; ok {
		return v
	}

	panic("Error symbol!")
	return 0
}

func Rpn(str string) string {
	stack := stack.Stack{}
	result := []string{}
	words := strings.Split(str, " ")
	for _, v := range words {
		if v == "(" {
			stack.Push(v)
		} else if v == ")" {
			for s := stack.Pop(); s.(string) != "("; s = stack.Pop() {
				result = append(result, s.(string))
			}
		} else {
			priority := getPriority(v)
			if priority == 0 {
				result = append(result, v)
			} else {
				top := stack.Get()
				if top != nil {
					prTop := getPriority(top.(string))
					if prTop >= priority {
						result = append(result, stack.Pop().(string))
					}
				}
				stack.Push(v)
			}
		}
	}

	for s := stack.Pop(); s != nil; s = stack.Pop() {
		result = append(result, s.(string))
	}

	return strings.Join(result, " ")
}

func ComputeRpn(str string) float64 {
	stack := stack.Stack{}
	result := float64(0)
	words := strings.Split(str, " ")
	for _, v := range words {
		n, err := strconv.ParseFloat(v, 64)
		if err == nil {
			stack.Push(n)
		} else {
			a := stack.Pop().(float64)
			b := stack.Pop().(float64)
			result = calc.ComputeExtend(extend.Calc{b, a}, v)
			stack.Push(result)
		}
	}

	return result
}

func Compute (str string) float64 {
	return ComputeRpn(Rpn(str))
}

package jsonata

import (
	"encoding/json"
	"github.com/blues/jsonata-go"
)

func Eval(input string, expression string) (string, error) {
	expr, err := jsonata.Compile(expression)
	if err != nil {
		return "", err
	}
	var inputData interface{}
	err = json.Unmarshal([]byte(input), &inputData)
	if err != nil {
		return "", err
	}
	output, err := expr.Eval(inputData)
	if err != nil {
		return "", err
	}
	outputJSON, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return "", err
	}
	return string(outputJSON), nil
}

package cmd

import (
	"bufio"
	"fmt"
	"github.com/footgod368/jsonata-cli/internal/jsonata"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var (
	verbose bool
)

func preRun(args []string) error {
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
	return nil
}

func run(args []string) error {
	var jsonataExpr string
	if len(args) == 0 {
		return fmt.Errorf("no arguments provided")
	}
	jsonataExpr = args[0]

	var jsonStr string
	if len(args) == 1 {
		scanner := bufio.NewScanner(os.Stdin)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		jsonStr = strings.Join(lines, "\n")
	} else {
		jsonStr = args[1]
	}

	logrus.Debugln("jsonataExpr: ", jsonataExpr)
	logrus.Debugln("jsonStr: ", jsonStr)

	result, err := jsonata.Eval(jsonStr, jsonataExpr)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

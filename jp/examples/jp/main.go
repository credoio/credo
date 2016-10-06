package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/credoio/credo/jp"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

func main() {
	var recover = func() { recover(); fmt.Println("ERROR") }
	defer recover()
	progname := os.Args[0]
	flag.Parse()
	args := flag.Args()
	if len(args) < 3 {
		fmt.Fprintf(os.Stderr, "%s filename OPERATION PATH...\n", progname)
		fmt.Fprintf(os.Stderr, "\t OPERATION: len, print, if, foreach\n")
		os.Exit(1)
	}
	jsonStr := ""
	var err error
	if args[0] == "-" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			jsonStr += scanner.Text()
		}
	} else {
		buf, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "can't read file %s, %v\n", args[0], err)
		}
		jsonStr = string(buf)
	}
	if jsonStr == "" {
		fmt.Fprintf(os.Stderr, "empty input\n")
		os.Exit(0)
	}
	var v interface{}
	err = json.Unmarshal([]byte(jsonStr), &v)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't parse json string, %v\n", err)
		os.Exit(1)
	}
	if args[1] == "if" {
		if len(args) < 5 {
			fmt.Fprintf(os.Stderr, "insufficent arguments, %v,%d\n", args, len(args))
			fmt.Fprintf(os.Stderr, "%s filename if EXPRESSION OPERATION PATH\n", progname)
			os.Exit(1)
		}
		exp := strings.Fields(args[2])
		if len(exp) != 3 {
			fmt.Fprintf(os.Stderr, "invalid expression %s\n", args[2])
			fmt.Fprintf(os.Stderr, " valid expression should be: path  COMPARATOR value\n")
			os.Exit(1)
		}
		kind, res, err := jp.Jsonpathos(v, exp[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "json parse error, cannot find %s, %v\n", exp[0], err)
			os.Exit(1)
		}
		if !validString(kind) {
			os.Exit(1)
		}
		comparator := exp[1]
		if comparator != "==" && comparator != "!=" {
			fmt.Fprintf(os.Stderr, "invalid comparator %s\n", comparator)
			os.Exit(1)
		}
		//TODO: more comparators => , > , <=, <
		if (comparator == "==" && res == exp[2]) ||
			(comparator == "!=" && res != exp[2]) {
			op := args[3]
			checkOp(op)
			jp.HandleExp(v, op, args[4:])
		}
		os.Exit(0)
	}
	if args[1] == "foreach" {
		if len(args) < 4 {
			fmt.Fprintf(os.Stderr, "insufficient arguments, %v,%d\n", args, len(args))
			fmt.Fprintf(os.Stderr, "%s foreach PATH OPERATION\n", progname)
			os.Exit(1)
		}
		path := args[2]
		kind, res, err := jp.Jsonpathos(v, path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "json parse error, cannot find %s, %v\n", path, err)
			os.Exit(1)
		}
		l := jp.FindLen(kind, res)
		if l < 0 {
			fmt.Fprintf(os.Stderr, "cannot find len for %s\n", path)
			os.Exit(1)
		}
		op := args[3]
		if op == "print" {
			for i := 0; i < l; i++ {
				p := fmt.Sprintf("%s[%d]", path, i)
				kind, res, err = jp.Jsonpathos(v, p)
				if err != nil {
					fmt.Fprintf(os.Stderr, "cannot find %s", p)
					os.Exit(1)
				}
				if kind == reflect.Invalid {
					fmt.Fprintf(os.Stderr, "unexpected invalid type %s", p)
					continue
				}
				jp.Output(kind, op, res)
			}
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, "unsupported operation %s", op)
		os.Exit(1)
	}
	op := args[1]
	checkOp(op)
	jp.HandleExp(v, op, args[2:])
	os.Exit(0)
}

func validString(kind reflect.Kind) bool {
	if kind == reflect.Invalid {
		fmt.Fprintf(os.Stderr, "unexpected invalid type")
		return false
	}
	if kind != reflect.String {
		fmt.Fprintf(os.Stderr, "expected string but got %v\n", kind)
		return false
	}
	return true
}
func checkOp(op string) {
	if op == "if" || op == "len" || op == "print" || op == "foreach" || op == "list" {
		return
	}
	fmt.Fprintf(os.Stderr, "invalid operation %s\n", op)
	os.Exit(1)
}

package main

import (
	"errors"
	"fmt"
)

func f1(value int) (int, error) {
	if value >= 100 {
		return -1, errors.New("Value is equal or more than 100")
	}
	return value * value, nil
}

type handleException struct {
	argument int
	problem  string
}

func (e *handleException) Error() string {
	return fmt.Sprintf("%d - %s", e.argument, e.problem)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &handleException{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 101} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*handleException); ok {
		fmt.Println(ae.argument)
		fmt.Println(ae.problem)
	}

}

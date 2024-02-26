package _defer

import "fmt"

func RecoverFromPanicWithDefer(s string) (p string, err error) {

	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
			err = fmt.Errorf("recovered from panic")
		}
	}()

	p = func(s string) string {

		if len(s) > 3 {
			panic("oops")
		}

		return s + s
	}(s)

	return p, err
}

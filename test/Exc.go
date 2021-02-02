package test

import "fmt"

func Student(names ...string) {

	for _, name := range names {

		fmt.Println(name)

	}

}

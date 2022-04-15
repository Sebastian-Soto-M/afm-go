package main

import (
	"log"
)

func prefixedCheck(prefix string) func(err error) (bool, error) {
	return func(err error) (bool, error) {
		if err != nil {
			log.Printf("ERROR:[%s]\t%s", prefix, err)
			return false, err
		}
		return true, nil
	}
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		// fmt.Printf("%v == %v\n", v, e)
		if v == e {
			return true
		}
	}
	return false
}

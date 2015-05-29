// +build debug

package main

import "log"

func debug(fmt string, args ...interface{}) {
	log.Printf(fmt, args...)
}

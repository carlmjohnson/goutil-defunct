package goutils

import "log"

func NilOrDie(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
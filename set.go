package goutils

type set map[Data] naught

func NewSet() set {
    return make(set)
}
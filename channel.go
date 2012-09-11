package goutils

import "fmt"

type Channel interface {
    Each(func(Data))
    Map(func(Data) Data) DataChannel
    String() string
    Any(func(Data) bool) bool
    All(func(Data) bool) bool
}

func NewChannel() dataChannel {}
func NewChannelFromSlice() dataChannel {}

type dataChannel chan Data

func (c dataChannel) Each(f func(Data)) {
    for v := range c {
        f(v)
    }
}

func (c dataChannel) Map(f func(Data) Data) dataChannel {
    r := make(DataChannel)
    go func() {
        for v := range c {
            c <- f(v)
        }
        close(r)
    }()
    return r
}

func (c dataChannel) String() string {
    s = ""
    for v := range c {
        s += fmt.Sprint(v)
    }
    return s
}

func (c dataChannel) Any() bool {
}

func (c dataChannel) All() bool {
}

type numberChannel chan Number

func NewNumberChannel() numberChannel {}
func NewNumberChannelFromSlice() numberChannel {}


type stringChannel chan string

func NewStringChannel() stringChannel {}
func NewStringChannelFromSlice() stringChannel {}

func NewStringChannelFromFile(filename string) stringChannel {}

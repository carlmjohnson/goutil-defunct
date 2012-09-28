package channel

import (
	"fmt"
	"github.com/earthboundkid/goutil"
)

type Channel interface {
	Each(func(goutil.Data))
	Map(func(goutil.Data) goutil.Data) Channel
	String() string
	Any(func(goutil.Data) bool) bool
	All(func(goutil.Data) bool) bool
}

type DataChannel chan goutil.Data

func NewChannel() Channel {
	return make(DataChannel)
}
func NewChannelFromSlice() Channel {
	return make(DataChannel)
}

func (c DataChannel) Each(f func(goutil.Data)) {
	for v := range c {
		f(v)
	}
}

func (c DataChannel) Map(f func(goutil.Data) goutil.Data) Channel {
	r := make(DataChannel)
	go func() {
		for v := range c {
			c <- f(v)
		}
		close(r)
	}()
	return r
}

func (c DataChannel) String() string {
	s := ""
	for v := range c {
		s += fmt.Sprint(v)
	}
	return s
}

func (c DataChannel) Any(func(goutil.Data) bool) bool {
	return true
}

func (c DataChannel) All(func(goutil.Data) bool) bool {
	return true
}

type NumberChannel chan goutil.Number

func NewNumberChannel() NumberChannel {
	return make(NumberChannel)
}
func NewNumberChannelFromSlice() NumberChannel {
	return make(NumberChannel)
}

type StringChannel chan string

func NewStringChannel() StringChannel {
	return make(StringChannel)
}
func NewStringChannelFromSlice() StringChannel {
	return make(StringChannel)
}

func NewStringChannelFromFile(filename string) StringChannel {
	return make(StringChannel)
}

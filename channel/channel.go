package channel

import "fmt"

type Channel interface {
	Each(func(Data))
	Map(func(Data) Data) Channel
	String() string
	Any(func(Data) bool) bool
	All(func(Data) bool) bool
}

func NewChannel() Channel {
	return make(dataChannel)
}
func NewChannelFromSlice() Channel {
	return make(dataChannel)
}

type dataChannel chan Data

func (c dataChannel) Each(f func(Data)) {
	for v := range c {
		f(v)
	}
}

func (c dataChannel) Map(f func(Data) Data) Channel {
	r := make(dataChannel)
	go func() {
		for v := range c {
			c <- f(v)
		}
		close(r)
	}()
	return r
}

func (c dataChannel) String() string {
	s := ""
	for v := range c {
		s += fmt.Sprint(v)
	}
	return s
}

func (c dataChannel) Any(func(Data) bool) bool {
	return true
}

func (c dataChannel) All(func(Data) bool) bool {
	return true
}

type numberChannel chan Number

func NewNumberChannel() numberChannel {
	return make(numberChannel)
}
func NewNumberChannelFromSlice() numberChannel {
	return make(numberChannel)
}

type stringChannel chan string

func NewStringChannel() stringChannel {
	return make(stringChannel)
}
func NewStringChannelFromSlice() stringChannel {
	return make(stringChannel)
}

func NewStringChannelFromFile(filename string) stringChannel {
	return make(stringChannel)
}

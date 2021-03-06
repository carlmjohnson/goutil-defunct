package ranging

import (
	"github.com/earthboundkid/goutil"
	"github.com/earthboundkid/goutil/channel"
)

func positiveRange(from, to, step goutil.Number, ch channel.NumberChannel) {
	for v := from; v < to; v += step {
		ch <- v
	}
	close(ch)
}

func negativeRange(from, to, step goutil.Number, ch channel.NumberChannel) {
	for v := from; to < v; v += step {
		ch <- v
	}
	close(ch)
}

func FromToStep(from, to, step goutil.Number) channel.NumberChannel {

	ch := channel.NewNumberChannel()

	if step > 0 {
		//If the Step is positive then wait for v to be greater than Stop
		go positiveRange(from, to, step, ch)
	} else {
		//And vice versa, if negative, we need to wait for v to be lesser
		go negativeRange(from, to, step, ch)
	}

	return ch
}

func To(to goutil.Number) channel.NumberChannel {

	ch := channel.NewNumberChannel()

	go positiveRange(0, to, 1, ch)

	return ch

}

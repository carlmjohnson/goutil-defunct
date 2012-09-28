package ranging

import (
	"github.com/earthboundkid/goutil"
	"github.com/earthboundkid/goutil/channel"
)

func FromToStep(from, to, step goutil.Number) channel.NumberChannel {

	ch := channel.NewNumberChannel()

	if step > 0 {
		//If the Step is positive then wait for v to be greater than Stop
		go func() {
			for v := from; v < to; v += step {
				ch <- v
			}
			close(ch)
		}()
	} else {
		//And vice versa, if negative, we need to wait for v to be lesser
		go func() {
			for v := from; v > to; v += step {
				ch <- v
			}
			close(ch)
		}()
	}

	return ch
}

func To(to goutil.Number) channel.NumberChannel {
	return RangeFromToStep(0, to, 1)
}

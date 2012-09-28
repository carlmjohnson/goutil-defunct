package range

import "github.com/earthboundkid/goutils/channel"

func RangeFromToStep(from, to, step Number) channel.numberChannel {

	ch := make(numberChannel)

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

func Range(to Number) numberChannel {
	return RangeFromToStep(0, to, 1)
}

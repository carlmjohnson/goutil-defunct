package ranging_test

import (
	"fmt"
	"github.com/earthboundkid/goutil/ranging"
	"testing"
)

func ExampleTo() {
	for i := range ranging.To(5) {
		fmt.Println(i)
	}
	//Output: 0
	//1
	//2
	//3
	//4
}

func ExampleFromToStep() {
	for i := range ranging.FromToStep(5, -5, -2) {
		fmt.Println(i)
	}
	//Output: 5
	//3
	//1
	//-1
	//-3
}

func TestSimpleRange(t *testing.T) {
	r := ranging.To(100)

	for i := 0; i < 100; i++ {
		v := <-r
		if i != int(v) {
			t.Errorf("Expected: %d Got: %d", i, v)
		}
	}
}

func TestUnevenStep(t *testing.T) {
    r := ranging.FromToStep(0, 5, 2)

    for i := 0; i < 5; i += 2 {
        v := <-r
        if i != int(v) {
            t.Errorf("Expected: %d Got: %d", i, v)
        }
    }
}

func TestNegativeStep(t *testing.T) {
	r := ranging.FromToStep(100, -100, -1)

	for i := 100; i > -100; i -= 1 {
		v := <-r
		if i != int(v) {
			t.Errorf("Expected: %d Got: %d", i, v)
		}
	}
}

func TestUnevenNegativeStep(t *testing.T) {
    r := ranging.FromToStep(100, -113, -20)

    for i := 100; i > -100; i -= 20 {
        v := <-r
        if i != int(v) {
            t.Errorf("Expected: %d Got: %d", i, v)
        }
    }
}

package ranging_test

import (
    "github.com/earthboundkid/goutil/ranging"
	_ "testing"
    "fmt"
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
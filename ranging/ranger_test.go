package ranging

import (
	_ "testing"
    "fmt"
)

func ExampleTo() {
    for i := range To(5) {
        fmt.Println(i)
    }
	//Output: 0
    //1
    //2
    //3
    //4
}

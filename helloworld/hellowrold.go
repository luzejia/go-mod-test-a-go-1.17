package helloworld

import "fmt"
import helloworldc "github.com/luzejia/go-mod-test-b/helloworldC"

func HelloWorld() {
    fmt.Println("I am a, I call on b's helloworldC")
    helloworldc.HelloWorld()
}

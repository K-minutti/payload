package payload

import "fmt"


func Hello(name string) string {
	msg := fmt.Sprintf("Hi, %v. Hello how are you?", name)
	return msg
}
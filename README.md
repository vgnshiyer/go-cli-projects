## Go lang cli projects for practice

#### 1. Loadbalancer

Concepts learnt:

- struct: a struct in go lang allows us to create a custom data structure. Similar to the ones in c++.
- An interface allows us to declare functions for the structs.
- We them define the implementation of the functions.
- '*' -> We use asterisk to pass a pointer to a variable or to dereference an address.
- '&' -> We use andpersand to pass the address of a variable.
- (If a function I want to call is asking for a pointer/reference to a variable, I must pass the address of that variable)
- Notice that we also pass pointers to functions that belong to a particular struct object. (Which is different that other languages like python and java)

Example for pass by value and pass by reference in go.
```
package main

import "fmt"

// int var
func makeChange(num *int) {
	*num = 11
}

// slice
func makeChangeInSlice(slice []int) {
	slice[0] = 1000
}

// struct
type animal struct {
	sound string
}

func (a *animal) changeSound() {
	a.sound = "meow"
}

func main() {
	var x int = 10
	fmt.Println(x) // 10
	makeChange(&x)
	fmt.Println(x) // 11

	var slice []int = []int{1, 2, 3}
	fmt.Println(slice) // [1 2 3]
	makeChangeInSlice(slice)
	fmt.Println(slice) // [1000 2 3]

	var cat = animal{}
	fmt.Println(cat) // {}
	cat.changeSound()
	fmt.Println(cat) // {meow}

}
```

Other concepts.

- A Proxy is a server which sits in front of a private network that forwards requests or responses between clients and servers.
- A forward proxy is for forwarding requests from clients sitting in a private network to the internet. It mainly blocks certain malicious websites from injesting bad data in the network. For example, your company blocking your traffic outside their network.
- A Reverse proxy does the opposite. It filters requests coming from the clients on the internet to the internal servers. It also acts as a load balancer to make sure no particular server is overloaded with requests.

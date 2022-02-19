package main

import "fmt"

func main() {
	a, b := 1, 2
	add(&a, &b)
	fmt.Println(a)
	c := &a
	d := &c
	fmt.Println("d=", d, "d=", *d, "**d=", **d)

	m := map[string]string{}
	mp1 := &m
	put(m)
	fmt.Println(*mp1)
}

func put(m map[string]string) {
	m["a"] = "AAA"
}

func add(a, b *int) {
	*a = *a + *b
}

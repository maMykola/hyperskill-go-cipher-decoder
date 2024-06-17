package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var g, p, alice int32

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "g is %d and p is %d", &g, &p)
	fmt.Println("Ok")

	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "A is %d", &alice)

	b := int32(rand.Intn(int(p)))
	bob := calcXMod(g, p, b)
	secret := calcXMod(alice, p, b)

	fmt.Printf("B is %d\n", bob)
	fmt.Println(encode("Will you marry me?", secret))

	scanner.Scan()
	answer := scanner.Text()

	switch decode(answer, secret) {
	case "Yeah, okay!":
		fmt.Println(encode("Great!", secret))
	case "Let's be friends.":
		fmt.Println(encode("What a pity!", secret))
	}
}

func calcXMod(g, p, b int32) int32 {
	var c int32 = 1
	var i int32

	for i = 0; i < b; i++ {
		c = (c * g) % p
	}

	return c
}

func encode(origin string, secret int32) string {
	return transform(origin, secret)
}

func decode(origin string, secret int32) string {
	return transform(origin, -secret)
}

func transform(origin string, secret int32) string {
	var message = make([]byte, 0, len(origin))

	for _, c := range origin {
		if c >= 'a' && c <= 'z' {
			message = append(message, shift(c, 'a', secret))
		} else if c >= 'A' && c <= 'Z' {
			message = append(message, shift(c, 'A', secret))
		} else {
			message = append(message, byte(c))
		}
	}

	return string(message)
}

func shift(v int32, c byte, secret int32) byte {
	z := v - int32(c)
	secret = 26 + secret%26

	return c + byte((z+secret)%26)
}

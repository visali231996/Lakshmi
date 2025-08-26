package multi

import (
	"fmt"
	"strings"
)

func Factorial() {
	var n int
	fact := 1
	fmt.Println("enter n factorial")
	fmt.Scan(&n)
	temp := n
	for i := n; n > 0; {
		fact = fact * i
		n = n - 1
		i = i - 1
	}
	fmt.Println("factorial of", temp, "is equal to", fact)

}
func Wordcount() {
	var p string
	p = "go lang go lang"
	m := make(map[string]int)
	words := strings.Fields(p)
	for _, i := range words {
		m[i]++
	}
	for i, j := range m {
		fmt.Println(i, j)
	}
}

func Fibonocci() {
	f1 := 0
	f2 := 1
	var n, f3 int
	a := []int{0, 1}
	fmt.Println("enter n in fibonocci")
	fmt.Scan(&n)
	for i := 2; i < n; i++ {
		f3 = f1 + f2
		a = append(a, f3)
		f1 = f2
		f2 = f3

	}
	fmt.Println("fibonocci series:", a)
}

func Palindrome() {
	var s string
	var l int
	fmt.Println("enter the string to check for palindrome")
	fmt.Scan(&s)
	l = len(s)
	temp := 0
	for i, j := 0, l-1; i < l/2 && i != j; i++ {
		if string(s[i]) == string(s[j]) {
			temp++
			continue
		} else {
			temp--
		}
		j--
	}
	if temp == l/2 {
		fmt.Println("s is  palindrome")

	} else {
		fmt.Println("s is not palindrome")

	}

}

func Gethindi(name string) {
	fmt.Println("Hello in Hindi", name)
}
func Gettamil(name string) {
	fmt.Println("Hello in Tamil", name)
}
func Gettelugu(name string) {
	fmt.Println("Hello in Telugu", name)
}
func Getspanish(name string) {
	fmt.Println("Hello in Spanish", name)
}

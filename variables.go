package main

import "fmt"

type peter struct {
   pi		float32
   counter	int16
   flag		bool
}

func incrementp(a *int) {
   println("\t\t\t\t\t\tEntering incrementp\n\n")
   // *a is the value held in the address of &a - the leter in the postbox
   // &a is the address of this local (a *int) - the new postbox address 
   // a is the value of what was passed (an address) from func point2 - the original postbox address that holds *a
   println("Before: ", *a, &a, a)
   *a = *a + 1
   println("After:  ", *a, &a, a)
   println("\t\t\t\t\t\tLeaving incrementp\n\n")
}

func increment(a int) int {
   println("\t\t\t\t\t\tEntering increment\n\n")
   println("Before: ", a, &a)
   a = a + 1
   println("After:  ", a, &a)
   println("\t\t\t\t\t\tLeaving increment\n\n")
   return a
}

func point2() {
   var a int

   a = 1
   println("\t\t\t\t\t\tEntering point2\n\n")
   println("Before: ", a, &a)
   a = increment(a)
   incrementp(&a)
   println("After:  ", a, &a)
   println("\t\t\t\t\t\tLeaving point2\n\n")
}

func point1() {
   var a int

   a = 10
   fmt.Printf("\t\t\t\t\t\tEntering point1\n\n")
   fmt.Printf("Before: \t %v at the address %v\n", a, &a)
   a++
   fmt.Printf("After: \t\t %v at the address %v\n", a, &a)
   fmt.Printf("\t\t\t\t\t\tLeaving point1\n\n")
}

func sub1 () {
   type going struct {
	pi	float32
	counter	int16
	flag	bool
   }
   var a   going

   a.pi = 3.22
   a.counter = 40
   a.flag = true

   fmt.Printf("\t\t\t\t\t\tEntering sub1\n\n")
   fmt.Printf("struct going \t %T \t [%+v]\n\n\n", a, a)
   fmt.Printf("\t\t\t\t\t\tLeaving sub1\n\n")
}

func main() {
   // Declare variables that are set to their zero values

   var a int
   var b string
   var c float64
   var d bool
   // This is the outside struct declaration
   var e peter

   // This is the inside struct declaration
   type fetter struct {
	pi	float32
	counter	int16
	flag	bool
   }
   var f  fetter

   g := struct {
	pi	float32
	counter	int16
	flag	bool
   }{
	pi:	6.28,
	counter:	30,
	flag:	true,
   }

   // This is assigning to the outside struct peter
   e.pi = 3.14159265358979323846
   e.counter = 12349
   e.flag = true

   fmt.Printf("\n\n\n\n\n\n\n\n\n\n")
   fmt.Printf("\t\t\t\t\t\tEntering Main\n\n")
   fmt.Printf("var a int \t %T [%v] \n", a, a)
   fmt.Printf("var b string \t %T [%s] [%v] \n", b, b, b)
   fmt.Printf("var c float64 \t %T [%f] [%v]\n", c, c, c)
   fmt.Printf("var d bool \t %T [%v] \n\n", d, d)
   fmt.Printf("struct e \t %T \t\t\t\t\t\t [%+v]\n", e, e)
   fmt.Printf("struct f \t %T \t [%+v]\n", f, f)
   fmt.Printf("struct g \t %T \t [%+v]\n\n", g, g)

   // Declare variables and initialize
   // Using the short variable declaration method 

   aa := 1000000000000000
   bb := "Hello, World!"
   cc := 3.14159265358979323846
   dd := true

   fmt.Printf("aa := 1000000000000000 \t %T [%v] \n", aa, aa)
   fmt.Printf("bb := \"Hello, World!\" \t %T [%s] [%v] \n", bb, bb, bb)
   fmt.Printf("cc := 3.14159265358979323846 \t %T [%0.8f] [%v] \n", cc, cc, cc)
   fmt.Printf("dd := true \t %T [%v] \n\n\n\n", dd, dd)

   sub1()
   point1()
   point2()
   fmt.Printf("\t\t\t\t\t\tLeaving Main\n\n")

}

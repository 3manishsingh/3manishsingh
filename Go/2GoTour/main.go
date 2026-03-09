package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"sync"
	"time"
)

func add(a, b int) int {
	return a + b
}

func swap(a, b string) (string, string) {
	return b, a
}

func digits(num int) (x, y int) {
	x = num / 10
	y = num % 10
	return
}

var c bool

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z1     complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

type Vertex struct {
	X, Y int
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func say1(s string) {
	for range 5 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func fibonacci1(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("fib2:", "quit")
			return
		}
	}
}

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

// ------------------------------------------------------------------------------------------------------

func main() {

	fmt.Printf("number: %g is sqrt of 7\n", math.Sqrt(7))
	fmt.Println("Pi:", math.Pi)

	fmt.Println("add:", add(1, 2))

	a, b := swap("hello", "world")

	fmt.Println("swapped:", a, b)

	fmt.Println(digits(12))

	var d int = 1
	e := 2

	fmt.Println(c, d, e)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z1, z1)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	fmt.Println(Small, needInt(Small))
	fmt.Println(Small, needFloat(Small))
	fmt.Println(needFloat(Big))

	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum:", sum)

	for sum < 100 {
		sum += sum
	}
	fmt.Println("sum:", sum)

	for sum < 1000 {
		sum += sum
	}
	fmt.Println("sum:", sum)

	if sum < 100 {
		fmt.Println("if 1 sum:", sum)
	} else {
		fmt.Println("if 2 sum:", sum)
	}

	today := time.Now().Weekday()
	fmt.Println("today:", today)
	fmt.Println("When's Saturday?")
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	switch {
	case sum < 100:
		fmt.Println("case1")
	case sum > 100:
		fmt.Println("case2")
	}

	fmt.Println("x:", x)

	p := &x
	fmt.Println("pointer x:", *p)

	*p = 5

	fmt.Println("x:", x)

	ver := Vertex{1, 2}

	fmt.Println("Vertex struct:", ver, ver.X, ver.Y)

	ver2 := Vertex{X: 1}

	fmt.Println("Vertex struct:", ver2)

	var str1 [2]string
	str1[0] = "hello"

	fmt.Println(str1)

	primes := [3]int{2, 3, 5}
	fmt.Println("primes", primes, primes[0])

	var prime1 []int = primes[0:2]
	fmt.Println("prime1", prime1)

	names := [4]string{
		"John",
		"Pinki",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a12 := names[0:2]
	b12 := names[2:4]
	fmt.Println(a12, b12)

	b12[0] = "XXX"
	fmt.Println(a12, b12)
	fmt.Println(names)

	struct1 := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
	}
	fmt.Println(struct1)

	fmt.Println(primes[0:3], primes[:3], primes[0:], primes[:])

	newstr := make([]int, 5)
	fmt.Println(newstr)

	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	fmt.Println(board)

	newstr = append(newstr, 1)
	fmt.Println(newstr)

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("i:%d, v:%d\n", i, v)
	}

	map1 := make(map[int]int)
	map1[1] = 1
	fmt.Println(map1)

	type Vertex struct {
		Lat, Long float64
	}

	var m12 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m12)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))

	fmt.Println(compute(hypot))

	pos, neg := adder(), adder()
	for i := 0; i < 3; i++ {
		fmt.Println(
			pos(i),
			neg(-i),
		)
	}

	go say1("hello")
	say1("world")

	s1 := []int{1, 2, 3}

	ch1 := make(chan int)
	go sum1(s1[:len(s1)/2], ch1)
	go sum1(s1[len(s1)/2:], ch1)
	x1, y1 := <-ch1, <-ch1
	fmt.Println(x1, y1, x1+y1)

	bufferCh := make(chan int, 2)

	bufferCh <- 1
	bufferCh <- 2
	fmt.Println(<-bufferCh)
	fmt.Println(<-bufferCh)

	ch2 := make(chan int, 2)
	go fibonacci1(cap(ch2), ch2)
	for i := range ch2 {
		fmt.Println("fib1:", i)
	}

	ch3 := make(chan int)
	quit := make(chan int)
	go func() {
		for range 3 {
			fmt.Println("fib2:", <-ch3)
		}
		quit <- 0
	}()
	fibonacci2(ch3, quit)

	ch4 := SafeCounter{v: make(map[string]int)}
	for range 1000 {
		go ch4.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(ch4.Value("somekey"))

}

package main

import (
	"errors"
	"fmt"
	"iter"
	"maps"
	"math"
	"slices"
	"sync"
	"time"
)

const s string = "constant"

func main() {
	fmt.Println("Hello!")

	fmt.Println("Hello! " + "golang")
	fmt.Println(1 + 2)
	fmt.Println(22.0 / 7.0)
	fmt.Println(true)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f) //The := syntax is shorthand for declaring and initializing a variable,
	// e.g. for var f string = "apple" in this case.
	// This syntax is only available inside functions.

	fmt.Println(s)

	const n = 5000000
	fmt.Println(n)

	const g = 3e10 / n
	fmt.Println(g)

	fmt.Println(int64(g))

	fmt.Println(math.Sin(g))

	i := 0
	for i < 5 {
		fmt.Println("i=", i)
		i = i + 1
	}

	for j := 0; j < 5; j++ {
		fmt.Println("j=", j)
	}

	for k := range 5 {
		fmt.Println("k=", k)
	}

	for {
		fmt.Println("loop")
		break
	}

	for l := range 5 {
		if l%2 == 0 {
			continue
		}
		fmt.Println("l=", l)
	}

	if f == "apple" {
		fmt.Println("f=", f, "true")
	} else {
		fmt.Println("false")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	m := 1
	fmt.Print("Write ", m, " as ")
	switch m {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("None")
	}

	day := time.Now().Weekday()
	fmt.Println("day=", day)
	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	timeNow := time.Now()
	fmt.Println("time=", timeNow)
	fmt.Println("time=", timeNow.Hour(), timeNow.Minute())

	whatAmI := func(i any) { //A type switch compares types instead of values.
		// You can use this to discover the type of an interface value.
		// In this example, the variable t will have the type corresponding to its clause.
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	var o [5]int
	fmt.Println("o= ", o)

	o[0] = 100

	fmt.Println("o= ", o)
	fmt.Println("o= ", len(o))

	p := [5]int{1, 2, 3, 4, 5}

	fmt.Println("p= ", p)

	q := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("q= ", q)

	r := [...]int{1, 2, 10: 8, 9}
	fmt.Println("r= ", r)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)

	var s []string
	fmt.Println(s, len(s), cap(s))

	s = make([]string, 2)
	s[0] = "s"
	fmt.Println(s, s[0], len(s), cap(s))

	s = append(s, "b")
	fmt.Println(s, s[0], len(s), cap(s))

	s = append(s, "c", "d")
	fmt.Println(s, s[0], len(s), cap(s))

	s = append(s, "c", "d")
	fmt.Println(s, s[0], len(s), cap(s))
	s = append(s, "c", "d")
	fmt.Println(s, s[0], len(s), cap(s))

	t := make([]string, len(s))
	fmt.Println(t, t[0], len(t), cap(t))
	copy(t, s)
	fmt.Println(t, t[0], len(t), cap(t))

	u := t[2:4]
	fmt.Println(u)

	v := t[2:]
	w := t[:5]
	fmt.Println(v, w)

	x := []string{"a", "b", "c"}
	fmt.Println(x)

	y := []string{"abc"}
	fmt.Println(y)

	if slices.Equal(x, y) {
		fmt.Println("x == y")
	} else {
		fmt.Println("x not = y")
	}

	twoD1 := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD1[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD1[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD1)

	a1 := make(map[string]int)
	a1["a"] = 1
	a1["b"] = 2
	fmt.Println(a1)

	a2 := a1["a"]
	fmt.Println(a2)

	fmt.Println(len(a1))

	delete(a1, "b")
	fmt.Println("map:", a1)
	clear(a1)
	fmt.Println("map:", a1)

	_, prs := a1["a"]
	fmt.Println("prs:", prs)

	n1 := map[string]int{"foo": 1, "bar": 2}
	n2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n1, n2)

	if maps.Equal(n1, n2) {
		fmt.Println("n1 == n2")
	}

	res := plus(1, 2)
	fmt.Println("res", res)

	val1, val2 := vals()
	fmt.Println(val1, val2)

	_, val3 := vals()
	fmt.Println(val3)

	sum(1, 2, 3)

	num1 := []int{1, 2, 3, 4}
	sum(num1...)

	ints := intSeq()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println(newInts())

	fmt.Println(fact(4))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(3))

	b1 := []int{1, 3, 5, 7, 9}

	sum := 0

	for _, num := range b1 {
		sum += num
	}
	fmt.Println("sum b1:", sum)

	for i, n := range b1 {
		fmt.Println(i, n)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for v := range kvs {
		fmt.Printf("%s\n", v)
	}

	for i, c := range "AZaz" {
		fmt.Println(i, c)
	}

	iptr := 1
	ptr(&iptr)
	fmt.Println("iptr: ", iptr)

	const c1 = "สวัสดี"

	for i := range len(c1) {
		fmt.Printf("%x ", c1[i])
	}
	fmt.Println(person{"Manish", 30})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	d1 := person{name: "Manish", age: 30}
	fmt.Println(d1.name)

	d2 := &d1

	fmt.Println(d2.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

	r1 := rect{width: 2, height: 4}
	fmt.Println("area:", r1.area())
	fmt.Println("perim:", r1.perim())

	rp := &r1
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	r12 := rect1{width: 3, height: 4}
	c12 := circle{radius: 5}

	measure(r12)
	measure(c12)

	d123 := Dog{}
	c123 := Cat{}
	MakeSound(d123)
	MakeSound(c123)

	ns := transition(StateIdle)
	fmt.Println(ns)

	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe())

	var s11 = []string{"foo", "bar", "zoo"}

	fmt.Println("index of zoo:", SlicesIndex(s11, "zoo"))

	_ = SlicesIndex(s11, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())

	for e := range lst.All() {
		fmt.Println(e)
	}

	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}
		fmt.Println("Tea is ready!")
	}

	f1("direct")

	go f1("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second * 2)
	fmt.Println("done")

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		chan1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "two"
	}()

	for range 2 {
		select {
		case msg1 := <-chan1:
			fmt.Println("received", msg1)
		case msg2 := <-chan2:
			fmt.Println("received", msg2)
		}
	}

	con1 := Container{

		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			con1.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(con1.counters)

	// ------------------------------------------------------------------------------------------------------

}

func plus(a, b int) int {

	return a + b
}

func vals() (int, int) {
	return 1, 2
}

func sum(nums ...int) {
	fmt.Println(nums, "")

	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		fmt.Println("intSeq")
		return i
	}

}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)

}

func ptr(iptr *int) {
	*iptr = 10000000
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 30
	return &p
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

type geometry interface {
	area1() float64
	perim1() float64
}

type rect1 struct {
	width, height float64
}

func (r rect1) area1() float64 {
	return r.width * r.height
}
func (r rect1) perim1() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area1() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim1() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("--")
	fmt.Println(g)
	fmt.Println(g.area1())
	fmt.Println(g.perim1())
}

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}

func (c Cat) Speak() string { return "Meow!" }

func MakeSound(s Speaker) {
	fmt.Println(s.Speak())
}

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

type container struct {
	base
	str string
}

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func f1(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {

	n, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		log.Fatalf("failed to parse arg %s: %v", args[1], err)
	}

	timeIt("normal  ", n, fib)
	timeIt("parallel", n, fip0)

	return nil
}

func timeIt(name string, n int64, f func(int64) int64) {

	t0 := time.Now()
	r := f(n)
	t1 := time.Now()

	e := t1.Sub(t0)

	fmt.Printf("%s: %d %s\n", name, r, e)
}

func fib(n int64) int64 {

	if n <= 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func fip0(n int64) int64 {

	if n <= 1 {
		return 1
	}

	return <-sum(
		bkgd(func() int64 { return fip1(n - 1) }),
		bkgd(func() int64 { return fip1(n - 2) }))
}

func fip1(n int64) int64 {

	if n <= 1 {
		return 1
	}

	return <-sum(
		bkgd(func() int64 { return fip2(n - 1) }),
		bkgd(func() int64 { return fip2(n - 2) }))
}

func fip2(n int64) int64 {

	if n <= 1 {
		return 1
	}

	return <-sum(
		bkgd(func() int64 { return fib(n - 1) }),
		bkgd(func() int64 { return fib(n - 2) }))
}

func sum(c1, c2 chan int64) chan int64 {

	c := make(chan int64)

	go func() {
		defer close(c)

		var r int64

		for c1 != nil || c2 != nil {
			select {
			case i, ok := <-c1:
				if !ok {
					c1 = nil
				}
				r += i
			case i, ok := <-c2:
				if !ok {
					c2 = nil
				}
				r += i
			}
		}

		c <- r
	}()

	return c
}

func bkgd(f func() int64) chan int64 {

	c := make(chan int64)

	go func() {
		defer close(c)

		c <- f()
	}()

	return c
}

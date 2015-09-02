// A concurrent prime sieve

package main

import "fmt"
import "flag"
import "os"
import "strconv"

import "github.com/stathat/go"

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [nprime]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)

}

// The prime sieve: Daisy-chain Filter processes.
func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("nprime is Missing.")
		os.Exit(1)
	}

	nprime, _ := strconv.Atoi(args[0])

	fmt.Fprintf(os.Stdout, "Input number of primes to compute %d\n", nprime)

	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < nprime; i++ {
		prime := <-ch
		fmt.Fprintf(os.Stdout, "Prime %d - %d\n", i+1, prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
		stathat.PostEZCount("sieveMetricDuration", "EPhIk5lK9fjKJwGT", 1)
	}
}

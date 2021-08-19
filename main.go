package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	var iterationCount uint64


	flag.Uint64Var( &iterationCount, "Iterations", 100000000, "Number of iterations" )

	flag.Parse()

	result1 := evaluateNotEqualStrings(  iterationCount )

	result2 := evaluateLenStrings( iterationCount )

	fmt.Printf( "Using == averaged %vns per iteration\n", float64(result1) / float64(iterationCount) )
	fmt.Printf( "Using len() averaged %vns per iteration\n", float64(result2) / float64(iterationCount) )

}

func evaluateNotEqualStrings( iterations uint64 ) int64 {
	var i uint64

	var lvalueString string  = "Mama don't let your babies grow up to be Cowboys"
	lvalueString = ""

	start := time.Now()

	for i = 0; i < iterations; i++ {

		//  Test for empty string
		if lvalueString == "" {
		}
	}

	finish := time.Since( start )

	return finish.Nanoseconds()
}

func evaluateLenStrings( iterations uint64 ) int64 {
	var i uint64

	var lvalueString string  = "Mama don't let your babies grow up to be Cowboys"
	lvalueString = ""

	start := time.Now()

	for i = 0; i < iterations; i++ {

		//  Test for empty string
		if len( lvalueString ) < 1 {
		}
	}

	finish := time.Since( start )

	return finish.Nanoseconds()
}
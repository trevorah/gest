package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

// TestEvent is the output from test2json
// https://golang.org/cmd/test2json/
type TestEvent struct {
	Time time.Time // The Time field holds the time the event happened. It is conventionally omitted for cached test results.
	// encodes as an RFC3339-format string

	// The Action field is one of a fixed set of action descriptions:
	//
	// run    - the test has started running
	// pause  - the test has been paused
	// cont   - the test has continued running
	// pass   - the test passed
	// bench  - the benchmark printed log output but did not fail
	// fail   - the test or benchmark failed
	// output - the test printed output
	// skip   - the test was skipped or the package contained no tests
	Action string
	// The Package field, if present, specifies the package being tested. When the go command runs parallel tests in -json mode, events from different tests are interlaced; the Package field allows readers to separate them.
	Package string
	// The Test field, if present, specifies the test, example, or benchmark function that caused the event. Events for the overall package test do not set Test.
	Test string
	// The Elapsed field is set for "pass" and "fail" events. It gives the time elapsed for the specific test or the overall package test that passed or failed.
	// seconds
	Elapsed float64
	// The Output field is set for Action == "output" and is a portion of the test's output (standard output and standard error merged together). The output is unmodified except that invalid UTF-8 output from a test is coerced into valid UTF-8 by use of replacement characters. With that one exception, the concatenation of the Output fields of all output events is the exact output of the test execution.
	Output string
}

func report(c chan TestEvent) {
	r := JestRunner{}

	for e := range c {
		r.onEvent(e)
	}

}

func main() {

	c := make(chan TestEvent)
	go report(c)

	ticker := time.NewTicker(time.Second / 2)
	go func() {
		for range ticker.C {
			c <- TestEvent{}
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		bytes, err := reader.ReadBytes("\n"[0])
		if err != nil {
			if err == io.EOF {
				ticker.Stop()
				break
			}
			fmt.Println(err)

		} else {
			event := TestEvent{}
			json.Unmarshal(bytes, &event)
			c <- event
		}
	}

	// reader := bufio.NewReader(os.Stdin)

	// r := JestRunner{}

	// for {
	// 	s, err := reader.ReadString(byte("\n"[0]))
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		fmt.Println(err)

	// 	} else {
	// 		event := TestEvent{}
	// 		json.Unmarshal([]byte(s), &event)
	// 		// fmt.Printf("%+v\n", event)
	// 		r.onEvent(event)
	// 	}
	// }
}

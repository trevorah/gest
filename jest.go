package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/logrusorgru/aurora"
)

type JestRunner struct {
	started       time.Time
	suitesPassed  int
	suitesFailed  int
	suitesSkipped int
	testsPassed   int
	testsFailed   int
	testsSkipped  int
	running       []string
}

func (j *JestRunner) clearStatus() {
	if len(j.running) > 0 {
		lines := 4
		fmt.Print(strings.Repeat("\r\x1B[K\r\x1B[1A", lines))
	}
}

func (j *JestRunner) onEvent(e TestEvent) {

	j.clearStatus()

	// runs := ansi.Color(" RUNS ", "yellow+ib")

	if e.Action == "run" {
		j.running = append(j.running, e.Package)
		// dir, name := filepath.Split(e.Package)
		// fmt.Printf("%s %s%s\n", runs, aurora.Gray(dir), name)
	} else if e.Action == "fail" && e.Test == "" {
		dir, name := filepath.Split(e.Package)
		fmt.Printf("%s %s%s\n", aurora.BgRed(" FAIL "), aurora.Gray(dir), name)
	} else if e.Action == "pass" && e.Test == "" {
		dir, name := filepath.Split(e.Package)
		fmt.Printf("%s %s%s\n", aurora.BgGreen(" PASS "), aurora.Gray(dir), name)
	}

	fmt.Printf(`
Test Suites: 3 passed, 3 total
Tests:       553 passed, 553 total
Time:        5s
`)

	// print status

	// $ test
	//
	// 	RUNS  packages/jest-haste-map/src/__tests__/index.test.js
	// 	RUNS  packages/jest-mock/src/__tests__/jest_mock.test.js
	// 	RUNS  packages/expect/src/__tests__/matchers.test.js
	//
	// Test Suites: 0 of 308 total
	// Tests:       0 total
	// Snapshots:   0 total
	// Time:        1s

	// ~/D/jest $ node ./packages/jest-cli/bin/jest.js
	//  PASS  packages/jest-mock/src/__tests__/jest_mock.test.js
	//  PASS  packages/expect/src/__tests__/matchers.test.js
	//  PASS  packages/jest-haste-map/src/__tests__/index.test.js
	//
	//  RUNS  packages/pretty-format/src/__tests__/Immutable.test.js
	//  RUNS  packages/expect/src/__tests__/spyMatchers.test.js
	//  RUNS  packages/jest-config/src/__tests__/normalize.test.js
	//
	// Test Suites: 3 passed, 3 of 308 total
	// Tests:       553 passed, 553 total
	// Snapshots:   456 passed, 456 total
	// Time:        5s
}

func (j *JestRunner) onFinish() {

}

package examples

import (
    "fmt"
    "time"
    "github.com/golibry/go-params/params"
)

// This example shows how to convert raw string inputs to typed values using the
// params strconv-style helpers. Each GetAs* function returns (value, defaultUsed):
//  - value: the parsed value OR the provided default
//  - defaultUsed: true when the default was used (input empty/invalid), false otherwise
func StrconvExamples() {
    // String
    s, used := params.GetAsString("  hello  ", "default")
    fmt.Printf("string: %q (defaultUsed=%v)\n", s, used) // => "hello", false

    s2, used2 := params.GetAsString("   ", "default")
    fmt.Printf("string-empty: %q (defaultUsed=%v)\n", s2, used2) // => "default", true

    // Int
    i, usedI := params.GetAsInt("42", 0)
    fmt.Printf("int: %d (defaultUsed=%v)\n", i, usedI) // => 42, false

    i2, usedI2 := params.GetAsInt("not-an-int", 7)
    fmt.Printf("int-invalid: %d (defaultUsed=%v)\n", i2, usedI2) // => 7, true

    // Bool (supports true/false/1/0 and case-insensitive)
    b, usedB := params.GetAsBool("TRUE", false)
    fmt.Printf("bool: %t (defaultUsed=%v)\n", b, usedB) // => true, false

    b2, usedB2 := params.GetAsBool(" maybe ", true)
    fmt.Printf("bool-invalid: %t (defaultUsed=%v)\n", b2, usedB2) // => true, true

    // Float
    f, usedF := params.GetAsFloat("3.1415", 0)
    fmt.Printf("float: %f (defaultUsed=%v)\n", f, usedF) // => 3.141500, false

    f2, usedF2 := params.GetAsFloat("nan?", 2.5)
    fmt.Printf("float-invalid: %f (defaultUsed=%v)\n", f2, usedF2) // => 2.500000, true

    // Duration (Go duration format, parsed by time.ParseDuration)
    d, usedD := params.GetAsDuration("2h30m", time.Minute)
    fmt.Printf("duration: %s (defaultUsed=%v)\n", d, usedD) // => 2h30m0s, false

    d2, usedD2 := params.GetAsDuration("soon", 10*time.Second)
    fmt.Printf("duration-invalid: %s (defaultUsed=%v)\n", d2, usedD2) // => 10s, true

    // The same conversions are also available as methods on params.RawVal
    var raw params.RawVal = "  123  "
    i3, usedI3 := raw.GetAsInt(0)
    fmt.Printf("raw-int: %d (defaultUsed=%v)\n", i3, usedI3) // => 123, false
}

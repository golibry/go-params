package examples

import (
    "fmt"
    "net/url"
    "time"
    "github.com/golibry/go-params/params"
)

// URLExamples demonstrates reading URL query parameters using params.QueryParams.
// For each GetAs* call you get (value, defaultUsed). If the key is missing or
// cannot be parsed to the requested type, the default is returned and
// defaultUsed is true.
func URLExamples() {
    // Build a URL with various query parameters
    u, _ := url.Parse("https://example.com/search?q=  hello  &limit=25&exact=true&ratio=0.8&ttl=1h15m")

    qp := params.NewQueryParamsFromUrl(*u)

    // String (trimming applied)
    q, used := qp.GetAsString("q", "*")
    fmt.Printf("query string: %q (defaultUsed=%v)\n", q, used) // => "hello", false

    // Missing key -> default
    sort, usedMissing := qp.GetAsString("sort", "relevance")
    fmt.Printf("missing string: %q (defaultUsed=%v)\n", sort, usedMissing) // => "relevance", true

    // Int
    limit, usedI := qp.GetAsInt("limit", 10)
    fmt.Printf("limit: %d (defaultUsed=%v)\n", limit, usedI) // => 25, false

    // Bool
    exact, usedB := qp.GetAsBool("exact", false)
    fmt.Printf("exact: %t (defaultUsed=%v)\n", exact, usedB) // => true, false

    // Float
    ratio, usedF := qp.GetAsFloat("ratio", 0.5)
    fmt.Printf("ratio: %f (defaultUsed=%v)\n", ratio, usedF) // => 0.800000, false

    // Duration
    ttl, usedD := qp.GetAsDuration("ttl", 30*time.Minute)
    fmt.Printf("ttl: %s (defaultUsed=%v)\n", ttl, usedD) // => 1h15m0s, false

    // Invalid value -> default
    u2, _ := url.Parse("https://example.com/search?limit=nope")
    qp2 := params.NewQueryParamsFromUrl(*u2)
    limit2, usedI2 := qp2.GetAsInt("limit", 5)
    fmt.Printf("invalid int -> default: %d (defaultUsed=%v)\n", limit2, usedI2) // => 5, true
}

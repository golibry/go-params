package examples

import (
    "fmt"
    "os"
    "time"
    "github.com/golibry/go-params/params"
)

// EnvExamples demonstrates reading environment variables as typed values using
// the params helpers. Each call returns (value, defaultUsed). When the env var
// is missing, empty, or invalid for the requested type, the default is returned
// and defaultUsed is true.
func EnvExamples() {
    // Prepare environment variables for demonstration
    _ = os.Setenv("APP_NAME", "  MyApp  ")
    _ = os.Setenv("MAX_CONN", "100")
    _ = os.Setenv("DEBUG", "1")        // truthy
    _ = os.Setenv("THRESHOLD", "2.75")
    _ = os.Setenv("TIMEOUT", "45s")
    _ = os.Unsetenv("MISSING_VAR")

    // String
    name, used := params.GetEnvAsString("APP_NAME", "default-app")
    fmt.Printf("env string: %q (defaultUsed=%v)\n", name, used) // => "MyApp", false

    missing, usedMissing := params.GetEnvAsString("MISSING_VAR", "fallback")
    fmt.Printf("env string missing: %q (defaultUsed=%v)\n", missing, usedMissing) // => "fallback", true

    // Int
    maxConn, usedI := params.GetEnvAsInt("MAX_CONN", 10)
    fmt.Printf("env int: %d (defaultUsed=%v)\n", maxConn, usedI) // => 100, false

    _ = os.Setenv("BAD_INT", "oops")
    badInt, usedBadInt := params.GetEnvAsInt("BAD_INT", 7)
    fmt.Printf("env int invalid: %d (defaultUsed=%v)\n", badInt, usedBadInt) // => 7, true

    // Bool
    debug, usedB := params.GetEnvAsBool("DEBUG", false)
    fmt.Printf("env bool: %t (defaultUsed=%v)\n", debug, usedB) // => true, false

    _ = os.Setenv("BAD_BOOL", "maybe")
    badBool, usedBadBool := params.GetEnvAsBool("BAD_BOOL", true)
    fmt.Printf("env bool invalid: %t (defaultUsed=%v)\n", badBool, usedBadBool) // => true, true

    // Float
    thr, usedF := params.GetEnvAsFloat("THRESHOLD", 0.5)
    fmt.Printf("env float: %f (defaultUsed=%v)\n", thr, usedF) // => 2.750000, false

    _ = os.Setenv("BAD_FLOAT", "NaN?")
    badFloat, usedBadFloat := params.GetEnvAsFloat("BAD_FLOAT", 1.25)
    fmt.Printf("env float invalid: %f (defaultUsed=%v)\n", badFloat, usedBadFloat) // => 1.250000, true

    // Duration (Go duration format)
    timeout, usedD := params.GetEnvAsDuration("TIMEOUT", time.Minute)
    fmt.Printf("env duration: %s (defaultUsed=%v)\n", timeout, usedD) // => 45s, false

    _ = os.Setenv("BAD_DURATION", "soon")
    badDuration, usedBadDuration := params.GetEnvAsDuration("BAD_DURATION", 2*time.Minute)
    fmt.Printf("env duration invalid: %s (defaultUsed=%v)\n", badDuration, usedBadDuration) // => 2m0s, true
}

#### Circuit Breaker Component

```golang
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"

    "github.com/clivern/drifter/core/component"
)

var cb *component.CircuitBreaker

func init() {
    var st component.Settings

    st.Name = "cb"

    st.ReadyToTrip = func(counts component.Counts) bool {
        failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
        return counts.Requests >= 20 && failureRatio >= 0.6
    }

    cb = component.NewCircuitBreaker(st)
}

func main() {
    for i := 0; i < 1000; i++ {
        Get("https://httpbi.org/status/500", i)
        time.Sleep(1 * time.Second)
    }
}

func Get(url string, count int) ([]byte, error) {
    body, err := cb.Execute(func() (interface{}, error) {
        if count > 30 {
            fmt.Println("URL is right\n")
            url = "https://httpbin.org/status/200"
        } else {
            fmt.Println("URL is wrong\n")
        }

        resp, err := http.Get(url)

        time.Sleep(1 * time.Second)

        if err != nil {
            return nil, err
        }

        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)

        if err != nil {
            return nil, err
        }

        return body, nil
    })

    if err != nil {
        return nil, err
    }

    return body.([]byte), nil
}
```

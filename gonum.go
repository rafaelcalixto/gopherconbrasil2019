package main

import (
    "net/http"
    "fmt"
    "log"
    "context"
    "time"
    "io/ioutil"
    "strings"
)

var (
    url        string
    kv_str     string
    key_values []string
    sc_schools map[string]string
    key        string
    value      string
)
func main() {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 30)
    defer cancel()

    url = "http://educacao.dadosabertosbr.com/api/cidades/SC"
    ans, err := http.NewRequest("GET", url, nil)
    if err != nil { log.Fatal("Request fail: %v", err) }

    ans = ans.WithContext(ctx)
    resp, _ := http.DefaultClient.Do(ans)
    defer resp.Body.Close()

    bytes, err := ioutil.ReadAll(resp.Body)
    sc_schools = make(map[string]string)

    for _, kv_str := range strings.Split(string(bytes), ",") {
        key_values = strings.Split(kv_str, ":")
        key = strings.Replace(key_values[0], `"`, "", -1)
        key = strings.Replace(key, "[", "", -1)
        value = strings.Replace(key_values[1], `"`, "", -1)
        value = strings.Replace(value, "[", "", -1)

        sc_schools[key] = value
    }

    for _, v := range sc_schools {
        fmt.Println(v)
    }

}

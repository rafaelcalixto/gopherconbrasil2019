package gopherconbr2019

import (
    "context"
    "time"
    "net/http"
    "log"
    "io/ioutil"
)

func GetData(url string) (bytes []byte) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 30)
    defer cancel()

    ans, err := http.NewRequest("GET", url, nil)
    if err != nil { log.Fatal("Request fail: %v", err) }

    ans = ans.WithContext(ctx)
    resp, _ := http.DefaultClient.Do(ans)
    defer resp.Body.Close()

    bytes, err = ioutil.ReadAll(resp.Body)
    if err != nil { log.Fatal("Request fail: %v", err) }

    return
}

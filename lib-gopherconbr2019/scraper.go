package gopherconbr2019

import (
    "context"
    "time"
    "net/http"
    "log"
    "io/ioutil"
    "sync"
    "fmt"
)

func GetData(wg *sync.WaitGroup, c chan []byte, url string) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 30)
    defer cancel()
    defer wg.Done()

    ans, err := http.NewRequest("GET", url, nil)
    if err != nil { log.Fatal("Request fail: %v", err) }

    ans = ans.WithContext(ctx)
    resp, err := http.DefaultClient.Do(ans)
    if err != nil {
        fmt.Println("Falha ao tentar obter dados da URL:", url)
        return
    }
    defer resp.Body.Close()

    bytes, err = ioutil.ReadAll(resp.Body)
    if err != nil { log.Fatal("Request fail: %v", err) }

    c <- bytes
}

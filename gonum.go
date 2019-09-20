package main

import (
    gcbr "lib-gopherconbr2019"
    "fmt"
    "sync"
)

var (
    mun_url    string
    stats_url  string
    info_type  string
    sc_schools map[string]string
    stats_sc   map[string][]float64
    infoLab    float64
    bytesChan  chan []byte
    api_ans    []byte
    again      bool
    wg         sync.WaitGroup
)
func main() {
    mun_url = "http://educacao.dadosabertosbr.com/api/cidades/SC"
    stats_url = "http://educacao.dadosabertosbr.com/api/estatisticas?codMunicipio="

    bytesChan := make(chan []byte, 300)
    defer close(bytesChan)
    wg.Add(1)
    go gcbr.GetData(&wg, bytesChan, mun_url)
    wg.Wait()
    api_ans = <- bytesChan
    sc_schools = gcbr.GetMunicipios(api_ans)

    stats_sc = make(map[string][]float64)
    info_type = "laboratorioInformatica"
    fmt.Print("loading data...")
    for k, _ := range sc_schools {
        wg.Add(1)
        go gcbr.GetData(&wg, bytesChan, stats_url + k)
    }
    wg.Wait()
    fmt.Print("calculating...")
    var count int
    count = 1

    for {
        api_ans, again = <- bytesChan
        fmt.Println(count)
        count++
        if again {
        stats_sc[info_type] = append(stats_sc[info_type],
                                     gcbr.GetStats(api_ans, info_type))
        } else { break }
    }

    fmt.Println(stats_sc[info_type])
}

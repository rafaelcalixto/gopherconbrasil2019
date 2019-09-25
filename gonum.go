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
    stats_sc   map[string]float64
    infoLab    float64
    bytesChan  chan []byte
    api_ans    []byte
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

    stats_sc = make(map[string]float64)
    fmt.Print("loading data...")
    for k, _ := range sc_schools {
        wg.Add(1)
        go gcbr.GetData(&wg, bytesChan, stats_url + k)
    }
    wg.Wait()
    fmt.Println("calculating...")

    for {
        api_ans = <- bytesChan
        st, ref := gcbr.GetStats(api_ans, "laboratorioInformatica", "nomeLocal")
        stats_sc[ref] = st

        if len(bytesChan) == 0 { break }
    }
    gcbr.API(stats_sc)
}

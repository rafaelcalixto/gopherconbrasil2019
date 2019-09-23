package main

import (
    gcbr "lib-gopherconbr2019"
    "fmt"
    "sync"
    "gonum.org/v1/gonum/stat"
    "gonum.org/v1/gonum/plot"
    "math"
    "sort"
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
    mean       float64
    median       float64
    variance   float64
    stddev     float64
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
    fmt.Println("calculating...")

    for {
        api_ans = <- bytesChan
        stats_sc[info_type] = append(stats_sc[info_type],
                                     gcbr.GetStats(api_ans, info_type))
        if len(bytesChan) == 0 { break }
    }

    sort.Float64s(stats_sc[info_type])
    mean = stat.Mean(stats_sc[info_type], nil)
    median = stat.Quantile(0.5, stat.Empirical, stats_sc[info_type], nil)
    variance = stat.Variance(stats_sc[info_type], nil)
    stddev = math.Sqrt(variance)

    fmt.Println("A média do ", info_type, "é: ", mean)
    fmt.Println("A mediana do ", info_type, "é: ", median)
    fmt.Println("A variança do ", info_type, "é: ", variance)
    fmt.Println("A desvio padrão do ", info_type, "é: ", stddev)

    goplot, err_plot := plot.New()
    if err != nil { log.Fatal("Error while plotting", err_plot) }


}

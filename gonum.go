package main

import (
    gcbr "gopherconbr2019"
    "fmt"
)

var (
    mun_url    string
    stats_url  string
    sc_schools map[string]string
    stats_sc   map[string]string
    infoLab    float64
)
func main() {
    mun_url = "http://educacao.dadosabertosbr.com/api/cidades/SC"
    stats_url = "http://educacao.dadosabertosbr.com/api/estatisticas?codMunicipio="

    sc_schools = gcbr.GetMunicipios(gcbr.GetData(mun_url))

    for k, _ := range sc_schools {
        stats_sc = gcbr.GetStats(gcbr.GetData(stats_url + k), "laboratorioInformatica")
        fmt.Println(infoLab)
    }
}

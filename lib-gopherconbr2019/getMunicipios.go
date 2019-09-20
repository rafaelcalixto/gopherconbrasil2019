package gopherconbr2019

import (
    "strings"
    "regexp"
    "log"
)

var (
    key_values []string
    sc_schools map[string]string
    stats_sc   map[string]string
    key        string
    value      string
    infoLab    float64
    err        error
    bytes      []byte
)

func GetMunicipios(bytes []byte) (sc_schools map[string]string) {

    sc_schools = make(map[string]string)
    for _, kv_str := range strings.Split(string(bytes), ",") {
        key_values = strings.Split(kv_str, ":")
        reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
        if err != nil { log.Fatal("Regex fail: %v", err) }

        key = reg.ReplaceAllString(key_values[0], "")
        value = reg.ReplaceAllString(key_values[1], "")

        sc_schools[key] = value
    }
    return

}

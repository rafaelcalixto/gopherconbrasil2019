package gopherconbr2019

import (
    "strings"
    "regexp"
    "log"
    "strconv"
)

func GetStats(bytes []byte, getIt, ref string) (float64, string) {

    stats_sc = make(map[string]string)

   for _, kv_str := range strings.Split(string(bytes), ",") {
        key_values = strings.Split(kv_str, ":")
        reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
        if err != nil { log.Fatal("Regex fail: %v", err) }

        key = reg.ReplaceAllString(key_values[0], "")
        value = key_values[1]
        stats_sc[key] = value

    }
    infoLab, err = strconv.ParseFloat(stats_sc[getIt], 64)
    if err != nil { log.Fatal("Regex fail: %v", err) }

    return infoLab, stats_sc[ref]
}

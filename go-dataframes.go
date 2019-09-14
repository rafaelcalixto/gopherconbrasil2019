package main

import (
    "fmt"
    "github.com/go-gota/gota/dataframe"
)

func main() {
    df := dataframe.LoadMaps(
        []map[string]interface{}{
            map[string]interface{}{
                "A" : "a",
                "B" : 1,
                "C" : true,
                "D" : 0,
            },
            map[string]interface{}{
                "A" : "b",
                "B" : 2,
                "C" : true,
                "D" : 0.5,
            },
        },
    )

    fmt.Println(df)
}

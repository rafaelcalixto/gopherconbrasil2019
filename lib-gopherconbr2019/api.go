package gopherconbr2019

import (
    "fmt"
    "net/http"
    "sync"
    "math"
    "sort"

    "gonum.org/v1/plot"
    "gonum.org/v1/gonum/stat"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/vg"
)

type server struct {
    to_labels []string
    to_stats  []float64
    sync.RWMutex
}

var (
    stats  map[string]float64
    msg    string
    s      server
)

// This is a simple package used just to call the API
func API(d map[string]float64) {
    for k, v := range d {
        s.to_stats = append(s.to_stats, v)
        s.to_labels = append(s.to_labels, k)
    }
    fmt.Println("starting the API...")
    http.HandleFunc("/", s.index)
    http.HandleFunc("/statz", s.statz)
    http.HandleFunc("/histogram", errorHandler(s.hist))
    http.ListenAndServe(":8000", nil)
}

func datascience(to_analyze []float64) {
    sort.Float64s(to_analyze)
    stats = make(map[string]float64)
    stats["média"] = stat.Mean(to_analyze, nil) * 100
    stats["mediana"] = stat.Quantile(0.5, stat.Empirical, to_analyze, nil) * 100
    stats["variança"] = stat.Variance(to_analyze, nil) * 100
    stats["desvio padrão"] = math.Sqrt(stats["variança"]) * 100
}

// This function returns for the Browsers some informations about the API
// This is mandary for some Browsers allows the access to the API
func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func errorHandler(h func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        err := h(w, r)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}

// This function returns a "Welcome message" to the API
func (s *server) index(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    fmt.Fprintf(w, "This is a simple page for GopherCon Brasil 2019")
}

func (s *server) statz(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    s.RLock()
    defer s.RUnlock()

    datascience(s.to_stats)
    msg = ""

    for k, v := range stats {
        msg = fmt.Sprintf("%sA %s de laboratórios de informática é: %.2f%.\n\n",
             msg, k, v)
    }
    fmt.Fprintf(w, msg)
}

func (s *server) hist(w http.ResponseWriter, r *http.Request) error {
    enableCors(&w)
    s.RLock()
    defer s.RUnlock()

    var values plotter.Values
    values = make(plotter.Values, len(s.to_stats))
    for i, v := range s.to_stats {
        values[i] = v
    }

    p, err := plot.New()
    if err != nil { return err }

    hist, err := plotter.NewBarChart(values, vg.Points(20))
    if err != nil { return err }

    p.Add(hist)
    p.Title.Text = "Histograma da estatística escolar de SC"
    p.Y.Label.Text = "Lab de Info na Escola"
    p.X.Label.Text = "cidades"
    for cid := range s.to_labels {
        // p.Legend.Add(cid, hist)
        fmt.Println(cid)
    }

    wt, err := p.WriterTo(512, 512, "png")
    if err != nil { return err }

    _, err = wt.WriteTo(w)
    return err

}

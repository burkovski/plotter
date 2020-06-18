package main

import (
    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/plotutil"
    "gonum.org/v1/plot/vg"
)

type Row struct {
    Name               string
    TotalClicks        float64
    TotalPaidClicks    float64
    TotalProductClicks float64
}

func toPercents(x, y float64) float64 {
    z := x - y
    return (z / y) * 100
}

func main() {

    data := []Row{
        {"Week 13", 5854147, 4421708, 2394563},
        {"Week 14", 3665790, 2324214, 1331593},
        {"Week 15", 4299656, 2677784, 1898413},
        {"Week 16", 3238947, 1810889, 1265588},
        {"Week 17", 3405983, 2008195, 1503263},
        {"Week 18", 2907680, 1842166, 1454557},
    }

    zaebumba := data[len(data)-1]
    var groupA plotter.Values
    var groupB plotter.Values
    var groupC plotter.Values

    for _, row := range data[0 : len(data)-1] {
        groupA = append(groupA, toPercents(row.TotalClicks, zaebumba.TotalClicks))
        groupB = append(groupB, toPercents(row.TotalPaidClicks, zaebumba.TotalPaidClicks))
        groupC = append(groupC, toPercents(row.TotalProductClicks, zaebumba.TotalProductClicks))
    }

    groupA = append(groupA, 0)
    groupB = append(groupB, 0)
    groupC = append(groupC, 0)

    groups := []plotter.Values{groupA, groupB, groupC}

    p, err := plot.New()
    if err != nil {
        panic(err)
    }
    p.Title.Text = "Bar chart"
    p.Y.Label.Text = "Heights"

    w := vg.Points(18)

    for i, group := range groups {
        bar, err := plotter.NewBarChart(group, w)
        if err != nil {
            panic(err)
        }
        bar.LineStyle.Width = vg.Length(0)
        bar.Color = plotutil.Color(i)
        bar.Offset = vg.Length(i-1) * w

        p.Add(bar)
    }

    p.Legend.Top = true
    p.NominalX("One", "Two", "Three", "Four", "Five")

    if err := p.Save(5*vg.Inch, 3*vg.Inch, "barchart.png"); err != nil {
        panic(err)
    }
}

package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"reflect"
)

type Row struct {
	Name               string
	TotalClicks        float64
	TotalPaidClicks    float64
	TotalProductClicks float64
}

type Rows []Row
type dataFrame []float64

func (rows Rows) totalClicks() dataFrame {
	var c dataFrame
	for _, row := range rows {
		c = append(c, row.TotalClicks)
	}
	return c
}

func (rows Rows) totalPaidClicks() dataFrame {
	var c dataFrame
	for _, row := range rows {
		c = append(c, row.TotalPaidClicks)
	}
	return c
}

func (rows Rows) totalProductClicks() dataFrame {
	var c dataFrame
	for _, row := range rows {
		c = append(c, row.TotalProductClicks)
	}
	return c
}

func (rows Rows) toPercents() Rows {
	methods := map[string]func() dataFrame{
		"TotalClicks":        rows.totalClicks,
		"TotalPaidClicks":    rows.totalPaidClicks,
		"TotalProductClicks": rows.totalProductClicks,
	}

	for field, method := range methods {
		clicks := method()
		last := len(clicks) - 1
		clicks = clicks.toPercents(clicks[last])

		for i, clickValue := range clicks {
			v := reflect.ValueOf(&rows[i])
			f := reflect.Indirect(v).FieldByName(field)
			f.SetFloat(clickValue)
		}
	}

	return rows
}

func (frame dataFrame) toPercents(v float64) dataFrame {
	for i := 0; i < len(frame); i++ {
		diff := frame[i] - v
		frame[i] = diff / frame[i] * 100
	}
	return frame
}

func main() {
	data := Rows{
		{"Week 13", 5854147, 4421708, 2394563},
		{"Week 14", 3665790, 2324214, 1331593},
		{"Week 15", 4299656, 2677784, 1898413},
		{"Week 16", 3238947, 1810889, 1265588},
		{"Week 17", 3405983, 2008195, 1503263},
		{"Week 18", 2907680, 1842166, 1454557},
	}

	dataPercents := data.toPercents()

	var groupA plotter.Values
	var groupB plotter.Values
	var groupC plotter.Values

	for _, row := range dataPercents {
		groupA = append(groupA, row.TotalClicks)
		groupB = append(groupB, row.TotalProductClicks)
		groupC = append(groupC, row.TotalPaidClicks)
	}

	groups := []plotter.Values{groupA, groupB, groupC}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Bar chart"
	p.Y.Label.Text = "Heights"

	w := vg.Points(15)

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

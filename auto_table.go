package glog

import "strings"

const (
	AUTO_PAD_LEFT = iota
	AUTO_PAD_CENTER
	AUTO_PAD_RIGHT
)

type AutoTableSeries struct {
	Name    string
	Values  []string
	maxLen  int
	padDir  int
	padChar rune
}

func (t *AutoTableSeries) Push(value ...interface{}) *AutoTableSeries {
	for _, v := range value {
		vs := Auto(v)
		vl := len(StripANSI(vs))
		t.Values = append(t.Values, vs)
		t.maxLen = Max(t.maxLen, vl)
	}
	return t
}

func (t *AutoTableSeries) PadValues() *AutoTableSeries {
	for i := range t.Values {
		vs := t.Values[i]
		if t.padDir == AUTO_PAD_LEFT {
			vs = PadLeft(vs, t.maxLen, t.padChar)
		} else if t.padDir == AUTO_PAD_CENTER {
			vs = PadCenter(vs, t.maxLen, t.padChar)
		} else if t.padDir == AUTO_PAD_RIGHT {
			vs = PadRight(vs, t.maxLen, t.padChar)
		}
		t.Values[i] = vs
	}
	return t
}

func (t *AutoTableSeries) Header() string {
	h := Auto(t.Name)
	if t.padDir == AUTO_PAD_LEFT {
		h = PadLeft(h, t.maxLen, t.padChar)
	} else if t.padDir == AUTO_PAD_CENTER {
		h = PadCenter(h, t.maxLen, t.padChar)
	} else if t.padDir == AUTO_PAD_RIGHT {
		h = PadRight(h, t.maxLen, t.padChar)
	}
	return h
}

func NewAutoTableSeries(name string, padDirection int, padChar rune) *AutoTableSeries {
	return &AutoTableSeries{
		Name:    name,
		Values:  []string{},
		maxLen:  len(name),
		padDir:  padDirection,
		padChar: padChar,
	}
}

func NewAutoTableSeriesLeft(name string, padChar rune) *AutoTableSeries {
	return NewAutoTableSeries(name, AUTO_PAD_LEFT, padChar)
}

func NewAutoTableSeriesRight(name string, padChar rune) *AutoTableSeries {
	return NewAutoTableSeries(name, AUTO_PAD_RIGHT, padChar)
}

func NewAutoTableSeriesCenter(name string, padChar rune) *AutoTableSeries {
	return NewAutoTableSeries(name, AUTO_PAD_CENTER, padChar)
}

type AutoTable struct {
	series []*AutoTableSeries
}

func (at *AutoTable) TableRows() []string {
	ls := len(at.series)
	topRow := make([]string, ls)
	headers := make([]string, ls)
	rows := [][]string{} // rows - columns
	for col, series := range at.series {
		headers[col] = series.Header()
		topRow[col] = strings.Repeat("─", series.maxLen)
		for row, v := range series.PadValues().Values {
			if len(rows) < row+1 {
				rows = append(rows, make([]string, ls))
			}
			rows[row][col] = v
		}
	}
	res := []string{
		"┌─" + strings.Join(topRow, "─┬─") + "─┐",
		"│ " + strings.Join(headers, " │ ") + " │",
		"├─" + strings.Join(topRow, "─┼─") + "─┤",
	}
	for _, r := range rows {
		for i, _ := range r {
			if r[i] == "" {
				if at.series[i].padDir == AUTO_PAD_LEFT {
					r[i] = PadLeft(Auto("N/A"), at.series[i].maxLen, at.series[i].padChar)

				} else if at.series[i].padDir == AUTO_PAD_CENTER {
					r[i] = PadCenter(Auto("N/A"), at.series[i].maxLen, at.series[i].padChar)

				} else if at.series[i].padDir == AUTO_PAD_RIGHT {
					r[i] = PadRight(Auto("N/A"), at.series[i].maxLen, at.series[i].padChar)

				}
			}
		}
		res = append(res, "│ "+strings.Join(r, " │ ")+" │")
	}
	res = append(res, "└─"+strings.Join(topRow, "─┴─")+"─┘")
	return res
}

func NewAutoTable(series ...*AutoTableSeries) *AutoTable {
	return &AutoTable{
		series: series,
	}
}

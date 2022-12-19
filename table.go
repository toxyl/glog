package glog

import (
	"strings"
)

const (
	PAD_LEFT = iota
	PAD_CENTER
	PAD_RIGHT
)

type TableColumn struct {
	Name        string
	values      []string
	maxLen      int
	padDir      int
	padChar     rune
	fnHighlight func(a ...interface{}) string
}

func (t *TableColumn) Push(value ...interface{}) *TableColumn {
	for _, v := range value {
		vs := t.fnHighlight(v)
		vl := len(StripANSI(vs))
		t.values = append(t.values, vs)
		t.maxLen = Max(t.maxLen, vl)
	}
	return t
}

func (t *TableColumn) padValues() *TableColumn {
	for i := range t.values {
		vs := t.values[i]
		if t.padDir == PAD_LEFT {
			vs = PadLeft(vs, t.maxLen, t.padChar)
		} else if t.padDir == PAD_CENTER {
			vs = PadCenter(vs, t.maxLen, t.padChar)
		} else if t.padDir == PAD_RIGHT {
			vs = PadRight(vs, t.maxLen, t.padChar)
		}
		t.values[i] = vs
	}
	return t
}

func (t *TableColumn) header() string {
	h := Auto(t.Name)
	if t.padDir == PAD_LEFT {
		h = PadLeft(h, t.maxLen, t.padChar)
	} else if t.padDir == PAD_CENTER {
		h = PadCenter(h, t.maxLen, t.padChar)
	} else if t.padDir == PAD_RIGHT {
		h = PadRight(h, t.maxLen, t.padChar)
	}
	return h
}

func NewTableColumnCustom(name string, padDirection int, padChar rune, highlighter func(a ...interface{}) string) *TableColumn {
	h := highlighter
	if highlighter == nil {
		h = Auto
	}
	return &TableColumn{
		Name:        name,
		values:      []string{},
		maxLen:      len(name),
		padDir:      padDirection,
		padChar:     padChar,
		fnHighlight: h,
	}
}

func NewTableColumn(name string, padDirection int) *TableColumn {
	return NewTableColumnCustom(name, padDirection, LoggerConfig.TablePadChar, nil)
}

func NewTableColumnLeftCustom(name string, padChar rune, highlighter func(a ...interface{}) string) *TableColumn {
	return NewTableColumnCustom(name, PAD_RIGHT, padChar, highlighter)
}

func NewTableColumnLeft(name string) *TableColumn {
	return NewTableColumn(name, PAD_RIGHT)
}

func NewTableColumnRightCustom(name string, padChar rune, highlighter func(a ...interface{}) string) *TableColumn {
	return NewTableColumnCustom(name, PAD_LEFT, padChar, highlighter)
}

func NewTableColumnRight(name string) *TableColumn {
	return NewTableColumn(name, PAD_LEFT)
}

func NewTableColumnCenterCustom(name string, padChar rune, highlighter func(a ...interface{}) string) *TableColumn {
	return NewTableColumnCustom(name, PAD_CENTER, padChar, highlighter)
}

func NewTableColumnCenter(name string) *TableColumn {
	return NewTableColumn(name, PAD_CENTER)
}

type Table struct {
	series []*TableColumn
}

func (at *Table) Rows() []string {
	ls := len(at.series)
	topRow := make([]string, ls)
	headers := make([]string, ls)
	rows := [][]string{} // rows - columns
	for col, series := range at.series {
		headers[col] = series.header()
		topRow[col] = strings.Repeat("─", series.maxLen)
		for row, v := range series.padValues().values {
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
				if at.series[i].padDir == PAD_LEFT {
					r[i] = PadLeft(Auto("N/A"), at.series[i].maxLen, at.series[i].padChar)

				} else if at.series[i].padDir == PAD_CENTER {
					r[i] = PadCenter(Auto("N/A"), at.series[i].maxLen, at.series[i].padChar)

				} else if at.series[i].padDir == PAD_RIGHT {
					r[i] = PadRight(Auto("N/A"), at.series[i].maxLen, at.series[i].padChar)

				}
			}
		}
		res = append(res, "│ "+strings.Join(r, " │ ")+" │")
	}
	res = append(res, "└─"+strings.Join(topRow, "─┴─")+"─┘")
	return res
}

func NewTable(series ...*TableColumn) *Table {
	return &Table{
		series: series,
	}
}

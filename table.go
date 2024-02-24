package glog

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	PAD_LEFT = iota
	PAD_CENTER
	PAD_RIGHT
)

type TableColumn struct {
	Name        string
	values      []string
	valuesRaw   []interface{}
	maxLen      int
	padDir      int
	padChar     rune
	fnHighlight func(a ...interface{}) string
}

func (t *TableColumn) Reset() {
	t.valuesRaw = []interface{}{}
	t.values = []string{}
	t.maxLen = len(t.Name)
}

func (t *TableColumn) Values() []interface{} {
	return t.valuesRaw
}

func (t *TableColumn) Push(value ...interface{}) *TableColumn {
	for _, v := range value {
		vs := t.fnHighlight(v)
		vl := len(StripANSI(vs))
		t.values = append(t.values, vs)
		t.valuesRaw = append(t.valuesRaw, v)
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

// Related config setting(s):
//
//   - `LoggerConfig.TablePadChar`
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

func (at *Table) RawData() [][]interface{} {
	// columns might have different row amounts,
	// let's find the maximum number of rows in a column
	numRows := 0
	for _, col := range at.series {
		numRows = Max(numRows, len(col.valuesRaw))
	}

	// now we pad every column with
	// less rows than the max with nil
	for _, col := range at.series {
		for len(col.valuesRaw) < numRows {
			col.Push(nil)
		}
	}

	headers := []interface{}{}
	for _, col := range at.series {
		headers = append(headers, col.Name)
	}

	// generate the output rows
	rows := [][]interface{}{}
	// add headers first
	rows = append(rows, headers)
	// then data
	for i := 0; i < numRows; i++ {
		row := []interface{}{}
		onlySeparators := true
		for _, col := range at.series {
			v := col.valuesRaw[i]
			if str, ok := v.(string); ok {
				v = strings.TrimSpace(StripANSI(str))
				if v == "---" {
					v = "" // strip separators as they are only used for visual representation
				} else if v != "" {
					onlySeparators = false
				}
			}
			row = append(row, v)
		}
		if !onlySeparators {
			rows = append(rows, row)
		}
	}

	return rows
}

func (at *Table) CSV(separator rune) string {
	rows := []string{}

	for _, cols := range at.RawData() {
		row := []string{}
		for _, col := range cols {
			row = append(row, fmt.Sprint(col))
		}
		rows = append(rows, strings.Join(row, string(separator)))
	}

	return strings.Join(rows, "\n")
}

func (at *Table) YAML() (string, error) {
	data, err := yaml.Marshal(at.RawData())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (at *Table) JSON() (string, error) {
	data, err := json.Marshal(at.RawData())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (at *Table) Print(logger *Logger) {
	rows := at.Rows()
	for _, line := range rows {
		logger.Blank("%s", line)
	}
}

func (at *Table) PrintWithoutHeader(logger *Logger) {
	rows := at.Rows()
	rows = append(rows[0:1], rows[3:]...)
	for _, line := range rows {
		logger.Blank("%s", line)
	}
}

func (at *Table) Rows() []string {
	const (
		NO_VAL    = ""
		CSEP      = "---"
		HCONN     = "─"
		HCONN_B   = "┬"
		HCONN_T   = "┴"
		HCONN_C   = "┼"
		VCONN     = "│"
		VCONN_R   = "├"
		VCONN_L   = "┤"
		VCONN_C   = "┼"
		CONN_TL   = "┌"
		CONN_TR   = "┐"
		CONN_BL   = "└"
		CONN_BR   = "┘"
		SPACE     = " "
		CTYPE_SEP = iota
		CTYPE_VAL
		CTYPE_NO_VAL
	)
	ls := len(at.series)
	topRow := make([]string, ls)
	headers := make([]string, ls)
	rows := [][]string{} // rows - columns
	for col, series := range at.series {
		headers[col] = series.header()
		topRow[col] = strings.Repeat(HCONN, series.maxLen)
		for row, v := range series.padValues().values {
			if len(rows) < row+1 {
				rows = append(rows, make([]string, ls))
			}
			rows[row][col] = v
		}
	}
	fnColType := func(col, cutset string) int {
		switch strings.Trim(StripANSI(col), cutset) {
		case NO_VAL:
			return CTYPE_NO_VAL
		case CSEP:
			return CTYPE_SEP
		}
		return CTYPE_VAL
	}

	res := []string{
		CONN_TL + HCONN + strings.Join(topRow, HCONN+HCONN_B+HCONN) + HCONN + CONN_TR,
		VCONN + SPACE + strings.Join(headers, SPACE+VCONN+SPACE) + SPACE + VCONN,
		VCONN_R + HCONN + strings.Join(topRow, HCONN+HCONN_C+HCONN) + HCONN + VCONN_L,
	}
	for _, row := range rows {
		colTypes := make([]int, len(row))
		for colIdx, col := range row {
			colTypes[colIdx] = fnColType(col, string(at.series[colIdx].padChar))
			switch colTypes[colIdx] {
			case CTYPE_NO_VAL:
				if len(StripANSI(col)) == 0 {
					if at.series[colIdx].padDir == PAD_LEFT {
						row[colIdx] = PadLeft(Auto("N/A"), at.series[colIdx].maxLen, at.series[colIdx].padChar)
					} else if at.series[colIdx].padDir == PAD_CENTER {
						row[colIdx] = PadCenter(Auto("N/A"), at.series[colIdx].maxLen, at.series[colIdx].padChar)
					} else if at.series[colIdx].padDir == PAD_RIGHT {
						row[colIdx] = PadRight(Auto("N/A"), at.series[colIdx].maxLen, at.series[colIdx].padChar)
					}
				}
			case CTYPE_SEP:
				row[colIdx] = PadLeft("", at.series[colIdx].maxLen, '─')
			}
		}

		rowStr := ""
		for colIdx, col := range row {
			currType := colTypes[colIdx]

			// start of row
			if colIdx == 0 {
				switch currType {
				case CTYPE_SEP:
					rowStr = VCONN_R + HCONN + col + HCONN
				default:
					rowStr = VCONN + SPACE + col + SPACE
				}
				continue
			}

			prevType := colTypes[colIdx-1]

			// end of row
			if colIdx == len(row)-1 {
				if prevType == CTYPE_SEP {
					if currType == CTYPE_SEP {
						rowStr += VCONN_C + HCONN + col + HCONN + VCONN_L
					} else {
						rowStr += VCONN_L + SPACE + col + SPACE + VCONN
					}
					continue
				}

				if currType == CTYPE_SEP {
					rowStr += VCONN_R + HCONN + col + HCONN + VCONN_L
				} else {
					rowStr += VCONN + SPACE + col + SPACE + VCONN
				}
				continue
			}

			// somewhere in the middle
			if prevType == CTYPE_SEP {
				if currType == CTYPE_SEP {
					rowStr += VCONN_C + HCONN + col + HCONN
				} else {
					rowStr += VCONN_L + SPACE + col + SPACE
				}
				continue
			}

			if currType == CTYPE_SEP {
				rowStr += VCONN_R + HCONN + col + HCONN
			} else {
				rowStr += VCONN + SPACE + col + SPACE
			}
		}

		res = append(res, rowStr)
	}
	res = append(res, CONN_BL+HCONN+strings.Join(topRow, HCONN+HCONN_T+HCONN)+HCONN+CONN_BR)
	return res
}

func NewTable(series ...*TableColumn) *Table {
	return &Table{
		series: series,
	}
}

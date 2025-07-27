package logger

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/toxyl/glog/colorizers"
	"github.com/toxyl/glog/config"
	"github.com/toxyl/glog/utils"
	"github.com/toxyl/math"
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
	valuesRaw   []any
	maxLen      int
	padDir      int
	padChar     rune
	fnHighlight func(a ...any) string
}

func (t *TableColumn) Reset() {
	t.valuesRaw = []any{}
	t.values = []string{}
	t.maxLen = len(t.Name)
}

func (t *TableColumn) Values() []any {
	return t.valuesRaw
}

func (t *TableColumn) Push(value ...any) *TableColumn {
	for _, v := range value {
		vs := t.fnHighlight(v)
		vl := len(utils.StripANSI(vs))
		t.values = append(t.values, vs)
		t.valuesRaw = append(t.valuesRaw, v)
		t.maxLen = math.Max(t.maxLen, vl)
	}
	return t
}

func (t *TableColumn) padValues() *TableColumn {
	for i := range t.values {
		vs := t.values[i]
		switch t.padDir {
		case PAD_LEFT:
			vs = utils.PadLeft(vs, t.maxLen, t.padChar)
		case PAD_CENTER:
			vs = utils.PadCenter(vs, t.maxLen, t.padChar)
		case PAD_RIGHT:
			vs = utils.PadRight(vs, t.maxLen, t.padChar)
		}
		t.values[i] = vs
	}
	return t
}

func (t *TableColumn) header() string {
	h := colorizers.Auto(t.Name)
	switch t.padDir {
	case PAD_LEFT:
		h = utils.PadLeft(h, t.maxLen, t.padChar)
	case PAD_CENTER:
		h = utils.PadCenter(h, t.maxLen, t.padChar)
	case PAD_RIGHT:
		h = utils.PadRight(h, t.maxLen, t.padChar)
	}
	return h
}

func NewTableColumnCustom(name string, padDirection int, padChar rune, highlighter func(a ...any) string) *TableColumn {
	h := highlighter
	if highlighter == nil {
		h = colorizers.Auto
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
	return NewTableColumnCustom(name, padDirection, config.LoggerConfig.TablePadChar, nil)
}

func NewTableColumnLeftCustom(name string, padChar rune, highlighter func(a ...any) string) *TableColumn {
	return NewTableColumnCustom(name, PAD_RIGHT, padChar, highlighter)
}

func NewTableColumnLeft(name string) *TableColumn {
	return NewTableColumn(name, PAD_RIGHT)
}

func NewTableColumnRightCustom(name string, padChar rune, highlighter func(a ...any) string) *TableColumn {
	return NewTableColumnCustom(name, PAD_LEFT, padChar, highlighter)
}

func NewTableColumnRight(name string) *TableColumn {
	return NewTableColumn(name, PAD_LEFT)
}

func NewTableColumnCenterCustom(name string, padChar rune, highlighter func(a ...any) string) *TableColumn {
	return NewTableColumnCustom(name, PAD_CENTER, padChar, highlighter)
}

func NewTableColumnCenter(name string) *TableColumn {
	return NewTableColumn(name, PAD_CENTER)
}

type Table struct {
	series []*TableColumn
}

func (t *Table) RawData() [][]any {
	// columns might have different row amounts,
	// let's find the maximum number of rows in a column
	numRows := 0
	for _, col := range t.series {
		numRows = math.Max(numRows, len(col.valuesRaw))
	}

	// now we pad every column with
	// less rows than the max with nil
	for _, col := range t.series {
		for len(col.valuesRaw) < numRows {
			col.Push(nil)
		}
	}

	headers := []any{}
	for _, col := range t.series {
		headers = append(headers, col.Name)
	}

	// generate the output rows
	rows := [][]any{}
	// add headers first
	rows = append(rows, headers)
	// then data
	for i := 0; i < numRows; i++ {
		row := []any{}
		onlySeparators := true
		for _, col := range t.series {
			v := col.valuesRaw[i]
			if str, ok := v.(string); ok {
				v = strings.TrimSpace(utils.StripANSI(str))
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

func (t *Table) CSV(separator rune) string {
	rows := []string{}

	for _, cols := range t.RawData() {
		row := []string{}
		for _, col := range cols {
			row = append(row, fmt.Sprint(col))
		}
		rows = append(rows, strings.Join(row, string(separator)))
	}

	return strings.Join(rows, "\n")
}

func (t *Table) YAML() (string, error) {
	data, err := yaml.Marshal(t.RawData())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (t *Table) JSON() (string, error) {
	data, err := json.Marshal(t.RawData())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (t *Table) Print(logger *Logger) {
	rows := t.Rows()
	for _, line := range rows {
		logger.Blank("%s", line)
	}
}

func (t *Table) PrintWithoutHeader(logger *Logger) {
	rows := t.Rows()
	rows = append(rows[0:1], rows[3:]...)
	for _, line := range rows {
		logger.Blank("%s", line)
	}
}

func (t *Table) Rows() []string {
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
	ls := len(t.series)
	topRow := make([]string, ls)
	headers := make([]string, ls)
	rows := [][]string{} // rows - columns
	for col, series := range t.series {
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
		switch strings.Trim(utils.StripANSI(col), cutset) {
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
			colTypes[colIdx] = fnColType(col, string(t.series[colIdx].padChar))
			switch colTypes[colIdx] {
			case CTYPE_NO_VAL:
				if len(utils.StripANSI(col)) == 0 {
					switch t.series[colIdx].padDir {
					case PAD_LEFT:
						row[colIdx] = utils.PadLeft(colorizers.Auto("N/A"), t.series[colIdx].maxLen, t.series[colIdx].padChar)
					case PAD_CENTER:
						row[colIdx] = utils.PadCenter(colorizers.Auto("N/A"), t.series[colIdx].maxLen, t.series[colIdx].padChar)
					case PAD_RIGHT:
						row[colIdx] = utils.PadRight(colorizers.Auto("N/A"), t.series[colIdx].maxLen, t.series[colIdx].padChar)
					}
				}
			case CTYPE_SEP:
				row[colIdx] = utils.PadLeft("", t.series[colIdx].maxLen, '─')
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

func NewTable(columns ...*TableColumn) *Table {
	return &Table{
		series: columns,
	}
}

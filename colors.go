package glog

const (
	DarkBlue     = 17
	Blue         = 21
	DarkGreen    = 22
	LightBlue    = 27
	OliveGreen   = 34
	Green        = 46
	Cyan         = 51
	Purple       = 53
	DarkOrange   = 130
	DarkYellow   = 142
	Lime         = 154
	DarkRed      = 160
	Red          = 196
	Pink         = 201
	Orange       = 208
	Yellow       = 220
	BrightYellow = 229
	DarkGray     = 234
	MediumGray   = 240
	Gray         = 250
)

// the Wrap* functions are generated using regexr.com
// input: the list of color constants defined above
// regex: /\s*(.*?)\s+=.*/g
// replacement: func Wrap$1(str string) string { return Wrap(str, $1) }\n

func WrapDarkBlue(str string) string     { return Wrap(str, DarkBlue) }
func WrapBlue(str string) string         { return Wrap(str, Blue) }
func WrapDarkGreen(str string) string    { return Wrap(str, DarkGreen) }
func WrapLightBlue(str string) string    { return Wrap(str, LightBlue) }
func WrapOliveGreen(str string) string   { return Wrap(str, OliveGreen) }
func WrapGreen(str string) string        { return Wrap(str, Green) }
func WrapCyan(str string) string         { return Wrap(str, Cyan) }
func WrapPurple(str string) string       { return Wrap(str, Purple) }
func WrapDarkOrange(str string) string   { return Wrap(str, DarkOrange) }
func WrapDarkYellow(str string) string   { return Wrap(str, DarkYellow) }
func WrapLime(str string) string         { return Wrap(str, Lime) }
func WrapDarkRed(str string) string      { return Wrap(str, DarkRed) }
func WrapRed(str string) string          { return Wrap(str, Red) }
func WrapPink(str string) string         { return Wrap(str, Pink) }
func WrapOrange(str string) string       { return Wrap(str, Orange) }
func WrapYellow(str string) string       { return Wrap(str, Yellow) }
func WrapBrightYellow(str string) string { return Wrap(str, BrightYellow) }
func WrapDarkGray(str string) string     { return Wrap(str, DarkGray) }
func WrapMediumGray(str string) string   { return Wrap(str, MediumGray) }
func WrapGray(str string) string         { return Wrap(str, Gray) }

package glog

var colorMap map[int]int = map[int]int{
	0: 16,
	1: 17,
	2: 18,
	3: 19,
	4: 20,
	5: 21,

	6:  27,
	7:  26,
	8:  25,
	9:  24,
	10: 23,
	11: 22,

	12: 28,
	13: 29,
	14: 30,
	15: 31,
	16: 32,
	17: 33,

	18: 39,
	19: 38,
	20: 37,
	21: 36,
	22: 35,
	23: 34,

	24: 40,
	25: 41,
	26: 42,
	27: 43,
	28: 44,
	29: 45,

	30: 51,
	31: 50,
	32: 49,
	33: 48,
	34: 47,
	35: 46,

	36: 82,
	37: 83,
	38: 84,
	39: 85,
	40: 86,
	41: 87,

	42: 81,
	43: 80,
	44: 79,
	45: 78,
	46: 77,
	47: 76,

	48: 70,
	49: 71,
	50: 72,
	51: 73,
	52: 74,
	53: 75,

	54: 69,
	55: 68,
	56: 67,
	57: 66,
	58: 65,
	59: 64,

	60: 58,
	61: 59,
	62: 60,
	63: 61,
	64: 62,
	65: 63,

	66: 57,
	67: 56,
	68: 55,
	69: 54,
	70: 53,
	71: 52,

	72: 88,
	73: 89,
	74: 90,
	75: 91,
	76: 92,
	77: 93,

	78: 99,
	79: 98,
	80: 97,
	81: 96,
	82: 95,
	83: 94,

	84: 100,
	85: 101,
	86: 102,
	87: 103,
	88: 104,
	89: 105,

	90: 111,
	91: 110,
	92: 109,
	93: 108,
	94: 107,
	95: 106,

	96:  112,
	97:  113,
	98:  114,
	99:  115,
	100: 116,
	101: 117,

	102: 123,
	103: 122,
	104: 121,
	105: 120,
	106: 119,
	107: 118,

	108: 154,
	109: 155,
	110: 156,
	111: 157,
	112: 158,
	113: 159,

	114: 153,
	115: 152,
	116: 151,
	117: 150,
	118: 149,
	119: 148,

	120: 142,
	121: 143,
	122: 144,
	123: 145,
	124: 146,
	125: 147,

	126: 141,
	127: 140,
	128: 139,
	129: 138,
	130: 137,
	131: 136,

	132: 130,
	133: 131,
	134: 132,
	135: 133,
	136: 134,
	137: 135,

	138: 129,
	139: 128,
	140: 127,
	141: 126,
	142: 125,
	143: 124,

	144: 160,
	145: 161,
	146: 162,
	147: 163,
	148: 164,
	149: 165,

	150: 171,
	151: 170,
	152: 169,
	153: 168,
	154: 167,
	155: 166,

	156: 172,
	157: 173,
	158: 174,
	159: 175,
	160: 176,
	161: 177,

	162: 183,
	163: 182,
	164: 181,
	165: 180,
	166: 179,
	167: 178,

	168: 184,
	169: 185,
	170: 186,
	171: 187,
	172: 188,
	173: 189,

	174: 195,
	175: 194,
	176: 193,
	177: 192,
	178: 191,
	179: 190,

	180: 226,
	181: 227,
	182: 228,
	183: 229,
	184: 230,
	185: 231,

	186: 225,
	187: 224,
	188: 223,
	189: 222,
	190: 221,
	191: 220,

	192: 214,
	193: 215,
	194: 216,
	195: 217,
	196: 218,
	197: 219,

	198: 213,
	199: 212,
	200: 211,
	201: 210,
	202: 209,
	203: 208,

	204: 202,
	205: 203,
	206: 204,
	207: 205,
	208: 206,
	209: 207,

	210: 201,
	211: 200,
	212: 199,
	213: 198,
	214: 197,
	215: 196,

	216: 232,
	217: 233,
	218: 234,
	219: 235,
	220: 236,
	221: 237,

	222: 238,
	223: 239,
	224: 240,
	225: 241,
	226: 242,
	227: 243,

	228: 244,
	229: 245,
	230: 246,
	231: 247,
	232: 248,
	233: 249,

	234: 250,
	235: 251,
	236: 252,
	237: 253,
	238: 254,
	239: 255,
}

// MapColor translates `index` from glog's color table to the corresponding ANSI color index.
// Glog uses it's own color table to make smooth (automated) color transitions easier to implement.
func MapColor(index int) int {
	if index <= 16 {
		return index
	}
	return colorMap[index-16]
}

const (
	DarkBlue     = 70
	Blue         = 33
	DarkGreen    = 39
	LightBlue    = 45
	OliveGreen   = 28
	Green        = 52
	Cyan         = 57
	Purple       = 93
	DarkOrange   = 171
	DarkYellow   = 147
	Lime         = 122
	DarkRed      = 160
	Red          = 231
	Pink         = 225
	Orange       = 219
	Yellow       = 196
	BrightYellow = 198
	DarkGray     = 232
	MediumGray   = 238
	Gray         = 250
	White        = 255
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
func WrapWhite(str string) string        { return Wrap(str, White) }

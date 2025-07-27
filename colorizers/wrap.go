package colorizers

import (
	"github.com/toxyl/glog/ansi"
	"github.com/toxyl/glog/colormap"
)

func WrapDarkBlue(str string) string     { return ansi.Wrap(str, colormap.DarkBlue).String() }
func WrapBlue(str string) string         { return ansi.Wrap(str, colormap.Blue).String() }
func WrapDarkGreen(str string) string    { return ansi.Wrap(str, colormap.DarkGreen).String() }
func WrapLightBlue(str string) string    { return ansi.Wrap(str, colormap.LightBlue).String() }
func WrapOliveGreen(str string) string   { return ansi.Wrap(str, colormap.OliveGreen).String() }
func WrapGreen(str string) string        { return ansi.Wrap(str, colormap.Green).String() }
func WrapCyan(str string) string         { return ansi.Wrap(str, colormap.Cyan).String() }
func WrapPurple(str string) string       { return ansi.Wrap(str, colormap.Purple).String() }
func WrapDarkOrange(str string) string   { return ansi.Wrap(str, colormap.DarkOrange).String() }
func WrapDarkYellow(str string) string   { return ansi.Wrap(str, colormap.DarkYellow).String() }
func WrapLime(str string) string         { return ansi.Wrap(str, colormap.Lime).String() }
func WrapDarkRed(str string) string      { return ansi.Wrap(str, colormap.DarkRed).String() }
func WrapRed(str string) string          { return ansi.Wrap(str, colormap.Red).String() }
func WrapPink(str string) string         { return ansi.Wrap(str, colormap.Pink).String() }
func WrapOrange(str string) string       { return ansi.Wrap(str, colormap.Orange).String() }
func WrapYellow(str string) string       { return ansi.Wrap(str, colormap.Yellow).String() }
func WrapBrightYellow(str string) string { return ansi.Wrap(str, colormap.BrightYellow).String() }
func WrapDarkGray(str string) string     { return ansi.Wrap(str, colormap.DarkGray).String() }
func WrapMediumGray(str string) string   { return ansi.Wrap(str, colormap.MediumGray).String() }
func WrapGray(str string) string         { return ansi.Wrap(str, colormap.Gray).String() }
func WrapWhite(str string) string        { return ansi.Wrap(str, colormap.White).String() }

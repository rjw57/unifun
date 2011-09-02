// A selection of fun unicode 'tricks' to apply to text
package uctricks

import (
	"fmt"
	"os"
	"utf8"
)

// A 'font' is a mapping from rune to rune and (optionally) a function to apply to the text before mapping.
type Font struct {
	runeMap map[int]int
	pre     func(string) string
}

var fonts = make(map[string]Font)

// Return a pointer to the font named.
func FontNamed(name string) (*Font, os.Error) {
	f, ok := fonts[name]
	if !ok {
		return nil, os.NewError("Unknown font: " + name)
	}
	return &f, nil
}

// Return a slice containing all font names.
func FontList() []string {
	var l []string
	for k, _ := range fonts {
		l = append(l, k)
	}
	return l
}

// initialise the mappings from vanilla runes -> blackletter
func initBlackLetter() {
	const (
		VANILLA          = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
		BLACKLETTER      = "𝔄𝔅ℭ𝔇𝔈𝔉𝔊ℌℑ𝔍𝔎𝔏𝔐𝔑𝔒𝔓𝔔ℜ𝔖𝔗𝔘𝔙𝔚𝔛𝔜ℨ𝔞𝔟𝔠𝔡𝔢𝔣𝔤𝔥𝔦𝔧𝔨𝔩𝔪𝔫𝔬𝔭𝔮𝔯𝔰𝔱𝔲𝔳𝔴𝔵𝔶𝔷"
		BLACKLETTER_BOLD = "𝕬𝕭𝕮𝕯𝕰𝕱𝕲𝕳𝕴𝕵𝕶𝕷𝕸𝕹𝕺𝕻𝕼𝕽𝕾𝕿𝖀𝖁𝖂𝖃𝖄𝖅𝖆𝖇𝖈𝖉𝖊𝖋𝖌𝖍𝖎𝖏𝖐𝖑𝖒𝖓𝖔𝖕𝖖𝖗𝖘𝖙𝖚𝖛𝖜𝖝𝖞𝖟"
	)

	var (
		vanillaIdx int
		runeMap    map[int]int
	)

	vanilla := utf8.NewString(VANILLA)

	vanillaIdx = 0
	runeMap = make(map[int]int)
	for _, r := range BLACKLETTER {
		runeMap[vanilla.At(vanillaIdx)] = r
		vanillaIdx++
	}
	fonts["blackletter"] = Font{runeMap, nil}

	vanillaIdx = 0
	runeMap = make(map[int]int)
	for _, r := range BLACKLETTER_BOLD {
		runeMap[vanilla.At(vanillaIdx)] = r
		vanillaIdx++
	}
	fonts["blackletterbold"] = Font{runeMap, nil}
}

// initialise the mappings from vanilla runes -> flipped
func initFlipped() {
	const (
		VANILLA = "zyxwvutsrqponmlkjihgfedcbaZYXWVUTSRQPONMLKJIHGFEDCBA0987654321&_?!\"'.,;"
		FLIPPED = "zʎxʍʌnʇsɹbdouɯlʞɾıɥɓɟǝpɔqɐZ⅄XMΛ∩⊥SᴚΌԀONW˥⋊ſIH⅁ℲƎ◖Ɔ𐐒∀068ㄥ9ގㄣƐᄅ⇂⅋‾¿¡„,˙'؛"
	)

	var (
		vanillaIdx int
		runeMap    map[int]int
	)

	vanilla := utf8.NewString(VANILLA)

	vanillaIdx = 0
	runeMap = make(map[int]int)
	for _, r := range FLIPPED {
		runeMap[vanilla.At(vanillaIdx)] = r
		vanillaIdx++
	}

	// a function to reverse the runes in a string
	reverse := func (s string) string {
		us := utf8.NewString(s)
		out := ""
		for i := us.RuneCount()-1; i >= 0; i-- {
			out += fmt.Sprintf("%c", us.At(i))
		}
		return out
	}

	fonts["flipped"] = Font{runeMap, reverse}
}

func (f *Font) Apply(s string) string {
	if f.pre != nil {
		s = f.pre(s)
	}

	out := ""
	for _, r := range s {
		trns, ok := f.runeMap[r]
		if ok {
			out += fmt.Sprintf("%c", trns)
		} else {
			out += fmt.Sprintf("%c", r)
		}
	}

	return out
}

func init() {
	initBlackLetter()
	initFlipped()
}

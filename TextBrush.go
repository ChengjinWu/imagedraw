package imagedraw

import (
	"container/list"
	"image"
	"io/ioutil"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/math/fixed"
)

type TextBrush struct {
	FontType  *truetype.Font
	FontSize  float64
	FontColor *image.Uniform
	TextWidth int
}

func NewTextBrush(FontFilePath string, FontSize float64, FontColor *image.Uniform, textWidth int) (*TextBrush, error) {
	fontFile, err := ioutil.ReadFile(FontFilePath)
	if err != nil {
		return nil, err
	}
	fontType, err := truetype.Parse(fontFile)
	if err != nil {
		return nil, err
	}
	if textWidth <= 0 {
		textWidth = 20
	}
	return &TextBrush{FontType: fontType, FontSize: FontSize, FontColor: FontColor, TextWidth: textWidth}, nil
}

// textLine 单行数据
type textLine struct {
	Text  string
	Index int
}

// EncodeImg 处理图片生成
func (fb *TextBrush) DrawFontOnRGBA(rgba *image.RGBA, pt image.Point, content string) {
	var dy int
	lineSize := int(fb.TextWidth / int(fb.FontSize))

	// 计算行数以及拆解字符串
	lines := list.New()
	tl := fb.sliptString(content, lines, lineSize)
	dy = tl * int(fb.FontSize)

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(fb.FontType)
	c.SetFontSize(fb.FontSize)
	c.SetClip(image.Rect(pt.X, pt.Y, pt.X+fb.TextWidth, pt.Y+dy))
	c.SetDst(rgba)
	c.SetSrc(fb.FontColor)

	intSize := int(fb.FontSize)

	var txtNext *list.Element
	for txt := lines.Front(); txt != nil; txt = txtNext {
		txtNext = txt.Next()
		t := txt.Value.(textLine)
		pt := freetype.Pt(pt.X, pt.Y+(t.Index+1)*intSize-int(c.PointToFixed(fb.FontSize))>>8)
		c.DrawString(t.Text, pt)
	}
}

// sliptString 字符串拆解
func (fb *TextBrush) sliptString(content string, lines *list.List, lineSize int) (countAll int) {
	//log.Println(tc.Text)
	countAll = 0
	text := strings.Replace(content, "\t", "    ", 0)
	texts := strings.Split(text, "\n")

	for _, v := range texts {

		runes := []rune(v)
		// 处理回车和换行
		runesLength := len(runes)

		opts := truetype.Options{}
		opts.Size = fb.FontSize
		face := truetype.NewFace(fb.FontType, &opts)
		text := ""
		var length fixed.Int26_6
		for j := 0; j < runesLength; j++ {
			faceWidth, _ := face.GlyphAdvance(runes[j])
			length += faceWidth
			if length.Ceil() > fb.TextWidth {
				countAll++
				lines.PushBack(textLine{text, countAll})
				text = ""
				length = 0
			} else {
				text += string(runes[j])
			}
		}
		countAll++
		lines.PushBack(textLine{text, countAll})
	}
	return
}

package avatar

import (
	"errors"
	"math/rand"
	"unicode"

	"github.com/fogleman/gg"
)

// Generator use to generate avatar
type Generator struct {
	radius  float64
	context *gg.Context
}

var bgColors = [...]string{`#F44336`, `#E91E63`, `#9C27B0`, `#673AB7`, `#3F51B5`, `#2196F3`, `#009688`, `#4CAF50`, `#F57F17`, `#795548`, `#424242`}

// New create new generator and load font to memory.
func New(radius float64, fontFacePath string, fontFacePoints float64) (*Generator, error) {
	c := gg.NewContext(int(radius)*2, int(radius)*2)
	err := c.LoadFontFace(fontFacePath, fontFacePoints)
	if err != nil {
		return nil, err
	}
	return &Generator{radius, c}, nil
}

// Gen generate avatar image. if success will return image file path.
func (c *Generator) Gen(name string) (string, error) {
	if c.context == nil {
		return ``, errors.New(`Generator must be initial by [func New(radius)]`)
	}

	r := []rune(name) // 可能有中文字
	l := len(r)

	c.context.SetRGBA(1, 1, 1, 0)
	c.context.Clear() // clear context

	// draw circle
	c.context.DrawCircle(c.radius, c.radius, c.radius)

	// random background color
	// rand.Seed(int64(l))
	bgColor := bgColors[rand.Intn(l)]
	c.context.SetHexColor(bgColor)
	c.context.Fill()

	// draw name
	// 巩祥啊 ---》 祥啊
	// Kevin.Gong --> Ke
	c.context.SetRGB(1, 1, 1)

	var rs []rune
	if l > 2 {
		if isChineseChar(name) {
			rs = r[l-2:]
		} else {
			rs = r[:2]
		}
	} else {
		rs = r
	}

	c.context.DrawStringAnchored(string(rs), c.radius, c.radius, 0.5, 0.5) // center
	c.context.Stroke()

	out := name + `.png`
	return out, c.context.SavePNG(out)
}

func isChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

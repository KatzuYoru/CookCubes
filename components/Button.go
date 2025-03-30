package components

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Button struct {
	X      int
	Y      int
	Width  int
	Height int
	Text   string
	Font   font.Face

	BackgroundColor      color.RGBA
	HoverBackgroundColor color.RGBA
	TextColor            color.RGBA

	isHovered  bool
	isPresed   bool
	wasClicled bool
	OnClick    func()
}

func NewButton(
	x, y, width, height int,
	text string, font font.Face, callback func(),
) *Button {
	return &Button{
		X:                    x,
		Y:                    y,
		Width:                width,
		Height:               height,
		Text:                 text,
		Font:                 font,
		BackgroundColor:      color.RGBA{60, 60, 60, 255},
		HoverBackgroundColor: color.RGBA{100, 100, 100, 255},
		TextColor:            color.RGBA{255, 255, 255, 255},
		OnClick:              callback,
	}
}

func (b *Button) Update() bool {
	x, y := ebiten.CursorPosition()

	b.isHovered = x >= b.X && x <= b.X+b.Width && y >= b.Y && y <= b.Y+b.Height

	if b.isHovered && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		b.isPresed = true
	}

	b.wasClicled = false
	if b.isHovered && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if b.isHovered {
			b.wasClicled = true
			if b.OnClick != nil {
				b.OnClick()
			}
		}
		b.isPresed = false
	}
	return b.wasClicled
}

func (b *Button) Draw(screen *ebiten.Image) {
	bgColor := b.BackgroundColor
	if b.isHovered {
		bgColor = b.HoverBackgroundColor
	}

	DrawRoundedRect(screen, float32(b.X), float32(b.Y),
		float32(b.Width), float32(b.Height),
		50, bgColor)

	bounds := text.BoundString(b.Font, b.Text)
	textWidth := bounds.Max.X - bounds.Min.X
	textHeight := bounds.Max.Y - bounds.Min.Y

	textX := b.X + (b.Width-textWidth)/2
	textY := b.Y + (b.Height+textHeight)/2

	text.Draw(screen, b.Text, b.Font, textX, textY, b.TextColor)
}

func (b *Button) IsClicked() bool {
	return b.wasClicled
}

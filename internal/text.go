package internal

import (
	"github.com/tracmo/maroto/internal/fpdf"
	"github.com/tracmo/maroto/pkg/consts"
	"github.com/tracmo/maroto/pkg/props"
)

// Text is the abstraction which deals of how to add text inside PDF.
type Text interface {
	Add(text string, cell Cell, textProp props.Text)
	GetLinesQuantity(text string, fontFamily props.Text, colWidth float64) int
	GetTextHeight(text string, textProp props.Text, colWidth float64) float64
}

type text struct {
	pdf  fpdf.Fpdf
	math Math
	font Font
}

// NewText create a Text.
func NewText(pdf fpdf.Fpdf, math Math, font Font) *text {
	return &text{
		pdf,
		math,
		font,
	}
}

// Add a text inside a cell.
func (s *text) Add(text string, cell Cell, textProp props.Text) {
	s.font.SetFont(textProp.Family, textProp.Style, textProp.Size)

	originalColor := s.font.GetColor()
	s.font.SetColor(textProp.Color)

	_, _, fontSize := s.font.GetFont()
	fontHeight := fontSize / s.font.GetScaleFactor()

	cell.Y += fontHeight

	// Apply Unicode before calc spaces
	stringWidth := s.pdf.GetStringWidth(text)
	accumulateOffsetY := 0.0

	// If should add one line
	if textProp.Extrapolate {
		s.addLine(textProp, cell.X, cell.Width, cell.Y, stringWidth, text)
	} else {
		lines := s.getLines(text, cell.Width)

		for index, line := range lines {
			lineWidth := s.pdf.GetStringWidth(line)

			s.addLine(textProp, cell.X, cell.Width, cell.Y+float64(index)*fontHeight+accumulateOffsetY, lineWidth, line)
			accumulateOffsetY += textProp.VerticalPadding
		}
	}

	s.font.SetColor(originalColor)
}

// GetLinesQuantity retrieve the quantity of lines which a text will occupy to avoid that text to extrapolate a cell.
func (s *text) GetLinesQuantity(text string, textProp props.Text, colWidth float64) int {
	// If should add one line.
	if textProp.Extrapolate {
		return 1
	}

	s.font.SetFont(textProp.Family, textProp.Style, textProp.Size)

	lines := s.getLines(text, colWidth)
	return len(lines)
}

// GetTextHeight calculate the height of lines with text
func (s *text) GetTextHeight(text string, textProp props.Text, colWidth float64) float64 {
	s.font.SetFont(textProp.Family, textProp.Style, textProp.Size)

	_, _, fontSize := s.font.GetFont()
	fontHeight := fontSize / s.font.GetScaleFactor()

	// If should add one line
	if textProp.Extrapolate {
		return fontHeight
	}

	lines := s.getLines(text, colWidth)
	numLines := float64(len(lines))

	return (numLines * fontHeight) + ((numLines - 1) * textProp.VerticalPadding)
}

func (s *text) getLines(unicodeText string, colWidth float64) []string {

	lines := []string{}
	oneLine := []rune{}
	currentlySize := 0.0
	lineHasSpace := -1   // the position of the last space in one line
	lineSpaceSize := 0.0 // the line size with last space

	for _, oneRune := range unicodeText {
		if oneRune == '\n' {
			lines = append(lines, string(oneLine))
			oneLine = []rune{}
			currentlySize = 0
			lineHasSpace = -1
			lineSpaceSize = 0.0
			continue
		}

		runeWidth := s.pdf.GetStringWidth(string(oneRune))
		if runeWidth+currentlySize < colWidth {
			// add the rune into the current line
			if oneRune == ' ' || oneRune == '\t' {
				// trace the position of last space rune in one line
				lineHasSpace = len(oneLine)
				lineSpaceSize = currentlySize + runeWidth
			}

			oneLine = append(oneLine, oneRune)
			currentlySize += runeWidth
		} else {
			// split into another new line
			if lineHasSpace < 1 {
				// no space or just prefix space in the line, do nothing
				// assert lineHasSpace == -1 && lineSpaceSize == 0.0
				lines = append(lines, string(oneLine))
				if oneRune == ' ' || oneRune == '\t' {
					// remove prefix with spaces
					oneLine = []rune{}
					currentlySize = 0
				} else {
					oneLine = []rune{oneRune}
					currentlySize = runeWidth
				}
			} else {
				// split current with space
				lines = append(lines, string(oneLine[:lineHasSpace]))

				// trim prefix space into a new line
				oneLine = oneLine[lineHasSpace+1:]
				currentlySize = currentlySize - lineSpaceSize

				// add the rune into new line
				if oneRune == ' ' || oneRune == '\t' {
					// trace the position of last space rune in one line
					lineHasSpace = len(oneLine)
					lineSpaceSize = currentlySize + runeWidth
				} else {
					lineHasSpace = -1
					lineSpaceSize = 0.0
				}

				oneLine = append(oneLine, oneRune)
				currentlySize += runeWidth
			}
		}
	}
	if len(oneLine) > 0 {
		lines = append(lines, string(oneLine))
	}

	return lines
}

func (s *text) addLine(textProp props.Text, xColOffset, colWidth, yColOffset, textWidth float64, text string) {
	left, top, _, _ := s.pdf.GetMargins()

	if textProp.Align == consts.Left {
		s.pdf.Text(xColOffset+left, yColOffset+top, text)
		return
	}

	var modifier float64 = 2

	if textProp.Align == consts.Right {
		modifier = 1
	}

	dx := (colWidth - textWidth) / modifier

	s.pdf.Text(dx+xColOffset+left, yColOffset+top, text)
}

func (s *text) textToUnicode(txt string, props props.Text) string {
	if props.Family == consts.Arial ||
		props.Family == consts.Helvetica ||
		props.Family == consts.Symbol ||
		props.Family == consts.ZapBats ||
		props.Family == consts.Courier {
		translator := s.pdf.UnicodeTranslatorFromDescriptor("")
		return translator(txt)
	}

	return txt
}

package widget

import (
	"image/color"
	"net/url"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

// Hyperlink widget is a text component with appropriate padding and layout.
// When clicked, the default web browser should open with a URL
type Hyperlink struct {
	textProvider
	Text      string
	Url       string
	Alignment fyne.TextAlign // The alignment of the Text
	TextStyle fyne.TextStyle // The style of the hyperlink text
}

// NewHyperlink creates a new hyperlink widget with the set text content
func NewHyperlink(text string, url string) *Hyperlink {
	return NewHyperlinkWithStyle(text, url, fyne.TextAlignLeading, fyne.TextStyle{})
}

// NewHyperlinkWithStyle creates a new hyperlink widget with the set text content
func NewHyperlinkWithStyle(text string, url string, alignment fyne.TextAlign, style fyne.TextStyle) *Hyperlink {
	hl := &Hyperlink{
		Text:      text,
		Url:       url,
		Alignment: alignment,
		TextStyle: style,
	}

	return hl
}

// SetText sets the text of the hyperlink
func (hl *Hyperlink) SetText(text string) {
	hl.Text = text
	hl.textProvider.SetText(text) // calls refresh
}

// SetUrl sets the URL of the hyperlink, taking in a string type
func (hl *Hyperlink) SetUrl(url string) {
	hl.Url = url
}

// SetUrl sets the URL of the hyperlink, taking in a URL type
func (hl *Hyperlink) SetUrlFromUrl(url *url.URL) {
	hl.Url = url.String()
}

// textAlign tells the rendering textProvider our alignment
func (hl *Hyperlink) textAlign() fyne.TextAlign {
	return hl.Alignment
}

// textStyle tells the rendering textProvider our style
func (hl *Hyperlink) textStyle() fyne.TextStyle {
	return hl.TextStyle
}

// textColor tells the rendering textProvider our color
func (hl *Hyperlink) textColor() color.Color {
	return theme.HyperlinkColor()
}

// password tells the rendering textProvider if we are a password field
func (hl *Hyperlink) password() bool {
	return false
}

// object returns the root object of the widget so it can be referenced
func (hl *Hyperlink) object() fyne.Widget {
	return hl
}

// OnMouseDown is called when a mouse down event is captured and triggers any change handler
func (hl *Hyperlink) OnMouseDown(*fyne.MouseEvent) {
	if hl.Url != "" {
		fyne.CurrentApp().OpenURL(hl.Url)
	}
}

// CreateRenderer is a private method to Fyne which links this widget to it's renderer
func (hl *Hyperlink) CreateRenderer() fyne.WidgetRenderer {
	hl.textProvider = newTextProvider(hl.Text, hl)
	return hl.textProvider.CreateRenderer()
}

// Resize sets a new size for a widget.
// Note this should not be used if the widget is being managed by a Layout within a Container.
func (hl *Hyperlink) Resize(size fyne.Size) {
	hl.resize(size, hl)
}

// Move the widget to a new position, relative to it's parent.
// Note this should not be used if the widget is being managed by a Layout within a Container.
func (hl *Hyperlink) Move(pos fyne.Position) {
	hl.move(pos, hl)
}

// MinSize returns the smallest size this widget can shrink to
func (hl *Hyperlink) MinSize() fyne.Size {
	return hl.minSize(hl)
}

// Show this widget, if it was previously hidden
func (hl *Hyperlink) Show() {
	hl.show(hl)
}

// Hide this widget, if it was previously visible
func (hl *Hyperlink) Hide() {
	hl.hide(hl)
}
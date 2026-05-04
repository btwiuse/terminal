package theme

import (
	"image/color"

	"charm.land/bubbles/v2/help"
	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
)

type Theme struct {
	border     color.Color
	background color.Color
	highlight  color.Color
	brand      color.Color
	error      color.Color
	body       color.Color
	accent     color.Color

	base lipgloss.Style
	form huh.Theme
}

func BasicTheme(highlight *string) Theme {
	base := Theme{}

	base.background = lipgloss.Color("#000000")
	base.border = lipgloss.Color("#3A3F42")
	base.body = lipgloss.Color("#889096")
	base.accent = lipgloss.Color("#FFFFFF")
	base.brand = lipgloss.Color("#FF5C00")
	if highlight != nil {
		base.highlight = lipgloss.Color(*highlight)
	} else {
		base.highlight = base.brand
	}
	base.error = lipgloss.Color("203")

	base.base = lipgloss.NewStyle().Foreground(base.body)
	base.form = HuhTheme(base)

	return base
}

func HuhTheme(theme Theme) huh.ThemeFunc {
	return func(_ bool) *huh.Styles {
		var t huh.Styles

		t.FieldSeparator = lipgloss.NewStyle().SetString("\n\n")

		f := &t.Focused
		f.Base = lipgloss.NewStyle().
			PaddingLeft(1).
			BorderStyle(lipgloss.ThickBorder()).
			BorderLeft(true).
			BorderForeground(theme.accent)
		f.Title = lipgloss.NewStyle().Foreground(theme.body)
		f.Description = lipgloss.NewStyle().Foreground(theme.body)
		f.TextInput.Cursor = lipgloss.NewStyle().Foreground(theme.brand)
		f.TextInput.Placeholder = lipgloss.NewStyle().Foreground(theme.body)
		f.TextInput.Prompt = lipgloss.NewStyle().Foreground(theme.accent)
		f.TextInput.Text = lipgloss.NewStyle().Foreground(theme.accent)
		f.ErrorIndicator = lipgloss.NewStyle().Foreground(theme.error)
		f.ErrorMessage = lipgloss.NewStyle().Foreground(theme.error)
		t.Help = help.New().Styles

		t.Blurred = copyFieldStyles(*f)
		t.Blurred.Base = t.Blurred.Base.BorderStyle(lipgloss.HiddenBorder())
		t.Blurred.Title = t.Blurred.Title.Foreground(theme.body)

		// TODO: add other huh form/input styles as needed

		return &t
	}
}

func (b Theme) Body() color.Color {
	return b.body
}

func (b Theme) Highlight() color.Color {
	return b.highlight
}

func (b Theme) Brand() color.Color {
	return b.brand
}

func (b Theme) Background() color.Color {
	return b.background
}

func (b Theme) Accent() color.Color {
	return b.accent
}

func (b Theme) Base() lipgloss.Style {
	return b.base.Copy()
}

func (b Theme) TextBody() lipgloss.Style {
	return b.Base().Foreground(b.body)
}

func (b Theme) TextAccent() lipgloss.Style {
	return b.Base().Foreground(b.accent)
}

func (b Theme) TextHighlight() lipgloss.Style {
	return b.Base().Foreground(b.highlight)
}

func (b Theme) TextBrand() lipgloss.Style {
	return b.Base().Foreground(b.brand)
}

func (b Theme) TextError() lipgloss.Style {
	return b.Base().Foreground(b.error)
}

func (b Theme) PanelError() lipgloss.Style {
	return b.Base().Background(b.error).Foreground(b.accent)
}

func (b Theme) Form() huh.Theme {
	return b.form
}

func (b Theme) Border() color.Color {
	return b.border
}

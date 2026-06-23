package fang

import (
	"image/color"
	"testing"

	"charm.land/lipgloss/v2"
)

func TestWithoutCodeblockFill(t *testing.T) {
	// Set the test width before calling makeStyles so that width()'s
	// sync.OnceValue cache is initialized at the same value golden tests use.
	t.Setenv("__FANG_TEST_WIDTH", "45")
	styles := makeStyles(DefaultColorScheme(lipgloss.LightDark(true)))

	// sanity: default codeblock styles DO carry a background.
	if isNoColor(styles.Codeblock.Base.GetBackground()) {
		t.Fatal("expected default Codeblock.Base to have a background")
	}

	stripped := styles.withoutCodeblockFill()

	checks := map[string]color.Color{
		"Base":                   stripped.Codeblock.Base.GetBackground(),
		"Text":                   stripped.Codeblock.Text.GetBackground(),
		"Comment":                stripped.Codeblock.Comment.GetBackground(),
		"Program.Name":           stripped.Codeblock.Program.Name.GetBackground(),
		"Program.Command":        stripped.Codeblock.Program.Command.GetBackground(),
		"Program.Flag":           stripped.Codeblock.Program.Flag.GetBackground(),
		"Program.Argument":       stripped.Codeblock.Program.Argument.GetBackground(),
		"Program.DimmedArgument": stripped.Codeblock.Program.DimmedArgument.GetBackground(),
		"Program.QuotedString":   stripped.Codeblock.Program.QuotedString.GetBackground(),
	}
	for name, bg := range checks {
		if !isNoColor(bg) {
			t.Errorf("%s background = %v, want unset", name, bg)
		}
	}

	// original must be unchanged (value semantics).
	if isNoColor(styles.Codeblock.Base.GetBackground()) {
		t.Error("original styles were mutated")
	}
}

func isNoColor(c color.Color) bool {
	_, ok := c.(lipgloss.NoColor)
	if ok {
		return true
	}
	return c == nil
}

func TestBoxMode(t *testing.T) {
	tests := []struct {
		name string
		opts helpOptions
		want boxMode
	}{
		{"default", helpOptions{}, boxFilled},
		{"without fill", helpOptions{noFill: true}, boxOutlined},
		{"without box", helpOptions{noBox: true}, boxNone},
		{"both, box wins", helpOptions{noFill: true, noBox: true}, boxNone},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.opts.boxMode(); got != tt.want {
				t.Errorf("boxMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

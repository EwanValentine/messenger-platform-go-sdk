package template

import "testing"

func TestButtonType(t *testing.T) {
	template := &ButtonTemplate{}
	if template.Type() != TemplateTypeButton {
		t.Error("Button template returned invalid type")
	}
}

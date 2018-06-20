package components

type TextStruct struct {
	ComponentData struct {
		Title string
		Text string
	}
}

func (this TextStruct) GetHtml() string {
	return this.ComponentData.Text
}

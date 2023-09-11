package regulartext

type RegularText struct {
	Value string
}

func ExtractRegularText(rawMD string) RegularText {
	return RegularText{Value: rawMD}
}

func (rt *RegularText) ToHTMLString() string {
	return rt.Value
}

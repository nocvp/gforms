package gforms

import (
	"bytes"
)

type fileInputWidget struct {
	Type  string
	Attrs map[string]string
	Widget
}

func (wg *fileInputWidget) html(f FieldInterface) string {
	var buffer bytes.Buffer
	err := Template.ExecuteTemplate(&buffer, "FileInputWidget", widgetContext{
		Type:  wg.Type,
		Field: f,
		Attrs: wg.Attrs,
		Value: f.GetV().RawStr,
	})
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

// Generate file input field: <input type="file" ...>
func FileInputWidget(attrs map[string]string) Widget {
	w := new(fileInputWidget)
	w.Type = "file"
	if attrs == nil {
		attrs = map[string]string{}
	}
	w.Attrs = attrs
	return w
}

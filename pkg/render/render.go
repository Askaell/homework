package render

type Renderer interface {
	Render(input interface{}) ([]byte, error)
}

//getRender is a fabric method returning new Renderer
func GetRenderer(renderType string) Renderer {
	switch renderType {
	case "html":
		return NewHTMLRenderer()
	case "json":
		return NewJSONRenderer()
	default:
		return NewJSONRenderer()
	}
}

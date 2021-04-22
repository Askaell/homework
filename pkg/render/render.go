package render

type Renderer interface {
	Render(input interface{}) ([]byte, error)
}

//getRender is a fabric method returning new Renderer
func getRenderer(renderType string) Renderer {
	switch renderType {
	case "json":
		return NewJSONRenderer()
	default:
		return NewJSONRenderer()
	}
}

package render

import (
	"encoding/json"
	"fmt"
)

type JSONRenderer struct {
}

func NewJSONRenderer() *JSONRenderer {
	return &JSONRenderer{}
}

func (r *JSONRenderer) Render(input interface{}) ([]byte, error) {
	output, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("json render invalid model")
	}

	return output, nil
}

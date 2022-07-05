package models

//TemplateData contains data that can be sent to the render from handler
type TemplateData struct {
	StringMap  map[string]string
	IntMap     map[int]int
	FloatMap   map[float32]float32
	DataMap    map[string]interface{}
	ActivePage string
	CSRFToken  string
	Flash      string
	Warning    string
	Error      string
}

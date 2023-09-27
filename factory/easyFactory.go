package factory

type IParser interface {
	parse(data []byte)
}
type jsonParser struct {
}

func (j jsonParser) parse(data []byte) {
	panic("implement me")
}

type yamlParser struct {
}

func (y yamlParser) parse(data []byte) {
	panic("implement me")
}
func NewIParser(str string) IParser {
	switch str {
	case "json":
		return jsonParser{}
	case "yaml":
		return yamlParser{}
	}
	return nil
}

package policy

type PolicyDefinition struct {
	Id      string
	Name    string
	Version string

	IsValid func(toMeasure interface{}) (bool, string)
}

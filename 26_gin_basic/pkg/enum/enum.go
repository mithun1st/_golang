package enum

type EnviromentE int

const (
	Development EnviromentE = iota + 1
	Staging
	Production
)

func (e EnviromentE) ToTitle() string {

	switch e {
	case Development:
		return "Development"

	case Staging:
		return "Staging"

	case Production:
		return "Production"

	default:
		return ""
	}
}

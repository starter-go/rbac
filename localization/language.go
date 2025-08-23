package localization

type Language string

func (l Language) String() string {
	return string(l)
}

////////////////////////////////////////////////////////////////////////////////

const (
	theLanguageZH = "zh"
)

const (
	LanguageZH Language = theLanguageZH
)

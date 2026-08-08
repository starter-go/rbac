package localization

// Locale = Language + Region
type Locale string

func (l Locale) String() string {
	return string(l)
}

////////////////////////////////////////////////////////////////////////////////

const (
	sep = "_"

	LocaleZhongCN Locale = theLanguageZH + sep + theRegionChina
)

func LocaleOf(language Language, region Region) Locale {
	str := language.String() + sep + region.String()
	return Locale(str)
}

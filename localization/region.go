package localization

type Region string

func (r Region) String() string {
	return string(r)
}

////////////////////////////////////////////////////////////////////////////////

const (
	theRegionChina = "CN"
	theRegionJapan = "JP"
	theRegionUSA   = "US"
)

const (
	RegionChina Region = theRegionChina
	RegionJapan Region = theRegionJapan
	RegionUSA   Region = theRegionUSA
)

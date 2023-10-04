package utils

type ProvinceInfo struct {
	Name     string
	Code     string
	CityInfo []CityInfo
}
type CityInfo struct {
	Name     string
	Code     string
	AreaInfo []AreaInfo
}
type AreaInfo struct {
	Name string
	Code string
}

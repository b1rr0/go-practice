package zoo

type SectorType int

const (
	SectorTypeEnam SectorType = iota
	SectorTypeWild
	SectorTypeCage
	SectorTypeAviary
)

const DefaultSectorType = SectorTypeEnam

func (s SectorType) String() string {
	return [...]string{"enam", "wild", "cage", "aviary"}[s]
}

func AllSectorTypeValues() []SectorType {
	return []SectorType{SectorTypeEnam, SectorTypeWild, SectorTypeCage, SectorTypeAviary}
}

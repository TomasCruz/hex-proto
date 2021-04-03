package scanner

type (
	Hexameter struct {
		HexID int32
		Para  bool
		Txt   string
	}

	WordToExport struct {
		ID  int32
		Txt string
	}

	HexWordToExport struct {
		HexID  int32
		Order  int32
		WordID int32
	}

	Translatable struct {
		Words    []WordToExport
		Hexes    []Hexameter
		HexWords []HexWordToExport
	}
)

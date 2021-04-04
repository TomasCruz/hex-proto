package scanner

type (
	Hexameter struct {
		HexID int32
		Para  bool
		Txt   string
	}

	Word struct {
		ID  int32
		Txt string
	}

	HexWord struct {
		HexID  int32
		Order  int32
		WordID int32
	}

	Translatable struct {
		Words    []Word
		Hexes    []Hexameter
		HexWords []HexWord
	}
)

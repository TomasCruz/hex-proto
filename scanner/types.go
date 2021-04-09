package scanner

type (
	Hexameter struct {
		HexID int64
		Para  bool
		Txt   string
	}

	Word struct {
		ID  int64
		Txt string
	}

	HexWord struct {
		HexID  int64
		Order  int64
		WordID int64
	}

	Translatable struct {
		Words    []Word
		Hexes    []Hexameter
		HexWords []HexWord
	}
)

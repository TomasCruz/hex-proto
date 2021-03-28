package scanner

type (
	Hexameter struct {
		ID   int32
		Para bool
		Txt  string
	}

	WordToExport struct {
		ID  int
		Txt string
	}

	HexWordToExport struct {
		ID     int
		Order  int
		WordID int
	}

	ParsedToExport struct {
		ID       int
		Para     bool
		Txt      string
		HexWords []HexWordToExport
	}

	Translatable struct {
		Words  []WordToExport
		Parsed []ParsedToExport
	}
)

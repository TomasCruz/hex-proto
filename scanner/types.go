package scanner

type (
	Hexameter struct {
		HexID int64  `json:"hex_id"`
		Para  bool   `json:"para"`
		Txt   string `json:"txt"`
	}

	Word struct {
		ID  int64  `json:"id"`
		Txt string `json:"txt"`
	}

	HexWord struct {
		HexID  int64 `json:"hex_id"`
		Order  int64 `json:"order"`
		WordID int64 `json:"word_id"`
	}

	Translatable struct {
		Words    []Word      `json:"words"`
		Hexes    []Hexameter `json:"hexes"`
		HexWords []HexWord   `json:"hex_words"`
	}

	ScannedHex struct {
		HexID           int64    `json:"hex_id"`
		Para            bool     `json:"para"`
		Txt             string   `json:"txt"`
		Scanned         string   `json:"scanned"`
		MusicalTemplate string   `json:"musical_template"`
		ChangeCount     int      `json:"change_count"`
		ScanningVariety int      `json:"scanning_variety"`
		Selected        bool     `json:"selected"`
		MainCaesura     int      `json:"main_caesura"`
		Caesurae        []int    `json:"caesurae"`
		Words           []string `json:"words"`
	}
)

package translator

type (
	Lang struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	HexLang struct {
		HexID  int64 `json:"hex_id"`
		LangID int64 `json:"lang_id"`
	}

	Count struct {
		Count int64 `json:"count"`
	}

	HWLang struct {
		HexLang HexLang `json:"hex_lang"`
		Order   int64   `json:"order"`
	}

	HWTrans struct {
		HWLang      HWLang `json:"hw_lang"`
		Grammar     string `json:"grammar"`
		Translation string `json:"translation"`
	}

	Trans struct {
		HexID       int64  `json:"hex_id"`
		Translation string `json:"translation"`
	}
)

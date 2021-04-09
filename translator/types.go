package translator

type (
	Lang struct {
		ID   int64
		Name string
	}

	HexLang struct {
		HexID  int64
		LangID int64
	}

	Count struct {
		Count int64
	}

	HWLang struct {
		HexLang HexLang
		Order   int64
	}

	HWTrans struct {
		HWLang      HWLang
		Grammar     string
		Translation string
	}

	Trans struct {
		HexID       int64
		Translation string
	}
)

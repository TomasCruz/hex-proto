package translator

type (
	Lang struct {
		ID   int
		Name string
	}

	HexLang struct {
		HexID  int32
		LangID int32
	}

	Count struct {
		Count int32
	}

	HWLang struct {
		HexLang HexLang
		Order   int32
	}

	HWTrans struct {
		HWLang      HWLang
		Grammar     string
		Translation string
	}

	Trans struct {
		HexID       int32
		Translation string
	}
)

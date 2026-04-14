package jsoncensor

type MatchMode int

const (
	MatchSubstring MatchMode = iota
	MatchWholeValue
)

type CensorConfig struct {
	Needles         []string
	CaseInsensitive bool
	Mode            MatchMode
}

func CensorJSON(jsonData []byte, cfg CensorConfig) ([]byte, error) {
	panic("implement me")
}

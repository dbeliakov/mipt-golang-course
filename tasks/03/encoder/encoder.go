package encoder

type EncodeMode string

const (
	DashMode       EncodeMode = "-"
	UnderScoreMode EncodeMode = "_"
)

type CustomEncoder string

func NewCustomEncoder(mode EncodeMode) *CustomEncoder {
	return nil
}

func (e *CustomEncoder) Encode(input string) string {
	return ""
}

func (e *CustomEncoder) Decode(input string) string {
	return ""
}

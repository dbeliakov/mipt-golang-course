package weekday

type Weekday any

const (
	Monday    = 0
	Tuesday   = 0
	Wednesday = 0
	Thursday  = 0
	Friday    = 0
	Saturday  = 0
	Sunday    = 0
)

func NextDay(day Weekday) Weekday {
	return Monday
}

package constants

const (
	ROWS    = 6
	COLUMNS = 7
)

type GameStatus string

const (
	NOT_STARTED GameStatus = "NOT_STARTED"
	IN_PROGRESS GameStatus = "IN_PROGRESS"
	ENDED       GameStatus = "ENDED"
)

const WIN_COUNT = 4

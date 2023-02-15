package logger

type WarnAndError interface {
	Warnw(msg string, args ...any)
	Errorw(msg string, err any, args ...any)
}

type Lite interface {
	Infow(msg string, args ...any)
	Warnw(msg string, args ...any)
	Errorw(msg string, err any, args ...any)
}
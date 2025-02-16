package model

type LimitHandling string

const (
	MoveLastToClosed LimitHandling = "moveLastToClosed"
	PushFront        LimitHandling = "pushFront"
)

// ShortlistConfig
type Config struct {
	MaxCount      int           `yaml:"maxCount"`
	LimitHandling LimitHandling `yaml:"limitHandling"`
}

func DefaultConfig() Config {
	return Config{
		MaxCount:      3,
		LimitHandling: MoveLastToClosed,
	}
}

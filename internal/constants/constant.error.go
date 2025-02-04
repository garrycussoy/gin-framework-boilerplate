package constants

import "errors"

var (
	// Global
	ErrUnexpected = errors.New("unexpected error")

	// Config
	ErrLoadConfig  = errors.New("failed to load config file")
	ErrParseConfig = errors.New("failed to parse env to config struct")
	ErrEmptyVar    = errors.New("required variabel environment is empty")
)

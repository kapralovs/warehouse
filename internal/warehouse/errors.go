package warehouse

import "errors"

var (
	ErrEmptyId       = errors.New("can't save cell with empty ID field")
	ErrNilCell       = errors.New("can't save current cell, because profile is nil")
	ErrCellNotExists = errors.New("cell does not exist")
)

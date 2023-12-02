package storage

import "errors"

// Общие ошибки для нашего хранилища
var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists   = errors.New("url exists")
)

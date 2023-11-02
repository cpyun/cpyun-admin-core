package source

import "errors"

var (
	// ErrWatcherStopped is returned when source watcher has been stopped
	ErrWatcherStopped = errors.New("watcher stopped")
)

type Source interface {
	Read() (*ChangeSet, error)
	Watch() (Watcher, error)
	String() string
}

// Watcher watches a source for changes
type Watcher interface {
	Next() (*ChangeSet, error)
	Stop() error
}

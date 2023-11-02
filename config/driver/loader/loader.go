package loader

import "github.com/cpyun/cpyun-admin-core/config/driver/source"

type Loader interface {
	Load(...source.Source) error
	// Snapshot A Snapshot of loaded config
	Snapshot() (*Snapshot, error)
	// Watch for changes
	Watch(...string) (Watcher, error)
	String() string
}
type Watcher interface {
	Next() (*Snapshot, error)
}

// Snapshot is a merged ChangeSet
type Snapshot struct {
	// The merged ChangeSet
	ChangeSet *source.ChangeSet
	// Version Deterministic and comparable version of the snapshot
	Version string
}

func Copy(s *Snapshot) *Snapshot {
	cs := *(s.ChangeSet)

	return &Snapshot{
		ChangeSet: &cs,
		Version:   s.Version,
	}
}

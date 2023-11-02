package file

import (
	"github.com/cpyun/cpyun-admin-core/config/driver/source"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type watcher struct {
	f    *file
	exit chan bool
}

func newWatcher(f *file) (source.Watcher, error) {
	w := &watcher{
		f:    f,
		exit: make(chan bool, 1),
	}
	go w.watch()

	return w, nil
}
func (w *watcher) Next() (set *source.ChangeSet, err error) {
	select {
	case <-w.exit:
		set, err = w.f.Read()
		return
	}
}

func (w *watcher) Stop() error {
	close(w.exit)
	return nil
}

func (w *watcher) watch() error {
	viper.OnConfigChange(func(in fsnotify.Event) {
		//log.Println(in.String())
		w.exit <- true
		return
	})
	viper.WatchConfig()
	return nil
}

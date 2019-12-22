package winsize

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// NewWatcher returns a new Watcher
func NewWatcher() *Watcher {
	return &Watcher{updateCH: make(chan Size)}
}

// Watcher holds a channel which can be used to receive the current winsize
// after starting the process.
type Watcher struct {
	updateCH chan Size
}

// Run will block until a window size change occurs using the os signal SIGWINCH.
// Afterwards the current window size is tried to be read out. If an error
// occurs, this is returned. If the context is done, <nil> is returned.
func (u *Watcher) Run(ctx context.Context) error {
	sigwinch := make(chan os.Signal)
	signal.Notify(sigwinch, syscall.SIGWINCH)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-sigwinch:
			s, err := Get()
			if err != nil {
				close(u.updateCH)
				return err
			}
			u.updateCH <- *s
		}
	}
}

// Subscribe returns a channel from which the winsize that occurred after a
// change can be read.
func (u *Watcher) Subscribe() <-chan Size {
	return u.updateCH
}

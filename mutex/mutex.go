package mutex

type (
	// ChanMutex собственный тип блокировки чтения/записи.
	// т.к. sync.Mutex приводит к deadlock
	ChanMutex interface {
		Lock()
		Unlock()
	}

	chanMutex struct {
		lockChan chan struct{}
	}
)

func (m *chanMutex) Lock() {
	m.lockChan <- struct{}{}
}

func (m *chanMutex) Unlock() {
	<-m.lockChan
}

func NewChanMutex() ChanMutex {
	return &chanMutex{
		lockChan: make(chan struct{}, 1),
	}
}

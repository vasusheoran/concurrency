package readerwriter

import "sync"

type ReaderWriter interface {
	Read() interface{}
	Write(data interface{})
}

type readerWriter struct {
	mutex *sync.RWMutex
	data  interface{}
	ch    chan int
	wg    *sync.WaitGroup
	rCh   chan int
	wCh   chan int
}

func (r *readerWriter) Read() interface{} {
	defer r.wg.Done()
	r.mutex.RLock()
	r.rCh <- 1
	val := r.data
	r.rCh <- -1
	r.mutex.RUnlock()
	return val
}

func (r *readerWriter) Write(data interface{}) {
	defer r.wg.Done()
	r.mutex.Lock()
	r.wCh <- 1
	r.data = data
	r.wCh <- -1
	r.mutex.Unlock()
}

func NewReaderWriter(data interface{}, wg *sync.WaitGroup, rCh, wCh chan int) ReaderWriter {
	if rCh == nil {
		rCh = make(chan int)
	}

	if wCh == nil {
		wCh = make(chan int)
	}

	return &readerWriter{
		mutex: &sync.RWMutex{},
		wg:    wg,
		rCh:   rCh,
		wCh:   wCh,
		data:  data,
	}
}

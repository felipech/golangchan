package worklist

type Entry struct {
	Path string
}

type Worklist struct {
	jobs chan Entry
}

func (w *Worklist) Add(work Entry) {
	w.jobs <- work
}

func (w *Worklist) Next() Entry {
	j := <-w.jobs
	return j
}

func New(bufferSize int) Worklist {
	return Worklist{make(chan Entry, bufferSize)}
}

func NewJob(path string) Entry {
	return Entry{path}
}

func (w *Worklist) Finalize(numeroWorkers int) {
	for i := 0; i < numeroWorkers; i++ {
		w.Add(Entry{""})
	}
}

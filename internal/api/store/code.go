package store

type CodeStore interface {
}

type codeStore struct {
}

func NewCodeStore(ds *datastore) CodeStore {
	return &codeStore{}
}

package memory

import (
	"bytes"
	"fmt"
	"io"

	"github.com/deifyed/post-service/pkg/stores"
)

func (receiver *store) StoreContent(id string, content io.Reader) error {
	raw, err := io.ReadAll(content)
	if err != nil {
		return fmt.Errorf("reading content: %w", err)
	}

	receiver.db[id] = raw

	return nil
}

func (receiver *store) RetrieveContent(id string) (io.Reader, error) {
	raw, ok := receiver.db[id]
	if !ok {
		return nil, stores.ErrNotFound
	}

	return bytes.NewReader(raw), nil
}

func (receiver *store) DeleteContent(id string) error {
	_, ok := receiver.db[id]
	if !ok {
		return stores.ErrNotFound
	}

	delete(receiver.db, id)

	return nil
}

func New() stores.ContentStore {
	return &store{
		db: make(map[string][]byte),
	}
}

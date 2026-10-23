package memstore_test

import (
	"context"

	"example.com/go-behavioral-spec/examples/04-in-memory-store"
)

// stubDB records DB calls for leaf-boundary asserts in Store specs.
type stubDB struct {
	getValue []byte
	getErr   error
	putErr   error
	delErr   error
	listKeys []string
	listVals [][]byte
	listNext string
	listErr  error

	getCalls    []string
	putCalls    []putCall
	deleteCalls []string
	listCalls   []listCall
}

type putCall struct {
	key   string
	value []byte
}

type listCall struct {
	prefix    string
	limit     int
	pageToken string
}

func (s *stubDB) Get(_ context.Context, key string) ([]byte, error) {
	s.getCalls = append(s.getCalls, key)
	return s.getValue, s.getErr
}

func (s *stubDB) Put(_ context.Context, key string, value []byte) error {
	cp := append([]byte(nil), value...)
	s.putCalls = append(s.putCalls, putCall{key: key, value: cp})
	return s.putErr
}

func (s *stubDB) Delete(_ context.Context, key string) error {
	s.deleteCalls = append(s.deleteCalls, key)
	return s.delErr
}

func (s *stubDB) List(_ context.Context, prefix string, limit int, pageToken string) ([]string, [][]byte, string, error) {
	s.listCalls = append(s.listCalls, listCall{prefix: prefix, limit: limit, pageToken: pageToken})
	return s.listKeys, s.listVals, s.listNext, s.listErr
}

// ensure stubDB implements memstore.DB
var _ memstore.DB = (*stubDB)(nil)

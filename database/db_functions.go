package database

import (
	"github.com/dgraph-io/badger/v4"
	"marckent/dictionary/domain"
	"marckent/dictionary/shared"
	"sort"
)

func (w *WordDb) iterate(it *badger.Iterator) (result []domain.WordEntry, err error) {
	for it.Rewind(); it.Valid(); it.Next() {
		item := it.Item()

		err = item.Value(func(v []byte) error {
			result = append(result, shared.DecodeStruct[domain.WordEntry](v))
			return nil
		})
	}
	return
}

func sortWords(entries []domain.WordEntry) []domain.WordEntry {
	sort.SliceStable(entries, func(i, j int) bool {
		return entries[i].Word < entries[j].Word
	})
	return entries
}

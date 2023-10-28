package database

import (
	"github.com/dgraph-io/badger/v4"
	"marckent/dictionary/domain"
	"marckent/dictionary/shared"
)

type WordDb struct {
	db *badger.DB
}

func (w *WordDb) Save(entry domain.WordEntry) (err error) {
	err = w.db.Update(func(txn *badger.Txn) error {
		entryBuff := shared.EncodeStruct(entry)
		err = txn.Set([]byte(entry.Word), entryBuff.Bytes())
		return err
	})

	return
}

func (w *WordDb) Remove(word string) (err error) {
	err = w.db.Update(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(word))
		if err != nil {
			return err
		}

		err = txn.Delete([]byte(word))
		return err
	})

	return
}

func (w *WordDb) GetOne(word string) (result domain.WordEntry, err error) {
	err = w.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(word))
		if err != nil {
			return err
		}

		err = item.Value(func(v []byte) error {
			result = shared.DecodeStruct[domain.WordEntry](v)
			return nil
		})

		return err
	})

	return
}

func (w *WordDb) GetAll() (result []domain.WordEntry, err error) {
	err = w.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()

		result, err = w.iterate(it)

		return err
	})

	result = sortWords(result)
	return
}

func NewWordRepository(bd *badger.DB) *WordDb {
	return &WordDb{db: bd}
}

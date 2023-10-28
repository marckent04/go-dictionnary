package domain

type WordRepository interface {
	Save(word WordEntry) (err error)
	GetAll() (result []WordEntry, err error)
	Remove(label string) (err error)
	GetOne(word string) (result WordEntry, err error)
}

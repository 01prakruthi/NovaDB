package mvcc

type Transaction struct {
	SnapshotTime int
}

func NewTransaction(
	snapshotTime int,
) *Transaction {

	return &Transaction{
		SnapshotTime: snapshotTime,
	}
}

func (t *Transaction) Read(
	store *Store,
	key string,
) (string, bool) {

	return store.GetAtTimestamp(
		key,
		t.SnapshotTime,
	)
}

package memtable

type MemTable struct {
	data map[string]string
}

func New() *MemTable {
	return &MemTable{
		data: make(map[string]string),
	}
}

func (m *MemTable) Put(key, value string) {
	m.data[key] = value
}

func (m *MemTable) Get(key string) (string, bool) {
	value, ok := m.data[key]
	return value, ok
}

func (m *MemTable) Delete(key string) {
	delete(m.data, key)
}

func (m *MemTable) Restore(data map[string]string) {
	for k, v := range data {
		m.data[k] = v
	}
}
func (m *MemTable) Data() map[string]string {
	return m.data
}

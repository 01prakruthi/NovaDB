package mvcc

type Version struct {
	VersionID int
	Timestamp int
	Value     string
}

type Store struct {
	data map[string][]Version
}

func New() *Store {
	return &Store{
		data: make(map[string][]Version),
	}
}

func (s *Store) Put(
	key string,
	value string,
	timestamp int,
) {

	versions := s.data[key]

	newVersion := Version{
		VersionID: len(versions) + 1,
		Timestamp: timestamp,
		Value:     value,
	}

	s.data[key] = append(
		versions,
		newVersion,
	)
}

func (s *Store) GetLatest(key string) (string, bool) {

	versions := s.data[key]

	if len(versions) == 0 {
		return "", false
	}

	return versions[len(versions)-1].Value, true
}

func (s *Store) GetVersion(
	key string,
	versionID int,
) (string, bool) {

	versions := s.data[key]

	for _, v := range versions {

		if v.VersionID == versionID {
			return v.Value, true
		}
	}

	return "", false
}
func (s *Store) GetAtTimestamp(
	key string,
	timestamp int,
) (string, bool) {

	versions := s.data[key]

	var result string
	found := false

	for _, v := range versions {

		if v.Timestamp <= timestamp {

			result = v.Value
			found = true
		}
	}

	return result, found
}

package internal

type DataStore struct {
	Store map[string]string
}

func InitDataStore() DataStore {
	s := make(map[string]string)
	s["version"] = "Sumir's KV Store : 1.0.0"
	return DataStore{ Store: s }
}

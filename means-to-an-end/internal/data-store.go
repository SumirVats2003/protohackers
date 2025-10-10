package internal

type DataStore struct {
	Store map[int32]int32
}

func InitDataStore() *DataStore {
	storeMap := make(map[int32]int32)
	dataStore := DataStore{Store: storeMap}
	return &dataStore
}

func (d *DataStore) Insert(timestamp, price int32) {
	d.Store[timestamp] = price
}

func (d *DataStore) Get(timestamp int32) int32 {
	return d.Store[timestamp]
}

func GetAvg(min, max int32) {}

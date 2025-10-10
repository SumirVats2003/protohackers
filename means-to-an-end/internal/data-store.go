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
	_, exists := d.Store[timestamp]
	if !exists {
		d.Store[timestamp] = price
	} else {
		delete(d.Store, timestamp)
	}
}

func (d *DataStore) Get(timestamp int32) int32 {
	return d.Store[timestamp]
}

func (d *DataStore) GetAvg(low, high int32) float64 {
	sum, count := int32(0), 0

	for k, v := range d.Store {
		if k <= high && k >= low {
			sum += v
			count++
		}
	}

	return float64(sum) / float64(count)
}

package test

import (
	"testing"

	"github.com/SumirVats2003/protohackers/means-to-an-end/internal"
)

func TestDataStore(t *testing.T) {
	t.Run("test creating a data store", func(t *testing.T) {
		dataStore := internal.InitDataStore()
		t.Log(dataStore)
	})

	t.Run("test inserting in a data store", func(t *testing.T) {
		dataStore := internal.InitDataStore()
		timeStamp := int32(12345)
		price := int32(101)
		got := insertAndGetValue(dataStore, timeStamp, price)

		assertEquals(t, got, int32(price))
	})

	t.Run("test inserting different values in a data store", func(t *testing.T) {
		dataStore := internal.InitDataStore()

		var price1 int32 = 101
		got1 := insertAndGetValue(dataStore, 12345, price1)

		var price2 int32 = 20001
		got2 := insertAndGetValue(dataStore, 17829, price2)

		assertEquals(t, got1, price1)
		assertEquals(t, got2, price2)
	})

	t.Run("test inserting several different values in a data store", func(t *testing.T) {
		dataStore := internal.InitDataStore()

		timeStamps := []int32{12345, 98283, 987839892, 29891, 908908032}
		prices := []int32{2321, 987, 78384, 29838, 00}

		for i := range 5 {
			got := insertAndGetValue(dataStore, timeStamps[i], prices[i])
			assertEquals(t, got, prices[i])
		}
	})

	t.Run("test inserting prices on the same timestamp results in invalid operation", func(t *testing.T) {
		dataStore := internal.InitDataStore()

		timeStamp := int32(445353)
		price1 := int32(934834)
		price2 := int32(479837)

		got := insertAndGetValue(dataStore, timeStamp, price1)
		assertEquals(t, got, price1)
		res := insertAndGetValue(dataStore, timeStamp, price2)
		assertNotEquals(t, res, price1)
		assertNotEquals(t, res, price2)
	})

	t.Run("test getting an average price over a range of time", func(t *testing.T) {
		dataStore := internal.InitDataStore()

		lowerLimit := int32(1000)
		upperLimit := int32(10000)
		timeStamps := []int32{3445, 9099, 2334, 6877, 1228}
		prices := []int32{43, 997, 345, 837591, 53422}

		sum := int32(0)
		for i := range timeStamps {
			sum += insertAndGetValue(dataStore, timeStamps[i], prices[i])
		}
		avg := float64(sum) / float64(5)
		actualAvg := dataStore.GetAvg(lowerLimit, upperLimit)

		assertFloatEquals(t, actualAvg, avg)
	})

	t.Run("test get average on duplicate insertion on timestamp", func(t *testing.T) {
		dataStore := internal.InitDataStore()

		lowerLimit := int32(1000)
		upperLimit := int32(10000)
		timeStamps := []int32{3445, 9099, 2334, 9099, 1228}
		prices := []int32{43, 997, 345, 837591, 53422}

		sum := int32(0)
		for i := range timeStamps {
			sum += insertAndGetValue(dataStore, timeStamps[i], prices[i])
		}
		sum -= (997)
		avg := float64(sum) / float64(3)
		actualAvg := dataStore.GetAvg(lowerLimit, upperLimit)

		assertFloatEquals(t, actualAvg, avg)
	})
}

func insertAndGetValue(dataStore *internal.DataStore, timeStamp int32, price int32) int32 {
	dataStore.Insert(timeStamp, price)
	got := dataStore.Get(timeStamp)
	return got
}

func assertFloatEquals(t testing.TB, got float64, price float64) {
	if got != price {
		t.Errorf("got %v but expected %v", got, price)
	}
}

func assertEquals(t testing.TB, got int32, price int32) {
	if got != price {
		t.Errorf("got %v but expected %v", got, price)
	}
}

func assertNotEquals(t testing.TB, got int32, price int32) {
	if got == price {
		t.Errorf("expected %v to be different than %v", got, price)
	}
}

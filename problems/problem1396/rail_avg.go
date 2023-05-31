package problem1396

/*
An underground railway system is keeping track of customer travel times between different stations.
They are using this data to calculate the average time it takes to travel from one station to another.
Implement the UndergroundSystem class:
    void checkIn(int id, string stationName, int t)
        A customer with a card ID equal to id, checks in at the station stationName at time t.
        A customer can only be checked into one place at a time.
    void checkOut(int id, string stationName, int t)
        A customer with a card ID equal to id, checks out from the station stationName at time t.
    double getAverageTime(string startStation, string endStation)
        Returns the average time it takes to travel from startStation to endStation.
        The average time is computed from all the previous traveling times from startStation to endStation that happened directly,
		meaning a check in at startStation followed by a check out from endStation.
        The time it takes to travel from startStation to endStation may be different from the time it takes to travel from endStation to startStation.
        There will be at least one customer that has traveled from startStation to endStation before getAverageTime is called.
You may assume all calls to the checkIn and checkOut methods are consistent.
If a customer checks in at time t1 then checks out at time t2, then t1 < t2.
All events happen in chronological order.
*/

type UndergroundSystem struct {
	Checked  map[int]CheckedIn  // People currently checked in
	Stations map[string]Station // Station by name
}

type Station struct {
	Rates map[string]Rate // Rates by destinations
}

type Rate struct {
	Total int // Total time travelled
	Count int // Number of trips
}

type CheckedIn struct {
	Start string // Starting station
	Time  int    // Starting time
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		Checked:  make(map[int]CheckedIn),
		Stations: make(map[string]Station),
	}
}

func (rail *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	// Adding person to checked-in list
	rail.Checked[id] = CheckedIn{Start: stationName, Time: t}
}

func (rail *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	// Getting start station and time
	var start = rail.Checked[id]

	// Removing from checked-in list
	delete(rail.Checked, id)

	// Creating start station if not exists
	if _, ok := rail.Stations[start.Start]; !ok {
		rail.Stations[start.Start] = Station{make(map[string]Rate)}
	}

	// Creating end station if not exists
	if _, ok := rail.Stations[stationName]; !ok {
		rail.Stations[stationName] = Station{make(map[string]Rate)}
	}

	// Updating source station rates
	sourceRates := rail.Stations[start.Start].Rates
	sourceRates[stationName] = Rate{
		Total: sourceRates[stationName].Total + (t - start.Time),
		Count: sourceRates[stationName].Count + 1,
	}
}

func (rail *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	rates := rail.Stations[startStation].Rates[endStation]
	return float64(rates.Total) / float64(rates.Count)
}

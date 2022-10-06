package problem0981

/*
Design a time-based key-value data structure that can store multiple values for the same key at different time stamps
and retrieve the key's value at a certain timestamp.

Implement the TimeMap class:

TimeMap() Initializes the object of the data structure.
void set(String key, String value, int timestamp) Stores the key key with the value value at the given time timestamp.
String get(String key, int timestamp) Returns a value such that set was called previously, with timestamp_prev <= timestamp.
If there are multiple such values, it returns the value associated with the largest timestamp_prev. If there are no values, it returns "".
*/

type TimeMap struct {
	KVs map[string]*KVTime
}

type KVTime struct {
	// Timestamps will be strictly increasing so we can
	// take advantage of binary search
	Timestamps []int
	// Corresponding timestamps and values share an index
	Values []string
}

func Constructor() TimeMap {
	return TimeMap{make(map[string]*KVTime)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if val, ok := this.KVs[key]; ok {
		// If key already exists, append the values
		val.Timestamps = append(val.Timestamps, timestamp)
		val.Values = append(val.Values, value)
	} else {
		// If key doesn't exist, create new
		this.KVs[key] = &KVTime{
			Timestamps: []int{timestamp},
			Values:     []string{value},
		}
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if val, ok := this.KVs[key]; ok {
		// If key found, we search for the closest but not over timestamp
		index := searchInsert(val.Timestamps, timestamp)
		// Edge cases when found number is at the start or beyond the array
		if index == len(val.Timestamps) || val.Timestamps[index] != timestamp {
			index--
			// If there is no smaller timestamp, return ""
			if index == -1 {
				return ""
			}
		}
		// Return corresponding value
		return val.Values[index]
	} else {
		return ""
	}
}

// Binary insert search, returns the index of the value
// as if it was to be added to the array
func searchInsert(nums []int, target int) int {
	var res, start, end = -1, 0, len(nums) - 1
	for end >= start {
		mid := ((end - start) / 2) + start
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			end = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid + 1
			}
			start = mid + 1
		}
	}
	return res
}

package problem2353

import "container/heap"

/*
Design a food rating system that can do the following:

Modify the rating of a food item listed in the system.
Return the highest-rated food item for a type of cuisine in the system.
Implement the FoodRatings class:

FoodRatings(String[] foods, String[] cuisines, int[] ratings) Initializes the system.
The food items are described by foods, cuisines and ratings, all of which have a length of n.
foods[i] is the name of the ith food,
cuisines[i] is the type of cuisine of the ith food, and
ratings[i] is the initial rating of the ith food.
void changeRating(String food, int newRating) Changes the rating of the food item with the name food.
String highestRated(String cuisine) Returns the name of the food item that has the highest rating for the given type of cuisine.
If there is a tie, return the item with the lexicographically smaller name.
*/

type FoodRatings struct {
	FoodToRate map[string]int    // Map containing rating of the actual foods (Only real foods)
	FoodToType map[string]string // Helper map to find cuisine of food
	MaxRated   map[string]*Heap  // Map of all foods/ratings by cuisine (Including old foods that changed)
}

type Food struct {
	Name   string
	Rating int
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	var ret = FoodRatings{
		FoodToRate: make(map[string]int),
		FoodToType: make(map[string]string),
		MaxRated:   make(map[string]*Heap),
	}

	// Building our maps and heaps
	for i := range foods {
		ret.FoodToRate[foods[i]] = ratings[i]
		ret.FoodToType[foods[i]] = cuisines[i]

		// Make heap for each cuisine type
		if exists := ret.MaxRated[cuisines[i]]; exists == nil {
			ret.MaxRated[cuisines[i]] = MakeHeap(MaxHeap)
		}
		heap.Push(ret.MaxRated[cuisines[i]], Food{
			Name:   foods[i],
			Rating: ratings[i],
		})
	}

	return ret
}

func (fr *FoodRatings) ChangeRating(food string, newRating int) {
	var c = fr.FoodToType[food]

	// Update the actual value
	fr.FoodToRate[food] = newRating

	// Add to heap
	heap.Push(fr.MaxRated[c], Food{
		Name:   food,
		Rating: newRating,
	})
}

func (fr *FoodRatings) HighestRated(cuisine string) string {
	var cuisineHeap = fr.MaxRated[cuisine]

	// Loop until you find the first real food
	for cuisineHeap.Len() > 0 {
		var food = cuisineHeap.Peek()
		// If the rating doesn't match the FoodToRate, it's an old food that changed
		if food.Rating == fr.FoodToRate[food.Name] {
			return food.Name
		}
		heap.Pop(cuisineHeap)
	}
	return ""
}

// Heap implementation
type Heap struct {
	Values   []Food
	LessFunc func(Food, Food) bool
}

func (h *Heap) Less(i, j int) bool { return h.LessFunc(h.Values[i], h.Values[j]) }
func (h *Heap) Swap(i, j int)      { h.Values[i], h.Values[j] = h.Values[j], h.Values[i] }
func (h *Heap) Len() int           { return len(h.Values) }
func (h *Heap) Peek() Food         { return h.Values[0] }
func (h *Heap) Pop() (v interface{}) {
	h.Values, v = h.Values[:h.Len()-1], h.Values[h.Len()-1]
	return v
}
func (h *Heap) Push(v interface{}) { h.Values = append(h.Values, v.(Food)) }

func MakeHeap(less func(Food, Food) bool) *Heap {
	return &Heap{LessFunc: less}
}

func MaxHeap(i, j Food) bool {
	if i.Rating == j.Rating {
		return i.Name < j.Name
	}
	return i.Rating > j.Rating
}
func MinHeap(i, j Food) bool {
	if i.Rating == j.Rating {
		return i.Name < j.Name
	}
	return i.Rating < j.Rating
}

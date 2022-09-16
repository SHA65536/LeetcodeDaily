package problem0307

type NumArray struct {
	Original  []int
	SumPrefix []int
}

func Constructor(nums []int) NumArray {
	var sum int
	res := NumArray{Original: nums, SumPrefix: make([]int, len(nums))}
	for idx, val := range nums {
		sum += val
		res.SumPrefix[idx] = sum
	}
	return res
}

func (this *NumArray) Update(index int, val int) {
	var value = val - this.Original[index]
	this.Original[index] = val
	for i := index; i < len(this.SumPrefix); i++ {
		this.SumPrefix[i] += value
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.SumPrefix[right]
	}
	return this.SumPrefix[right] - this.SumPrefix[left-1]
}

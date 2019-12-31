package problem0001


func twoSumBruteForceOptimzed(nums []int, target int) []int {
    for i:= 0; i < len(nums); i++ {
		for j:= 1; j < len(nums); j++ {
			if i == j {
				break
			}
			if nums[j]+nums[i] == target {
				return []int{j, i}
			}
		}
	}
	return []int{}
}

func twoSumHashMap(nums []int, target int) []int {
    m := make(map[int][]int)
    for i:= 0; i < len(nums); i++ {
		m[nums[i]] = append(m[nums[i]], i)
		if complimentary_indicies, ok := m[target - nums[i]]; ok{
			for _, index:= range complimentary_indicies {
				if index != i {
					return []int{index, i}
				}
			}
		}
	}
	return []int{}
}

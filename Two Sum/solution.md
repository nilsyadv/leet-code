## Intuition
<!-- Describe your first thoughts on how to solve this problem. -->

## Approach
<!-- Describe your approach to solving the problem. -->

## Complexity
- Time complexity:
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity:
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

## Code
```go
func twoSum(nums []int, target int) []int {
    for index, firstVal := range nums {
        for secIndex := index + 1; secIndex < len(nums);secIndex++ {
            if firstVal + nums[secIndex] == target {
                return []int{index,secIndex}
            }
        }
    }
    return []int{}
}
```
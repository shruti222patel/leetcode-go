# Two Sum

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

## Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

## Related Topics:
- Array
- Hashtable

## Algo
### Brute Force - Simple for-loop
- 2 loops (i & j), one for each of the addends
- if i == j, continue inner loop
- not optimial, as you only need to process half of the total sums this brute force method will calculate to find there is no answer

*Big-O* n^2

### Brute Force - optimized
- 2 loops (i & j), one for each of the addends
- if i == j, break inner loop
- start j with index 1 as we can't return the same index to sum to the target
- this should cover all the sums once rather than twice like the previous method
*Big-O* (n^2)/2
## Complexity
- Time: O(n^2) - we have a loop within a loop
- Space: O(1) -  we're not saving anything
## Results
- Runtime: 20 ms, faster than 35.66% of Go online submissions for Two Sum.
- Memory Usage: 3 MB, less than 100.00% of Go online submissions for Two Sum.


### Change Array Shape
- 1 loops (i) over the nums indexes
- add {nums[i], i} as map entry
- look to see if map has target - nums[i] val & it's not nums[i
## Complexity
- Time: O(n) - we only loop over nums once
- Space: O(n) - we are temporarily storing data
## Results
- Runtime: 4 ms, faster than 95.00% of Go online submissions for Two Sum.
- Memory Usage: 4.6 MB, less than 7.69% of Go online submissions for Two Sum.

## Takeaway
- to optimize looping over arrays, consider saving the loop as an inverted array via hash map; meaning, save the value of the array as the map key and the index of the array as map value
    - beware of different array indicies having the same value. (solved by having a list of indicies as the map value)

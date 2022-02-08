//1. Two Sum
//Easy
//
//28970
//
//926
//
//Add to List
//
//Share
//Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//You can return the answer in any order.
//
// 
//
//Example 1:
//
//Input: nums = [2,7,11,15], target = 9
//Output: [0,1]
//Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
//Example 2:
//
//Input: nums = [3,2,4], target = 6
//Output: [1,2]
//Example 3:
//
//Input: nums = [3,3], target = 6
//Output: [0,1]
// 
//
//Constraints:
//
//2 <= nums.length <= 104
//-109 <= nums[i] <= 109
//-109 <= target <= 109
//Only one valid answer exists.
//
//Runtime: 4 ms, faster than 97.13% of Go online submissions for Two Sum.
//Memory Usage: 4.2 MB, less than 61.10% of Go online submissions for Two Sum.
//
//

func twoSum(nums []int, target int) []int {
    for i:=0; i<len(nums); i+=1 {
        for j:=i+1;j<len(nums);j+=1 {
            if nums[i]+nums[j]==target {
                return []int{i,j}
            }
        }
    }
    return []int{0,0}
}


func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i,val:=range(nums) {
        j,check:=m[target-val]
        if check {
            return[]int{j,i}   
        }
        m[nums[i]]=i
    }
    return []int{}
}

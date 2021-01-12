# 二分查找问题
**二分搜索，思想很简单，细节是魔鬼。让我们搞定这个魔鬼**<br>
参考  [leetcode二分查找专题](https://leetcode-cn.com/explore/learn/card/binary-search)
## 问题
对于已排序的数组nums，查找到给定元素target，返回其索引，如果不存在返回-1
## 三种模板
### 模板一：查找单个索引
```
func binarySearch(nums []int, target) int {
    left, right := 0, len(nums)-1
    for left <= right {
        // prevent (right + left) overflow
        mid := left + (right - left) / 2
        swith {
        case nums[mid] == target:
            return mid
        case nums[mid] < target:
            left = mid + 1
        default:
            right = mid - 1
        }
    }
    // end condition: left > right
    return -1
}
```
关键属性
* 二分查找的最基础和最基本的形式。
* 查找条件可以在不与元素的两侧进行比较的情况下确定（或使用它周围的特定元素）。
* 不需要后处理，因为每一步中，你都在检查是否找到了元素。如果到达末尾，则知道未找到该元素。
### 模板二：查找单个索引及其右邻居
```
func binarySearch(nums []int, target) int {
    left, right := 0, len(nums)
    for left < right {
        // prevent (right + left) overflow
        mid := left + (right - left) / 2
        swith {
        case nums[mid] == target:
            return mid
        case nums[mid] < target:
            left = mid + 1
        default:
            right = mid
        }
    }
    // end condithon:  left == right
    if left < len(nums) && nums[left] == target {
        return left
    }
    return -1
}
```
关键属性
* 一种实现二分查找的高级方法。
* 查找条件需要访问元素的直接右邻居。
* 使用元素的右邻居来确定是否满足条件，并决定是向左还是向右。
* 保证查找空间在每一步中至少有 2 个元素。
* 需要进行后处理。 当你剩下 1 个元素时，循环 / 递归结束。 需要评估剩余元素是否符合条件。
### 模板三： 查找单个索引及其左右邻居
```
func binarySearch(nums []int, target) int {
    if len(nums) == 0 {
        return -1
    }
    left, right := 0, len(nums)-1
    for left+1 < right {
        // prevent (right + left) overflow
        mid := left + (right - left) / 2
        swith {
        case nums[mid] == target:
            return mid
        case nums[mid] < target:
            left = mid
        default:
            right = mid
        }        
    }
    // end condithon:  left + 1 == right
    if nums[left] == target {
        return left
    }
    if nums[right] == target {
        return right
    }
    return -1
}
```
关键属性
* 实现二分查找的另一种方法。
* 搜索条件需要访问元素的直接左右邻居。
* 使用元素的邻居来确定它是向右还是向左。
* 保证查找空间在每个步骤中至少有 3 个元素。
* 需要进行后处理。 当剩下 2 个元素时，循环 / 递归结束。 需要评估其余元素是否符合条件。
### 如果有多个target，如何找到左右target元素的索引
参考[34. 在排序数组中查找元素的第一个和最后一个位置](../solutions/find-first-and-last-position-of-element-in-sorted-array/d.go)
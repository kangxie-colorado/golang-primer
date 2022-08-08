def twoSum(nums, target):
    i, j = 0, len(nums)-1
    res = []

    while i < j:
        S = nums[i]+nums[j]
        if S == target:
            res.append([nums[i], nums[j]])
            while i < j and nums[i] == nums[i+1]:
                i += 1
            while i < j and nums[j] == nums[j-1]:
                j -= 1
            i += 1
            j -= 1

        elif S < target:
            i += 1
        else:
            j -= 1
    return res


nums = [1, 1, 2, 2, 3, 4, 5, 5, 6]
print(twoSum(nums, 700))

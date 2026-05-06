def find_max_consecutive_ones(nums):
    count = 0
    maxFound = 0
    for i in range(len(nums)):
        if nums[i] != 0 and nums[i] == 1:
            count += 1
        else:
            count = 0
        maxFound = max(count, maxFound)
    return maxFound

def main():
    print(find_max_consecutive_ones([1,1,0,1,1,1]))

if __name__ == "__main__":
    main()

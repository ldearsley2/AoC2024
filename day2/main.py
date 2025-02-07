def part1(file):
    safe_count = 0
    with open(file) as f:
        lines = f.read().strip().split('\n')

    def report_safe(nums):
        inc = nums[1] > nums[0]
        if inc:
            for i in range(1, len(nums)):
                diff = nums[i] - nums[i-1]
                if not 1 <= diff <= 3:
                    return False
            return True
        else:
            for i in range(1, len(nums)):
                diff = nums[i] - nums[i-1]
                if not -3 <= diff <= -1:
                    return False
            return True

    for line in lines:
        nums = [int(i) for i in line.split()]
        safe_count += report_safe(nums)

    print(safe_count)




part1("input.txt")
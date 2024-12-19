def part1(file):
    left_list, right_list = [], []
    with open(file) as f:
        lines = f.readlines()

    for line in lines:
        split_line = line.split()
        left_list.append(int(split_line[0]))
        right_list.append(int(split_line[1]))

    left_list.sort()
    right_list.sort()

    total = 0
    for i in range(len(left_list)):
        total += abs(left_list[i] - right_list[i])
    return total


def part2(file):
    left_list, right_list = [], []
    with open(file) as f:
        lines = f.readlines()

    for line in lines:
        split_line = line.split()
        left_list.append(int(split_line[0]))
        right_list.append(int(split_line[1]))

    total = 0
    for i in left_list:
        occurrences = 0
        for j in right_list:
            if i == j:
                occurrences += 1
        total += i * occurrences

    return total


print(part2("input.txt"))
from copy import deepcopy


def get_column(table, column):
    return [table[l][column] for l in range(9)]


def get_line(table, line):
    return table[line][:]


def get_block(table, ib, jb):
    res = []
    for l in range(3):
        for c in range(3):
            res.append(table[ib * 3 + l][jb * 3 + c])
    return res


def sudoku(puzzle):
    puzzle, is_changed, n = deepcopy(puzzle), True, 9
    while is_changed:
        is_changed = False
        for i in range(n):
            for j in range(n):
                if puzzle[i][j] != 0:
                    continue
                avs = set(range(1, 10))
                avs -= set(get_column(puzzle, j) + get_line(puzzle, i) + get_block(puzzle, ib=i // 3, jb=j // 3))
                if len(avs) == 1:
                    is_changed, puzzle[i][j] = True, avs.pop()

    return puzzle

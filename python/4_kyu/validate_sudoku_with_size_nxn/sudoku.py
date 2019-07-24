from math import sqrt


class Sudoku(object):
    def __init__(self, data):
        self.data = data

    def is_valid(self):
        n = len(self.data)
        lsn = int(sqrt(n))
        if sqrt(n) - lsn != 0:
            return False

        def get_column(cn):
            for i in range(n):
                yield self.data[i][cn]

        def get_square(sn):
            for i in range(sn % lsn * lsn, sn % lsn * lsn + lsn):
                for j in range(sn // lsn * lsn, sn // lsn * lsn + lsn):
                    yield self.data[i][j]

        for i in range(n):
            for d in self.data[i], list(get_column(i)), list(get_square(i)):
                if len(self.data[i]) is not n or list(filter(
                        lambda x: type(x) is not int or not 1 <= x <= n or d.count(x) > 1, d)):
                    return False
        return True

import itertools


def permutations(string):
    return [''.join(x) for x in set(itertools.permutations(string))]

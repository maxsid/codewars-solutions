import re


def count_smileys(arr):
    count = 0
    for a in arr:
        count += bool(re.fullmatch(r'(:|;){1}(\-|\~)?(\)|D){1}', a))
    return count

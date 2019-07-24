import unittest
import re


def parse_int(string):
    in_words = {'zero': 0, 'one': 1, 'two': 2, 'three': 3, 'four': 4, 'five': 5, 'six': 6, 'seven': 7,
                'eight': 8, 'nine': 9, 'ten': 10, 'eleven': 11, 'twelve': 12, 'thirteen': 13,
                'fourteen': 14, 'fifteen': 15, 'sixteen': 16, 'seventeen': 17, 'eighteen': 18,
                'nineteen': 19, 'twenty': 20, 'thirty': 30, 'forty': 40, 'fifty': 50,
                'sixty': 60, 'seventy': 70, 'eighty': 80, 'ninety': 90}

    big_in_words = {'million': 10 ** 6, 'thousand': 10 ** 3, 'hundred': 10 ** 2}
    last_int, number = 0, 0
    for wi in map(lambda w: in_words.get(w, w), re.findall(r'\w+', string)):
        if wi == 'and':
            continue
        if isinstance(wi, int):
            last_int += wi
        else:
            last_int *= big_in_words[wi]
            if last_int >= 1000:
                number += last_int
                last_int = 0
    return number + last_int


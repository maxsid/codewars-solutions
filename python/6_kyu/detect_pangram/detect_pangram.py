import string

def is_pangram(s):
    s = s.lower()
    try:
        next(filter(lambda x: x not in s, string.ascii_lowercase))
        return False
    except StopIteration:
        return True
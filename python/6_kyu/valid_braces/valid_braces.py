def validBraces(string):
    brackets = {'(': ')', '[': ']', '{': '}'}
    opened = []
    for s in string:
        if s in brackets:
            opened.append(s)
        elif not opened or brackets[opened.pop()] != s:
            return False
    return not opened

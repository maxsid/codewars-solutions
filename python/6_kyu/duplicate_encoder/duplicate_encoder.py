def duplicate_encode(word):
    word = word.lower()
    return ''.join(map(lambda l: '(' if word.count(l) == 1 else ')', word))

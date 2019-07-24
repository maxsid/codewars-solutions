import re

def generate_hashtag(s):
    if not s: return False
    hashtag = '#' + ''.join([w.capitalize() for w in re.findall(r'\w+', s)])
    if len(hashtag) > 140: return False
    return hashtag
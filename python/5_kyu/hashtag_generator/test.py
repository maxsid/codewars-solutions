import unittest
from .hashtag_generator import generate_hashtag


class MyTestCase(unittest.TestCase):
    def test_something(self):
        self.assertEqual(generate_hashtag(''), False, 'Expected an empty string to return False')
        self.assertEqual(generate_hashtag('Do We have A Hashtag')[0], '#', 'Expeted a Hashtag (#) at the beginning.')
        self.assertEqual(generate_hashtag('Codewars'), '#Codewars', 'Should handle a single word.')
        self.assertEqual(generate_hashtag('Codewars      '), '#Codewars', 'Should handle trailing whitespace.')
        self.assertEqual(generate_hashtag('Codewars Is Nice'), '#CodewarsIsNice', 'Should remove spaces.')
        self.assertEqual(generate_hashtag('codewars is nice'), '#CodewarsIsNice',
                           'Should capitalize first letters of words.')
        self.assertEqual(generate_hashtag('CodeWars is nice'), '#CodewarsIsNice',
                           'Should capitalize all letters of words - all lower case but the first.')
        self.assertEqual(generate_hashtag('c i n'), '#CIN',
                           'Should capitalize first letters of words even when single letters.')
        self.assertEqual(generate_hashtag('codewars  is  nice'), '#CodewarsIsNice',
                           'Should deal with unnecessary middle spaces.')
        self.assertEqual(generate_hashtag(
            'Looooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooong Cat'),
                           False, 'Should return False if the final word is longer than 140 chars.')


if __name__ == '__main__':
    unittest.main()

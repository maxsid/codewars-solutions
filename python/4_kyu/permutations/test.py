import unittest
from .permutations import permutations


class PermutationsUnittest(unittest.TestCase):
    def test_first(self):
        self.assertEqual(permutations('a'), ['a'])

    def test_second(self):
        p = sorted(permutations('ab'))
        self.assertEqual(p, ['ab', 'ba'])

    def test_third(self):
        p = sorted(permutations('aabb'))
        self.assertEqual(p, ['aabb', 'abab', 'abba', 'baab', 'baba', 'bbaa'])


if __name__ == '__main__':
    unittest.main()

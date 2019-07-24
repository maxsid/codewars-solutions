import unittest
from .parseint_reloaded import parse_int


class ParseIntReloadedUnittest(unittest.TestCase):
    def test_first(self):
        self.assertEqual(parse_int('one'), 1)

    def test_second(self):
        self.assertEqual(parse_int('twenty'), 20)

    def test_third(self):
        self.assertEqual(parse_int('two hundred forty-six'), 246)

    def test_zero(self):
        self.assertEqual(parse_int('zero'), 0)

    def test_million(self):
        self.assertEqual(parse_int('one million'), 1000000)


if __name__ == '__main__':
    unittest.main()

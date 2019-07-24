import unittest
from .detect_pangram import is_pangram


class KataTestCase(unittest.TestCase):
    def test_sample(self):
        pangram = "The quick, brown fox jumps over the lazy dog!"
        self.assertTrue(is_pangram(pangram))


if __name__ == '__main__':
    unittest.main()

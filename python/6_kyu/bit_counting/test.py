import unittest
from .bit_counting import countBits


class KataTestCase(unittest.TestCase):
    def test_sample(self):
        self.assertEqual(countBits(0), 0)
        self.assertEqual(countBits(4), 1)
        self.assertEqual(countBits(7), 3)
        self.assertEqual(countBits(9), 2)
        self.assertEqual(countBits(10), 2)


if __name__ == '__main__':
    unittest.main()

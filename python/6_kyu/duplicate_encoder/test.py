import unittest
from .duplicate_encoder import duplicate_encode


class KataTestCase(unittest.TestCase):
    def test_something(self):
        self.assertEqual(duplicate_encode("din"), "(((")
        self.assertEqual(duplicate_encode("recede"), "()()()")
        self.assertEqual(duplicate_encode("Success"), ")())())", "should ignore case")
        self.assertEqual(duplicate_encode("(( @"), "))((")


if __name__ == '__main__':
    unittest.main()

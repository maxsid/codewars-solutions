import unittest
from .valid_braces import validBraces


class KotaTest(unittest.TestCase):
    def test_sample(self):
        self.assertTrue(validBraces(r"(){}[]"))
        self.assertTrue(validBraces(r"([{}])"))
        self.assertFalse(validBraces(r"(}"))
        self.assertFalse(validBraces(r"[(])"))
        self.assertFalse(validBraces(r"[({})](]"))


if __name__ == '__main__':
    unittest.main()

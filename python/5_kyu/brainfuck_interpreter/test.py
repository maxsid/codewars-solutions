import unittest
from .brainfuck_interpreter import brain_luck


class KataTestCase(unittest.TestCase):
    def test_something(self):
        # Echo until byte(255) encountered
        self.assertEqual(
            brain_luck(',+[-.,+]', 'Codewars' + chr(255)),
            'Codewars'
        )

        # Echo until byte(0) encountered
        self.assertEqual(
            brain_luck(',[.[-],]', 'Codewars' + chr(0)),
            'Codewars'
        )

        # Two numbers multiplier
        self.assertEqual(
            brain_luck(',>,<[>[->+>+<<]>>[-<<+>>]<<<-]>>.', chr(8) + chr(9)),
            chr(72)
        )


if __name__ == '__main__':
    unittest.main()

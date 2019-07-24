import unittest
from .pick_peaks import pick_peaks


class MyTestCase(unittest.TestCase):
    def test_something(self):
        self.assertEqual(pick_peaks([1, 2, 3, 6, 4, 1, 2, 3, 2, 1]), {"pos": [3, 7], "peaks": [6, 3]},
                         'should support finding peaks')
        self.assertEqual(pick_peaks([3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3]), {"pos": [3, 7], "peaks": [6, 3]},
                         'should support finding peaks, but should ignore peaks on the edge of the array')
        self.assertEqual(pick_peaks([3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1]),
                         {"pos": [3, 7, 10], "peaks": [6, 3, 2]},
                         'should support finding peaks; if the peak is a plateau, it should only return '
                         'the position of the first element of the plateau')
        self.assertEqual(pick_peaks([2, 1, 3, 1, 2, 2, 2, 2, 1]), {"pos": [2, 4], "peaks": [3, 2]},
                         'should support finding peaks; if the peak is a plateau, it should only return '
                         'the position of the first element of the plateau')
        self.assertEqual(pick_peaks([2, 1, 3, 1, 2, 2, 2, 2]), {"pos": [2], "peaks": [3]},
                         'should support finding peaks, but should ignore peaks on the edge of the array')


if __name__ == '__main__':
    unittest.main()

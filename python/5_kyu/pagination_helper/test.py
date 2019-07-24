import unittest
from .pagination_helper import PaginationHelper


class MyTestCase(unittest.TestCase):
    def test_something(self):
        collection = range(1, 25)
        helper = PaginationHelper(collection, 10)

        self.assertEqual(helper.page_count(), 3, 'page_count is returning incorrect value.')
        self.assertEqual(helper.page_index(23), 2, 'page_index returned incorrect value')
        self.assertEqual(helper.item_count(), 24, 'item_count returned incorrect value')


if __name__ == '__main__':
    unittest.main()

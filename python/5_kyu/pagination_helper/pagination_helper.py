class PaginationHelper:
    def __init__(self, collection, items_per_page):
        self._count = len(collection)
        self._items_per_page = items_per_page

    def item_count(self):
        return self._count

    def page_count(self):
        return self._count / self._items_per_page + bool(self._count % self._items_per_page)

    def page_item_count(self, page_index):
        if page_index < 0 or page_index >= self.page_count(): return -1
        if page_index == self.page_count() - 1: return self.item_count() % self._items_per_page
        return self._items_per_page

    def page_index(self, item_index):
        if item_index < 0 or item_index >= self.item_count(): return -1
        return item_index % self.item_count() / self._items_per_page


collection = range(1,25)
helper = PaginationHelper(collection, 10)

print(helper.page_count() == 3)
print(helper.page_index(23) == 2)
print(helper.item_count() == 24)
print(helper.page_item_count(-42) == -1)
print(helper.page_item_count(300) == -1)
print(helper.page_item_count(3))
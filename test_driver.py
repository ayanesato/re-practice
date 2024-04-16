import unittest
from isleapyear import is_leap_year
import unittest
class TestIsLeapYear(unittest.TestCase):
    def test_leap_year(self):
        # うるう年の場合
        self.assertEqual(is_leap_year(2000), True)
        self.assertEqual(is_leap_year(2024), True)
        self.assertEqual(is_leap_year(2020), True)

    def test_non_leap_year(self):
        # うるう年でない場合 同地分割
        self.assertEqual(is_leap_year(1999), False)
        self.assertEqual(is_leap_year(2001), False)
        self.assertEqual(is_leap_year(2023), False)
        self.assertEqual(is_leap_year(2025), False)
        self.assertEqual(is_leap_year(2019), False)
        self.assertEqual(is_leap_year(2021), False)

if __name__ == '__main__':
    unittest.main()

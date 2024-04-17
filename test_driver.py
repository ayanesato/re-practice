from isleapyear import is_leap_year
import unittest
class TestIsLeapYear(unittest.TestCase):
    def test_leap_year(self):
        #4で割り切れるが100で割り切れない場合とその前後
        self.assertEqual(is_leap_year(2019), False)
        self.assertEqual(is_leap_year(2020), True)
        self.assertEqual(is_leap_year(2021), False)

        #4と100で割り切れる場合(平年)とその前後
        self.assertEqual(is_leap_year(1899), False)
        self.assertEqual(is_leap_year(1900), False)
        self.assertEqual(is_leap_year(1901), False)

        #4でも100でも400でも割り切れる場合とその前後
        self.assertEqual(is_leap_year(1999), False)
        self.assertEqual(is_leap_year(2000), True)
        self.assertEqual(is_leap_year(2001), False)


if __name__ == '__main__':
    unittest.main()

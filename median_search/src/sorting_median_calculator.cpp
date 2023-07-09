#include <algorithm>
#include "sorting_median_calculator.hpp"

SortingMedianCalculator::SortingMedianCalculator() {}

SortingMedianCalculator::~SortingMedianCalculator() {}

float
SortingMedianCalculator::median( std::vector<float> & set ) {
  auto n = set.size();
  auto i1 = n / 2, i2 = (n - 1) / 2;
  auto sorted_set(set);
  std::sort(sorted_set.begin(), sorted_set.end());
  auto v1 = sorted_set[i1], v2 = sorted_set[i2];
  return (v1 + v2) / 2.f;
}
#include "quick_select_median_calculator.hpp"

#include <algorithm>

QuickSelectMedianCalculator::QuickSelectMedianCalculator() {}
QuickSelectMedianCalculator::~QuickSelectMedianCalculator() {}

// Return the median number of a float-type number set.
// It is based on quick select method, find the n/2-th number and 
// (n+1)/2-th number for a set with n number and the median is mean
// of them.
float QuickSelectMedianCalculator::median(std::vector<float> &set)
{

  int n = set.size();
  int i1 = (n + 1) / 2, i2 = (n + 2) / 2;

  float v1 = quickSelect(set, i1);
  float v2 = quickSelect(set, i2);

  return (v1 + v2) / 2.f;
}

// Return the k-th smallest number of a given float-type number set.
// Be careful that k should not out of range.
float QuickSelectMedianCalculator::quickSelect(std::vector<float> &set, int k)
{
  int n = set.size();
  float pivot = set[set.size() / 2];

  // A vector used to keep numbers which are smaller than the pivot
  auto less_than_pivot = std::vector<float>(); 

  // A vector used to keep numbers which are greater than the pivot
  auto more_than_pivot = std::vector<float>(); 

  // A vector used to keep numbers which are equal to the pivot
  auto pivots = std::vector<float>();         
  
  float curr_number;

  // Map each number to its vector.
  for (int i = 0; i < n; i++)
  {
    curr_number = set[i];
    if (curr_number < pivot)
    {
      less_than_pivot.push_back(curr_number);
    }
    else if (curr_number > pivot)
    {
      more_than_pivot.push_back(curr_number);
    }
    else
    {
      pivots.push_back(curr_number);
    }
  }

  // Decide which vector contains the k-th smallest number and 
  // recursively find k-th number or return the valur of k-th number.
  if (less_than_pivot.size() >= k)
  {
    return quickSelect(less_than_pivot, k);
  }
  else if (less_than_pivot.size() + pivots.size() >= k)
  {
    return pivot;
  }
  else
  {
    return quickSelect(more_than_pivot, k - less_than_pivot.size() - pivots.size());
  }
}


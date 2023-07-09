#include "randomized_median_calculator.hpp"

#include <algorithm>
#include <random>

#include <cmath>

RandomizedMedianCalculator::RandomizedMedianCalculator() {}
RandomizedMedianCalculator::~RandomizedMedianCalculator() {}

// Return median of a float number vector using randomized algorithm
float RandomizedMedianCalculator::median(std::vector<float> &set)
{

  // Randomly collect number from input set with replacement 
  // to form a new set with size ceil(n^(3/4))
  int size_set_s = set.size();
  float var_n = pow(size_set_s, 3.f / 4.f);
  int size = ceil(var_n);
  
  std::random_device random;    // Set random engine region
  std::default_random_engine random_engine(random());
  std::uniform_int_distribution<int> uniform_random(0, size_set_s - 1);   // Set random range

  int random_num;
  auto subset_R = std::vector<float>();
  for (int i = 0; i < size; i++)
  {
    random_num = uniform_random(random);
    subset_R.push_back(set[random_num]);
  }

  // Sort this new set and then find two specific number of the sorted set
  auto sorted_subset_R(subset_R);
  std::sort(sorted_subset_R.begin(), sorted_subset_R.end());
  auto index_d = floor((1.f / 2.f) * var_n - sqrt(size_set_s)) - 1.f;
  auto index_u = ceil((1.f / 2.f) * var_n + sqrt(size_set_s)) - 1.f;

  
  if (index_d < 0)    // Ensure that index not out of range
    index_d = 0;
  if (index_u >= size)    // Ensure that index not out of range
    index_u = size - 1;
  auto element_d = sorted_subset_R[index_d];
  auto element_u = sorted_subset_R[index_u];

  // Collect all number in range [element_d, emelent_u] from input set S to 
  // a new set C.
  auto set_c = std::vector<float>();
  int ld = 0;   // used to count numbers smaller than the range
  int lu = 0;   // used to count numbers greater than the range 
  for (auto num : set)
  {
    if (num < element_d)
    {
      ld++;
    }
    else if (num >= element_d && num <= element_u)
    {
      set_c.push_back(num);
    }
    else
    {
      lu++;
    }
  }

  // If ld > n/2 or lu > n/2, median is not in set C. Repeat this function until success.
  // If size of set C > 4n^(3/4), set C can't be sorted in 
  // linear time. Repeat this function until success.
  // Else sort set C and return needed value. 
  if (ld >= (size_set_s / 2) || lu >= (size_set_s / 2))
  {
    return median(set);
  }
  else if (set_c.size() > (4 * var_n))
  {
    return median(set);
  }
  else
  {
    int i1 = floor((size_set_s + 1) / 2.f) - ld -1;   // Be careful that index should minus 1.
    int i2 = floor((size_set_s + 2) / 2.f) - ld -1;   // Be careful that index should minus 1.

    std::sort(set_c.begin(), set_c.end());
    return (set_c[i1] + set_c[i2]) / 2.f;
  }
}

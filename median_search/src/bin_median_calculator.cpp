#include "bin_median_calculator.hpp"

#include <numeric>
#include <algorithm>

#include <cmath>

BinMedianCalculator::BinMedianCalculator() {}
BinMedianCalculator::~BinMedianCalculator() {}

// Return the median of a float-type number vector using bin median method.
float BinMedianCalculator::median(std::vector<float> &set)
{
  int n = set.size();
  int i1 = (n + 1) / 2;
  int i2 = (n + 2) / 2;

  // Calculate the mean of this number set
  float sum = std::accumulate(set.begin(), set.end(), 0.f);
  float mean = sum / n;

  // Calculate the variance of this number set
  float temp_accumulate = 0.f;
  for (auto num : set)
  {
    temp_accumulate += (num - mean) * (num - mean);
  }
  float variance = temp_accumulate / n;
  float sigma = sqrt(variance);

  // When variance is 0, all number are same, median is anyone of them
  if (sigma == 0)
  {
    return mean;
  }

  // Find the region of median bin, median is also in this bin
  float left_band = mean - sigma;
  float right_band = mean + sigma;

  // Collect all numbers in the median bin region,
  // and count how many number smaller than the median bin.
  int smaller_than_left_band = 0;
  auto subset_bins = std::vector<float>();
  for (auto num : set)
  {
    if (num < left_band)
    {
      smaller_than_left_band++;
    }
    else if (num <= right_band)
    {
      subset_bins.push_back(num);
    }
  }

  // Call findNthNumberWithBand to find the i1-th number and i2-th number of the total set,
  // calculate and return its median.
  if (i1 == i2)
  {
    float val = findNthNumberWithBand(subset_bins, i1 - smaller_than_left_band, left_band, right_band);
    return val;
  }
  else
  {
    float val_i1 = findNthNumberWithBand(subset_bins, i1 - smaller_than_left_band, left_band, right_band);
    float val_i2 = findNthNumberWithBand(subset_bins, i2 - smaller_than_left_band, left_band, right_band);
    return (val_i1 + val_i2) / 2.f;
  }
}

// Return k-th smallest number of a float number set with value interval [left_band, right_band]
float BinMedianCalculator::findNthNumberWithBand(std::vector<float> &set, int k, float left_band, float right_band)
{
  int bin_number;

  // K decides the number of bins of the interval
  int K = 100;
  float bin_length = (right_band - left_band) / K;

  // Set a vector to record how many number each bin contains
  auto bin_number_counts = std::vector<int>(K, 0);

  // Count each bin's number by a counter
  for (auto num : set)
  {
    bin_number = (int)((num - left_band) / bin_length);
    bin_number_counts[bin_number]++;
  }

  int count = 0;
  int bin_of_kth;

  // Calculate which bin contions kth number, and find new k respect to this bin  
  for (int i = 0; i < K; i++)
  {
    count += bin_number_counts[i];
    if (count >= k)
    {
      k = k - (count - bin_number_counts[i]);
      bin_of_kth = i;
      break;
    }
  }

  // Calculate new interval region of the subbin 
  float tight_left_band = left_band + bin_of_kth * bin_length;
  float tight_right_band = tight_left_band + bin_length;

  int left_number_count = 0;
  auto bin_set = std::vector<float>();
  
  // Collect numbers in the interval
  for (auto num : set)
  {
    if (num < tight_left_band)
    {
      left_number_count++;
    }
    else if (num <= tight_right_band)
    {
      bin_set.push_back(num);
    }
  }

  // If not many numbers, sort it and find the result.
  // Else recursively find the result with new bin, k and interval
  if (bin_set.size() < 10)
  {
    std::sort(bin_set.begin(), bin_set.end());

    return bin_set[k - 1];
  }
  else
  {
    return findNthNumberWithBand(bin_set, k, tight_left_band, tight_right_band);
  }
}

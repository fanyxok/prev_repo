#ifndef CS133_SORTINGMEDIANCALCULATOR_HPP_
#define CS133_SORTINGMEDIANCALCULATOR_HPP_

#include "median_calculator_base.hpp"


class SortingMedianCalculator : public MedianCalculatorBase {
 public:
  SortingMedianCalculator();
  virtual ~SortingMedianCalculator();

  virtual float median( std::vector<float> & set );
};

#endif
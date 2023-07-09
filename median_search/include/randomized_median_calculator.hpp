#ifndef CS133_RANDOMIZEDMEDIANCALCULATOR_HPP_
#define CS133_RANDOMIZEDMEDIANCALCULATOR_HPP_

#include "median_calculator_base.hpp"

class RandomizedMedianCalculator : public MedianCalculatorBase {
  public:
  RandomizedMedianCalculator();
  virtual ~RandomizedMedianCalculator();

  virtual float median( std::vector<float> &set );
};

#endif 
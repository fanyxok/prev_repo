#ifndef CS133_BINMEDIANCALCULATOR_HPP_
#define CS133_BINMEDIANCALCULATOR_HPP_

#include "median_calculator_base.hpp"

class BinMedianCalculator: public MedianCalculatorBase{
  public: 
    BinMedianCalculator();
    virtual ~BinMedianCalculator();

    virtual float median( std::vector<float> &set );

    virtual float findNthNumberWithBand( std::vector<float> &set, int k, float left_band, float right_band);
};



#endif
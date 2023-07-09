#ifndef CS133_MEDIANCALCULATORBASE_HPP_
#define CS133_MEDIANCALCULATORBASE_HPP_

#include <vector>

class MedianCalculatorBase {
 public:
  MedianCalculatorBase();
  virtual ~MedianCalculatorBase();

  virtual float median( std::vector<float> & set ) = 0;
};

#endif
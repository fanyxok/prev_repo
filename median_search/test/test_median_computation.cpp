#include <stdlib.h>
#include <sstream>
#include <iostream>
#include <fstream>
#include <string>

#include "sorting_median_calculator.hpp"
#include "quick_select_median_calculator.hpp"
#include "bin_median_calculator.hpp"
#include "randomized_median_calculator.hpp"

int main(int argc, char **argv) {
  // Get index of median calculator
  std::stringstream ss;
  ss << std::string(argv[1]);
  int calculator;
  ss >> calculator;

  // initialize calculator
  MedianCalculatorBase *calc = NULL;
  switch (calculator) {
    case 0: // simple sorting calculator
      calc = new SortingMedianCalculator();
      break;
    // TODO: specify the `calculator` index of your implementation (e.g. `case 1:`),
    //   and instantiation your calculator here. Note that the index here should
    //   be consistent with the first argument of `test_median` function in
    //   CMakeList.txt.
    case 1: // quick select median calculator
      calc = new QuickSelectMedianCalculator();
      break;
    case 2: // binmedian calculator
      calc = new BinMedianCalculator();
      break;
    case 3: // randomized median calculator
      calc = new RandomizedMedianCalculator();
      break;
    default:
      std::cout << "Error: the requested median calculator has not yet been implemented\n";
      break;
  }

  // load sequence
  std::ifstream inputfile(argv[2]);
  std::vector<float> set;
  float temp;

  while (inputfile >> temp) {
    set.push_back(temp);
  }

  // call median calculator
  float median = calc->median(set);

  // report
  std::cout << median << "\n";

  delete calc;
  return 0;
}

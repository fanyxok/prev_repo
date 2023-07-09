#ifndef CS133_HW3_POLYNOMIALPARSER_HPP_
#define CS133_HW3_POLYNOMIALPARSER_HPP_

#include "polynomial.hpp"
#include <string>
#include <map>
#include <functional>

class PolynomialParser {
 public:
  //procedure to compute a polynomial function as a new polynomial
  //takes path to a function file fxx.txt
  //takes a map of polynomials that have been loaded and are used in the function
  //  (key=name of polynomial)
  //  (value=the polynomial)
  Polynomial compute_polynomial( const std::string& filename,
                                 std::map<std::string,Polynomial>& polys);

  //procedure to compute a lambda function that represents the polynomial function
  //parameters are same than for the previous function
  //NOTE: this function cannot make use of any of the operators defined in Polynomial (including compose)
  typedef std::function<float(float)> scalarFct;
  scalarFct compute_lambda( const std::string& filename, 
                            std::map<std::string, Polynomial> & polys);
};

#endif
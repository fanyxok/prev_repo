#include <string>
#include <map>
#include <iostream>
#include "polynomial.hpp"
#include "polynomial_parser.hpp"


int main(int argc, char** argv) {
  std::map<std::string, Polynomial> polys;
  PolynomialParser parser;

  // load all polynomials
  char polyname[80];
  for (int i = 1; i <= 10; ++i) {
    sprintf(polyname, "p%02d", i);
    std::string filename = std::string(argv[1]) + polyname + ".txt";
    Polynomial poly(filename);
    polys[std::string(polyname)] = poly;
  }  
  
  auto lambda_func = parser.compute_lambda(std::string(argv[2]), polys);
  std::cout << lambda_func(1.33) << std::endl;
  return 0;
}
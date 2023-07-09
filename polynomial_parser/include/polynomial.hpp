#ifndef CS133_HW3_POLYNOMIAL_HPP_
#define CS133_HW3_POLYNOMIAL_HPP_

#include <vector>

class Polynomial {
public:
  //default constructor
  Polynomial();

  //constructor from coefficient vectors (both copy and move versions)
  Polynomial( const std::vector<float> & coeffs );
  Polynomial( std::vector<float> && coeffs );
  //constructor from initializer list
  Polynomial( std::initializer_list<float> coeffs );

  //constructor from a path to a pXX.txt file of coefficients
  Polynomial( const std::string & path );

  //copy constructor
  Polynomial( const Polynomial & p );
  //move constructor
  Polynomial( Polynomial && p );

  //copy and move assignments
  Polynomial & operator=(const Polynomial & p);
  Polynomial & operator=( Polynomial && p );

  //destructor
  virtual ~Polynomial();

  //access to polynomial coefficients
  float & operator[](int index);
  const float & operator[](int index) const;

  //get number of coefficients (equals degree)
  int size() const;

  //print the polynomial into the console (may use for debugging)
  void print() const;

  //functional operations (self-explanatory)
  Polynomial   operator+ (const Polynomial & p) const;
  Polynomial & operator+=(const Polynomial & p);
  Polynomial   operator- (const Polynomial & p) const;
  Polynomial & operator-=(const Polynomial & p);
  Polynomial   operator* (const Polynomial & p) const;
  Polynomial   operator* (float factor) const;
  Polynomial & operator*=(const Polynomial & p);

  //function (polynomial) composition
  Polynomial compose( const Polynomial & p ) const;

  //function (polynomial) evaluation
  float operator()( float x ) const;

private:
  std::vector<float> _coeffs;
};

#endif
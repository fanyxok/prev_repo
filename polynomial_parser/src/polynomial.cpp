#include <string>
#include <fstream>
#include <iostream>

#include "polynomial.hpp"

Polynomial::Polynomial() {

}

Polynomial::Polynomial( const std::vector<float> & coeffs ) {
  _coeffs = coeffs;
}

Polynomial::Polynomial( std::vector<float> && coeffs ) {
  _coeffs = std::move( coeffs );
}

Polynomial::Polynomial( std::initializer_list<float> coeffs ) {
  _coeffs = std::vector<float>(coeffs.begin(), coeffs.end());
}

Polynomial::Polynomial( const std::string & path ) {
  std::ifstream file(path);
  _coeffs = std::vector<float>();
  std::string coeff_str;
  
  while( getline(file, coeff_str, ' ') ) {
    _coeffs.push_back(stof(coeff_str));
  }
}

Polynomial::Polynomial( const Polynomial & p) {
  *this = p;
}

Polynomial::Polynomial( Polynomial && p) {
  *this = std::move(p);
}

Polynomial & Polynomial::operator=( const Polynomial & p ) {
  _coeffs = p._coeffs;
  return *this;
}

Polynomial & Polynomial::operator=( Polynomial && p ) {
  *this = std::move(p);
  return *this;
}

Polynomial::~Polynomial(void) {

}

float & Polynomial::operator[]( int index ) {
  return _coeffs[ index ];
}

const float & Polynomial::operator[]( int index ) const {
  return _coeffs[ index ];
}

int Polynomial::size() const {
  return _coeffs.size();
}

void Polynomial::print() const {
  for( auto i : _coeffs ) {
    std::cout<< i <<" "; 
  }
  std::cout<<std::endl;
}

Polynomial Polynomial::operator+( const Polynomial & p ) const {
  std::vector<float> ret = std::vector<float>();
  // Case size p1 > p2
  if ( size() >= p.size() ) {
    for ( int i = 0; i < p.size(); i++ ) {
      ret.push_back( _coeffs[i] + p._coeffs[i]);
    }
    ret.insert(ret.end(), _coeffs.begin()+p.size(), _coeffs.end());
  // Case size p1 <= p2
  }else {
    for ( int i = 0; i < size(); i++ ) {
      ret.push_back( _coeffs[i] + p._coeffs[i]);
    }
    ret.insert(ret.end(), p._coeffs.begin()+size(), p._coeffs.end());
  }
  return Polynomial(ret);
}

Polynomial & Polynomial::operator+=( const Polynomial & p ) {
  // Case size p1 >= p2
  if ( size() >= p.size() ) {
    for ( int i = 0; i < p.size(); i++ ) {
      _coeffs[i] += p._coeffs[i];
    }
  // Case size p1 < p2
  }else {
    for ( int i = 0; i < size(); i++ ) {
      _coeffs[i] += p._coeffs[i];
    }
    for ( int i = size(); i < p.size(); i++ ) {
      _coeffs.push_back(p._coeffs[i]);
    }
  }
  return *this;
}

Polynomial Polynomial::operator-( const Polynomial & p ) const {
  std::vector<float> ret = std::vector<float>();
  // Case size p1 >= p2
  if ( size() >= p.size() ) {
    for ( int i = 0; i < p.size(); i++ ) {
      ret.push_back( _coeffs[i] - p._coeffs[i]);
    }
    ret.insert(ret.end(), _coeffs.begin()+p.size(), _coeffs.end());
  // Case size p1 < p2
  }else {
    for ( int i = 0; i < size(); i++ ) {
      ret.push_back( _coeffs[i] - p._coeffs[i]);
    }
    for ( int i = size(); i < p.size(); i++ ) {
      ret.push_back( -p._coeffs[i]);
    }
  }
  return Polynomial(ret);
}

Polynomial & Polynomial::operator-=( const Polynomial & p ) {
  // Case size p1 >= p2
  if ( size() >= p.size() ) {
    for ( int i = 0; i < p.size(); i++ ) {
      _coeffs[i] -= p._coeffs[i];
    }
  // Case size p1 < p2
  }else {
    for ( int i = 0; i < size(); i++ ) {
      _coeffs[i] -= p._coeffs[i];
    }
    for ( int i = size(); i < p.size(); i++ ) {
      _coeffs.push_back(-p._coeffs[i]);
    }
  }
  return *this;
}

Polynomial Polynomial::operator*( const Polynomial & p ) const {
  // The result vector should has known size
  std::vector<float> ret = std::vector<float>(size()+p.size()-1, 0.0);
  // Apply each coefficient to the polynomial and combine them all
  for ( int i = 0; i < size(); i ++ ) {
    for ( int j = 0; j < p.size(); j++ ) {
      ret[i+j] += _coeffs[i]*p[j];
    }
  }
  return Polynomial(ret);
}

Polynomial & Polynomial::operator*=( const Polynomial & p ) {
  // The result vector should has known size
  std::vector<float> ret = std::vector<float>(size()+p.size()-1, 0.0);
  // Apply each coefficient to the polynomial and combine them all
  for ( int i = 0; i < size(); i ++ ) {
    for ( int j = 0; j < p.size(); j++ ) {
      ret[i+j] += _coeffs[i]*p._coeffs[j];
    }
  }
  _coeffs = ret;
  return *this;
}

Polynomial Polynomial::operator*( float factor ) const {
  std::vector<float> ret = std::vector<float>();
  for ( auto i : _coeffs ) {
    ret.push_back( i * factor );
  }
  return Polynomial(ret);
}

Polynomial Polynomial::compose( const Polynomial & p ) const {
  std::vector<float> ret;
  // Set the 1 power of p
  Polynomial p_pow_i = Polynomial(p);
  // Init composed poly with coefficient of 0 power
  Polynomial composed_p = Polynomial(std::vector<float>{_coeffs[0]});
  // From 1 power to n power, calculate power of p and add to composed_p
  for ( int i =1; i < size(); i++ ) {
    composed_p += (p_pow_i * _coeffs[i]) ;
    p_pow_i *= p;
  }
  return composed_p;
}

float Polynomial::operator()( float x ) const {
  // Init result with constant of polynomial
  float result = _coeffs[0];
  // Calculate n power of x
  float x_pow_i = x;
  for ( int i = 1; i < size(); i++ ) {
    result += _coeffs[i] * x_pow_i;
    x_pow_i*= x;
  }
  return result;
}

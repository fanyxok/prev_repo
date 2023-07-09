#include <fstream>
#include <algorithm>

#include "polynomial_parser.hpp"

Polynomial PolynomialParser::compute_polynomial( const std::string & filename,
                                                 std::map<std::string, Polynomial> & polys ) {
  // Read polynomial expression and store them as "(", ")", "+/-/*", "pxx"
  std::ifstream file(filename);
  std::istreambuf_iterator<char> it(file),end;
  std::vector<std::string> expr_stack = std::vector<std::string>();
  std::string pol_str = std::string();
  while(it != end) {
    if ( *it == '(' ||
         *it == ')' ||
         *it == '+' ||
         *it == '-' ||
         *it == '*') {
      if ( pol_str.size() != 0 ) {
        expr_stack.push_back( pol_str );
        pol_str.clear();
      }
      expr_stack.push_back(std::string(1, *it));
    }else {
      pol_str.append(1, *it);
    }
    it++;
  }
  // Reverse expression for parsing stage
  std::reverse(expr_stack.begin(), expr_stack.end());
  // Create a stack for parsing stage
  std::vector<std::string> stack_stack = std::vector<std::string>();
  // Parse expression from left to right. As reversed expression, from back to front
  while ( !expr_stack.empty()) {
    // Case top is ")", reduction a "(...)"
    if (expr_stack.back() == ")" ) {
      std::vector<std::string>::reverse_iterator it = stack_stack.rbegin();
      it++;
      // Case (p) with stack: ...(p
      //     with expr_stack: ...)
      if ( *it == "(" ) {
        // Pop ) from expr
        expr_stack.pop_back();
        std::string p2_str = stack_stack.back();
        // Pop p from stack
        stack_stack.pop_back();
        // Pop ( from stack
        stack_stack.pop_back();

        if ( stack_stack.empty() ) {
          // Push p2 into stack
          stack_stack.push_back(p2_str);
        }else if ( stack_stack.back().size() == 1 ) {
          stack_stack.push_back(p2_str);
        // Case p1(p2) with stack: ...p1
        }else {
          // Compose p1 p2
          std::string p1_str = stack_stack.back();
          Polynomial p1 = polys[p1_str];
          Polynomial p2 = polys[p2_str];
          std::string p3_str = p1_str + p2_str;
          Polynomial p3 = p1.compose(p2);
          polys[p3_str] = p3;
          // Pop p1 from stack
          stack_stack.pop_back();
          // Push p3 into stack
          stack_stack.push_back(p3_str);
        }
      }
    // Case top is a poly
    }else if (expr_stack.back().size() > 1){
      // Case p1 op p2 with stack: ...p1 op
      //                     expr: ...p2
      if ( stack_stack.back() == "+" ) {
        // Pop + from stack
        stack_stack.pop_back();

        std::string p1_str = stack_stack.back();
        Polynomial p1 = polys[p1_str];

        std::string p2_str = expr_stack.back();
        Polynomial p2 = polys[p2_str];
        // Pop p2 from expr
        expr_stack.pop_back();

        std::string p3_str = p1_str + "+" + p2_str;
        Polynomial p3 = p1 + p2;
        polys[p3_str] = p3;
        // Change p1 to p3 in stack
        stack_stack.back() = p3_str;
      }else if ( stack_stack.back() == "-" ) {
        stack_stack.pop_back();
        std::string p1_str = stack_stack.back();
        Polynomial p1 = polys[p1_str];
        std::string p2_str = expr_stack.back();
        Polynomial p2 = polys[p2_str];
        expr_stack.pop_back();
        std::string p3_str = p1_str + "-" + p2_str;
        Polynomial p3 = p1 - p2;
        polys[p3_str] = p3;
        stack_stack.back() = p3_str;
      }else if ( stack_stack.back() == "*" ) {
        stack_stack.pop_back();
        std::string p1_str = stack_stack.back();
        Polynomial p1 = polys[p1_str];
        std::string p2_str = expr_stack.back();
        Polynomial p2 = polys[p2_str];
        expr_stack.pop_back();
        std::string p3_str = p1_str + "*" + p2_str;
        Polynomial p3 = p1 * p2;
        polys[p3_str] = p3;
        stack_stack.back() = p3_str;
      // Normal case 
      }else {
        stack_stack.push_back( expr_stack.back());
        expr_stack.pop_back();
      } 
    // Normal case     
    }else {
      stack_stack.push_back( expr_stack.back());
      expr_stack.pop_back();
    }  
  }
  // The one left in stack is the result polynomial
  return polys[stack_stack.front()];
}

std::function<float(float)> PolynomialParser::compute_lambda( const std::string & filename, std::map<std::string, Polynomial> & polys) {
  scalarFct s;
    
  // Load expression string from path
  std::ifstream file(filename);
  std::istreambuf_iterator<char> it(file),end;
  std::vector<std::string> expr_stack = std::vector<std::string>();
  std::string pol_str = std::string();
  while(it != end) {
    if ( *it == '(' ||
        * it == ')' ||
        * it == '+' ||
        * it == '-' ||
        * it == '*') {
      if ( pol_str.size() != 0 ) {
        expr_stack.push_back( pol_str );
        pol_str.clear();
      }
      expr_stack.push_back(std::string(1, *it));
    }else {
      pol_str.append(1, *it);
    }
    it++;
  }
  
  std::reverse(expr_stack.begin(), expr_stack.end());  
  std::vector<std::string> stack_stack = std::vector<std::string>();

  s = [&, expr_stack, stack_stack](float x)mutable->float {
    // Define some lambda function for calculate value, +, -, *, compose 
    std::function<float(Polynomial &, float)> calculateFct = [&](Polynomial p, float x){
      float result = p[0];
      float x_pow_i = x;
      for (int i = 1; i < p.size(); i++ ) {
        result += (p[i]* x_pow_i);
        x_pow_i *= x;
      }
      return result;
    };
    std::function<Polynomial(Polynomial &, Polynomial &, float)> plusFct = [&](Polynomial p1, Polynomial p2, float x)->Polynomial{
      float result = calculateFct(p1, x) + calculateFct(p2, x);
      return Polynomial(std::vector<float>{result});
    };
    std::function<Polynomial(Polynomial &, Polynomial &, float)> minusFct = [&](Polynomial p1, Polynomial p2, float x)->Polynomial{
      float result = calculateFct(p1, x) - calculateFct(p2, x);
      return Polynomial(std::vector<float>{result});
    };
    std::function<Polynomial(Polynomial &, Polynomial &, float)> mulFct = [&](Polynomial p1, Polynomial p2, float x)->Polynomial{
      float result = calculateFct(p1, x) * calculateFct(p2, x);
      return Polynomial(std::vector<float>{result});
    }; 
    std::function<Polynomial(Polynomial &, Polynomial &, float)> composeFct = [&](Polynomial p1, Polynomial p2, float x)->Polynomial{
      float result = calculateFct(p1, calculateFct(p2, x));
      return Polynomial(std::vector<float>{result});
    }; 
    // Parse expression
    while ( !expr_stack.empty()) {   
      if (expr_stack.back() == ")" ) {  
        std::vector<std::string>::reverse_iterator it = stack_stack.rbegin();
        it++;  
        // Case (p) with stack: ...(p
        //     with expr_stack: ...)
        if ( *it == "(" ) {
          // Pop ) from expr
          expr_stack.pop_back();
          std::string p2_str = stack_stack.back();
          // Pop p from stack
          stack_stack.pop_back();
          // Pop ( from stack
          stack_stack.pop_back();
          if ( stack_stack.empty() ) {
            stack_stack.push_back(p2_str);
          }else if ( stack_stack.back().size() == 1 ) {
            stack_stack.push_back(p2_str);
          // Case p1(p2) with stack: ...p1 
          }else {
            std::string p1_str = stack_stack.back();
            Polynomial p1 = polys[p1_str];
            Polynomial p2 = polys[p2_str];
            std::string p3_str = p1_str + p2_str;
            Polynomial p3 = composeFct(p1, p2, x);
            polys[p3_str] = p3;
            // Pop p1 from stack
            stack_stack.pop_back();
            // Push p3 into stack
            stack_stack.push_back(p3_str);
          }
        }
      }else if (expr_stack.back().size() > 1){
        // Case (p1-p2)(p3) with stack: ...(p1-
        //                        expr: ...)p3()p2
        // change stacks to stack: ...(p13-
        //                   expr: ...)p23
        if ( expr_stack.size() >= 5 && stack_stack.size() >=3 ) {
          if ( expr_stack[expr_stack.size()-2] == ")" && expr_stack[expr_stack.size()-3]== "(") {
            std::string p1_str = stack_stack[stack_stack.size()-2];
            std::string p2_str = expr_stack.back();
            std::string op = stack_stack.back();
            std::string p3_str = expr_stack[expr_stack.size()-4];
            
            std::string p13_str = p1_str + p3_str;
            Polynomial p13 = composeFct(polys[p1_str], polys[p3_str], x);
            polys[p13_str] = p13;
            std::string p23_str = p2_str + p3_str;
            Polynomial p23 = composeFct(polys[p2_str], polys[p3_str], x);
            polys[p23_str] = p23;
            stack_stack[stack_stack.size()-2] = p13_str;
            
            for ( int i : {1,2,3,4} ) {
              expr_stack.pop_back();
            }
            expr_stack.push_back(p23_str);
            continue;
          }
        }
        // Case p1 op p2 with stack: ...p1 op
        //                     expr: ...p2
        if ( stack_stack.back() == "+" ) {
          // Pop + from stack
          stack_stack.pop_back();
          std::string p1_str = stack_stack.back();
          Polynomial p1 = polys[p1_str];

          std::string p2_str = expr_stack.back();
          Polynomial p2 = polys[p2_str];
          // Pop p2 from expr
          expr_stack.pop_back();
          // Calculate p3
          std::string p3_str = p1_str + "+" + p2_str;
          Polynomial p3 = plusFct(p1, p2, x);
          polys[p3_str] = p3;
          // Change p1 to p3 in stack
          stack_stack.back() = p3_str;
        }else if ( stack_stack.back() == "-" ) {
          stack_stack.pop_back();
          std::string p1_str = stack_stack.back();
          Polynomial p1 = polys[p1_str];
          std::string p2_str = expr_stack.back();
          Polynomial p2 = polys[p2_str];
          expr_stack.pop_back();
          std::string p3_str = p1_str + "-" + p2_str;
          Polynomial p3 = minusFct(p1, p2, x);
          polys[p3_str] = p3;
          stack_stack.back() = p3_str;
        }else if ( stack_stack.back() == "*" ) {
          stack_stack.pop_back();
          std::string p1_str = stack_stack.back();
          Polynomial p1 = polys[p1_str];
          std::string p2_str = expr_stack.back();
          Polynomial p2 = polys[p2_str];
          expr_stack.pop_back();
          std::string p3_str = p1_str + "*" + p2_str;
          Polynomial p3 = mulFct(p1, p2, x);
          polys[p3_str] = p3;
          stack_stack.back() = p3_str;
        }else {
          stack_stack.push_back( expr_stack.back());
          expr_stack.pop_back();
        }      
      }else {
        stack_stack.push_back( expr_stack.back());
        expr_stack.pop_back();
      }
    }  
    return polys[stack_stack.back()][0];
  };
  return s;
}
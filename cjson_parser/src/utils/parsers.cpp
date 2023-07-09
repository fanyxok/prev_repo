//
// Created by lirundong on 2019-03-25.
//
#include <locale>
#include <codecvt>
#include <iostream>
#include "utils/parsers.hpp"

namespace toy_json {

// TODO: Implement parsing helpers in this file. If any parsing error occurs,
//   e.g., meet with invalid input string, just throw `std::runtime_error` with
//   bellowing error information (without surrounding ` ):
//   - fail parsing `bool`:    `parsing bool type failed`
//   - fail parsing `number`:  `parsing number type failed`
//   - fail parsing `string`:  `parsing string type failed`
//   - fail parsing `array`:   `parsing array type failed`
//   - fail parsing `object`:  `parsing object type failed`

bool parse_bool(std::string::const_iterator &str_it) {
  bool result = false;
  // Parse "true"
  if ( *str_it == 't' ) {
    std::string true_str = "true";
    int len_true = 4;
    for ( int i = 0; *str_it != '\0' && i < len_true ; i++ ) {
      if ( *str_it != true_str[i] ) {
        throw std::runtime_error("parsing bool type failed");
        break;
      } else {
        str_it++;
      }
    }
    result = true;

  // Parse "false"
  }else if ( *str_it == 'f' ) {
    std::string false_str = "false"; 
    int len_false = 5;
    for ( int i = 0; *str_it != '\0' && i < len_false ; i++ ) {
      if ( *str_it != false_str[i] ) {
        throw std::runtime_error("parsing bool type failed");
        break;
      } else {
        str_it++;
      }
    }
    result = false;
  } else {
    throw std::runtime_error("parsing bool type failed");
  }
  return result;
}


double parse_number(std::string::const_iterator &str_it) {
  double result = 0.0;
  std::string number_str = std::string();
  int len_number = 0;
  // Parse optional '-'
  if ( *str_it == '-' ) {
    len_number++;
    str_it++;
  }
  // Parse two case '0' | ('1-9'+'0-9'*)
  if ( isdigit(*str_it) ) {
    if ( *str_it == '0' ) {
      len_number++;
      str_it++;
    }else {
      len_number++;
      str_it++;
      parse_digits_in_number(str_it, len_number);
    }
  }else {
    throw std::runtime_error("parsing number type failed");
  }
  // Parse optional '.'+'0-9'+'0-9'* 
  if ( *str_it == '.' ) {
    len_number++;
    str_it++;
    if ( isdigit(*str_it) ) {
      len_number++;
      str_it++;
      parse_digits_in_number(str_it, len_number);
    }else {
      throw std::runtime_error("parsing number type failed");
    }
  }
  // Parse optional 'e | E' + '+ | | -' + '0-9' + '0-9'*
  if ( *str_it == 'e' || *str_it == 'E' ) {
    len_number++;
    str_it++;
    if ( *str_it == '+' || *str_it == '-' || isdigit(*str_it) ) {
      len_number++;
      str_it++;
      parse_digits_in_number(str_it, len_number);
    }else {
      throw std::runtime_error("parsing number type failed");
    }
  }
  // Convert string to double
  number_str = std::string(str_it-len_number, str_it);
  // Ensure automata when single unit parse 
  if ( *str_it == '\0' ) {
    result = std::stod(number_str);
  // Ensure automate when parse in Json text
  }else if ( *str_it == ',' || *str_it == ']' || *str_it == '}' ) {
    result = std::stod(number_str);
  }
  else {
    throw std::runtime_error("parsing number type failed");
  }
  return result;
}

std::string parse_string(std::string::const_iterator &str_it) {
  std::string result = std::string();
  bool end_with_quote = false;
  // String start with '"'
  if ( *str_it == '"' ) {
    str_it++;
    // Deal with string obj name
    while ( *str_it != '\0' ){
      if ( *str_it == '\\' ) {
        str_it++;
        switch ( *str_it ) {
          case '"' :
            result.append("\"");
            str_it++;
            break;
          case '\\':
            result.append("\\");
            str_it++;
            break;
          case '/' :
            result.append("/");
            str_it++;
            break;
          case 'b' :
            result.append("\b");
            str_it++;
            break;
          case 'f' :
            result.append("\f");
            str_it++;
            break;
          case 'n' :
            result.append("\n");
            str_it++;
            break;
          case 'r' :
            result.append("\r");
            str_it++;
            break;
          case 't' :
            result.append("\t");
            str_it++;
            break;
          // Unicode string convert to uft-8
          case 'u' : {
            str_it++;
            std::string utf8_str = parse_unicode_in_string( str_it );
            result.append(utf8_str);     
            break;
          }
          default:
            break;
        }
      // Ensure that end with '"'
      }else if ( *str_it == '"') {
        end_with_quote = true;
        str_it++;
        break;
      // Normal char, put in result string
      }else {
        result.append(str_it, str_it+1);
        str_it++;
      }
    }
  }else {
    throw std::runtime_error("parsing string type failed");
  }
  if ( end_with_quote ){
    return result;
  }else {
    throw std::runtime_error("parsing string type failed");
  }
}


JsonNode::array parse_array(std::string::const_iterator &str_it) {
  JsonNode::array result;
  bool end_with_square_bracket = false;
  // Array start with '['
  if ( *str_it != '[') {
    throw std::runtime_error("parsing array type failed");
  }
  str_it++;

  // Parser value
  while ( *str_it != '\0' ) {
    ignore_space( str_it );
    // Case string
    if ( *str_it == '"'){
      std::string value_str = parse_string(str_it); 
      std::unique_ptr<std::string> value_unique_str = std::make_unique<std::string>(value_str);
      std::unique_ptr<JsonNode> value = std::make_unique<JsonNode>(std::move(value_unique_str));  
      result.push_back(*value);
    // Case number
    }else if ( isdigit(*str_it) || *str_it == '-' || *str_it == '+') {
      double value_number = parse_number(str_it);
      std::unique_ptr<JsonNode> value = std::make_unique<JsonNode>(value_number); 
      result.push_back(*value);
    // Case bool
    }else if ( *str_it == 't' || *str_it == 'f'
             ||*str_it == 'T' || *str_it == 'F') {
      bool value_bool = parse_bool(str_it);
      std::unique_ptr<JsonNode> value = std::make_unique<JsonNode>(value_bool);
      result.push_back(*value);  
    // Case null
    }else if ( *str_it == 'n') {
      std::string null_str = "null";
      int len_null = 4;
      for (int i = 0; i < len_null; i++) {
        if ( *str_it == '\0') {
          throw std::runtime_error("parsing null type failed");
        }
        if ( *str_it != null_str[i]){    
          throw std::runtime_error("parsing null type failed");
        }
        str_it++;
      }
      std::unique_ptr<JsonNode> value = std::make_unique<JsonNode>();
      result.push_back(*value);
    // Case array
    }else if ( *str_it == '[') {
      JsonNode::array value_array = parse_array(str_it);
      std::unique_ptr<JsonNode::array> value_array_ptr = std::make_unique<JsonNode::array>(value_array);
      std::unique_ptr<JsonNode> value = std::make_unique<JsonNode>(std::move(value_array_ptr));
      result.push_back(*value);
    // Case object
    }else if ( *str_it == '{') {
      JsonNode::object value_object = parse_object(str_it);
      std::unique_ptr<JsonNode::object> value_object_ptr = std::make_unique<JsonNode::object>(value_object);
      std::unique_ptr<JsonNode> value = std::make_unique<JsonNode>(std::move(value_object_ptr));
      result.push_back(*value);
    }else{
        throw std::runtime_error("parsing array type failed");
    }
    // Next value exist?
    if ( *str_it == ',' ) {
      str_it++;
      continue;
    // Array end
    }else if ( *str_it == ']' ) {
      str_it++;
      end_with_square_bracket = true;
      break;
    }
  }
  if ( end_with_square_bracket ) {
    return result;
  }else {
    throw std::runtime_error("parsing array type failed");
  }
  
}

JsonNode::object parse_object(std::string::const_iterator &str_it) {
  JsonNode::object result;
  bool end_with_brace = false;
  ignore_space( str_it );
  // Object start with '{'
  if ( *str_it != '{') {
    throw std::runtime_error("parsing object type failed");
  }
  str_it++;
  std::string key = std::string();
  while( *str_it != '\0') {
    ignore_space( str_it);
    // Deal with (string : value) pair
    if ( *str_it == '"') {
      key = parse_string(str_it);
    }else {
      throw std::runtime_error("parsing object type failed");
    }

    // Ensure ':' between key and value
    ignore_space( str_it );
    if (*str_it != ':') {
        throw std::runtime_error("parsing object type failed");
    }
    str_it++;

    ignore_space( str_it );
    // Deal with obj value
    // Case string
    if ( *str_it == '"') {
      std::string value_str = parse_string(str_it);
      std::unique_ptr<std::string> value_ptr = std::make_unique<std::string>(value_str);
      std::unique_ptr<JsonNode> value = std::make_unique<JsonNode>(std::move(value_ptr));
      result.insert(std::make_pair(key, *value));
    // Case object
    }else if ( *str_it == '{') {
      JsonNode::object value_obj = parse_object(str_it);
      std::unique_ptr<JsonNode::object> value_ptr = std::make_unique<JsonNode::object>(value_obj);
      std::unique_ptr<JsonNode> value = std::make_unique<JsonNode>(std::move(value_ptr));
      result.insert(std::make_pair(key, *value));
    // Case array
    }else if ( *str_it == '[') {
      JsonNode::array value_array =  parse_array(str_it);
      std::unique_ptr<JsonNode::array> value_ptr = std::make_unique<JsonNode::array>(value_array);
      std::unique_ptr<JsonNode> value = std::make_unique<JsonNode>(std::move(value_ptr));
      result.insert(std::make_pair(key, *value));
    // Case bool 
    }else if ( *str_it == 't' || *str_it == 'f'
            || *str_it == 'T' || *str_it == 'F') {
      bool value_bool = parse_bool(str_it);
      std::unique_ptr<JsonNode> value = std::make_unique<JsonNode>(value_bool);
      result.insert(std::make_pair(key, *value));
    // Case null
    }else if ( *str_it == 'n' ) {
      std::string null_str = "null";
      int len_null = 4;
      for (int i = 0; i < len_null; i++) {
        if ( *str_it == '\0' ) {
          throw std::runtime_error("parsing null type failed");
        }
        if ( *str_it != null_str[i] ){
          
          throw std::runtime_error("parsing null type failed");
        }
        str_it++;
      }
      std::unique_ptr<JsonNode> value = std::make_unique<JsonNode>();
      result.insert(std::make_pair(key, *value));
    // Case number
    }else if ( *str_it == '-' || *str_it == '+' || isdigit(*str_it) ){
      double value_number = parse_number(str_it);
      std::unique_ptr<JsonNode> value = std::make_unique<JsonNode>(value_number);
      result.insert(std::make_pair(key, *value));
    }else {
      throw std::runtime_error("parsing object type failed");
    }
    // Object end
    if ( *str_it == '}' ) {
      str_it++;
      end_with_brace =true;
      break;
    // Next pair exist
    }else if ( *str_it == ',' ) {
      str_it++;
      continue;
    }else {
      throw std::runtime_error("parsing object type failed");
    }   
  }
  return result;
}


// Helper Function Implementation
//
// Input: str_it, counter
// Output: Void
// Function: iterate the str_it until there is not a digit, 
//           for each iteration add 1 to counter 
void parse_digits_in_number( std::string::const_iterator &str_it, int &len ) {
  while ( *str_it != '0' ) {
    if ( isdigit(*str_it) ) {
      len++;
      str_it++;
    }else {
      break;
    }
  }
}

// Input: str_it
// Output: Bool
// Function: iterate the str_it until there is not a space
bool ignore_space(std::string::const_iterator &str_it) {
  while ( *str_it != '\0') {
    if ( *str_it == ' ') {
      str_it++;
    }else {
      return true;
    }
  }
  return false;
}

// Input: str_it
// Output: string
// Function: convert unicode string to utf-8 string
std::string parse_unicode_in_string( std::string::const_iterator &str_it ) {
  // Create a string with "0x" to store unicode
  std::string unicode_str = std::string("0x");
  for (int i = 0; *(str_it+i) != '\0' && i < 4; i++ ) {
    if ( !isalpha(*(str_it+i)) && !isdigit(*(str_it+i)) ) {
      throw std::runtime_error("parsing string type failed");
    }
  }
  unicode_str.append(str_it, str_it+4);
  str_it+=4;
  // Convert hex string to hex unsigned
  std::wstring_convert<std::codecvt_utf8<char32_t>, char32_t> converter;
  unsigned hex;
  sscanf(unicode_str.c_str(), "%x", &hex);
  // If is surrogate pair, read one more unicode
  if ( 0xd800 < hex && hex < 0xdfff ) {
    if ( *str_it == '\\' && *(str_it+1) == 'u' ) {
      str_it+=2;
      std::string unicode_str_2 = std::string("0x");
      for (int i = 0; *(str_it+i) != '\0' && i < 4 ; i++ ) {
        if ( !isalpha(*(str_it+i)) && !isdigit(*(str_it+i)) ) {
          throw std::runtime_error("parsing string type failed");
        }
      }
      unicode_str_2.append(str_it, str_it+4);
      str_it+=4;
      unsigned hex_2;
      sscanf(unicode_str_2.c_str(), "%x", &hex_2);
      // Convert surrogate pair to a single unicode point
      if ( 0xd800 < hex_2 && hex_2 < 0xdfff ) {
        unsigned hex_high = hex & 0x3ff;
        unsigned hex_low = hex_2 & 0x3ff;
        hex = ((hex_high << 10) | hex_low ) + 0x10000; 
      }else {
        throw std::runtime_error("parsing string type failed");
      }
    }else {
      throw std::runtime_error("parsing string type failed");
    }
  }
  // Convert hex unsigned to utf-8
  std::string utf8_str = converter.to_bytes(hex);
  return utf8_str;
}

}

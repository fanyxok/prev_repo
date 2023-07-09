#include <iostream>
#include <algorithm>
#include <fstream>

#include "toy_json.hpp"
#include "utils/parsers.hpp"

namespace toy_json
{

JsonNode::JsonNode(const JsonNode &rhs)
{
  release_union();
  type_ = rhs.type();
  switch (rhs.type())
  {
  case JSON_NULL:
    break;
  case JSON_BOOL:
    v_bool_ = rhs.get_bool();
    break;
  case JSON_NUMBER:
    v_number_ = rhs.get_number();
    break;
  case JSON_STRING:
    new (&v_string_) std::unique_ptr<std::string>(new std::string{*rhs.v_string_});
    break;
  case JSON_ARRAY:
    new (&v_array_) std::unique_ptr<array>(new array{*rhs.v_array_});
    break;
  case JSON_OBJECT:
    new (&v_object_) std::unique_ptr<object>(new object{*rhs.v_object_});
    break;
  }
}

JsonNode::JsonNode(JsonNode &&rhs) noexcept
{
  release_union();
  type_ = rhs.type();
  switch (rhs.type())
  {
  case JSON_NULL:
    break;
  case JSON_BOOL:
    v_bool_ = rhs.get_bool();
    break;
  case JSON_NUMBER:
    v_number_ = rhs.get_number();
    break;
  case JSON_STRING:
    new (&v_string_) std::unique_ptr<std::string>(std::move(rhs.v_string_));
    break;
  case JSON_ARRAY:
    new (&v_array_) std::unique_ptr<array>(std::move(rhs.v_array_));
    break;
  case JSON_OBJECT:
    new (&v_object_) std::unique_ptr<object>(std::move(rhs.v_object_));
    break;
  }
}

void JsonNode::release_union()
{
  switch (type_)
  {
  case JSON_STRING:
    v_string_.~unique_ptr<std::string>();
    break;
  case JSON_ARRAY:
    v_array_.~unique_ptr<array>();
    break;
  case JSON_OBJECT:
    v_object_.~unique_ptr<object>();
    break;
  default:
    break;
  }
}

std::string Json::error_info_;

std::unique_ptr<JsonNode> Json::parse(const std::string &fin) noexcept
{
  // TODO: Implement main parsing procedure here, some notices:
  //   1. You may find helper functions in `utils/parsers.hpp` useful. Note that
  //      we did not provide `parse_null`, why?
  //   2. Note that actual data that larger than 8 bytes is hold by `shared_ptr`
  //      you should construct `string`, `array` and `object` types by passing
  //      r-valued `unique_ptr`s;
  //   3. DO NOT throw exceptions out of this function. If an internal exception
  //      occurs, e.g. get invalid input, just catch that exception, modify the
  //      `error_info_` and return a `nullptr`;
  //   4. Expected `error_info` for each kind of parsing errors: (these
  //      messages, without surrounding "`", will be examined in unit tests,
  //      make sure you throw the correct internal errors)
  //      - fail parsing `null`:    `parsing null type failed`
  //      - fail parsing `bool`:    `parsing bool type failed`
  //      - fail parsing `number`:  `parsing number type failed`
  //      - fail parsing `string`:  `parsing string type failed`
  //      - fail parsing `array`:   `parsing array type failed`
  //      - fail parsing `object`:  `parsing object type failed`

  // Read file and save it in a string
  std::ifstream file(fin);
  std::string json = std::string((std::istreambuf_iterator<char>(file)), std::istreambuf_iterator<char>());

  // Create a const iterator
  std::string::const_iterator root_it = json.cbegin();
  
  // Try parse to deal with exception
  try
  {
    JsonNode::object root = parse_object(root_it);
    std::unique_ptr<JsonNode::object> root_ptr = std::make_unique<JsonNode::object>(root);
    std::unique_ptr<JsonNode> root_node_ptr = std::make_unique<JsonNode>(std::move(root_ptr));
    return root_node_ptr;
  }
  // Catch the exception and save info
  catch (std::runtime_error &e)
  {
    error_info_ = e.what();
    return nullptr;
  }
}
} // namespace toy_json

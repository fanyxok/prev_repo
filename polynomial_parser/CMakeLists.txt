cmake_minimum_required(VERSION 3.5)

project(cs133hw3)
set(CMAKE_CXX_STANDARD 14)

#version number
set( CS133_HW3_VERSION_MAJOR 1 )
set( CS133_HW3_VERSION_MINOR 0 )

#set build type
set(CMAKE_BUILD_TYPE Debug)

#add path for find scripts
set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} "${PROJECT_SOURCE_DIR}/cmake/")

#include CTest so we can run unit tests with CMake
include(CTest)

#find libraries

#add directory with include files of this project
include_directories( "${PROJECT_SOURCE_DIR}/include" )

add_subdirectory( src )
add_subdirectory( test )
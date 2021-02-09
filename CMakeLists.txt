cmake_minimum_required(VERSION 3.0.0)
project(vdc-session VERSION 0.1.0)


include(./cmake/common.cmake)
include(./cmake/go.cmake)

add_subdirectory(client)
add_subdirectory(cmd/sessiond)
add_subdirectory(cmd/sessionctl)
# cmake build file for sessiond.
# Assumes protobuf and gRPC have been installed using cmake.

find_program(_GO go)
message(STATUS "Using golang ${_GO}")
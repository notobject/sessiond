# Copyright 2021 SANGFOR.COM xxxx.
#
# cmake build file for vdc-session.
# Assumes protobuf and gRPC have been installed using cmake.

find_program(_GO go)
message(STATUS "Using golang ${_GO}")
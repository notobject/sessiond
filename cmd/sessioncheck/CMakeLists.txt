# Copyright 2021 SANGFOR.COM xxxx.
#
# cmake build file for vdc-session.
# Assumes protobuf and gRPC have been installed using cmake.

add_custom_command(
    OUTPUT sessioncheck-service
    COMMAND ${_GO} 
    ARGS build 
         -o ${CMAKE_CURRENT_BINARY_DIR}/sessioncheck
         ${CMAKE_CURRENT_SOURCE_DIR}/main.go
)

add_custom_target(sessioncheck ALL
    DEPENDS sessioncheck-service
)

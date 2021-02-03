# cmake build file for sessiond.
# Assumes protobuf and gRPC have been installed using cmake.

add_custom_command(
    OUTPUT sessionctl-service
    COMMAND ${_GO} 
    ARGS build 
         -o ${CMAKE_CURRENT_BINARY_DIR}/sessionctl
         ${CMAKE_CURRENT_SOURCE_DIR}/main.go
)

add_custom_target(sessionctl ALL
    DEPENDS sessionctl-service
)

# cmake build file for sessiond.
# Assumes protobuf and gRPC have been installed using cmake.

# Proto file
get_filename_component(proto "../../api/rpc/session.proto" ABSOLUTE)
get_filename_component(proto_path "${proto}" PATH)

# Generated sources
set(proto_srcs "${CMAKE_CURRENT_BINARY_DIR}/session.pb.go")
set(grpc_srcs "${CMAKE_CURRENT_BINARY_DIR}/session_grpc.pb.go")
add_custom_command(
    OUTPUT "${proto_srcs}" "${grpc_srcs}"
    COMMAND ${_PROTOBUF_PROTOC}
    ARGS --go_out="${PROJECT_SOURCE_DIR}/api/rpc"
         --go-grpc_out="${PROJECT_SOURCE_DIR}/api/rpc"
         -I ${proto_path}
         "${proto}"
    DEPENDS "${proto}")

add_custom_command(
    OUTPUT session-service
    COMMAND ${_GO} 
    ARGS build 
         -o ${CMAKE_CURRENT_BINARY_DIR}/sessiond
         ${CMAKE_CURRENT_SOURCE_DIR}/main.go
    DEPENDS "${proto_srcs}" 
    DEPENDS "${grpc_srcs}"
)

add_custom_target(sessiond ALL
    DEPENDS session-service
)


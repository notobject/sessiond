# cmake build file for sessiond.
# Assumes protobuf and gRPC have been installed using cmake.

# Proto file
get_filename_component(proto "../api/rpc/session.proto" ABSOLUTE)
get_filename_component(proto_path "${proto}" PATH)

# Generated sources of pb
set(proto_srcs "${CMAKE_CURRENT_SOURCE_DIR}/session.pb.cc")
set(proto_hdrs "${CMAKE_CURRENT_SOURCE_DIR}/session.pb.h")
set(grpc_srcs "${CMAKE_CURRENT_SOURCE_DIR}/session.grpc.pb.cc")
set(grpc_hdrs "${CMAKE_CURRENT_SOURCE_DIR}/session.grpc.pb.h")
add_custom_command(
      OUTPUT "${proto_srcs}" "${proto_hdrs}" "${grpc_srcs}" "${grpc_hdrs}"
      COMMAND ${_PROTOBUF_PROTOC}
      ARGS --grpc_out "${CMAKE_CURRENT_SOURCE_DIR}"
        --cpp_out "${CMAKE_CURRENT_SOURCE_DIR}"
        -I "${proto_path}"
        --plugin=protoc-gen-grpc="${_GRPC_CPP_PLUGIN_EXECUTABLE}"
        "${proto}"
      DEPENDS "${proto}")

# Include generated *.pb.h files
include_directories(
  "${CMAKE_CURRENT_SOURCE_DIR}"
  "${CMAKE_CURRENT_SOURCE_DIR}/include"
)

# libsession
add_library(session SHARED
  ${grpc_srcs}
  ${grpc_hdrs}
  ${proto_srcs}
  ${proto_hdrs}
  "${CMAKE_CURRENT_SOURCE_DIR}/include/client_debug.h"
  "${CMAKE_CURRENT_SOURCE_DIR}/include/session.h"
  "${CMAKE_CURRENT_SOURCE_DIR}/lib/session_client.h"
  "${CMAKE_CURRENT_SOURCE_DIR}/lib/session.cpp"
)
target_link_libraries(session
  ${_REFLECTION}
  ${_GRPC_GRPCPP}
  ${_PROTOBUF_LIBPROTOBUF}
)

# client_test
add_executable(client_test "${CMAKE_CURRENT_SOURCE_DIR}/test/client_test.c")
target_link_libraries(client_test
  session
)
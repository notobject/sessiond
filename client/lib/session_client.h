// Created by longduping 2021-02-09
// 会话管理接口 调用层

#include <iostream>
#include <memory>
#include <grpcpp/grpcpp.h>

#include "session.pb.h"
#include "session.grpc.pb.h"

using rpc::Session;

class SessionClient
{
private:
    SessionClient(std::shared_ptr<grpc::Channel> channel)
        : stub_(rpc::Session::NewStub(channel)) {}

public:
    static SessionClient *GetInstance()
    {
        static SessionClient s(grpc::CreateChannel("127.0.0.1:8080", grpc::InsecureChannelCredentials()));
        return &s;
    }

    int CreateSession(const rpc::CreateSessionReq &req, rpc::CreateSessionRsp &rsp)
    {
        grpc::ClientContext context;
        grpc::Status status = stub_->CreateSession(&context, req, &rsp);
        if (!status.ok())
        {
            std::cout << "CreateSession failed: ["
                      << status.error_code() << ": " << status.error_message()
                      << "]" << std::endl;
            return status.error_code();
        }
        std::cout << "CreateSession success: \n"
                  << rsp.DebugString() << std::endl;
        return rsp.ret();
    }

    int GetSessionInfoByID(const char *session_id, online_t *online)
    {
        grpc::ClientContext context;
        rpc::GetSessionInfoByIDReq req;
        rpc::GetSessionInfoRsp rsp;

        // 封装请求对象
        req.set_sessionid(session_id);

        // RPC调用
        grpc::Status status = stub_->GetSessionInfoByID(&context, req, &rsp);
        if (!status.ok())
        {
            std::cout << "GetSessionInfoByID failed: ["
                      << status.error_code() << ": " << status.error_message()
                      << "]" << std::endl;
            return status.error_code();
        }

        // TODO online_t 反序列化
        // online = (online_t *)rsp.userinfo();

        std::cout << "GetSessionInfoByID success: \n"
                  << rsp.DebugString() << std::endl;
        return rsp.ret();
    }

private:
    std::unique_ptr<rpc::Session::Stub> stub_;
};
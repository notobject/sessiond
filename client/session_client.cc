#include <iostream>
#include <memory>
#include <grpcpp/grpcpp.h>

#include "session.pb.h"
#include "session.grpc.pb.h"

using rpc::Session;

class SessionClient
{
public:
    SessionClient(std::shared_ptr<grpc::Channel> channel)
        : stub_(rpc::Session::NewStub(channel)) {}

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
        return 0;
    }

    int CheckSession(const rpc::CheckSessionReq &req, rpc::CheckSessionRsp &rsp)
    {
        grpc::ClientContext context;
        grpc::Status status = stub_->CheckSession(&context, req, &rsp);
        if (!status.ok())
        {
            std::cout << "CheckSession failed: ["
                      << status.error_code() << ": " << status.error_message()
                      << "]" << std::endl;
            return status.error_code();
        }
        std::cout << "CheckSession success: \n"
                  << rsp.DebugString() << std::endl;
        return 0;
    }

    int GetSessionInfoByID(const rpc::GetSessionInfoByIDReq &req, rpc::GetSessionInfoRsp &rsp)
    {
        grpc::ClientContext context;
        grpc::Status status = stub_->GetSessionInfoByID(&context, req, &rsp);
        if (!status.ok())
        {
            std::cout << "GetSessionInfoByID failed: ["
                      << status.error_code() << ": " << status.error_message()
                      << "]" << std::endl;
            return status.error_code();
        }
        std::cout << "GetSessionInfoByID success: \n"
                  << rsp.DebugString() << std::endl;
        return 0;
    }

private:
    std::unique_ptr<rpc::Session::Stub> stub_;
};

int main(void)
{
    int ret = 0; 
    SessionClient session(grpc::CreateChannel("127.0.0.1:8080", grpc::InsecureChannelCredentials()));

    // 创建会话
    rpc::CreateSessionReq createSessionReq;
    rpc::CreateSessionRsp createSessionRsp;
    createSessionReq.set_username("ldp");
    createSessionReq.set_userinfo("online_t");
    ret = session.CreateSession(createSessionReq, createSessionRsp);

    // 检查会话
    rpc::CheckSessionReq checkSessionReq;
    rpc::CheckSessionRsp checkSessionRsp;
    checkSessionReq.set_sessionid(createSessionRsp.sessionid());
    ret = session.CheckSession(checkSessionReq, checkSessionRsp);

    // 获取会话信息
    rpc::GetSessionInfoByIDReq getSessionInfoByIDReq;
    rpc::GetSessionInfoRsp getSessionInfoRsp;
    getSessionInfoByIDReq.set_sessionid(createSessionRsp.sessionid());
    ret = session.GetSessionInfoByID(getSessionInfoByIDReq, getSessionInfoRsp);
}
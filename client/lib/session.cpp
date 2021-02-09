// Created by longduping 2021-02-09
// 会话管理接口 Wrapper层

#include "session.h"
#include "session_client.h"

/**
 * 创建会话
 * 
 * 这里演示的是将包装和解包的逻辑放到Wrapper层做，使调用层更清晰
 */
int32_t session_create(const online_t *online, char *session_id)
{
    rpc::CreateSessionReq req;
    rpc::CreateSessionRsp rsp;

    // TODO 处理请求对象：online_t 序列化
    req.set_userinfo("online_t");

    // 通过C++客户端进程RPC调用
    int ret = SessionClient::GetInstance()->CreateSession(req, rsp);
    if (ret != 0){
        return ret;
    }

    // TODO 处理返回对象
    strncpy(session_id, rsp.sessionid().c_str(), 16);
    return ret;
}

/**
 * 获取会话
 * 
 * 这里演示的是将包装和解包的逻辑都放到最终RPC调用处做，使包装层更清晰
 */
int32_t session_get(const char *session_id, online_t *online)
{
    return SessionClient::GetInstance()->GetSessionInfoByID(session_id, online);
}
// Created by longduping 2021-02-09
// 会话管理 C语言接口
#ifndef _SESSION_20210209_H
#define _SESSION_20210209_H

#include <stdio.h>

// TODO client_debug.h just for test
#include "client_debug.h"

__BEGIN_DECLS

/**
 * 获取错误提示信息
 * @param[IN] errcode 错误码
 * @return 错误提示
 */
char* session_errmsg(int32_t errcode);

/**
 * 创建会话
 * @param[IN] online 会话对象
 * @param[OUT] session_id 会话ID
 * @return 错误码
 */
int32_t session_create(const online_t *online, char *session_id);

/**
 * 获取会话
 * @param[IN] session_id 会话ID
 * @param[OUT] online 会话对象
 * @return 错误码
 */
int32_t session_get(const char *session_id, online_t *online);

__END_DECLS

#endif
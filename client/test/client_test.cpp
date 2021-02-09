// Created by longduping 2021-02-09
// C++客户端测试

#include <stdio.h>

#include "session.h"

int main(void)
{
    online_t online;
    char session_id[16];
    int ret = session_create(&online, session_id);
    printf("session_create: ret=%d, session_id=%s\n", ret, session_id);

    return 0;
}
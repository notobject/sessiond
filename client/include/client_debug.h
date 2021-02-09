#ifndef _CLIENT_DEBUG_H 
#define _CLIENT_DEBUG_H 

// Just For Test
#define NAME_SIZE 96
typedef struct _online{
    char user_name[NAME_SIZE + 1];
    int is_login;
    int64_t login_time;
} online_t;

#endif
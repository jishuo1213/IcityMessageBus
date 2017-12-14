#ifndef MQ_CLI_H
#define MQ_CLI_H
#include<pthread.h>

#ifdef __cplusplus
extern "C"
{
#endif
/* mqCli.c */
int connectMq(char *ip, short int port, char *username, char *password);
int disConnectMq(int mqId);
int putMq(int mqId, char *msg, unsigned int msgLen);
int putMqAsync(int mqId, char *msg, unsigned int msgLen);
int putMqAsyncResult (int mqId);
int putMsgToGroup(char *msg, unsigned int msgLen, char **msgGroup);
int putMqGroup(int mqId, char *msgGroup);
int putMqGroupAsync(int mqId, char *msgGroup);
int putMqGroupAsyncResult (int mqId);
int putFileToMq (int mqId, char *fileName);
int putFileMsgToMq (int mqId, char *fileName,char *msg,int msgLen);
int getMq(int mqId, char **retMsg, int *retLen);
int getMqInfo(int mqId, char **retMsg, int *retLen);
int getTopicInfo(int mqId, char *topic,char **retMsg, int *retLen);
int getMqGroup (int mqId, char **msgGroup, unsigned int *retLen);
int getMsgFromGroup (char **retMsg, unsigned int *msgLen, char **msgGroup);
int getMqAsync(int mqId);
int getMqAsyncForAgent(int mqId, int cliSockfd);
int getMqAsyncResult(int mqId, char **retMsg, int *retLen);
/* mqTransCli.c */
int getMqTrans(int mqId, char **retMsg, int *retLen);
int getMqGroupTrans(int mqId, char **retMsg, int *retLen);
int getMqCommit(int mqId);
int getMqRollback(int mqId);

#ifdef __cplusplus
}
#endif

#endif

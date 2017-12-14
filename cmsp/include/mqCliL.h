#ifndef MQ_CLI_L_H
#define MQ_CLI_L_H
#include<pthread.h>
#include <sys/types.h>
#include <unistd.h>


#define MQ_ERR_LSEQ	-1001	//比预期小的错误的消息序号
#define MQ_ERR_BSEQ	-1002	//比预期大的错误的消息序号
#define MQ_ERR_MQID	-1003	//无效的消息队列Id

#ifdef __cplusplus
extern "C"
{
#endif
  typedef struct ms_queue_msg
  {
    unsigned int t;             //消息加入时间
    unsigned int id;            //消息Id,系统自动产生
    unsigned int size;          //消息大小
    unsigned short int mqmId;   //消息队列Id,记录源消息是从哪个队列产生的
    unsigned short int reserve;
    unsigned long nOffset;      //下一条消息偏移量
  } MS_QUEUE_MSG;

  int openCMSP (char *userName,char *password,char *topicName, char **queue);
  int openCmsp (char *topicName, char **queue);
  int openMq (char *queueFileName, char **queue);
  int closeQueue (char *queue, size_t size);
  int getMq (char *queue, char **msg, MS_QUEUE_MSG * mqmHead);
  int putMq (char *queue, char *msg, unsigned int msgSize);
  int getMqGroup (char *queue, char **msg, unsigned int *retSize);
  int getMqGroupTrans (char *queue, char **msg, unsigned int *retSize);
  int getMqCommit(char *queue);
  int getMqRollback(char *queue);
  int getMsgFromGroup(char **retMsg, unsigned int *msgLen, char **msgGroup);
  int getMqInfo (char *queueFileName, char **retMqInfo, int *retLen);
  int putMsgToGroup(char *msg, unsigned int msgLen, char **msgGroup);
  int putMqGroup (char *queue, char *msgGroup);
  int putMqLoop (char *queue, char *msg, unsigned int msgSize);
  int getMqByOffset (char *queue, char **msg, unsigned int *size,unsigned int *id,long *offset);
  int getNextOffsetByOffset (char *queue,long offset,long *nextOffset);
int putMsgToGroup(char *msg, unsigned int msgLen, char **msgGroup);
int getMsgFromGroup(char **retMsg, unsigned int *msgLen, char **msgGroup);

#ifdef __cplusplus
}
#endif

#endif

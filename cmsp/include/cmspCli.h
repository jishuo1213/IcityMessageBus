#ifndef CMSP_CLI_DEV_H
#define CMSP_CLI_DEV_H

#ifdef __cplusplus
extern "C"
{
#endif

typedef struct cli_node_struct
{
char ip[16];
unsigned short int port; 
}CLI_NODE_STRUCT;
/* cliConnect.c */
int connectCMSP(char *ip, unsigned short int port, char *username, char *password, char *topic, int cloud);
int connectCmspByDomain (char *domainName, unsigned short int port, char *username, char *password, char *topic,int cloud);
int disConnectCmsp(int cmspId);
int getIpFromDomain (char *domain, char **ip);
int getQueueFileName(int cmspId, char *topic, char **queueFileName);
int createTopic(int cmspId, char *topic);
int dropTopic(int cmspId, char *topic);
int resizeTopic(int cmspId, char *topic, unsigned int newSize);
int getAllTopic(int cmspId, char **topic, int *size);
#ifdef __cplusplus
}
#endif

#endif

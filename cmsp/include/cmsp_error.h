#ifndef CMSP_ERROR_H
#define CMSP_ERROR_H

#ifdef __cplusplus
extern "C"
{
#endif

#define CMSP_OK		0
#define CMSP_ERROR	-1

#define ERR_OVERFLOW			-99
#define ERR_UNKNOWN			-100
#define ERR_PARAMETER_NULL		-1001
#define ERR_MEMORY_NOT_ENOUGH		-1002
#define ERR_NOT_FOUND			-1003
#define ERR_MAP				-1004
#define ERR_SOCKET_WRITE		-1005
#define ERR_SOCKET_READ			-1006
#define ERR_FILENAME_TOO_LONG		-1007
#define ERR_OPEN_FILE			-1008
#define ERR_FILE_SIZE			-1009
#define ERR_SENDFILE			-1010
#define ERR_NO_CONFIG			-1011
#define ERR_LOGIC			-1012
#define ERR_READ_TRANS			-1013
#define ERR_CONNECT_LN			-1014
#define ERR_REGISTER_IP_INVALID		-1015
#define ERR_TOO_MANY_CONNECT		-1016
#define ERR_LOGIN			-1017
#define ERR_SERVER			-1018
#define ERR_USER_EXIST			-1019
#define ERR_USER_NOT_EXIST		-1020
#define ERR_PARAMETER			-1021
#define ERR_IP_DENY			-1022
#define ERR_IP_NOT_ALLOW		-1023

#define ERR_TOPIC_EXIST			-2001
#define ERR_TOPIC_NOT_EXIST		-2002
#define ERR_TOPIC_OVER_MAX		-2003
#define ERR_TOPIC_EXIST_IN_LN		-2004
#define ERR_TOPIC_AUTHORITY_INVALID	-2005
#define ERR_TOPIC_NOT_ONLINE		-2006
#define ERR_QUEUEFILE_NOT_EXIST		-2010

#define ERR_STACKFILE_NULL		-2101
#define ERR_STACKFILE_ZERO		-2102
#define ERR_STACKFILE_FULL		-2103

#define ERR_QUEUE_EXTEND_TOO_LOW	-2201
#define ERR_QUEUE_EXTEND_TOO_HIGH	-2202

#define ERR_CMS_NOT_EXIST		-2301

#define ERR_CQ_FULL			-2401

#define ERR_DOMAIN_NAME			-9001

#ifdef __cplusplus
}
#endif

#endif

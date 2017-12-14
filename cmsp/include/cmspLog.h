#ifndef CMSP_LOG_H
#define CMSP_LOG_H
#ifdef __cplusplus
extern "C" {
#endif
#define CMSP_LOG                0
#define CMSPLN_LOG              1
#define CMSPSN_LOG              2
#define CMSPRN_LOG              3
#define CMSPDN_LOG              4
#define CMSPPN_LOG              5
#define CMSPTN_LOG              6
#define CMSP_MQTT_LOG            7
#define NC_LOG		        8
#define CMSP_ORACLESYNC_LOG     9
#define CMSP_FILESYNC_LOG       10
#define CMSP_HTTP_LOG          	11
#define CMSP_HTTPS_LOG         	12
#define CMSPBN_LOG              13
#define CMSP_FTP_LOG             14
#define APP_LOG              	20
#define OTH_LOG                 1000

#define ERROR_LOG               0
#define INFORMATION_LOG         1
#define NORMAL_LOG              2
#define NET_LOG 	        3
#define LOG3 		        3
#define DEBUG_LOG4 	        4
#define LOG4 		        4
#define DEBUG_LOG5 	        5
#define LOG5 		        5
#define DEBUG_IMPORTANT	        6
#define LOG6 		        6
#define DEBUG_NORMAL 	        7
#define LOG7 		        7
#define DEBUG_LOG8 	        8
#define DEBUG_LOG9 	        9
#define DEBUG_LOG10 	        10
#define DEBUG_LOG11 	        11
#define DEBUG_LOG12 	        12
#define DEBUG_LOG13 	        13
#define DEBUG_LOG14 	        14
#define DEBUG_LOG15 	        15
#define DEBUG_LOG16 	        16
#define DEBUG_LOG17 	        17
#define DEBUG_LOG18 	        18
#define DEBUG_LOG19 	        19
#define DEBUG_LOG20 	        20
#define DEBUG_USER 	        100
#define DEBUG_EVERYTHING        999

/* cmspLog.c */
int writeLog(int type, int level, char *msg);
int write_log(int type, int level, char *msg);
int appLog(int level, char *msg);
int cmspLog(int level, char *msg);
int cmspLNLog(int level, char *msg);
int cmspPNLog(int level, char *msg);
int cmspSNLog(int level, char *msg);
int cmspRNLog(int level, char *msg);
int cmspdNLog(int level, char *msg);
int cmspTNLog(int level, char *msg);
int cmspMQTTLog(int level, char *msg);
int cmspOracleSyncLog(int level, char *msg);
int cmspFileSyncLog(int level, char *msg);
int cmspBNLog(int level, char *msg);
int cmspHttpLog(int level, char *msg);
int cmspHttpsLog(int level, char *msg);
#ifdef __cplusplus
}
#endif
#endif

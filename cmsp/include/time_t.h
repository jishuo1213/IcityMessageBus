#ifndef MMS_TIME_H
#define MMS_TIME_H
#include<time.h>
#include<sys/time.h>
#include<stdio.h>
//#include<debug.h>

#ifdef __cplusplus
extern "C"
{
#endif

#define T_D struct timeval LCHM_tv1, LCHM_tv2;  
extern struct timeval LCHM_tv1, LCHM_tv2;
extern float TimeDiff (struct timeval LCHM_tv1, struct timeval LCHM_tv2);

#define T_B gettimeofday(&LCHM_tv1,NULL);
#define T_E gettimeofday(&LCHM_tv2,NULL);
//#define T_P {DEBUG_INFO; printf("\n\nTotal cost time is %f seconds!\n\n",TimeDiff(LCHM_tv1,LCHM_tv2));}
#define T_P printf("\n\nTotal cost time is %f seconds!\n\n",TimeDiff(LCHM_tv1,LCHM_tv2));

/*date:YYYY.MM.DD HH:MM:SS,事先需分配好内存*/
char *getDatetimeFromSeconds1(unsigned int time,char *date);

/*date:YYYY-MM-DD HH:MM:SS */
unsigned int getDatetimeSeconds1(char *date);

#ifdef __cplusplus
}
#endif

#endif

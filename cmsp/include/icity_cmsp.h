//
// Created by fan on 17-12-14.
//

#ifndef TESTCMSPNEW_ICITY_CMSP_H_H
#define TESTCMSPNEW_ICITY_CMSP_H_H

char *openQueue(char *topic);

int putOneMessageToQueue(char *queue, char *message, unsigned int size);

char *getOneMessageFromQueue(char *queue, void *readLength);

#endif //TESTCMSPNEW_ICITY_CMSP_H_H

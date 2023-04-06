#ifndef POSAPI_H
#define POSAPI_H

#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

char* checkApi();
char* getInformation();
char* callFunction(char* funcName, char* param);
char* put(char* param);
char* returnBill(char* param);
char* sendData();

#ifdef __cplusplus
}
#endif

#endif /* POSAPI_H */

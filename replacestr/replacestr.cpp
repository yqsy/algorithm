#include <string>
#include <iostream>
#include <string.h>

#include "replacestr.h"

using namespace std;



// replace blank with %20
// data一个有extra capacity的空间,内部含有一个\0末尾的origin字符串
// capacity表示空间的大小,不思考\0
bool replaceBlank(char* data, int capacity, Error* err) {
    if (data == nullptr || capacity < 1) {
        err->err = "input data error";
        return false;
    }

    // 获得原始长度
    // 不包括\0
    int originLen = 0;
    for (int i = 0; i < capacity && data[i]!='\0'; ++i) {
        originLen += 1;
    }

    if (originLen < 1) {
        return true;
    }

    // 空格数量
    int blankCount = 0;
    for (int i = 0; i < originLen; ++i) {
        if (data[i] == ' ') {
            blankCount += 1;
        }
    }

    if (blankCount < 1) {
        return true;
    }

    char replaceStr[] = "%20";
    int newSize = originLen - blankCount + blankCount * strlen(replaceStr);

    if (newSize > capacity) {
        err->err = "capacity size dissatisfy";
        return false;
    }

    int pEnd = newSize - 1;
    for (int i = originLen - 1; i >= 0; --i) {
        if (data[i] == ' ') {
            data[pEnd--] = '0';
            data[pEnd--] = '2';
            data[pEnd--] = '%';
        } else {
            data[pEnd--] = data[i];
        }
    }

    // 补末尾\0
    if (newSize == capacity) {
        // 末尾由capacity之外的\0补充
    } else {
        data[newSize] = '\0';
    }

    return true;
}


#include <string.h>
#include <iostream>

#include "replacestr.h"

using namespace std;


void TestSimple1() {
    // 最前面
    char* buf1 = new char[33];
    strcpy(buf1, " 123123123");

    Error error;
    bool ok = replaceBlank(buf1, 32, &error);

    if (error.err != "" || !ok || strcmp(buf1, "%20123123123") != 0) {
        std::cout << "error: " << error.err << __LINE__ << std::endl;
    }
}

void TestSimple2() {
    // 中间
    char* buf1 = new char[33];
    strcpy(buf1, "123 123123");

    Error error;
    bool ok = replaceBlank(buf1, 32, &error);

    if (error.err != "" || !ok || strcmp(buf1, "123%20123123") != 0) {
        std::cout << "error: " << error.err << __LINE__ << std::endl;
    }
}

void TestSimple3() {
    // 最后面
    char* buf1 = new char[33];
    strcpy(buf1, "123123123 ");

    Error error;
    bool ok = replaceBlank(buf1, 32, &error);

    if (error.err != "" || !ok || strcmp(buf1, "123123123%20") != 0) {
        std::cout << "error: " << error.err << __LINE__ << std::endl;
    }
}

void TestSimple4() {
    // 连续多个空格
    char* buf1 = new char[33];
    strcpy(buf1, "123123   123");

    Error error;
    bool ok = replaceBlank(buf1, 32, &error);

    if (error.err != "" || !ok || strcmp(buf1, "123123%20%20%20123") != 0) {
        std::cout << "error: " << error.err << __LINE__ << std::endl;
    }
}

void TestSimple5() {
    // 没有空格
    char* buf1 = new char[33];
    strcpy(buf1, "123123123");

    Error error;
    bool ok = replaceBlank(buf1, 32, &error);

    if (!ok || strcmp(buf1, "123123123") != 0) {
        std::cout << "error: " << error.err << __LINE__ << std::endl;
    }
}

void TestSimple6() {
    // 空指针
    char* buf1;

    Error error;
    bool ok = replaceBlank(buf1, 0, &error);

    if (ok) {
        std::cout << "error: " << error.err << __LINE__ << std::endl;
    }
}

void TestSimple7() {
    // 空字符串
    char* buf1 = "";

    Error error;
    bool ok = replaceBlank(buf1, 0, &error);

    if (ok) {
        std::cout << "error: " << error.err << __LINE__ << std::endl;
    }
}

void TestSimple8() {
    // 只有一个空格
    char* buf1 = new char[33];
    strcpy(buf1, " ");

    Error error;
    bool ok = replaceBlank(buf1, 32, &error);

    if (!ok || strcmp(buf1, "%20") != 0) {
        std::cout << "error: " << error.err << __LINE__ << std::endl;
    }
}

void TestSimple9() {
    // 极值
    char* buf1 = new char[17];
    strcpy(buf1, "123456789123456 ");

    Error error;
    bool ok = replaceBlank(buf1, 16, &error);

    if (ok || strcmp(buf1, "123456789123456 ") != 0) {
        std::cout << "error: " << error.err << __LINE__ << std::endl;
    }
}

void TestSimple10() {
    // 极值
    char* buf1 = new char[17];
    strcpy(buf1, "1234567891234 ");

    Error error;
    bool ok = replaceBlank(buf1, 16, &error);

    if (!ok || strcmp(buf1, "1234567891234%20") != 0) {
        std::cout << "error: " << error.err << __LINE__ << std::endl;
    }
}

void TestSimple11() {
    // 极值2
    char* buf1 = new char[17];
    strcpy(buf1, "123456789123 ");

    Error error;
    bool ok = replaceBlank(buf1, 16, &error);

    if (!ok || strcmp(buf1, "123456789123%20") != 0) {
        std::cout << "error: " << error.err << __LINE__ << std::endl;
    }
}


int main() {
    // 不考虑内存泄露

    TestSimple1();
    TestSimple2();
    TestSimple3();
    TestSimple4();
    TestSimple5();
    TestSimple6();
    TestSimple7();
    TestSimple8();
    TestSimple9();
    TestSimple10();

    return 0;
}
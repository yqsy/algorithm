#include <string>

struct Error {
    std::string err;
};

bool replaceBlank(char* data, int capacity, Error* err);

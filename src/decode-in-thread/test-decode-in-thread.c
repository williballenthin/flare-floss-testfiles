#include <stdio.h>
#include <string.h>
#include <stdint.h>
#include <stdlib.h>

#ifdef _WIN32
#include <windows.h>
#define THREAD_PROC DWORD WINAPI
#define THREAD_DONE 0
#else
#include <pthread.h>
#define THREAD_PROC void *
#define THREAD_DONE NULL
#endif

// "http://evil.example", each byte xor 0x42, including the terminator.
uint8_t encoded[] = {
    'h' ^ 0x42, 't' ^ 0x42, 't' ^ 0x42, 'p' ^ 0x42, ':' ^ 0x42, '/' ^ 0x42, '/' ^ 0x42,
    'e' ^ 0x42, 'v' ^ 0x42, 'i' ^ 0x42, 'l' ^ 0x42, '.' ^ 0x42,
    'e' ^ 0x42, 'x' ^ 0x42, 'a' ^ 0x42, 'm' ^ 0x42, 'p' ^ 0x42, 'l' ^ 0x42, 'e' ^ 0x42,
    0x42,
};
char url[sizeof(encoded)];

// a thread start routine that decodes into a global buffer.
// nothing in the program calls it: the operating system does, when the thread starts.
THREAD_PROC decode_url(void *param) {
    for (size_t i = 0; i < sizeof(encoded); i++) {
        url[i] = encoded[i] ^ 0x42;
    }
    return THREAD_DONE;
}

int main(int argc, char **argv) {
#ifdef _WIN32
    HANDLE thread = CreateThread(NULL, 0, decode_url, NULL, 0, NULL);
    if (!thread) {
        perror("failed to start thread.\n");
        return -1;
    }
    WaitForSingleObject(thread, INFINITE);
    CloseHandle(thread);
#else
    pthread_t thread;
    if (pthread_create(&thread, NULL, decode_url, NULL)) {
        perror("failed to start thread.\n");
        return -1;
    }
    pthread_join(thread, NULL);
#endif
    printf("%s\n", url);
    return 0;
}

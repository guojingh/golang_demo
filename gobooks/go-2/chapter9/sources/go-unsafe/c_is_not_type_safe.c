#include <stdio.h>

// c语言是一门静态类型语言，但它不是类型安全的语言，因为我们可以通过像下面的示例通过合法的语法轻易“刺透”其类型系统
int main() {
    int a = 0x12345678;
    unsigned char *p = (unsigned char*)&a;
    printf("0x%x\n", a);  //0x12345678
    *p = 0x23;
    *(p+1) = 0x45;
    *(p+2) = 0x67;
    *(p+3) = 0x8a;
    printf("0x%x\n", a);  //0x8a674523（注：在小端字节序列系统中输出此值）
}
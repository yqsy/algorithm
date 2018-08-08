#include <math.h>
#include <iostream>
#include <stdio.h>

// q: 攻击者找到下一区块的概率
// z: 块数
double AttackerSuccessProbability(double q, int z) {
    // p: 诚实者找到下一区块的概率
    double p = 1.0 - q;

    // 泊松分布的期望值
    double lambda = z * (q / p);

    double sum = 1.0;

    int i, k;
    for (k = 0; k <= z; k++) {
        // 泊松 
        double poisson = exp(-lambda);

        for (i = 1; i <= k; i++) {
            poisson *= lambda / i;
        }

        sum -= poisson * (1 - pow(q / p, z - k));
    }

    return sum;
}   


int main(int argc, char* argv[]) {

    printf("q=0.1\n");
    for (int i = 0; i < 11; i++) {
        printf("z=%d P=%.7f\n", i, AttackerSuccessProbability(0.1, i));
    }

    printf("\n");

    printf("q=0.3\n");
    for (int i = 0; i < 51; i+=5) {
        printf("z=%d P=%.7f\n", i, AttackerSuccessProbability(0.3, i));
    }
    
    printf("\n");

    printf("P < 0.001\n");
    for (double q=0.10; q < 0.451; q += 0.05) {
        
        // loop forever
        for(int z = 0; ; z++) {
            
            double P = AttackerSuccessProbability(q, z);
            
            if (P < 0.001) {
                printf("q=%.2f z=%d\n", q, z);
                break;
            }
        }
    }

    return 0;
}

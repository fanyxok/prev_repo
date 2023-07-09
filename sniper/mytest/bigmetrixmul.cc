
int main() {
    int a[512][1024];
    int b[1024][512];

    for (int i = 0; i < 512; i++) {
        for (int j = 0; j < 1024; j++ ) {
            a[i][j] = i%2;
            b[i][j] = j%2;
        }
    }
    for (int i = 0; i < 1024; i++) {
        for (int j = 0; j < 512; j++) {
            int c = a[j][i] * b[i][j];
            if ( c == 0 ) {
                break;
            }
        }
    }

}
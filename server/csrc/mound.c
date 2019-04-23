#include <stdio.h>

// 堆排序算法关键
// 节点点个数： |length / 2|
// 叶子结点： length - |length / 2|
// 从下往上，从右往左遍历节点
// 节点与叶子节点进行比较大小：m与2m + 1 （左节点） 2m+ 2（右节点）
// 出现交换，交换完成需要飙孙子节点，循环交换，直至2m + 1 不大于length
void swap(int arr[], int m, int n){
    int temp = arr[m];
    arr[m] = arr[n];
    arr[n] = temp;
};

void print(int a[], int len){
    int i;
    for (i = 0; i < len; i++) {
        printf("%d,", a[i]);
    }
    printf("\n");
}

void MaxMound(int arr[], int m, int length){
    int temp = arr[m];
    for(int k = m * 2 + 1; k < length; k = k * 2 + 1) { // k * 2 + 1 比较孙子节点
        if(k + 1 < length && arr[k] < arr[k+1]) {
            k++;
        }
        if(arr[k] > temp){
            arr[m] = arr[k];
            m = k; 
        } else {
            break;
        }
    }
    arr[m] = temp;
};

void MinMound(int arr[], int m, int length){
    int temp = arr[m];
    for (int k = m * 2 + 1; k < length; k = k * 2 + 1){
        if (k + 1 < length && arr[k] > arr[k+1]) {
            k++;
        }
        if(arr[k] < temp){
            arr[m] = arr[k];
            m =k;
        } else {
            break;
        }
    }
    arr[m] = temp;
}

void heap_sort(int arr[], int len, void (*hans)(int arr[], int m, int l)){
    int i;

    for(i = (len / 2) - 1; i >= 0; i--) {
        hans(arr, i, len);
    }
    for (i = len - 1; i > 0; i--){
        swap(arr, 0, i);
        hans(arr, 0, i);
    }
};
void max_heapify(int arr[], int start, int end) {
    //建立父节点指标和子节点指标
    int dad = start;
    int son = dad * 2 + 1;
    while (son <= end) { //若子节点指标在范围内才做比较
        if (son + 1 <= end && arr[son] < arr[son + 1]) //先比较两个子节点大小，选择最大的
            son++;
        if (arr[dad] > arr[son]) //如果父节点大於子节点代表调整完毕，直接跳出函数
            return;
        else { //否则交换父子内容再继续子节点和孙节点比较
            swap(arr,dad, son);
            dad = son;
            son = dad * 2 + 1;
        }
    }
}
void hedp_sort(int arr[], int len) {
    int i;
    //初始化，i从最後一个父节点开始调整
    for (i = len / 2 - 1; i >= 0; i--)
        max_heapify(arr, i, len - 1);
    //先将第一个元素和已排好元素前一位做交换，再重新调整，直到排序完毕
    for (i = len - 1; i > 0; i--) {
        swap(arr, 0, i);
        max_heapify(arr, 0, i - 1);
    }
}
 

int main() {
    int arr[] = {3,5,1,7,9,6,1,3,4};
    int brr[] = {3,5,1,7,9,6,1,3,4};
    int len = (int)sizeof(arr) / sizeof(*arr);
    hedp_sort(arr, len);
    print(arr, len);

    heap_sort(arr, len, &MinMound);
    print(arr, len);
    return 0;
}
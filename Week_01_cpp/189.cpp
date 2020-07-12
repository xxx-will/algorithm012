
void rotate(vector<int>& nums, int k) {

    int len = nums.size();
    if (k <= 0 || len == 0) return;

    k = k % len;

    int tmp = 0;
    for（int j= 0; j<k; j++ ）{
        int pre = nums[len -1];
        for (int i = 0; i < len; i++) {
            tmp = nums[i];
            nums[i]=pre;
            pre = tmp;
        }
    }
   

    int n = nums.size();
    k = k % n;
    int cnt;

    for (int start = 0; cnt < n; start++) {
        int cur = start;
        int pre = nums[start];
        do{
            int next = (cur+k)%n; 
            int temp = nums[next];
            nums[next]= pre;
            pre = temp;
            cur = next;
            cnt++;
        }while(start != cur);
    }


    reverse([]int nums, start, end)

}
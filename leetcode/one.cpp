#include<iostream>
#include<vector>
using namespace std;

int valueAfterKSeconds(int n, int k) {
        
        vector<long long> arr(n,1);
        long long l= 1000000000+7;
        k--;
        while(k--){
            
            for(int i=n-1; i>=0; i--){
                long long sum=arr[i];
                
                for(int j=0; j<i; j++){
                    sum += arr[j]%l;
                    sum=sum%l;
                }
                arr[i]=sum;
            
            }
        }
        long long ans=0;

        for(int i=0; i<n; i++){
            ans+=arr[i]%l;
        }
        return ans;
        
}
int main(){

    int n,k;
    cin>>n>>k;

    cout<<valueAfterKSeconds(n,k);

}
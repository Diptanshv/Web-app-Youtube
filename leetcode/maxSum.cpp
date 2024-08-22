#include<iostream>
#include<vector>
#include<math.h>
#include<array>
using namespace std;

int main(){

    int t;
    cin>>t;
    int n,k;
    cin>>n>>k;
    int m = 1000000007;
    int a[n];
    vector<int> ans(n+k,0);

    //int total=0;
    int max_so_far=INT_MIN;
    int max_end_here=0;

    for(int i=0; i<n;i++){
        cin>>a[i];
        ans[i]=a[i];
        max_end_here=max_end_here+a[i];
        if(max_so_far<max_end_here) max_so_far=max_end_here;
        if(max_end_here<0) max_end_here=0;
        
    }

    while(k)
}
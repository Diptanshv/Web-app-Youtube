        vector<int> arr(n,1);
        int l= 1000000000+7;
        k--;
        while(k--){
            
            for(int i=n-1; i>=0; i--){
                int sum=arr[i];
                
                for(int j=0; j<i; j++){
                    sum+=arr[j]%l;
                }
                arr[i]=sum;
            
            }
        }
        int ans=0;

        for(int i=0; i<n; i++){
            ans+=arr[i]%l;
        }
        return ans;
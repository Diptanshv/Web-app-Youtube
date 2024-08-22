#include <iostream>
#include<stack>
using namespace std;
int main() {
   int hash=0;
   int haash=0;

   string s1="bvvv";
   string s2="xxxw";

   for(int i=0; i<s1.length();i++){
       hash+=(i+1)*(s1[i]-'a'+1);
   }
   for(int i=0; i<s2.length();i++){
       haash+=(i+1)*(s2[i]-'a'+1);
   }

    if(haash==hash){
        cout<<"true";
    }else{
        cout<<"false";
    }
}
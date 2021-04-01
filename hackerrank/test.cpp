#include
#define ll long long int
using namespace std;

ll ans;
void dfs(vector<vector>&adj,ll u,ll ht){
ans=max(ans,ht);
for(auto x:adj[u]){
dfs(adj,x,ht+1);
}
}
int main() {
ios_base::sync_with_stdio(false);
cin.tie(0);
cout.tie(0);
ll t,n;
cin>>t;
while(t–){
cin>>n;
vectorvec(n);
vector<vector>adj(n);
vectorroot;
for(ll i=0;i>vec[i];
if(vec[i]==-1){
root.push_back(i);
}
else{
adj[vec[i]].push_back(i);
}
}
ans=0;
for(ll i=0;i<root.size();i++){
dfs(adj,root[i],1);
}
cout<<ans<<endl;
}

}
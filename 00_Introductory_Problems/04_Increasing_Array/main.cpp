#include <bits/stdc++.h>
using namespace std;
int main()
{
    int n;
    cin >> n;
    vector<int> a(n, 0);
    long long int ans = 0;

    for (int i = 0; i < n; i++)
        cin >> a[i];

    for (int i = 1; i < n; i++)
    {
        if (a[i - 1] > a[i])
        {
            ans += a[i - 1] - a[i];
            a[i] = a[i - 1];
        }
    }
    cout << ans << "\n";
}
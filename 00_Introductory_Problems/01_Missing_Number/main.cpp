#include <bits/stdc++.h>
using namespace std;
int main()
{
    long n;
    cin >> n;

    vector<long> a(n - 1, 0);
    long long int prod = n * (n + 1) / 2;

    for (int i = 0; i < n - 1; i++)
    {
        cin >> a[i];
    }

    prod -= accumulate(a.begin(), a.end(), 0LL);
    cout << prod << "\n";
}
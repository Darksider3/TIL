# std::sort(it.begin(), it, end()
Sorts, for an example an integer, vector and places the highest value at the end of it.

```
int birthdayCakeCandles(vector<int> ar) {
    std::sort(ar.begin(), ar.end());
    return std::count(ar.begin(), ar.end(), ar[ar.size()-1]);
}
```

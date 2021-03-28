### 문제
https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

### 아이디어
어렵지 않게...생각하려 했다.
그냥 이중 for문으로 찾아봤다.

### 풀이
```javascript
var twoSum = function(numbers, target) {
    for(let i=0;i<numbers.length;i++){
        for(let j=i+1;j<numbers.length;j++) {
            if(numbers[i]+numbers[j]==target) {
                return [i+1,j+1] ;
            }
        }
    }
};
```

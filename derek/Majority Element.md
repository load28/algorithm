### 문제
https://leetcode.com/problems/majority-element/

### 아이디어
숫자들을 순회하면서 각 숫자들의 수를 세어준다.
그리고 가장 큰 숫자를 찾아주면 된다.

### 풀이

```javascript
var majorityElement = function(nums) {
  // 해시 맵을 이용 할 수도 있지만 자바스크립트의 객체를 이용해서 동일한 속도가 나온다.
  let mapObj = {};
  for (let i=0 ; i < nums.length ; i++){
      if(!mapObj[nums[i]]){
          mapObj[nums[i]] = 0;
      }
      mapObj[nums[i]] += 1;
  }
    
  return Object.keys(mapObj).reduce((a, b) => mapObj[a] > mapObj[b] ? a : b);
};
```

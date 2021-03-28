### 문제
https://leetcode.com/problems/longest-substring-without-repeating-characters/

### 아이디어
초반에는 이중 for문으로 해결하려 했다.  
그렇게 작성하니깐 코드안에서 조건이 너무 복잡해 졌다.  
한번만 순회하면서 값을 찾는 방법을 생각하다가 생각이 나지 않아서, 다른 사람들의 코드를 보며 분석했다.  
포인트는 **순회 하면서 각 값을 Map에 대입하며** 만약 기존에 존재하는 값이 있는 경우 이미 존재하던 값의 인덱스를 가져와서  
이 이후부터 길이를 체크하는 방법이였다.  


### 풀이
```javascript
var lengthOfLongestSubstring = function(s) {
    let map=new Map();
    let d=0;
    let result=0;
    
    for(let p=0; p<s.length;end++){
        // 존재한다면 기존에 체크한 존재하던 인덱스 값과 비교해서 더 큰 값을 넣어서 체크하는 인덱스 값의 포지션을 이동시킨다
        if (map.has(s[p])) {
            d=Math.max(d, map.get(s[p])+1);
        }
        // 중복된 값이 있다면 나중에 확인된 중복된 인덱스의 위치로 덮어씌워 진다
        map.set(s[p], p);
        // 이 곳에서 p-d+1을 통해서 길이를 계산하여 기존 길이와 비교하여 더 큰값을 대입한다
        // 결국 p는 현재 바라보고 있는 문자이며, d라는 값은 중복여부를 통해서 포지션이 재계산 되고 p-d를 통해서 길이를 측정 할 수 있다.
        result=Math.max(result, p-d+1);
    }
    return result;
};
```

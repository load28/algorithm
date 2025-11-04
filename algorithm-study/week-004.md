# 4주차: 이진 탐색 - 1부

## 학습 목표
- 선형 탐색과 이진 탐색의 차이를 이해한다
- 이진 탐색 알고리즘을 구현할 수 있다
- Lower Bound와 Upper Bound 개념을 이해한다
- 이진 탐색의 시간복잡도를 분석할 수 있다
- 이진 탐색을 활용한 문제를 풀 수 있다

> 💡 **왜 중요한가?** 이진 탐색은 O(n)을 O(log n)으로 줄이는 강력한 알고리즘입니다!

---

## 📑 목차

1. [탐색이란?](#1-탐색이란)
2. [선형 탐색 (Linear Search)](#2-선형-탐색-linear-search)
3. [이진 탐색 (Binary Search)](#3-이진-탐색-binary-search)
4. [이진 탐색 구현 패턴](#4-이진-탐색-구현-패턴)
5. [Lower Bound와 Upper Bound](#5-lower-bound와-upper-bound)
6. [이진 탐색 활용](#6-이진-탐색-활용)
7. [실전 예제](#7-실전-예제)
8. [코딩테스트 팁](#8-코딩테스트-팁)

---

## 1. 탐색이란?

### 1.1 탐색의 정의

**탐색(Search)**: 데이터 집합에서 원하는 값을 찾는 것

```python
# 예시: 배열에서 5 찾기
arr = [1, 3, 5, 7, 9, 11, 13]
target = 5
# 어떻게 찾을까?
```

### 1.2 탐색의 중요성

**데이터베이스**:
- 사용자 검색
- 상품 검색
- 로그 검색

**알고리즘**:
- 많은 문제가 탐색을 포함
- 탐색 속도가 전체 성능 결정

**코딩테스트**:
- 매우 자주 출제
- 다른 알고리즘과 결합

### 1.3 탐색 방법

**선형 탐색 (Linear Search)**:
- 처음부터 끝까지 순차적으로 찾기
- O(n)

**이진 탐색 (Binary Search)**:
- 정렬된 데이터에서 절반씩 줄여가며 찾기
- O(log n)

**해시 탐색 (Hash Search)**:
- 해시 테이블 사용
- O(1)

> 이번 주는 **이진 탐색**에 집중합니다!

---

## 2. 선형 탐색 (Linear Search)

### 2.1 선형 탐색이란?

**원리**: 배열의 처음부터 끝까지 하나씩 비교

```
[1, 3, 5, 7, 9, 11, 13]  target = 7
 ↓
[1, 3, 5, 7, 9, 11, 13]  1 ≠ 7
    ↓
[1, 3, 5, 7, 9, 11, 13]  3 ≠ 7
       ↓
[1, 3, 5, 7, 9, 11, 13]  5 ≠ 7
          ↓
[1, 3, 5, 7, 9, 11, 13]  7 == 7 ✅ 찾음!
```

### 2.2 선형 탐색 구현

```python
def linear_search(arr, target):
    """
    선형 탐색
    시간복잡도: O(n)
    """
    for i in range(len(arr)):
        if arr[i] == target:
            return i  # 인덱스 반환
    return -1  # 못 찾음

# 테스트
arr = [1, 3, 5, 7, 9, 11, 13]
print(linear_search(arr, 7))   # 3
print(linear_search(arr, 10))  # -1
```

**파이썬 내장 방법**:
```python
# in 연산자 (True/False)
target = 7
if target in arr:
    print("존재함")

# index() 메서드 (인덱스 반환)
try:
    idx = arr.index(target)
    print(f"인덱스: {idx}")
except ValueError:
    print("없음")
```

### 2.3 선형 탐색의 특징

**장점**:
- ✅ 구현이 매우 간단
- ✅ 정렬되지 않은 데이터에도 사용 가능
- ✅ 작은 데이터셋에서는 빠름

**단점**:
- ❌ 데이터가 많으면 느림 O(n)
- ❌ 최악의 경우 모든 요소를 확인해야 함

**복잡도**:
- 시간복잡도: O(n)
- 공간복잡도: O(1)

### 2.4 선형 탐색이 필요한 경우

```python
# 1. 정렬되지 않은 데이터
arr = [5, 2, 9, 1, 7, 3]  # 정렬 안 됨
# 이진 탐색 불가, 선형 탐색 사용

# 2. 데이터가 적은 경우 (n < 100)
arr = [1, 2, 3, 4, 5]
# 이진 탐색보다 선형 탐색이 오히려 빠를 수 있음

# 3. 링크드 리스트
# 랜덤 접근 불가 → 이진 탐색 불가
```

---

## 3. 이진 탐색 (Binary Search)

### 3.1 이진 탐색이란?

**원리**: 정렬된 배열에서 **중간값**과 비교하여 탐색 범위를 절반씩 줄임

**전제 조건**: ⚠️ **배열이 정렬되어 있어야 함!**

```
[1, 3, 5, 7, 9, 11, 13]  target = 7
         ↑ mid = 7
         7 == 7 ✅ 찾음! (1번 만에!)
```

```
[1, 3, 5, 7, 9, 11, 13]  target = 11
         ↑ mid = 7
         11 > 7 → 오른쪽 절반만 탐색
                  [9, 11, 13]
                      ↑ mid = 11
                      11 == 11 ✅ 찾음! (2번 만에!)
```

### 3.2 이진 탐색 과정

```
[1, 3, 5, 7, 9, 11, 13]  target = 3
 L           M        R
         3 < 7 → 왼쪽 절반

[1, 3, 5]
 L  M  R
    3 == 3 ✅ 찾음!
```

### 3.3 이진 탐색 구현 (반복문)

```python
def binary_search(arr, target):
    """
    이진 탐색 (반복문 버전)
    시간복잡도: O(log n)
    """
    left = 0
    right = len(arr) - 1

    while left <= right:
        mid = (left + right) // 2

        if arr[mid] == target:
            return mid  # 찾음!
        elif arr[mid] < target:
            left = mid + 1  # 오른쪽 절반 탐색
        else:
            right = mid - 1  # 왼쪽 절반 탐색

    return -1  # 못 찾음

# 테스트
arr = [1, 3, 5, 7, 9, 11, 13]
print(binary_search(arr, 7))   # 3
print(binary_search(arr, 3))   # 1
print(binary_search(arr, 10))  # -1
```

**단계별 실행 과정**:
```python
# arr = [1, 3, 5, 7, 9, 11, 13], target = 3
#
# 1단계:
#   left=0, right=6, mid=3
#   arr[3]=7, 3 < 7 → right = mid - 1 = 2
#
# 2단계:
#   left=0, right=2, mid=1
#   arr[1]=3, 3 == 3 → return 1
```

### 3.4 이진 탐색 구현 (재귀)

```python
def binary_search_recursive(arr, target, left, right):
    """
    이진 탐색 (재귀 버전)
    시간복잡도: O(log n)
    공간복잡도: O(log n) - 재귀 스택
    """
    # 기저 조건: 범위가 유효하지 않음
    if left > right:
        return -1

    mid = (left + right) // 2

    if arr[mid] == target:
        return mid
    elif arr[mid] < target:
        return binary_search_recursive(arr, target, mid + 1, right)
    else:
        return binary_search_recursive(arr, target, left, mid - 1)

# 테스트
arr = [1, 3, 5, 7, 9, 11, 13]
print(binary_search_recursive(arr, 7, 0, len(arr) - 1))   # 3
print(binary_search_recursive(arr, 10, 0, len(arr) - 1))  # -1
```

### 3.5 파이썬 내장 이진 탐색

```python
import bisect

arr = [1, 3, 5, 7, 9, 11, 13]

# 값의 위치 찾기 (정확히 일치하는 값)
# bisect 모듈은 인덱스 반환이 아니라 삽입 위치를 반환
idx = bisect.bisect_left(arr, 7)
if idx < len(arr) and arr[idx] == 7:
    print(f"찾음: 인덱스 {idx}")  # 찾음: 인덱스 3

# 값이 들어갈 위치 찾기
pos = bisect.bisect_left(arr, 6)
print(f"6이 들어갈 위치: {pos}")  # 6이 들어갈 위치: 3
```

### 3.6 이진 탐색의 시간복잡도

**매번 절반씩 줄어듦**:
```
n → n/2 → n/4 → n/8 → ... → 1
```

**연산 횟수**:
```
n = 1,000,000
선형 탐색: 최악 1,000,000번
이진 탐색: 최악 20번 (log₂ 1,000,000 ≈ 20)
```

**복잡도**:
- 시간복잡도: O(log n)
- 공간복잡도:
  - 반복문: O(1)
  - 재귀: O(log n)

**비교**:
| n | 선형 탐색 | 이진 탐색 |
|---|----------|----------|
| 10 | 10 | 4 |
| 100 | 100 | 7 |
| 1,000 | 1,000 | 10 |
| 1,000,000 | 1,000,000 | 20 |
| 1,000,000,000 | 1,000,000,000 | 30 |

→ **엄청난 차이!**

---

## 4. 이진 탐색 구현 패턴

### 4.1 mid 계산 방법

**일반적인 방법**:
```python
mid = (left + right) // 2
```

**오버플로우 방지** (큰 수에서):
```python
mid = left + (right - left) // 2
```

> 💡 파이썬은 int 크기 제한이 없어서 오버플로우 걱정 없지만, 다른 언어(C++, Java)에서는 중요!

### 4.2 경계 조건

**while left <= right** (일반적):
```python
while left <= right:
    mid = (left + right) // 2
    if arr[mid] == target:
        return mid
    elif arr[mid] < target:
        left = mid + 1
    else:
        right = mid - 1
return -1
```

**while left < right** (특수한 경우):
```python
while left < right:
    mid = (left + right) // 2
    if arr[mid] < target:
        left = mid + 1
    else:
        right = mid
return left
```

### 4.3 이진 탐색 템플릿

**기본 템플릿**:
```python
def binary_search_template(arr, target):
    left, right = 0, len(arr) - 1

    while left <= right:
        mid = (left + right) // 2

        if arr[mid] == target:
            return mid  # 또는 원하는 처리
        elif arr[mid] < target:
            left = mid + 1
        else:
            right = mid - 1

    return -1  # 또는 left, right
```

### 4.4 자주 하는 실수

**❌ 실수 1: 무한 루프**
```python
# 잘못된 코드
while left < right:
    mid = (left + right) // 2
    if arr[mid] < target:
        left = mid  # ❌ mid + 1이 아니라 mid
    else:
        right = mid - 1
```

**✅ 올바른 코드**:
```python
while left < right:
    mid = (left + right) // 2
    if arr[mid] < target:
        left = mid + 1  # ✅
    else:
        right = mid
```

**❌ 실수 2: 범위 초과**
```python
# 잘못된 코드
while left <= right:
    mid = (left + right) // 2
    if arr[mid] == target:
        return mid
    elif arr[mid] < target:
        left = mid  # ❌ +1 빠짐
    else:
        right = mid  # ❌ -1 빠짐
# 무한 루프!
```

**❌ 실수 3: 정렬 안 된 배열**
```python
arr = [5, 2, 9, 1, 7]  # ❌ 정렬 안 됨
binary_search(arr, 7)  # 잘못된 결과!

# ✅ 먼저 정렬
arr.sort()
binary_search(arr, 7)  # 올바른 결과
```

---

## 5. Lower Bound와 Upper Bound

### 5.1 Lower Bound란?

**Lower Bound**: target 이상인 값이 처음 나타나는 위치

```
[1, 3, 3, 3, 5, 7, 9]  target = 3
    ↑
    Lower Bound = 1 (3이 처음 나타나는 위치)

[1, 3, 3, 3, 5, 7, 9]  target = 4
                ↑
    Lower Bound = 4 (4 이상인 값이 처음 나타나는 위치)
```

### 5.2 Lower Bound 구현

```python
def lower_bound(arr, target):
    """
    target 이상인 값이 처음 나타나는 인덱스
    없으면 len(arr) 반환
    """
    left, right = 0, len(arr)

    while left < right:
        mid = (left + right) // 2

        if arr[mid] < target:
            left = mid + 1
        else:
            right = mid

    return left

# 테스트
arr = [1, 3, 3, 3, 5, 7, 9]
print(lower_bound(arr, 3))   # 1
print(lower_bound(arr, 4))   # 4
print(lower_bound(arr, 0))   # 0
print(lower_bound(arr, 10))  # 7 (len(arr))
```

**파이썬 내장 함수**:
```python
import bisect

arr = [1, 3, 3, 3, 5, 7, 9]
print(bisect.bisect_left(arr, 3))   # 1
print(bisect.bisect_left(arr, 4))   # 4
```

### 5.3 Upper Bound란?

**Upper Bound**: target을 초과하는 값이 처음 나타나는 위치

```
[1, 3, 3, 3, 5, 7, 9]  target = 3
             ↑
    Upper Bound = 4 (3을 초과하는 값이 처음 나타나는 위치)

[1, 3, 3, 3, 5, 7, 9]  target = 4
                ↑
    Upper Bound = 4 (4를 초과하는 값이 처음 나타나는 위치)
```

### 5.4 Upper Bound 구현

```python
def upper_bound(arr, target):
    """
    target을 초과하는 값이 처음 나타나는 인덱스
    없으면 len(arr) 반환
    """
    left, right = 0, len(arr)

    while left < right:
        mid = (left + right) // 2

        if arr[mid] <= target:
            left = mid + 1
        else:
            right = mid

    return left

# 테스트
arr = [1, 3, 3, 3, 5, 7, 9]
print(upper_bound(arr, 3))   # 4
print(upper_bound(arr, 4))   # 4
print(upper_bound(arr, 0))   # 0
print(upper_bound(arr, 10))  # 7 (len(arr))
```

**파이썬 내장 함수**:
```python
import bisect

arr = [1, 3, 3, 3, 5, 7, 9]
print(bisect.bisect_right(arr, 3))  # 4
print(bisect.bisect(arr, 3))        # 4 (bisect_right와 동일)
```

### 5.5 Lower vs Upper Bound 차이

```python
arr = [1, 3, 3, 3, 5, 7, 9]

# Lower Bound (≥)
print(bisect.bisect_left(arr, 3))   # 1 (3 이상)

# Upper Bound (>)
print(bisect.bisect_right(arr, 3))  # 4 (3 초과)

# 개수 세기
count = bisect.bisect_right(arr, 3) - bisect.bisect_left(arr, 3)
print(f"3의 개수: {count}")  # 3의 개수: 3
```

### 5.6 Lower/Upper Bound 활용

**특정 값의 개수 세기**:
```python
def count_range(arr, target):
    """
    정렬된 배열에서 target의 개수
    시간복잡도: O(log n)
    """
    left = bisect.bisect_left(arr, target)
    right = bisect.bisect_right(arr, target)
    return right - left

arr = [1, 3, 3, 3, 5, 7, 9]
print(count_range(arr, 3))  # 3
print(count_range(arr, 5))  # 1
print(count_range(arr, 10)) # 0
```

**범위 내 값의 개수**:
```python
def count_range_between(arr, left_val, right_val):
    """
    정렬된 배열에서 [left_val, right_val] 범위의 개수
    """
    left = bisect.bisect_left(arr, left_val)
    right = bisect.bisect_right(arr, right_val)
    return right - left

arr = [1, 3, 3, 3, 5, 7, 9]
print(count_range_between(arr, 3, 5))  # 4 (3, 3, 3, 5)
print(count_range_between(arr, 4, 8))  # 2 (5, 7)
```

**값 삽입 위치 찾기**:
```python
arr = [1, 3, 5, 7, 9]

# 4를 삽입할 위치
pos = bisect.bisect_left(arr, 4)
arr.insert(pos, 4)
print(arr)  # [1, 3, 4, 5, 7, 9]
```

---

## 6. 이진 탐색 활용

### 6.1 정렬된 배열에서 탐색

```python
# 학생 점수 리스트에서 특정 점수 찾기
scores = [45, 67, 72, 85, 88, 92, 95]
target_score = 85

idx = binary_search(scores, target_score)
if idx != -1:
    print(f"점수 {target_score}는 {idx}번 인덱스에 있습니다.")
else:
    print("해당 점수가 없습니다.")
```

### 6.2 중복된 값 처리

```python
# 중복된 값이 있을 때 첫 번째 위치 찾기
def find_first(arr, target):
    idx = bisect.bisect_left(arr, target)
    if idx < len(arr) and arr[idx] == target:
        return idx
    return -1

# 중복된 값이 있을 때 마지막 위치 찾기
def find_last(arr, target):
    idx = bisect.bisect_right(arr, target) - 1
    if idx >= 0 and arr[idx] == target:
        return idx
    return -1

arr = [1, 3, 3, 3, 5, 7, 9]
print(find_first(arr, 3))  # 1
print(find_last(arr, 3))   # 3
```

### 6.3 가장 가까운 값 찾기

```python
def find_closest(arr, target):
    """
    정렬된 배열에서 target에 가장 가까운 값 찾기
    """
    if not arr:
        return None

    # Lower Bound 위치 찾기
    idx = bisect.bisect_left(arr, target)

    # 경계 처리
    if idx == 0:
        return arr[0]
    if idx == len(arr):
        return arr[-1]

    # 앞뒤 값 중 더 가까운 값
    before = arr[idx - 1]
    after = arr[idx]

    if abs(before - target) <= abs(after - target):
        return before
    return after

arr = [1, 3, 5, 7, 9, 11, 13]
print(find_closest(arr, 6))   # 5 또는 7 (여기선 5)
print(find_closest(arr, 8))   # 7 또는 9 (여기선 7)
print(find_closest(arr, 0))   # 1
print(find_closest(arr, 100)) # 13
```

### 6.4 2D 배열에서 이진 탐색

```python
def search_2d_matrix(matrix, target):
    """
    정렬된 2D 배열에서 탐색
    각 행이 정렬되어 있고, 다음 행의 첫 값이 이전 행의 마지막 값보다 큼
    """
    if not matrix or not matrix[0]:
        return False

    rows, cols = len(matrix), len(matrix[0])
    left, right = 0, rows * cols - 1

    while left <= right:
        mid = (left + right) // 2
        # 1D 인덱스를 2D 인덱스로 변환
        row = mid // cols
        col = mid % cols
        value = matrix[row][col]

        if value == target:
            return True
        elif value < target:
            left = mid + 1
        else:
            right = mid - 1

    return False

# 테스트
matrix = [
    [1, 3, 5, 7],
    [10, 11, 16, 20],
    [23, 30, 34, 60]
]
print(search_2d_matrix(matrix, 3))   # True
print(search_2d_matrix(matrix, 13))  # False
```

---

## 7. 실전 예제

### 예제 1: 회전된 정렬 배열에서 탐색

```python
def search_rotated_array(arr, target):
    """
    회전된 정렬 배열에서 target 찾기
    예: [4, 5, 6, 7, 0, 1, 2] (원래 [0, 1, 2, 4, 5, 6, 7])
    """
    left, right = 0, len(arr) - 1

    while left <= right:
        mid = (left + right) // 2

        if arr[mid] == target:
            return mid

        # 왼쪽 절반이 정렬되어 있는 경우
        if arr[left] <= arr[mid]:
            if arr[left] <= target < arr[mid]:
                right = mid - 1
            else:
                left = mid + 1
        # 오른쪽 절반이 정렬되어 있는 경우
        else:
            if arr[mid] < target <= arr[right]:
                left = mid + 1
            else:
                right = mid - 1

    return -1

# 테스트
arr = [4, 5, 6, 7, 0, 1, 2]
print(search_rotated_array(arr, 0))  # 4
print(search_rotated_array(arr, 3))  # -1
```

### 예제 2: 첫 번째와 마지막 위치 찾기

```python
def search_range(arr, target):
    """
    정렬된 배열에서 target의 첫 번째와 마지막 위치 찾기
    """
    def find_left():
        left, right = 0, len(arr)
        while left < right:
            mid = (left + right) // 2
            if arr[mid] < target:
                left = mid + 1
            else:
                right = mid
        return left

    def find_right():
        left, right = 0, len(arr)
        while left < right:
            mid = (left + right) // 2
            if arr[mid] <= target:
                left = mid + 1
            else:
                right = mid
        return left - 1

    left_idx = find_left()

    # target이 존재하지 않는 경우
    if left_idx >= len(arr) or arr[left_idx] != target:
        return [-1, -1]

    right_idx = find_right()

    return [left_idx, right_idx]

# 테스트
arr = [5, 7, 7, 8, 8, 8, 10]
print(search_range(arr, 8))   # [3, 5]
print(search_range(arr, 7))   # [1, 2]
print(search_range(arr, 6))   # [-1, -1]
```

### 예제 3: 피크 찾기

```python
def find_peak_element(arr):
    """
    배열에서 피크 요소 찾기
    피크: arr[i-1] < arr[i] > arr[i+1]
    """
    left, right = 0, len(arr) - 1

    while left < right:
        mid = (left + right) // 2

        # 오른쪽이 더 크면 오른쪽에 피크가 있음
        if arr[mid] < arr[mid + 1]:
            left = mid + 1
        # 왼쪽이 더 크거나 같으면 왼쪽에 피크가 있음
        else:
            right = mid

    return left

# 테스트
arr = [1, 2, 3, 1]
print(find_peak_element(arr))  # 2 (값 3)

arr = [1, 2, 1, 3, 5, 6, 4]
print(find_peak_element(arr))  # 1 또는 5 (값 2 또는 6)
```

### 예제 4: 제곱근 구하기

```python
def my_sqrt(x):
    """
    정수 x의 제곱근 (소수점 버림)
    이진 탐색으로 구하기
    """
    if x < 2:
        return x

    left, right = 1, x // 2

    while left <= right:
        mid = (left + right) // 2
        square = mid * mid

        if square == x:
            return mid
        elif square < x:
            left = mid + 1
        else:
            right = mid - 1

    return right  # right가 답

# 테스트
print(my_sqrt(4))   # 2
print(my_sqrt(8))   # 2
print(my_sqrt(16))  # 4
print(my_sqrt(17))  # 4
```

### 예제 5: 최솟값 찾기 (회전된 배열)

```python
def find_min_rotated(arr):
    """
    회전된 정렬 배열에서 최솟값 찾기
    """
    left, right = 0, len(arr) - 1

    while left < right:
        mid = (left + right) // 2

        # 오른쪽이 회전된 부분
        if arr[mid] > arr[right]:
            left = mid + 1
        # 왼쪽이 회전된 부분
        else:
            right = mid

    return arr[left]

# 테스트
arr = [4, 5, 6, 7, 0, 1, 2]
print(find_min_rotated(arr))  # 0

arr = [3, 4, 5, 1, 2]
print(find_min_rotated(arr))  # 1
```

---

## 8. 코딩테스트 팁

### 8.1 이진 탐색 사용 판단

**이진 탐색을 사용해야 하는 신호**:
- ✅ "정렬된 배열"
- ✅ "O(log n) 시간복잡도"
- ✅ "n ≤ 10^9" (선형 탐색 불가능)
- ✅ "최솟값, 최댓값 찾기"
- ✅ "특정 조건을 만족하는 값"

**문제 예시**:
```
"정렬된 배열에서 K를 찾아라"
"배열에서 K 이상인 첫 번째 값을 찾아라"
"배열에서 K의 개수를 구하라"
```

### 8.2 구현 체크리스트

```python
# ✅ 1. 배열이 정렬되어 있는가?
if not is_sorted(arr):
    arr.sort()

# ✅ 2. 경계 조건
# - left, right 초기값
# - while 조건 (<=? <?)
# - mid 계산

# ✅ 3. 무한 루프 체크
# - left = mid + 1 (또는 mid)
# - right = mid - 1 (또는 mid)

# ✅ 4. 반환값
# - 인덱스? 값? -1?
```

### 8.3 Lower/Upper Bound 사용법

```python
import bisect

# Lower Bound (≥)
idx = bisect.bisect_left(arr, target)

# Upper Bound (>)
idx = bisect.bisect_right(arr, target)

# 개수 세기
count = bisect.bisect_right(arr, target) - bisect.bisect_left(arr, target)

# 값 삽입
bisect.insort(arr, value)  # 정렬 유지하며 삽입
```

### 8.4 자주 나오는 패턴

**패턴 1: 정확한 값 찾기**
```python
def find_exact(arr, target):
    left, right = 0, len(arr) - 1
    while left <= right:
        mid = (left + right) // 2
        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    return -1
```

**패턴 2: Lower Bound**
```python
def lower_bound(arr, target):
    left, right = 0, len(arr)
    while left < right:
        mid = (left + right) // 2
        if arr[mid] < target:
            left = mid + 1
        else:
            right = mid
    return left
```

**패턴 3: Upper Bound**
```python
def upper_bound(arr, target):
    left, right = 0, len(arr)
    while left < right:
        mid = (left + right) // 2
        if arr[mid] <= target:
            left = mid + 1
        else:
            right = mid
    return left
```

### 8.5 디버깅 팁

```python
def binary_search_debug(arr, target):
    left, right = 0, len(arr) - 1
    step = 0

    while left <= right:
        mid = (left + right) // 2
        step += 1

        # 디버그 출력
        print(f"Step {step}: left={left}, mid={mid}, right={right}, arr[mid]={arr[mid]}")

        if arr[mid] == target:
            print(f"찾음! {step}번 만에")
            return mid
        elif arr[mid] < target:
            left = mid + 1
        else:
            right = mid - 1

    print(f"못 찾음! {step}번 시도")
    return -1

# 테스트
arr = [1, 3, 5, 7, 9, 11, 13, 15, 17, 19]
binary_search_debug(arr, 13)
```

### 8.6 실수하지 않기

**1) 정렬 확인**:
```python
# ❌ 정렬 안 된 배열에 이진 탐색
arr = [5, 2, 9, 1, 7]
binary_search(arr, 7)  # 잘못된 결과!

# ✅ 먼저 정렬
arr.sort()
binary_search(arr, 7)
```

**2) 경계 조건**:
```python
# ❌ 잘못된 경계
left, right = 0, len(arr) - 1
while left < right:  # ← 마지막 요소를 놓칠 수 있음
    ...

# ✅ 올바른 경계
left, right = 0, len(arr) - 1
while left <= right:  # ← 모든 요소 확인
    ...
```

**3) 무한 루프**:
```python
# ❌ 무한 루프 가능
while left < right:
    mid = (left + right) // 2
    if arr[mid] < target:
        left = mid  # ← +1 빠짐
    else:
        right = mid

# ✅ 올바른 코드
while left < right:
    mid = (left + right) // 2
    if arr[mid] < target:
        left = mid + 1  # ← +1 추가
    else:
        right = mid
```

---

## 9. 다음 주차 예고

다음 주에는 **이진 탐색 - 2부**를 배웁니다!

**5주차: 이진 탐색 - 2부**
- 매개변수 탐색 (Parametric Search)
- 결정 문제로 변환하기
- 최솟값의 최댓값, 최댓값의 최솟값
- 이진 탐색 심화 문제

이진 탐색을 응용한 고급 기법을 배웁니다!

---

## 핵심 정리

✅ **이진 탐색**: 정렬된 배열에서 O(log n)으로 탐색
✅ **전제 조건**: 배열이 반드시 정렬되어 있어야 함
✅ **선형 vs 이진**: O(n) vs O(log n) - 엄청난 차이!
✅ **Lower Bound**: target ≥ 인 첫 번째 위치
✅ **Upper Bound**: target > 인 첫 번째 위치
✅ **bisect 모듈**: 파이썬 내장 이진 탐색
✅ **활용**: 탐색, 개수 세기, 범위 찾기, 삽입 위치

---

## 🎯 실습 과제

이번 주 과제:
1. ✅ 이진 탐색 직접 구현 (반복문, 재귀)
2. ✅ Lower Bound, Upper Bound 구현
3. ✅ bisect 모듈 사용법 익히기
4. ✅ 추천 문제:
   - 백준 1920 (수 찾기 - 이진 탐색 기초)
   - 백준 10816 (숫자 카드 2 - Lower/Upper Bound)
   - 백준 1654 (랜선 자르기 - 다음 주 예습)
   - 백준 2110 (공유기 설치 - 다음 주 예습)
   - 백준 2805 (나무 자르기 - 다음 주 예습)
   - 프로그래머스 "입국심사"

**완료하셨나요?** README.md의 "4주차: 이진 탐색 1부"에 체크하고 5주차로!

---

**질문이 있거나 도움이 필요하면 언제든 물어보세요!**

**다음 주에 만나요! 이진 탐색 마스터! 🚀**

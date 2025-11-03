# 2주차: 시간복잡도와 공간복잡도

## 학습 목표
- 알고리즘 효율성의 개념을 이해한다
- Big-O 표기법을 정확히 알고 사용할 수 있다
- 코드를 보고 시간복잡도를 분석할 수 있다
- 공간복잡도를 계산할 수 있다
- 문제의 제약 조건을 보고 적절한 알고리즘을 선택할 수 있다

> 💡 **왜 중요한가?** 복잡도 분석 능력이 있어야 시간 초과 없이 문제를 풀 수 있습니다!

---

## 📑 목차

1. [알고리즘 효율성이란?](#1-알고리즘-효율성이란)
2. [Big-O 표기법](#2-big-o-표기법)
3. [시간복잡도 분석](#3-시간복잡도-분석)
4. [주요 시간복잡도 비교](#4-주요-시간복잡도-비교)
5. [공간복잡도](#5-공간복잡도)
6. [실전 예제](#6-실전-예제)
7. [복잡도 분석 연습](#7-복잡도-분석-연습)
8. [코딩테스트 팁](#8-코딩테스트-팁)

---

## 1. 알고리즘 효율성이란?

### 1.1 같은 문제, 다른 해결법

**문제**: 1부터 n까지의 합을 구하세요.

**방법 1: 반복문**
```python
def sum_1_to_n(n):
    total = 0
    for i in range(1, n + 1):
        total += i
    return total

print(sum_1_to_n(100))  # 5050
```
- n번 반복
- n이 커지면 느려짐

**방법 2: 수학 공식**
```python
def sum_1_to_n(n):
    return n * (n + 1) // 2

print(sum_1_to_n(100))  # 5050
```
- 한 번의 계산
- n이 아무리 커도 빠름!

💡 **핵심**: 같은 결과를 내지만 **효율성이 다릅니다**.

### 1.2 효율성을 측정하는 두 가지

**시간 효율성** (Time Complexity)
- 알고리즘이 실행되는 데 걸리는 시간
- "얼마나 빠른가?"

**공간 효율성** (Space Complexity)
- 알고리즘이 사용하는 메모리 양
- "얼마나 많은 메모리를 쓰는가?"

### 1.3 왜 정확한 시간이 아니라 복잡도인가?

실행 시간은 다음에 따라 달라집니다:
- 컴퓨터 성능
- 프로그래밍 언어
- 컴파일러/인터프리터
- 운영체제

→ **상대적인 효율성**을 비교하기 위해 **복잡도**를 사용!

---

## 2. Big-O 표기법

### 2.1 Big-O란?

**Big-O 표기법**은 알고리즘의 **최악의 경우** 효율성을 나타냅니다.

**표기**: O(f(n))
- n: 입력 크기
- f(n): n에 대한 함수

**의미**: "입력 크기 n에 대해 최대 f(n)에 비례하는 시간이 걸린다"

### 2.2 Big-O의 특징

**1) 최고차항만 남김**
```
3n² + 5n + 10 → O(n²)
```
- n이 커지면 n²이 지배적
- 낮은 차수는 무시

**2) 상수 계수 제거**
```
5n → O(n)
100n → O(n)
```
- 상수배는 의미 없음
- 성장률만 중요

**3) 최악의 경우 분석**
```python
def search(arr, target):
    for i in range(len(arr)):
        if arr[i] == target:
            return i
    return -1
```
- 최선: O(1) - 첫 번째에서 발견
- 평균: O(n/2) = O(n)
- **최악: O(n) - 끝까지 찾아야 함** ← Big-O는 이것!

### 2.3 주요 Big-O 표기

빠름 → 느림 순서:

```
O(1) < O(log n) < O(n) < O(n log n) < O(n²) < O(n³) < O(2ⁿ) < O(n!)
```

**시각화**:
```
실행 시간
    ↑
    |                                    O(n!)
    |                               O(2ⁿ)
    |                          O(n³)
    |                     O(n²)
    |              O(n log n)
    |         O(n)
    |    O(log n)
    |O(1)
    └────────────────────────────────→ n (입력 크기)
```

---

## 3. 시간복잡도 분석

### 3.1 O(1) - 상수 시간

입력 크기와 **무관**하게 일정한 시간

```python
# 예제 1: 배열 접근
arr = [1, 2, 3, 4, 5]
print(arr[2])  # O(1)

# 예제 2: 딕셔너리 접근
d = {'a': 1, 'b': 2}
print(d['a'])  # O(1)

# 예제 3: 수학 계산
def get_first_and_last(arr):
    return arr[0], arr[-1]  # O(1)
```

**특징**:
- 가장 빠름
- n이 얼마든 상관없음
- 해시맵, 배열 인덱스 접근

### 3.2 O(log n) - 로그 시간

입력이 반으로 줄어듦

```python
# 예제 1: 이진 탐색
def binary_search(arr, target):
    left, right = 0, len(arr) - 1

    while left <= right:
        mid = (left + right) // 2
        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            left = mid + 1  # 오른쪽 절반만 탐색
        else:
            right = mid - 1  # 왼쪽 절반만 탐색

    return -1
# 매번 절반씩 줄어듦 → O(log n)

# 예제 2: 거듭제곱 (분할정복)
def power(x, n):
    if n == 0:
        return 1
    if n % 2 == 0:
        half = power(x, n // 2)
        return half * half  # O(log n)
    else:
        return x * power(x, n - 1)
```

**특징**:
- 매우 빠름
- n = 1,000,000 → 약 20번 연산
- 이진 탐색, 트리 높이

### 3.3 O(n) - 선형 시간

입력 크기에 **비례**

```python
# 예제 1: 배열 순회
def find_max(arr):
    max_val = arr[0]
    for num in arr:  # n번 반복
        if num > max_val:
            max_val = num
    return max_val
# O(n)

# 예제 2: 리스트 합계
def sum_array(arr):
    total = 0
    for num in arr:  # n번
        total += num
    return total
# O(n)

# 예제 3: 선형 탐색
def linear_search(arr, target):
    for i in range(len(arr)):  # 최악 n번
        if arr[i] == target:
            return i
    return -1
# O(n)
```

**특징**:
- 가장 흔함
- 모든 요소를 한 번씩 봐야 할 때
- 단순 반복문

### 3.4 O(n log n) - 선형 로그 시간

효율적인 정렬의 시간복잡도

```python
# 예제 1: 병합 정렬
def merge_sort(arr):
    if len(arr) <= 1:
        return arr

    mid = len(arr) // 2
    left = merge_sort(arr[:mid])      # log n 깊이
    right = merge_sort(arr[mid:])

    return merge(left, right)          # n번 병합
# O(n log n)

# 예제 2: Python 내장 정렬
arr = [3, 1, 4, 1, 5, 9, 2, 6]
arr.sort()  # O(n log n)
sorted_arr = sorted(arr)  # O(n log n)
```

**특징**:
- 효율적인 정렬 알고리즘
- 병합 정렬, 퀵 정렬, 힙 정렬
- 대부분의 정렬 문제

### 3.5 O(n²) - 이차 시간

중첩 반복문

```python
# 예제 1: 버블 정렬
def bubble_sort(arr):
    n = len(arr)
    for i in range(n):          # n번
        for j in range(n - i - 1):  # n번
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
    return arr
# O(n²)

# 예제 2: 2D 배열 순회
def sum_2d_array(matrix):
    total = 0
    for i in range(len(matrix)):      # n번
        for j in range(len(matrix[0])):  # m번
            total += matrix[i][j]
    return total
# O(n × m) → 정사각 행렬이면 O(n²)

# 예제 3: 모든 쌍 확인
def count_pairs(arr):
    count = 0
    for i in range(len(arr)):         # n번
        for j in range(i + 1, len(arr)):  # n번
            if arr[i] + arr[j] == 10:
                count += 1
    return count
# O(n²)
```

**특징**:
- 느림
- n = 10,000이면 100,000,000번 연산
- 이중 반복문

### 3.6 O(n³) - 삼차 시간

삼중 중첩 반복문

```python
# 예제: 3D 배열 순회
def sum_3d_array(cube):
    total = 0
    for i in range(len(cube)):        # n번
        for j in range(len(cube[0])):  # n번
            for k in range(len(cube[0][0])):  # n번
                total += cube[i][j][k]
    return total
# O(n³)
```

**특징**:
- 매우 느림
- n = 1,000이면 1,000,000,000번 연산
- 거의 사용 안 함

### 3.7 O(2ⁿ) - 지수 시간

입력이 1 증가할 때마다 2배 증가

```python
# 예제 1: 피보나치 (순수 재귀)
def fib(n):
    if n <= 1:
        return n
    return fib(n - 1) + fib(n - 2)  # 2개로 분기
# O(2ⁿ)

# 예제 2: 부분집합 생성
def generate_subsets(arr):
    if not arr:
        return [[]]

    subsets = generate_subsets(arr[1:])
    return subsets + [[arr[0]] + s for s in subsets]
# O(2ⁿ) - 2ⁿ개의 부분집합
```

**특징**:
- 극도로 느림
- n = 30이면 1,000,000,000번 연산
- n ≤ 20~25 정도만 가능

### 3.8 O(n!) - 팩토리얼 시간

모든 순열을 생성

```python
# 예제: 순열 생성
def permutations(arr):
    if len(arr) <= 1:
        return [arr]

    result = []
    for i in range(len(arr)):
        rest = arr[:i] + arr[i+1:]
        for p in permutations(rest):
            result.append([arr[i]] + p)
    return result
# O(n!) - n!개의 순열
```

**특징**:
- 가장 느림
- n = 10이면 3,628,800번 연산
- n ≤ 10~11 정도만 가능

---

## 4. 주요 시간복잡도 비교

### 4.1 연산 횟수 비교 표

| n | O(log n) | O(n) | O(n log n) | O(n²) | O(2ⁿ) | O(n!) |
|---|----------|------|------------|-------|-------|-------|
| 10 | 3 | 10 | 30 | 100 | 1,024 | 3,628,800 |
| 100 | 7 | 100 | 700 | 10,000 | 1.3×10³⁰ | 9.3×10¹⁵⁷ |
| 1,000 | 10 | 1,000 | 10,000 | 1,000,000 | ∞ | ∞ |
| 10,000 | 13 | 10,000 | 130,000 | 100,000,000 | ∞ | ∞ |
| 100,000 | 17 | 100,000 | 1,700,000 | 10,000,000,000 | ∞ | ∞ |

### 4.2 입력 크기별 허용 복잡도

**1초 안에 실행 가능한 복잡도** (대략 1억 번 연산 기준)

| 입력 크기 n | 가능한 시간복잡도 |
|-------------|-------------------|
| n ≤ 11 | O(n!), O(n²ⁿ) |
| n ≤ 20~25 | O(2ⁿ) |
| n ≤ 100 | O(n⁴) |
| n ≤ 500 | O(n³) |
| n ≤ 5,000 | O(n²) |
| n ≤ 100,000 | O(n log n) |
| n ≤ 10,000,000 | O(n) |
| n > 10,000,000 | O(log n), O(1) |

**실전 팁**:
```python
# 문제에서 n ≤ 100,000이라면?
# → O(n log n) 또는 O(n) 알고리즘 필요!
# → O(n²)은 시간 초과!

# 문제에서 n ≤ 1,000이라면?
# → O(n²)까지 가능
# → 이중 반복문 OK
```

---

## 5. 공간복잡도

### 5.1 공간복잡도란?

알고리즘이 **실행 중 사용하는 메모리**의 양

**포함되는 것**:
- 입력 데이터
- 변수
- 배열, 리스트
- 재귀 호출 스택

### 5.2 공간복잡도 분석

**O(1) - 상수 공간**
```python
def sum_array(arr):
    total = 0  # O(1) 공간
    for num in arr:
        total += num
    return total
# 공간복잡도: O(1)
```

**O(n) - 선형 공간**
```python
def copy_array(arr):
    new_arr = []  # O(n) 공간
    for num in arr:
        new_arr.append(num)
    return new_arr
# 공간복잡도: O(n)

def fibonacci_dp(n):
    dp = [0] * (n + 1)  # O(n) 공간
    dp[1] = 1
    for i in range(2, n + 1):
        dp[i] = dp[i-1] + dp[i-2]
    return dp[n]
# 공간복잡도: O(n)
```

**O(n²) - 이차 공간**
```python
def create_2d_array(n):
    matrix = [[0] * n for _ in range(n)]  # O(n²) 공간
    return matrix
# 공간복잡도: O(n²)
```

**재귀 호출 스택**
```python
def factorial(n):
    if n <= 1:
        return 1
    return n * factorial(n - 1)
# 공간복잡도: O(n) - 재귀 깊이만큼 스택 사용
```

### 5.3 시간 vs 공간 트레이드오프

**피보나치 비교**

```python
# 방법 1: 재귀 (메모이제이션 없음)
def fib_recursive(n):
    if n <= 1:
        return n
    return fib_recursive(n-1) + fib_recursive(n-2)
# 시간: O(2ⁿ), 공간: O(n) - 재귀 스택

# 방법 2: DP (배열 사용)
def fib_dp(n):
    dp = [0] * (n + 1)
    dp[1] = 1
    for i in range(2, n + 1):
        dp[i] = dp[i-1] + dp[i-2]
    return dp[n]
# 시간: O(n), 공간: O(n)

# 방법 3: DP (공간 최적화)
def fib_optimized(n):
    if n <= 1:
        return n
    prev, curr = 0, 1
    for _ in range(2, n + 1):
        prev, curr = curr, prev + curr
    return curr
# 시간: O(n), 공간: O(1)
```

→ **공간을 희생**하면 **시간을 절약**할 수 있고, 그 반대도 가능!

---

## 6. 실전 예제

### 예제 1: 배열에서 중복 찾기

**방법 1: 이중 반복문 - O(n²)**
```python
def has_duplicate_1(arr):
    for i in range(len(arr)):
        for j in range(i + 1, len(arr)):
            if arr[i] == arr[j]:
                return True
    return False
# 시간: O(n²), 공간: O(1)
```

**방법 2: 정렬 사용 - O(n log n)**
```python
def has_duplicate_2(arr):
    arr.sort()  # O(n log n)
    for i in range(len(arr) - 1):  # O(n)
        if arr[i] == arr[i + 1]:
            return True
    return False
# 시간: O(n log n), 공간: O(1)
```

**방법 3: 집합 사용 - O(n)**
```python
def has_duplicate_3(arr):
    seen = set()  # O(n) 공간
    for num in arr:  # O(n)
        if num in seen:  # O(1)
            return True
        seen.add(num)
    return False
# 시간: O(n), 공간: O(n)
```

→ **가장 빠른 방법**: 집합 사용 (공간을 더 쓰지만 시간 절약)

### 예제 2: 두 수의 합

**문제**: 배열에서 합이 target인 두 수의 인덱스 찾기

**방법 1: 이중 반복문 - O(n²)**
```python
def two_sum_1(arr, target):
    for i in range(len(arr)):
        for j in range(i + 1, len(arr)):
            if arr[i] + arr[j] == target:
                return [i, j]
    return []
# 시간: O(n²), 공간: O(1)
```

**방법 2: 해시맵 사용 - O(n)**
```python
def two_sum_2(arr, target):
    seen = {}  # {값: 인덱스}
    for i, num in enumerate(arr):
        complement = target - num
        if complement in seen:
            return [seen[complement], i]
        seen[num] = i
    return []
# 시간: O(n), 공간: O(n)
```

### 예제 3: 배열 회전

**문제**: 배열을 k번 오른쪽으로 회전

**방법 1: 실제 회전 - O(n×k)**
```python
def rotate_1(arr, k):
    for _ in range(k):
        arr.insert(0, arr.pop())
    return arr
# 시간: O(n×k), 공간: O(1)
```

**방법 2: 슬라이싱 - O(n)**
```python
def rotate_2(arr, k):
    k = k % len(arr)  # k가 len(arr)보다 클 수 있음
    return arr[-k:] + arr[:-k]
# 시간: O(n), 공간: O(n)
```

**방법 3: 역순 활용 - O(n)**
```python
def rotate_3(arr, k):
    def reverse(arr, start, end):
        while start < end:
            arr[start], arr[end] = arr[end], arr[start]
            start += 1
            end -= 1

    k = k % len(arr)
    reverse(arr, 0, len(arr) - 1)  # 전체 역순
    reverse(arr, 0, k - 1)          # 앞부분 역순
    reverse(arr, k, len(arr) - 1)   # 뒷부분 역순
    return arr
# 시간: O(n), 공간: O(1)
```

---

## 7. 복잡도 분석 연습

### 연습 1: 다음 코드의 시간복잡도는?

```python
# 코드 1
def func1(n):
    for i in range(n):
        print(i)
```
<details>
<summary>답 보기</summary>
O(n) - 단일 반복문
</details>

```python
# 코드 2
def func2(n):
    for i in range(n):
        for j in range(n):
            print(i, j)
```
<details>
<summary>답 보기</summary>
O(n²) - 이중 반복문
</details>

```python
# 코드 3
def func3(n):
    for i in range(n):
        for j in range(i):
            print(i, j)
```
<details>
<summary>답 보기</summary>
O(n²)
- 총 연산: 0 + 1 + 2 + ... + (n-1) = n(n-1)/2
- n²의 계수이므로 O(n²)
</details>

```python
# 코드 4
def func4(n):
    i = 1
    while i < n:
        print(i)
        i *= 2
```
<details>
<summary>답 보기</summary>
O(log n)
- i: 1 → 2 → 4 → 8 → ... → n
- 2^k = n → k = log₂ n
</details>

```python
# 코드 5
def func5(arr):
    arr.sort()
    for num in arr:
        print(num)
```
<details>
<summary>답 보기</summary>
O(n log n)
- sort(): O(n log n)
- for 루프: O(n)
- 합: O(n log n + n) = O(n log n)
</details>

```python
# 코드 6
def func6(n):
    for i in range(n):
        for j in range(n):
            for k in range(n):
                print(i, j, k)
```
<details>
<summary>답 보기</summary>
O(n³) - 삼중 반복문
</details>

### 연습 2: 문제별 적절한 복잡도

**문제 1**: n = 100,000, 시간 제한 1초
- 가능한 복잡도: O(n), O(n log n)
- 불가능: O(n²) 이상

**문제 2**: n = 1,000, 시간 제한 2초
- 가능한 복잡도: O(n²), O(n³)까지도 가능
- 불가능: O(2ⁿ)

**문제 3**: n = 20, 시간 제한 1초
- 가능한 복잡도: O(2ⁿ), O(n!)까지도 가능

---

## 8. 코딩테스트 팁

### 8.1 문제 보자마자 체크할 것

**1) 입력 크기 확인**
```
n ≤ 100? → O(n³)까지 가능
n ≤ 10,000? → O(n²)까지 가능
n ≤ 100,000? → O(n log n) 필요
n ≤ 1,000,000? → O(n) 필요
```

**2) 시간 제한 확인**
```
1초 → 대략 1억 번 연산
2초 → 대략 2억 번 연산
```

**3) 메모리 제한 확인**
```
128MB → 약 3천만 개 정수
256MB → 약 6천만 개 정수
```

### 8.2 알고리즘 선택 가이드

**n ≤ 11**: 순열/조합, 백트래킹
**n ≤ 20**: 비트마스킹, 2ⁿ 알고리즘
**n ≤ 500**: O(n³) 가능 (플로이드-워셜)
**n ≤ 5,000**: O(n²) 가능
**n ≤ 100,000**: 정렬, 이진 탐색, DP
**n ≤ 1,000,000**: 선형 알고리즘, 해시맵

### 8.3 시간 초과 해결 방법

**1) 알고리즘 바꾸기**
- O(n²) → O(n log n) 또는 O(n)
- 이중 반복문 → 해시맵

**2) 입출력 최적화**
```python
import sys
input = sys.stdin.readline
```

**3) 불필요한 연산 제거**
```python
# ❌ 느림
for i in range(len(arr)):
    for j in range(len(arr)):  # len(arr) 매번 계산

# ✅ 빠름
n = len(arr)  # 한 번만 계산
for i in range(n):
    for j in range(n):
```

**4) 자료구조 바꾸기**
- 리스트 → 집합 (탐색 O(n) → O(1))
- 리스트 → 딕셔너리 (탐색 O(n) → O(1))
- 리스트 → 덱 (앞 삽입/삭제 O(n) → O(1))

### 8.4 복잡도 계산 체크리스트

```python
def analyze_complexity():
    """코드 작성 후 체크"""

    # ✅ 1. 반복문 개수 세기
    # - 단일 반복문: O(n)
    # - 이중 반복문: O(n²)
    # - 삼중 반복문: O(n³)

    # ✅ 2. 재귀 깊이 확인
    # - 재귀 깊이 n: O(n) 공간
    # - 이진 재귀: 시간복잡도 확인 필요

    # ✅ 3. 내장 함수 복잡도
    # - sort(): O(n log n)
    # - in (list): O(n)
    # - in (set): O(1)

    # ✅ 4. 예상 연산 횟수 계산
    # - n × 복잡도 < 100,000,000 이면 OK

    # ✅ 5. 공간 사용량 확인
    # - 배열 크기 × 데이터 크기 < 메모리 제한
```

---

## 9. 다음 주차 예고

다음 주에는 **정렬 알고리즘**을 배웁니다!

**3주차: 정렬 알고리즘**
- 버블 정렬, 선택 정렬, 삽입 정렬
- 병합 정렬, 퀵 정렬
- 파이썬 내장 정렬 활용
- 정렬 문제 풀이

정렬은 많은 알고리즘의 기초가 되므로 중요합니다!

---

## 핵심 정리

✅ **Big-O 표기법**: 최악의 경우 시간복잡도
✅ **주요 복잡도**: O(1) < O(log n) < O(n) < O(n log n) < O(n²) < O(2ⁿ) < O(n!)
✅ **입력 크기별 가능한 복잡도**:
   - n ≤ 100: O(n³)
   - n ≤ 5,000: O(n²)
   - n ≤ 100,000: O(n log n)
   - n > 1,000,000: O(n) 또는 O(log n)
✅ **시간 vs 공간 트레이드오프**: 공간을 더 쓰면 시간 절약 가능
✅ **코딩테스트**: 문제 보자마자 입력 크기로 복잡도 판단

---

## 🎯 실습 과제

이번 주 과제:
1. ✅ 자신이 작성한 코드의 시간복잡도 분석해보기
2. ✅ 백준 문제 풀 때 예상 복잡도 계산하기
3. ✅ O(n²)를 O(n)으로 최적화하는 연습
4. ✅ 추천 문제:
   - 백준 2750, 2751 (정렬 - 복잡도 차이 체험)
   - 백준 1920 (이진 탐색 vs 선형 탐색)
   - 백준 1927 (힙 - O(log n) 삽입/삭제)

**완료하셨나요?** README.md의 "2주차: 시간/공간복잡도"에 체크하고 3주차로!

---

**질문이 있거나 도움이 필요하면 언제든 물어보세요!**

**다음 주에 만나요! 복잡도 분석 마스터! 🚀**

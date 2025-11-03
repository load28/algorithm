# 3주차: 정렬 알고리즘

## 학습 목표
- 다양한 정렬 알고리즘의 원리와 특징을 이해한다
- 각 정렬 알고리즘의 시간복잡도와 공간복잡도를 분석할 수 있다
- 문제에 따라 적절한 정렬 알고리즘을 선택할 수 있다
- 파이썬 내장 정렬을 효과적으로 활용할 수 있다
- 정렬을 활용한 문제 해결 능력을 기른다

> 💡 **왜 중요한가?** 정렬은 가장 기본적이면서도 강력한 알고리즘입니다. 많은 문제가 정렬로 시작됩니다!

---

## 📑 목차

1. [정렬이란?](#1-정렬이란)
2. [기초 정렬 알고리즘 (O(n²))](#2-기초-정렬-알고리즘-on)
3. [고급 정렬 알고리즘 (O(n log n))](#3-고급-정렬-알고리즘-on-log-n)
4. [파이썬 내장 정렬](#4-파이썬-내장-정렬)
5. [정렬 알고리즘 비교](#5-정렬-알고리즘-비교)
6. [정렬 활용 패턴](#6-정렬-활용-패턴)
7. [실전 예제](#7-실전-예제)
8. [코딩테스트 팁](#8-코딩테스트-팁)

---

## 1. 정렬이란?

### 1.1 정렬의 정의

**정렬(Sorting)**: 데이터를 특정 순서대로 배열하는 것

```python
# 오름차순 정렬
[3, 1, 4, 1, 5, 9, 2] → [1, 1, 2, 3, 4, 5, 9]

# 내림차순 정렬
[3, 1, 4, 1, 5, 9, 2] → [9, 5, 4, 3, 2, 1, 1]
```

### 1.2 정렬이 필요한 이유

**1) 탐색 효율성**
```python
# 정렬되지 않은 배열: 선형 탐색 O(n)
arr = [3, 1, 4, 1, 5, 9, 2]
target = 5
# 모든 요소를 확인해야 함

# 정렬된 배열: 이진 탐색 O(log n)
arr = [1, 1, 2, 3, 4, 5, 9]
target = 5
# 절반씩 줄여가며 빠르게 찾음
```

**2) 데이터 분석**
- 최댓값/최솟값 찾기
- 중앙값 계산
- 상위 k개 찾기

**3) 문제 해결**
- 많은 알고리즘 문제의 시작점
- 투 포인터, 그리디 등에서 활용

### 1.3 정렬의 종류

**안정 정렬 (Stable Sort)**
- 같은 값의 순서가 유지됨
- 예: 병합 정렬, 삽입 정렬, 버블 정렬

**불안정 정렬 (Unstable Sort)**
- 같은 값의 순서가 바뀔 수 있음
- 예: 퀵 정렬, 선택 정렬, 힙 정렬

```python
# 안정 정렬 예시
원본: [(1, 'a'), (2, 'b'), (1, 'c')]
정렬: [(1, 'a'), (1, 'c'), (2, 'b')]  # 'a'가 'c'보다 앞에 유지

# 불안정 정렬 예시
원본: [(1, 'a'), (2, 'b'), (1, 'c')]
정렬: [(1, 'c'), (1, 'a'), (2, 'b')]  # 'a'와 'c'의 순서가 바뀔 수 있음
```

---

## 2. 기초 정렬 알고리즘 (O(n²))

### 2.1 버블 정렬 (Bubble Sort)

**원리**: 인접한 두 요소를 비교하여 교환, 큰 값이 뒤로 "버블"처럼 올라감

**과정**:
```
[5, 3, 8, 4, 2]
↓ 1단계: 5와 3 비교 → 교환
[3, 5, 8, 4, 2]
↓ 5와 8 비교 → 유지
[3, 5, 8, 4, 2]
↓ 8과 4 비교 → 교환
[3, 5, 4, 8, 2]
↓ 8과 2 비교 → 교환
[3, 5, 4, 2, 8]  ← 8이 제자리 찾음

↓ 2단계 반복...
[3, 4, 2, 5, 8]
↓ 계속...
[2, 3, 4, 5, 8]  ← 완료!
```

**구현**:
```python
def bubble_sort(arr):
    n = len(arr)

    # n번 반복
    for i in range(n):
        # 마지막 i개는 이미 정렬됨
        for j in range(n - i - 1):
            # 인접한 요소 비교
            if arr[j] > arr[j + 1]:
                # 교환
                arr[j], arr[j + 1] = arr[j + 1], arr[j]

    return arr

# 테스트
arr = [5, 3, 8, 4, 2]
print(bubble_sort(arr))  # [2, 3, 4, 5, 8]
```

**최적화 버전** (이미 정렬된 경우 조기 종료):
```python
def bubble_sort_optimized(arr):
    n = len(arr)

    for i in range(n):
        swapped = False  # 교환 여부 체크

        for j in range(n - i - 1):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
                swapped = True

        # 교환이 없었으면 이미 정렬됨
        if not swapped:
            break

    return arr
```

**복잡도**:
- 시간복잡도: O(n²) - 최선 O(n), 최악 O(n²)
- 공간복잡도: O(1)
- 안정 정렬

**장단점**:
- ✅ 구현이 간단
- ✅ 안정 정렬
- ❌ 매우 느림

### 2.2 선택 정렬 (Selection Sort)

**원리**: 매번 최솟값을 찾아 앞으로 이동

**과정**:
```
[5, 3, 8, 4, 2]
↓ 1단계: 최솟값 2 찾아서 첫 번째와 교환
[2, 3, 8, 4, 5]
      ↑
↓ 2단계: 나머지에서 최솟값 3 (이미 제자리)
[2, 3, 8, 4, 5]
         ↑
↓ 3단계: 나머지에서 최솟값 4 찾아서 교환
[2, 3, 4, 8, 5]
            ↑
↓ 4단계: 나머지에서 최솟값 5 찾아서 교환
[2, 3, 4, 5, 8]  ← 완료!
```

**구현**:
```python
def selection_sort(arr):
    n = len(arr)

    # i번째 위치에 들어갈 최솟값 찾기
    for i in range(n):
        min_idx = i

        # i+1부터 끝까지 최솟값 찾기
        for j in range(i + 1, n):
            if arr[j] < arr[min_idx]:
                min_idx = j

        # 최솟값과 i번째 교환
        arr[i], arr[min_idx] = arr[min_idx], arr[i]

    return arr

# 테스트
arr = [5, 3, 8, 4, 2]
print(selection_sort(arr))  # [2, 3, 4, 5, 8]
```

**복잡도**:
- 시간복잡도: O(n²) - 항상 O(n²)
- 공간복잡도: O(1)
- 불안정 정렬

**장단점**:
- ✅ 구현이 간단
- ✅ 메모리 효율적
- ❌ 느림
- ❌ 불안정 정렬

### 2.3 삽입 정렬 (Insertion Sort)

**원리**: 카드 정렬처럼, 하나씩 꺼내서 적절한 위치에 삽입

**과정**:
```
[5, 3, 8, 4, 2]
↓ 1단계: 3을 5 앞에 삽입
[3, 5, 8, 4, 2]
↓ 2단계: 8은 제자리
[3, 5, 8, 4, 2]
↓ 3단계: 4를 5와 8 사이에 삽입
[3, 4, 5, 8, 2]
↓ 4단계: 2를 맨 앞에 삽입
[2, 3, 4, 5, 8]  ← 완료!
```

**구현**:
```python
def insertion_sort(arr):
    n = len(arr)

    # 1번째부터 시작 (0번째는 이미 정렬됨)
    for i in range(1, n):
        key = arr[i]  # 삽입할 값
        j = i - 1

        # key보다 큰 값들을 오른쪽으로 이동
        while j >= 0 and arr[j] > key:
            arr[j + 1] = arr[j]
            j -= 1

        # key를 적절한 위치에 삽입
        arr[j + 1] = key

    return arr

# 테스트
arr = [5, 3, 8, 4, 2]
print(insertion_sort(arr))  # [2, 3, 4, 5, 8]
```

**복잡도**:
- 시간복잡도: O(n²) - 최선 O(n), 최악 O(n²)
- 공간복잡도: O(1)
- 안정 정렬

**장단점**:
- ✅ 구현이 간단
- ✅ 안정 정렬
- ✅ 거의 정렬된 배열에서 매우 빠름 O(n)
- ✅ 온라인 알고리즘 (데이터가 들어오는 대로 정렬 가능)
- ❌ 일반적으로 느림

---

## 3. 고급 정렬 알고리즘 (O(n log n))

### 3.1 병합 정렬 (Merge Sort)

**원리**: 분할 정복 - 배열을 반으로 나누고, 각각 정렬한 후 병합

**과정**:
```
[5, 3, 8, 4, 2, 7, 1, 6]
        ↓ 분할
[5, 3, 8, 4]    [2, 7, 1, 6]
        ↓ 분할
[5, 3]  [8, 4]  [2, 7]  [1, 6]
        ↓ 분할
[5] [3] [8] [4] [2] [7] [1] [6]
        ↓ 병합
[3, 5]  [4, 8]  [2, 7]  [1, 6]
        ↓ 병합
[3, 4, 5, 8]    [1, 2, 6, 7]
        ↓ 병합
[1, 2, 3, 4, 5, 6, 7, 8]  ← 완료!
```

**구현**:
```python
def merge_sort(arr):
    # 기저 조건: 길이가 1 이하면 이미 정렬됨
    if len(arr) <= 1:
        return arr

    # 분할
    mid = len(arr) // 2
    left = merge_sort(arr[:mid])    # 왼쪽 정렬
    right = merge_sort(arr[mid:])   # 오른쪽 정렬

    # 병합
    return merge(left, right)

def merge(left, right):
    result = []
    i = j = 0

    # 두 배열을 비교하며 병합
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1

    # 남은 요소 추가
    result.extend(left[i:])
    result.extend(right[j:])

    return result

# 테스트
arr = [5, 3, 8, 4, 2, 7, 1, 6]
print(merge_sort(arr))  # [1, 2, 3, 4, 5, 6, 7, 8]
```

**In-place 버전** (공간복잡도 최적화):
```python
def merge_sort_inplace(arr):
    if len(arr) <= 1:
        return arr

    def merge_sort_helper(arr, left, right):
        if left >= right:
            return

        mid = (left + right) // 2
        merge_sort_helper(arr, left, mid)
        merge_sort_helper(arr, mid + 1, right)
        merge_helper(arr, left, mid, right)

    def merge_helper(arr, left, mid, right):
        # 임시 배열에 복사
        temp = arr[left:right + 1]

        i = 0  # 왼쪽 시작
        j = mid - left + 1  # 오른쪽 시작
        k = left  # 원본 배열 인덱스

        # 병합
        while i <= mid - left and j < len(temp):
            if temp[i] <= temp[j]:
                arr[k] = temp[i]
                i += 1
            else:
                arr[k] = temp[j]
                j += 1
            k += 1

        # 남은 요소 복사
        while i <= mid - left:
            arr[k] = temp[i]
            i += 1
            k += 1

    merge_sort_helper(arr, 0, len(arr) - 1)
    return arr
```

**복잡도**:
- 시간복잡도: O(n log n) - 항상 O(n log n)
- 공간복잡도: O(n)
- 안정 정렬

**장단점**:
- ✅ 안정적인 O(n log n) 성능
- ✅ 안정 정렬
- ✅ 링크드 리스트에 효율적
- ❌ 추가 메모리 필요

### 3.2 퀵 정렬 (Quick Sort)

**원리**: 피벗을 선택하고, 피벗보다 작은 값은 왼쪽, 큰 값은 오른쪽으로 분할

**과정**:
```
[5, 3, 8, 4, 2, 7, 1, 6]  피벗=5
↓ 분할 (5보다 작은 것 | 5 | 5보다 큰 것)
[3, 4, 2, 1]  [5]  [8, 7, 6]
       ↓               ↓
[3, 2, 1] [4] []    [6] [8, 7]
    ↓                      ↓
[1] [3, 2]              [7] [8]
       ↓
    [2] [3]
       ↓
[1, 2, 3, 4, 5, 6, 7, 8]  ← 완료!
```

**구현 1: 간단한 버전** (추가 메모리 사용):
```python
def quick_sort(arr):
    # 기저 조건
    if len(arr) <= 1:
        return arr

    # 피벗 선택 (보통 중간값)
    pivot = arr[len(arr) // 2]

    # 분할
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]

    # 재귀적으로 정렬 및 병합
    return quick_sort(left) + middle + quick_sort(right)

# 테스트
arr = [5, 3, 8, 4, 2, 7, 1, 6]
print(quick_sort(arr))  # [1, 2, 3, 4, 5, 6, 7, 8]
```

**구현 2: In-place 버전** (Lomuto Partition):
```python
def quick_sort_inplace(arr):
    def quick_sort_helper(arr, low, high):
        if low < high:
            # 분할
            pi = partition(arr, low, high)

            # 재귀적으로 정렬
            quick_sort_helper(arr, low, pi - 1)
            quick_sort_helper(arr, pi + 1, high)

    def partition(arr, low, high):
        # 마지막 요소를 피벗으로
        pivot = arr[high]
        i = low - 1  # 작은 요소의 인덱스

        for j in range(low, high):
            # 현재 요소가 피벗보다 작으면
            if arr[j] < pivot:
                i += 1
                arr[i], arr[j] = arr[j], arr[i]

        # 피벗을 제자리에
        arr[i + 1], arr[high] = arr[high], arr[i + 1]
        return i + 1

    quick_sort_helper(arr, 0, len(arr) - 1)
    return arr

# 테스트
arr = [5, 3, 8, 4, 2, 7, 1, 6]
print(quick_sort_inplace(arr))  # [1, 2, 3, 4, 5, 6, 7, 8]
```

**구현 3: Hoare Partition** (더 효율적):
```python
def quick_sort_hoare(arr):
    def quick_sort_helper(arr, low, high):
        if low < high:
            pi = partition(arr, low, high)
            quick_sort_helper(arr, low, pi)
            quick_sort_helper(arr, pi + 1, high)

    def partition(arr, low, high):
        pivot = arr[low]
        i = low - 1
        j = high + 1

        while True:
            # 피벗보다 큰 값 찾기
            i += 1
            while arr[i] < pivot:
                i += 1

            # 피벗보다 작은 값 찾기
            j -= 1
            while arr[j] > pivot:
                j -= 1

            if i >= j:
                return j

            # 교환
            arr[i], arr[j] = arr[j], arr[i]

    quick_sort_helper(arr, 0, len(arr) - 1)
    return arr
```

**복잡도**:
- 시간복잡도: O(n log n) - 최악 O(n²)
- 공간복잡도: O(log n) - 재귀 스택
- 불안정 정렬

**장단점**:
- ✅ 평균적으로 가장 빠름
- ✅ 메모리 효율적 (in-place)
- ✅ 캐시 효율성 좋음
- ❌ 최악의 경우 O(n²)
- ❌ 불안정 정렬

**피벗 선택 전략**:
```python
# 1. 첫 번째 요소
pivot = arr[low]

# 2. 마지막 요소
pivot = arr[high]

# 3. 중간 요소
pivot = arr[(low + high) // 2]

# 4. 랜덤 (최악의 경우 회피)
import random
pivot_idx = random.randint(low, high)
pivot = arr[pivot_idx]

# 5. 중앙값 of 3 (가장 안정적)
mid = (low + high) // 2
pivot = sorted([arr[low], arr[mid], arr[high]])[1]
```

### 3.3 힙 정렬 (Heap Sort) - 소개

**원리**: 힙 자료구조를 이용한 정렬

```python
import heapq

def heap_sort(arr):
    # 최소 힙으로 변환
    heap = []
    for num in arr:
        heapq.heappush(heap, num)

    # 힙에서 하나씩 꺼내기
    result = []
    while heap:
        result.append(heapq.heappop(heap))

    return result

# 테스트
arr = [5, 3, 8, 4, 2, 7, 1, 6]
print(heap_sort(arr))  # [1, 2, 3, 4, 5, 6, 7, 8]
```

**복잡도**:
- 시간복잡도: O(n log n)
- 공간복잡도: O(1) - in-place 구현 시
- 불안정 정렬

> 💡 힙 정렬은 28주차 "트리 기초"에서 자세히 배웁니다!

---

## 4. 파이썬 내장 정렬

### 4.1 sort() vs sorted()

**sort()**: 리스트를 **제자리에서** 정렬 (원본 변경)
```python
arr = [3, 1, 4, 1, 5, 9, 2]
arr.sort()
print(arr)  # [1, 1, 2, 3, 4, 5, 9]
```

**sorted()**: **새로운** 정렬된 리스트 반환 (원본 유지)
```python
arr = [3, 1, 4, 1, 5, 9, 2]
sorted_arr = sorted(arr)
print(arr)        # [3, 1, 4, 1, 5, 9, 2] (원본 유지)
print(sorted_arr) # [1, 1, 2, 3, 4, 5, 9]
```

### 4.2 기본 사용법

```python
# 오름차순 (기본)
arr = [3, 1, 4, 1, 5, 9, 2]
arr.sort()
print(arr)  # [1, 1, 2, 3, 4, 5, 9]

# 내림차순
arr = [3, 1, 4, 1, 5, 9, 2]
arr.sort(reverse=True)
print(arr)  # [9, 5, 4, 3, 2, 1, 1]

# sorted() 사용
arr = [3, 1, 4, 1, 5, 9, 2]
print(sorted(arr))              # [1, 1, 2, 3, 4, 5, 9]
print(sorted(arr, reverse=True)) # [9, 5, 4, 3, 2, 1, 1]
```

### 4.3 key 함수 활용

**절댓값 기준 정렬**:
```python
arr = [-4, -1, 0, 3, 5, -2]
arr.sort(key=abs)
print(arr)  # [0, -1, -2, 3, -4, 5]
```

**문자열 길이 기준 정렬**:
```python
words = ["apple", "pie", "banana", "cat"]
words.sort(key=len)
print(words)  # ['pie', 'cat', 'apple', 'banana']
```

**대소문자 무시 정렬**:
```python
words = ["banana", "Apple", "cherry", "Banana"]
words.sort(key=str.lower)
print(words)  # ['Apple', 'banana', 'Banana', 'cherry']
```

**튜플의 특정 요소 기준 정렬**:
```python
students = [
    ("Alice", 25, 85),
    ("Bob", 20, 90),
    ("Charlie", 23, 85),
    ("David", 22, 95)
]

# 나이 기준
students.sort(key=lambda x: x[1])
# [('Bob', 20, 90), ('David', 22, 95), ('Charlie', 23, 85), ('Alice', 25, 85)]

# 성적 기준 (내림차순)
students.sort(key=lambda x: x[2], reverse=True)
# [('David', 22, 95), ('Bob', 20, 90), ('Alice', 25, 85), ('Charlie', 23, 85)]
```

**여러 조건으로 정렬**:
```python
students = [
    ("Alice", 25, 85),
    ("Bob", 20, 90),
    ("Charlie", 23, 85),
    ("David", 22, 85)
]

# 성적 내림차순, 같으면 나이 오름차순
students.sort(key=lambda x: (-x[2], x[1]))
print(students)
# [('Bob', 20, 90), ('David', 22, 85), ('Charlie', 23, 85), ('Alice', 25, 85)]
```

### 4.4 다양한 자료형 정렬

**딕셔너리 정렬**:
```python
# 키로 정렬
d = {'banana': 3, 'apple': 5, 'cherry': 2}
sorted_items = sorted(d.items())
print(sorted_items)  # [('apple', 5), ('banana', 3), ('cherry', 2)]

# 값으로 정렬
sorted_items = sorted(d.items(), key=lambda x: x[1])
print(sorted_items)  # [('cherry', 2), ('banana', 3), ('apple', 5)]

# 딕셔너리로 다시 변환
sorted_dict = dict(sorted_items)
```

**2D 리스트 정렬**:
```python
matrix = [
    [3, 4],
    [1, 2],
    [5, 1]
]

# 첫 번째 열 기준
matrix.sort(key=lambda x: x[0])
# [[1, 2], [3, 4], [5, 1]]

# 두 번째 열 기준
matrix.sort(key=lambda x: x[1])
# [[5, 1], [1, 2], [3, 4]]
```

**객체 정렬**:
```python
from dataclasses import dataclass

@dataclass
class Student:
    name: str
    age: int
    score: int

students = [
    Student("Alice", 25, 85),
    Student("Bob", 20, 90),
    Student("Charlie", 23, 85)
]

# 나이 기준
students.sort(key=lambda s: s.age)

# 여러 조건
students.sort(key=lambda s: (-s.score, s.age))
```

### 4.5 파이썬 정렬의 특징

**Tim Sort**:
- 파이썬의 내장 정렬 알고리즘
- 병합 정렬 + 삽입 정렬의 하이브리드
- 시간복잡도: O(n log n)
- 안정 정렬
- 실제 데이터에서 매우 빠름

```python
# 거의 정렬된 데이터에서 특히 빠름
arr = [1, 2, 3, 4, 5, 7, 6, 8, 9, 10]  # 거의 정렬됨
arr.sort()  # 매우 빠르게 실행!
```

---

## 5. 정렬 알고리즘 비교

### 5.1 시간복잡도 비교

| 알고리즘 | 최선 | 평균 | 최악 | 공간복잡도 | 안정성 |
|---------|------|------|------|-----------|--------|
| 버블 정렬 | O(n) | O(n²) | O(n²) | O(1) | 안정 |
| 선택 정렬 | O(n²) | O(n²) | O(n²) | O(1) | 불안정 |
| 삽입 정렬 | O(n) | O(n²) | O(n²) | O(1) | 안정 |
| 병합 정렬 | O(n log n) | O(n log n) | O(n log n) | O(n) | 안정 |
| 퀵 정렬 | O(n log n) | O(n log n) | O(n²) | O(log n) | 불안정 |
| 힙 정렬 | O(n log n) | O(n log n) | O(n log n) | O(1) | 불안정 |
| Tim Sort | O(n) | O(n log n) | O(n log n) | O(n) | 안정 |

### 5.2 언제 어떤 정렬을 사용할까?

**거의 정렬된 데이터**:
- ✅ 삽입 정렬 O(n)
- ✅ Tim Sort (파이썬 내장)

**안정 정렬이 필요한 경우**:
- ✅ 병합 정렬
- ✅ Tim Sort (파이썬 내장)

**메모리가 제한적인 경우**:
- ✅ 힙 정렬 O(1) 공간
- ✅ 퀵 정렬 O(log n) 공간

**평균적으로 가장 빠른 정렬**:
- ✅ 퀵 정렬
- ✅ Tim Sort (실제 데이터)

**코딩테스트에서**:
- ✅ **파이썬 내장 정렬 (sort/sorted)**
- 직접 구현은 거의 필요 없음
- 병합 정렬, 퀵 정렬 원리만 이해

### 5.3 성능 비교 실험

```python
import time
import random

def measure_time(sort_func, arr):
    start = time.time()
    sort_func(arr.copy())
    end = time.time()
    return end - start

# 테스트 데이터
n = 10000
random_data = [random.randint(1, 10000) for _ in range(n)]
sorted_data = sorted(random_data)
reverse_data = sorted_data[::-1]

print("=== 랜덤 데이터 ===")
print(f"버블 정렬: {measure_time(bubble_sort, random_data):.4f}초")
print(f"선택 정렬: {measure_time(selection_sort, random_data):.4f}초")
print(f"삽입 정렬: {measure_time(insertion_sort, random_data):.4f}초")
print(f"병합 정렬: {measure_time(merge_sort, random_data):.4f}초")
print(f"퀵 정렬: {measure_time(quick_sort, random_data):.4f}초")
print(f"파이썬 정렬: {measure_time(sorted, random_data):.6f}초")

print("\n=== 정렬된 데이터 ===")
print(f"삽입 정렬: {measure_time(insertion_sort, sorted_data):.6f}초")  # 매우 빠름!
print(f"퀵 정렬: {measure_time(quick_sort, sorted_data):.4f}초")        # 느림 (최악)
print(f"파이썬 정렬: {measure_time(sorted, sorted_data):.6f}초")       # 매우 빠름!
```

---

## 6. 정렬 활용 패턴

### 6.1 중복 제거

```python
# 방법 1: set 사용 (순서 유지 안 됨)
arr = [3, 1, 4, 1, 5, 9, 2, 6, 5]
unique = list(set(arr))
print(sorted(unique))  # [1, 2, 3, 4, 5, 6, 9]

# 방법 2: 정렬 후 제거 (순서 유지)
arr = [3, 1, 4, 1, 5, 9, 2, 6, 5]
arr.sort()
unique = [arr[0]]
for i in range(1, len(arr)):
    if arr[i] != arr[i-1]:
        unique.append(arr[i])
print(unique)  # [1, 2, 3, 4, 5, 6, 9]

# 방법 3: dict (Python 3.7+ 순서 유지)
arr = [3, 1, 4, 1, 5, 9, 2, 6, 5]
unique = list(dict.fromkeys(arr))
print(unique)  # [3, 1, 4, 5, 9, 2, 6] (원래 순서)
```

### 6.2 두 배열 합병

```python
def merge_two_sorted_arrays(arr1, arr2):
    result = []
    i = j = 0

    while i < len(arr1) and j < len(arr2):
        if arr1[i] <= arr2[j]:
            result.append(arr1[i])
            i += 1
        else:
            result.append(arr2[j])
            j += 1

    result.extend(arr1[i:])
    result.extend(arr2[j:])

    return result

# 테스트
arr1 = [1, 3, 5, 7]
arr2 = [2, 4, 6, 8]
print(merge_two_sorted_arrays(arr1, arr2))
# [1, 2, 3, 4, 5, 6, 7, 8]
```

### 6.3 k번째 큰/작은 수

```python
# 방법 1: 정렬
def kth_largest(arr, k):
    arr.sort(reverse=True)
    return arr[k-1]

# 방법 2: 힙 (더 효율적)
import heapq

def kth_largest_heap(arr, k):
    return heapq.nlargest(k, arr)[-1]

def kth_smallest_heap(arr, k):
    return heapq.nsmallest(k, arr)[-1]

# 테스트
arr = [3, 1, 4, 1, 5, 9, 2, 6]
print(kth_largest(arr, 3))       # 5
print(kth_largest_heap(arr, 3))  # 5
```

### 6.4 정렬 후 투 포인터

```python
# 두 수의 합
def two_sum_sorted(arr, target):
    arr.sort()  # 먼저 정렬
    left, right = 0, len(arr) - 1

    while left < right:
        current_sum = arr[left] + arr[right]
        if current_sum == target:
            return [left, right]
        elif current_sum < target:
            left += 1
        else:
            right -= 1

    return []

# 세 수의 합
def three_sum(arr, target):
    arr.sort()
    result = []

    for i in range(len(arr) - 2):
        # 중복 건너뛰기
        if i > 0 and arr[i] == arr[i-1]:
            continue

        left, right = i + 1, len(arr) - 1

        while left < right:
            current_sum = arr[i] + arr[left] + arr[right]

            if current_sum == target:
                result.append([arr[i], arr[left], arr[right]])
                left += 1
                right -= 1

                # 중복 건너뛰기
                while left < right and arr[left] == arr[left-1]:
                    left += 1
                while left < right and arr[right] == arr[right+1]:
                    right -= 1
            elif current_sum < target:
                left += 1
            else:
                right -= 1

    return result
```

### 6.5 구간 병합

```python
def merge_intervals(intervals):
    if not intervals:
        return []

    # 시작점 기준으로 정렬
    intervals.sort(key=lambda x: x[0])

    merged = [intervals[0]]

    for i in range(1, len(intervals)):
        current = intervals[i]
        last_merged = merged[-1]

        # 겹치면 병합
        if current[0] <= last_merged[1]:
            merged[-1] = [last_merged[0], max(last_merged[1], current[1])]
        else:
            merged.append(current)

    return merged

# 테스트
intervals = [[1, 3], [2, 6], [8, 10], [15, 18]]
print(merge_intervals(intervals))
# [[1, 6], [8, 10], [15, 18]]
```

### 6.6 정렬 후 이진 탐색

```python
def binary_search(arr, target):
    # 먼저 정렬
    arr.sort()

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

---

## 7. 실전 예제

### 예제 1: 배열의 교집합 (정렬 활용)

```python
def intersection(arr1, arr2):
    # 방법 1: set 사용
    return list(set(arr1) & set(arr2))

# 방법 2: 정렬 + 투 포인터
def intersection_sorted(arr1, arr2):
    arr1.sort()
    arr2.sort()

    result = []
    i = j = 0

    while i < len(arr1) and j < len(arr2):
        if arr1[i] == arr2[j]:
            if not result or result[-1] != arr1[i]:
                result.append(arr1[i])
            i += 1
            j += 1
        elif arr1[i] < arr2[j]:
            i += 1
        else:
            j += 1

    return result

# 테스트
arr1 = [1, 2, 2, 1]
arr2 = [2, 2]
print(intersection(arr1, arr2))  # [2]
```

### 예제 2: 가장 큰 수 만들기

```python
def largest_number(nums):
    # 문자열로 변환
    nums_str = list(map(str, nums))

    # 사용자 정의 정렬
    # 두 수를 합쳤을 때 더 큰 것이 앞으로
    from functools import cmp_to_key

    def compare(x, y):
        if x + y > y + x:
            return -1
        elif x + y < y + x:
            return 1
        else:
            return 0

    nums_str.sort(key=cmp_to_key(compare))

    # 0만 있는 경우 처리
    result = ''.join(nums_str)
    return '0' if result[0] == '0' else result

# 테스트
print(largest_number([3, 30, 34, 5, 9]))  # "9534330"
print(largest_number([10, 2]))            # "210"
```

### 예제 3: 회의실 배정

```python
def max_meetings(meetings):
    # 끝나는 시간 기준으로 정렬
    meetings.sort(key=lambda x: x[1])

    count = 0
    last_end_time = 0

    for start, end in meetings:
        # 이전 회의가 끝난 후 시작 가능하면
        if start >= last_end_time:
            count += 1
            last_end_time = end

    return count

# 테스트
meetings = [(1, 4), (3, 5), (0, 6), (5, 7), (3, 8), (5, 9), (6, 10), (8, 11), (8, 12), (2, 13), (12, 14)]
print(max_meetings(meetings))  # 4
```

### 예제 4: H-Index

```python
def h_index(citations):
    citations.sort(reverse=True)

    h = 0
    for i, citation in enumerate(citations):
        # h번 이상 인용된 논문이 h편 이상
        if citation >= i + 1:
            h = i + 1
        else:
            break

    return h

# 테스트
print(h_index([3, 0, 6, 1, 5]))  # 3
# 3번 이상 인용된 논문이 3편 이상 (6, 5, 3)
```

### 예제 5: 문자열 정렬

```python
# 아나그램 그룹화
def group_anagrams(words):
    from collections import defaultdict

    groups = defaultdict(list)

    for word in words:
        # 정렬된 문자열을 키로
        key = ''.join(sorted(word))
        groups[key].append(word)

    return list(groups.values())

# 테스트
words = ["eat", "tea", "tan", "ate", "nat", "bat"]
print(group_anagrams(words))
# [['eat', 'tea', 'ate'], ['tan', 'nat'], ['bat']]
```

---

## 8. 코딩테스트 팁

### 8.1 정렬 사용 판단

**정렬을 사용해야 하는 신호**:
- "k번째 큰/작은" → 정렬 또는 힙
- "중복 제거" → 정렬 또는 set
- "구간, 범위" → 정렬
- "최적 배치" → 정렬 (그리디)
- "이진 탐색" → 정렬 필수

**정렬을 피해야 하는 경우**:
- n이 매우 큼 (> 10^7)이고 O(n log n)도 느린 경우
- 원래 순서를 유지해야 하는 경우
- 이미 정렬된 데이터

### 8.2 정렬 최적화

**1) 정렬 횟수 줄이기**:
```python
# ❌ 나쁜 예: 매번 정렬
for _ in range(n):
    arr.append(new_value)
    arr.sort()  # O(n log n) × n = O(n² log n)

# ✅ 좋은 예: 한 번만 정렬
for _ in range(n):
    arr.append(new_value)
arr.sort()  # O(n log n)

# ✅ 더 좋은 예: 힙 사용
import heapq
heap = []
for _ in range(n):
    heapq.heappush(heap, new_value)  # O(log n) × n = O(n log n)
```

**2) 부분 정렬**:
```python
# k번째만 필요한 경우
import heapq

# 상위 k개
top_k = heapq.nlargest(k, arr)  # O(n log k)

# 하위 k개
bottom_k = heapq.nsmallest(k, arr)  # O(n log k)

# 전체 정렬보다 빠름!
```

**3) 안정성이 필요 없으면**:
```python
# 안정성이 필요 없으면 파이썬 정렬 그대로 사용
# (어차피 Tim Sort가 가장 빠름)
arr.sort()
```

### 8.3 자주 나오는 정렬 문제 유형

**1) 정렬 + 그리디**:
```python
# 회의실 배정, 작업 스케줄링
meetings.sort(key=lambda x: x[1])  # 끝나는 시간 순
```

**2) 정렬 + 투 포인터**:
```python
# 두 수의 합, 세 수의 합
arr.sort()
left, right = 0, len(arr) - 1
```

**3) 정렬 + 이진 탐색**:
```python
# Lower Bound, Upper Bound
arr.sort()
left, right = 0, len(arr) - 1
```

**4) 사용자 정의 정렬**:
```python
# 특수한 기준
arr.sort(key=lambda x: (조건1, 조건2, ...))
```

### 8.4 실수하기 쉬운 포인트

**1) 원본 배열 변경**:
```python
# ❌ 원본이 변경됨
def func(arr):
    arr.sort()
    return arr

# ✅ 원본 유지
def func(arr):
    return sorted(arr)
```

**2) 문자열 정렬**:
```python
# 주의: 문자열은 사전순
arr = ['10', '2', '20']
arr.sort()
print(arr)  # ['10', '2', '20'] (사전순)

# 숫자로 정렬하려면
arr.sort(key=int)
print(arr)  # ['2', '10', '20']
```

**3) 튜플 정렬 순서**:
```python
# 튜플은 첫 번째 요소부터 비교
arr = [(1, 3), (1, 2), (2, 1)]
arr.sort()
print(arr)  # [(1, 2), (1, 3), (2, 1)]
```

**4) 내림차순 실수**:
```python
# ❌ 잘못된 방법
arr.sort()
arr.reverse()

# ✅ 올바른 방법
arr.sort(reverse=True)
```

### 8.5 복잡도 체크리스트

```python
# 정렬 사용 시 복잡도
# - sort/sorted: O(n log n)
# - key 함수: 각 요소마다 한 번 호출
# - 중첩 정렬: 주의!

# 예시
arr = [[1, 2], [3, 1], [2, 3]]
arr.sort(key=lambda x: sorted(x))  # O(n × m log m) where m = len(x)
```

---

## 9. 다음 주차 예고

다음 주에는 **이진 탐색 - 1부**를 배웁니다!

**4주차: 이진 탐색 - 1부**
- 선형 탐색 vs 이진 탐색
- 이진 탐색 구현
- Lower Bound, Upper Bound
- 이진 탐색 기초 문제

정렬된 데이터에서 빠르게 탐색하는 방법을 배웁니다!

---

## 핵심 정리

✅ **기초 정렬**: 버블, 선택, 삽입 - O(n²), 이해용
✅ **고급 정렬**: 병합, 퀵 - O(n log n), 실전용
✅ **파이썬 정렬**: `sort()`, `sorted()` + `key` 함수 마스터
✅ **Tim Sort**: 파이썬 내장, 가장 빠르고 안정적
✅ **정렬 활용**: 투 포인터, 그리디, 이진 탐색의 전처리
✅ **코딩테스트**: 대부분 파이썬 내장 정렬 사용, 원리만 이해

---

## 🎯 실습 과제

이번 주 과제:
1. ✅ 병합 정렬, 퀵 정렬 직접 구현해보기
2. ✅ 파이썬 내장 정렬 key 함수 다양하게 연습
3. ✅ 추천 문제:
   - 백준 2750 (정렬 기초)
   - 백준 2751 (정렬 - 빠른 정렬 필요)
   - 백준 10814 (나이순 정렬 - 안정 정렬)
   - 백준 11650 (좌표 정렬)
   - 백준 1181 (단어 정렬 - 여러 조건)
   - 백준 11399 (ATM - 정렬 + 그리디)
   - 프로그래머스 "가장 큰 수"
   - 프로그래머스 "H-Index"

**완료하셨나요?** README.md의 "3주차: 정렬 알고리즘"에 체크하고 4주차로!

---

**질문이 있거나 도움이 필요하면 언제든 물어보세요!**

**다음 주에 만나요! 정렬 마스터! 🚀**

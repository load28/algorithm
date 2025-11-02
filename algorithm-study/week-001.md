# 1주차: 파이썬으로 알고리즘 풀기 - 완전 정복

## 학습 목표
- 알고리즘 문제 풀이에 필요한 **모든** 파이썬 문법 마스터
- 고급 알고리즘까지 풀 수 있는 **모든** 자료구조 이해
- 코딩테스트 실전 팁과 최적화 방법 습득

> 💡 **이 문서는 파이썬 완전 가이드입니다**. 천천히 읽고 모든 예제를 직접 실행해보세요!

---

## 📑 목차

1. [기본 문법](#1-기본-문법)
2. [입출력](#2-입출력)
3. [문자열 다루기](#3-문자열-다루기)
4. [자료구조](#4-자료구조)
5. [고급 기능](#5-고급-기능)
6. [실전 팁](#6-실전-팁)

---

## 1. 기본 문법

### 1.1 변수와 자료형

```python
# 변수 선언 (타입 명시 불필요)
a = 10              # 정수 (int)
b = 3.14            # 실수 (float)
c = "Hello"         # 문자열 (str)
d = True            # 불린 (bool)

# 여러 변수 동시 할당
x, y, z = 1, 2, 3
a = b = c = 0

# 변수 교환 (파이썬만의 특징!)
a, b = b, a

# 타입 확인
print(type(a))      # <class 'int'>
```

### 1.2 연산자

```python
# 산술 연산자
10 + 3    # 13 (더하기)
10 - 3    # 7  (빼기)
10 * 3    # 30 (곱하기)
10 / 3    # 3.333... (나누기)
10 // 3   # 3  (몫)
10 % 3    # 1  (나머지)
10 ** 3   # 1000 (거듭제곱)

# 비교 연산자
a == b    # 같다
a != b    # 다르다
a < b     # 작다
a > b     # 크다
a <= b    # 작거나 같다
a >= b    # 크거나 같다

# 논리 연산자
a and b   # 그리고
a or b    # 또는
not a     # 부정

# 비트 연산자 (나중에 비트마스크에서 중요!)
a & b     # AND
a | b     # OR
a ^ b     # XOR
~a        # NOT
a << 1    # 왼쪽 시프트 (2배)
a >> 1    # 오른쪽 시프트 (//2)
```

### 1.3 조건문

```python
# 기본 if문
if a > 10:
    print("크다")
elif a > 5:
    print("중간")
else:
    print("작다")

# 한 줄 if문 (삼항 연산자)
result = "양수" if a > 0 else "음수"

# 여러 조건
if a > 0 and b > 0:
    print("둘 다 양수")

# in 연산자
if a in [1, 2, 3, 4, 5]:
    print("1~5 사이")
```

### 1.4 반복문

```python
# for문
for i in range(5):          # 0, 1, 2, 3, 4
    print(i)

for i in range(2, 10):      # 2부터 9까지
    print(i)

for i in range(0, 10, 2):   # 0, 2, 4, 6, 8 (step=2)
    print(i)

for i in range(10, 0, -1):  # 10, 9, 8, ..., 1 (역순)
    print(i)

# 리스트 순회
arr = [1, 2, 3, 4, 5]
for num in arr:
    print(num)

# 인덱스와 함께 순회
for idx, num in enumerate(arr):
    print(f"arr[{idx}] = {num}")

# while문
i = 0
while i < 10:
    print(i)
    i += 1

# break와 continue
for i in range(10):
    if i == 5:
        break           # 반복문 종료
    if i % 2 == 0:
        continue        # 다음 반복으로
    print(i)

# 이중 반복문
for i in range(3):
    for j in range(3):
        print(f"({i}, {j})")
```

### 1.5 함수

```python
# 기본 함수
def add(a, b):
    return a + b

result = add(3, 5)  # 8

# 기본값 매개변수
def greet(name="Guest"):
    print(f"Hello, {name}")

greet()           # Hello, Guest
greet("Alice")    # Hello, Alice

# 여러 값 반환
def get_stats(arr):
    return min(arr), max(arr), sum(arr)

minimum, maximum, total = get_stats([1, 2, 3, 4, 5])

# 가변 인자 (*args)
def sum_all(*args):
    return sum(args)

print(sum_all(1, 2, 3, 4, 5))  # 15

# 키워드 인자 (**kwargs)
def print_info(**kwargs):
    for key, value in kwargs.items():
        print(f"{key}: {value}")

print_info(name="Alice", age=25)

# 람다 함수 (익명 함수)
square = lambda x: x ** 2
print(square(5))  # 25

# 정렬에서 자주 사용
arr = [(1, 3), (2, 1), (3, 2)]
arr.sort(key=lambda x: x[1])  # 두 번째 요소로 정렬
```

### 1.6 재귀 함수

```python
# 팩토리얼
def factorial(n):
    if n <= 1:
        return 1
    return n * factorial(n - 1)

# 피보나치
def fib(n):
    if n <= 1:
        return n
    return fib(n - 1) + fib(n - 2)

# 재귀 깊이 늘리기 (필요시)
import sys
sys.setrecursionlimit(10**6)
```

---

## 2. 입출력

### 2.1 기본 입출력

```python
# 입력
s = input()              # 한 줄 입력 (문자열)
n = int(input())         # 정수 입력
f = float(input())       # 실수 입력

# 여러 값 입력
a, b = map(int, input().split())        # 공백으로 구분
arr = list(map(int, input().split()))   # 리스트로

# 출력
print("Hello")
print(a, b)              # 공백으로 구분
print(a, b, sep=", ")    # 구분자 지정
print(a, end=" ")        # 줄바꿈 대신 공백
```

### 2.2 빠른 입출력 (중요!)

```python
# 많은 입력이 있을 때 (10만 줄 이상)
import sys
input = sys.stdin.readline

# 사용법은 동일
n = int(input())
arr = list(map(int, input().split()))

# 주의: 끝에 \n이 포함되므로 문자열은 .strip() 사용
s = input().strip()
```

### 2.3 출력 포매팅

```python
# f-string (Python 3.6+, 추천!)
name = "Alice"
age = 25
print(f"My name is {name} and I'm {age} years old.")
print(f"{10 / 3:.2f}")  # 소수점 2자리: 3.33

# format()
print("My name is {} and I'm {} years old.".format(name, age))

# % 포매팅 (옛날 방식)
print("My name is %s and I'm %d years old." % (name, age))

# 정수 자릿수 맞추기
print(f"{42:5d}")   # "   42" (5자리, 오른쪽 정렬)
print(f"{42:05d}")  # "00042" (0으로 채우기)
```

---

## 3. 문자열 다루기

### 3.1 문자열 기초

```python
s = "Hello World"

# 인덱싱
s[0]        # 'H'
s[-1]       # 'd' (뒤에서 첫 번째)

# 슬라이싱
s[0:5]      # "Hello"
s[6:]       # "World"
s[:5]       # "Hello"
s[:]        # "Hello World" (전체 복사)
s[::-1]     # "dlroW olleH" (역순)

# 길이
len(s)      # 11

# 문자열 연결
a = "Hello"
b = "World"
c = a + " " + b     # "Hello World"
d = a * 3           # "HelloHelloHello"

# 문자열 비교
"abc" == "abc"      # True
"abc" < "abd"       # True (사전순)
```

### 3.2 문자열 메서드

```python
s = "Hello World"

# 대소문자
s.lower()           # "hello world"
s.upper()           # "HELLO WORLD"
s.capitalize()      # "Hello world"
s.title()           # "Hello World"

# 검색
s.find("World")     # 6 (인덱스 반환, 없으면 -1)
s.index("World")    # 6 (인덱스 반환, 없으면 에러)
"World" in s        # True
s.count("l")        # 3

# 치환
s.replace("World", "Python")    # "Hello Python"

# 분할
s.split()           # ['Hello', 'World']
s.split("o")        # ['Hell', ' W', 'rld']
"a,b,c".split(",")  # ['a', 'b', 'c']

# 결합
" ".join(['Hello', 'World'])    # "Hello World"
",".join(['a', 'b', 'c'])       # "a,b,c"

# 제거
"  hello  ".strip()     # "hello"
"  hello  ".lstrip()    # "hello  "
"  hello  ".rstrip()    # "  hello"

# 판별
"123".isdigit()         # True (숫자인지)
"abc".isalpha()         # True (알파벳인지)
"abc123".isalnum()      # True (알파벳+숫자인지)

# 시작/끝 확인
s.startswith("Hello")   # True
s.endswith("World")     # True
```

### 3.3 문자열 <-> 리스트

```python
# 문자열을 리스트로
s = "Hello"
arr = list(s)       # ['H', 'e', 'l', 'l', 'o']

# 리스트를 문자열로
arr = ['H', 'e', 'l', 'l', 'o']
s = "".join(arr)    # "Hello"

# 문자 코드
ord('A')            # 65
ord('a')            # 97
chr(65)             # 'A'
```

---

## 4. 자료구조

### 4.1 리스트 (List) - 배열

```python
# 생성
arr = []
arr = [1, 2, 3, 4, 5]
arr = [0] * 10              # [0, 0, 0, ..., 0] (10개)
arr = list(range(5))        # [0, 1, 2, 3, 4]

# 접근
arr[0]      # 첫 번째 요소
arr[-1]     # 마지막 요소
arr[1:4]    # 슬라이싱

# 추가
arr.append(6)               # 끝에 추가
arr.insert(0, 10)           # 특정 위치에 추가
arr.extend([7, 8, 9])       # 여러 개 추가
arr += [10, 11]             # 확장

# 삭제
arr.pop()                   # 마지막 요소 제거 및 반환
arr.pop(0)                  # 특정 인덱스 제거
arr.remove(5)               # 값으로 제거 (첫 번째 발견)
del arr[0]                  # 인덱스로 제거
arr.clear()                 # 전체 제거

# 정렬
arr.sort()                  # 오름차순 정렬 (원본 변경)
arr.sort(reverse=True)      # 내림차순 정렬
sorted(arr)                 # 정렬된 새 리스트 반환 (원본 유지)

# 역순
arr.reverse()               # 원본 역순으로 변경
arr[::-1]                   # 역순 새 리스트 반환

# 기타
len(arr)                    # 길이
arr.count(5)                # 값 개수
arr.index(5)                # 값의 인덱스
min(arr), max(arr), sum(arr)
```

### 4.2 2D 배열

```python
# 생성
# ❌ 잘못된 방법
arr = [[0] * 3] * 3     # 주의! 같은 리스트가 3번 참조됨

# ✅ 올바른 방법
arr = [[0] * 3 for _ in range(3)]
# [[0, 0, 0],
#  [0, 0, 0],
#  [0, 0, 0]]

# 접근
arr[0][0] = 1
arr[1][2] = 5

# 순회
for i in range(len(arr)):
    for j in range(len(arr[0])):
        print(arr[i][j], end=" ")
    print()

# 또는
for row in arr:
    for val in row:
        print(val, end=" ")
    print()
```

### 4.3 튜플 (Tuple)

```python
# 생성 (불변!)
t = (1, 2, 3)
t = 1, 2, 3             # 괄호 생략 가능
t = (1,)                # 단일 요소 튜플

# 사용
a, b, c = t             # 언패킹
x, y = y, x             # 교환

# 튜플은 변경 불가
# t[0] = 10             # ❌ 에러!

# 리스트의 요소로 자주 사용
arr = [(1, 2), (3, 4), (5, 6)]
for x, y in arr:
    print(x, y)
```

### 4.4 딕셔너리 (Dictionary) - 해시맵

```python
# 생성
d = {}
d = dict()
d = {'apple': 100, 'banana': 200}

# 추가/수정
d['orange'] = 150
d['apple'] = 120        # 값 수정

# 접근
d['apple']              # 120
d.get('apple')          # 120
d.get('grape', 0)       # 없으면 0 반환 (기본값)

# 삭제
del d['apple']
d.pop('banana')
d.clear()

# 존재 확인
'apple' in d            # True/False

# 순회
for key in d:
    print(key, d[key])

for key, value in d.items():
    print(key, value)

for key in d.keys():
    print(key)

for value in d.values():
    print(value)

# 유용한 메서드
d.keys()                # dict_keys(['apple', 'banana', ...])
d.values()              # dict_values([100, 200, ...])
d.items()               # dict_items([('apple', 100), ...])

# defaultdict (자주 사용!)
from collections import defaultdict

d = defaultdict(int)    # 기본값 0
d['a'] += 1             # KeyError 없이 바로 사용 가능

d = defaultdict(list)   # 기본값 빈 리스트
d['a'].append(1)
```

### 4.5 집합 (Set)

```python
# 생성
s = set()
s = {1, 2, 3, 4, 5}
s = set([1, 2, 2, 3, 3])    # {1, 2, 3} (중복 제거)

# 추가/삭제
s.add(6)
s.remove(3)         # 없으면 에러
s.discard(3)        # 없어도 에러 없음
s.clear()

# 집합 연산
a = {1, 2, 3, 4, 5}
b = {3, 4, 5, 6, 7}

a | b               # 합집합 {1, 2, 3, 4, 5, 6, 7}
a & b               # 교집합 {3, 4, 5}
a - b               # 차집합 {1, 2}
a ^ b               # 대칭차집합 {1, 2, 6, 7}

# 존재 확인 O(1)
3 in s              # True

# 집합은 순서가 없음!
# s[0]              # ❌ 에러
```

### 4.6 스택 (Stack)

```python
# 리스트로 구현
stack = []

# push
stack.append(1)
stack.append(2)
stack.append(3)

# pop
stack.pop()         # 3
stack.pop()         # 2

# top
stack[-1]           # 맨 위 요소

# empty check
len(stack) == 0
```

### 4.7 큐 (Queue)

```python
# collections.deque 사용 (리스트보다 빠름!)
from collections import deque

queue = deque()

# enqueue
queue.append(1)
queue.append(2)
queue.append(3)

# dequeue
queue.popleft()     # 1
queue.popleft()     # 2

# front
queue[0]            # 맨 앞 요소

# empty check
len(queue) == 0
```

### 4.8 덱 (Deque) - 양방향 큐

```python
from collections import deque

dq = deque()

# 뒤에 추가/제거
dq.append(1)        # 오른쪽에 추가
dq.pop()            # 오른쪽에서 제거

# 앞에 추가/제거
dq.appendleft(0)    # 왼쪽에 추가
dq.popleft()        # 왼쪽에서 제거

# 회전
dq = deque([1, 2, 3, 4, 5])
dq.rotate(1)        # [5, 1, 2, 3, 4]
dq.rotate(-1)       # [1, 2, 3, 4, 5]

# 덱의 장점: 양쪽 끝 삽입/삭제가 O(1)
```

### 4.9 우선순위 큐 (힙) - heapq

```python
import heapq

heap = []

# 삽입 (자동으로 최소 힙 유지)
heapq.heappush(heap, 5)
heapq.heappush(heap, 3)
heapq.heappush(heap, 7)
heapq.heappush(heap, 1)

# 최솟값 제거 및 반환
min_val = heapq.heappop(heap)   # 1

# 최솟값 확인 (제거 X)
heap[0]             # 최솟값

# 리스트를 힙으로 변환
arr = [5, 3, 7, 1]
heapq.heapify(arr)  # O(n)

# 최댓값 힙 (음수로 변환)
heap = []
heapq.heappush(heap, -5)
heapq.heappush(heap, -3)
max_val = -heapq.heappop(heap)  # 5

# n개의 최소/최대값
arr = [5, 3, 7, 1, 9, 2]
heapq.nsmallest(3, arr)     # [1, 2, 3]
heapq.nlargest(3, arr)      # [9, 7, 5]
```

### 4.10 카운터 (Counter)

```python
from collections import Counter

# 생성
arr = [1, 2, 2, 3, 3, 3, 4]
counter = Counter(arr)      # Counter({3: 3, 2: 2, 1: 1, 4: 1})

s = "hello world"
counter = Counter(s)        # Counter({'l': 3, 'o': 2, ...})

# 접근
counter[1]          # 1의 개수
counter['l']        # 'l'의 개수

# 최빈값
counter.most_common(2)      # 가장 많은 2개 [('l', 3), ('o', 2)]

# 산술 연산
c1 = Counter(['a', 'b', 'c'])
c2 = Counter(['a', 'b', 'd'])
c1 + c2             # Counter({'a': 2, 'b': 2, 'c': 1, 'd': 1})
c1 - c2             # Counter({'c': 1})
```

---

## 5. 고급 기능

### 5.1 리스트 컴프리헨션

```python
# 기본
arr = [i for i in range(10)]                # [0, 1, 2, ..., 9]
arr = [i * 2 for i in range(10)]            # [0, 2, 4, ..., 18]

# 조건부
arr = [i for i in range(10) if i % 2 == 0] # [0, 2, 4, 6, 8]

# 2D 배열
arr = [[0] * 3 for _ in range(3)]

# if-else
arr = ["짝수" if i % 2 == 0 else "홀수" for i in range(5)]

# 중첩
arr = [(i, j) for i in range(3) for j in range(3)]
# [(0, 0), (0, 1), (0, 2), (1, 0), (1, 1), (1, 2), (2, 0), (2, 1), (2, 2)]
```

### 5.2 정렬

```python
arr = [3, 1, 4, 1, 5, 9, 2, 6]

# 기본 정렬
arr.sort()                      # 오름차순 (원본 변경)
arr.sort(reverse=True)          # 내림차순
sorted(arr)                     # 정렬된 새 리스트 (원본 유지)

# key 함수로 정렬
arr = [(1, 3), (2, 1), (3, 2)]
arr.sort(key=lambda x: x[1])    # 두 번째 요소로 정렬

# 문자열 정렬
arr = ["banana", "apple", "cherry"]
arr.sort()                      # 사전순
arr.sort(key=len)               # 길이순
arr.sort(key=lambda x: x[::-1]) # 역순 문자열 기준

# 여러 조건으로 정렬
arr = [(1, 3), (2, 1), (1, 2)]
arr.sort(key=lambda x: (x[0], x[1]))    # 첫 번째, 두 번째 순서대로

# 안정 정렬 (같은 값의 순서 유지)
# Python의 sort()는 stable sort (Tim Sort)
```

### 5.3 itertools 모듈

```python
from itertools import *

# 순열 (Permutations)
list(permutations([1, 2, 3], 2))        # [(1, 2), (1, 3), (2, 1), (2, 3), (3, 1), (3, 2)]

# 조합 (Combinations)
list(combinations([1, 2, 3], 2))        # [(1, 2), (1, 3), (2, 3)]

# 중복 순열
list(product([1, 2], repeat=2))         # [(1, 1), (1, 2), (2, 1), (2, 2)]

# 중복 조합
list(combinations_with_replacement([1, 2], 2))  # [(1, 1), (1, 2), (2, 2)]

# 누적 합 (Accumulate)
list(accumulate([1, 2, 3, 4, 5]))       # [1, 3, 6, 10, 15]

# 체인 (여러 이터러블 연결)
list(chain([1, 2], [3, 4], [5, 6]))     # [1, 2, 3, 4, 5, 6]

# 그룹화
from itertools import groupby
arr = [(1, 'a'), (1, 'b'), (2, 'c'), (2, 'd')]
for key, group in groupby(arr, key=lambda x: x[0]):
    print(key, list(group))
# 1 [(1, 'a'), (1, 'b')]
# 2 [(2, 'c'), (2, 'd')]
```

### 5.4 비트 연산

```python
# 비트마스크 활용
n = 5       # 0101 (2진수)
m = 3       # 0011 (2진수)

n & m       # 0001 = 1 (AND)
n | m       # 0111 = 7 (OR)
n ^ m       # 0110 = 6 (XOR)
~n          # -(n+1) (NOT)
n << 1      # 1010 = 10 (왼쪽 시프트)
n >> 1      # 0010 = 2 (오른쪽 시프트)

# i번째 비트 확인
def check_bit(n, i):
    return (n >> i) & 1

# i번째 비트 켜기
def set_bit(n, i):
    return n | (1 << i)

# i번째 비트 끄기
def clear_bit(n, i):
    return n & ~(1 << i)

# i번째 비트 토글
def toggle_bit(n, i):
    return n ^ (1 << i)

# 모든 부분집합 순회
n = 3
for mask in range(1 << n):      # 0부터 2^n-1까지
    subset = []
    for i in range(n):
        if mask & (1 << i):
            subset.append(i)
    print(subset)
```

### 5.5 기타 유용한 함수

```python
# zip (여러 이터러블 묶기)
a = [1, 2, 3]
b = ['a', 'b', 'c']
list(zip(a, b))         # [(1, 'a'), (2, 'b'), (3, 'c')]

# map (모든 요소에 함수 적용)
arr = [1, 2, 3, 4, 5]
list(map(lambda x: x * 2, arr))     # [2, 4, 6, 8, 10]

# filter (조건에 맞는 요소만)
list(filter(lambda x: x % 2 == 0, arr))     # [2, 4]

# all (모두 True인지)
all([True, True, True])     # True
all([True, False, True])    # False

# any (하나라도 True인지)
any([False, False, True])   # True
any([False, False, False])  # False

# eval (문자열을 코드로 실행)
result = eval("2 + 3 * 4")  # 14

# divmod (몫과 나머지를 동시에)
q, r = divmod(10, 3)        # q=3, r=1
```

---

## 6. 실전 팁

### 6.1 입출력 최적화

```python
# 📌 필수! 많은 입력이 있을 때
import sys
input = sys.stdin.readline

# 주의: strip() 사용
s = input().strip()

# 여러 테스트 케이스
T = int(input())
for _ in range(T):
    n = int(input())
    # 처리...
```

### 6.2 무한대 표현

```python
INF = float('inf')      # 무한대
INF = 10**9             # 충분히 큰 수
INF = int(1e9)          # 10억

# 최솟값 찾기
min_val = INF
for val in arr:
    min_val = min(min_val, val)
```

### 6.3 자주 쓰는 패턴

```python
# 구간 합 (누적 합)
arr = [1, 2, 3, 4, 5]
prefix = [0] * (len(arr) + 1)
for i in range(len(arr)):
    prefix[i + 1] = prefix[i] + arr[i]

# 구간 [L, R]의 합
sum_range = prefix[R + 1] - prefix[L]

# 카운트 배열
cnt = [0] * 100001      # 0~100000까지 카운트
for num in arr:
    cnt[num] += 1

# 방향 벡터 (상하좌우)
dx = [-1, 1, 0, 0]
dy = [0, 0, -1, 1]
for i in range(4):
    nx = x + dx[i]
    ny = y + dy[i]

# 대각선 포함 (8방향)
dx = [-1, -1, -1, 0, 0, 1, 1, 1]
dy = [-1, 0, 1, -1, 1, -1, 0, 1]
```

### 6.4 시간복잡도 체크리스트

```python
# 입력 크기에 따른 시간복잡도 가이드
# n <= 11: O(n!)
# n <= 25: O(2^n)
# n <= 100: O(n^4)
# n <= 500: O(n^3)
# n <= 3000: O(n^2 log n)
# n <= 5000: O(n^2)
# n <= 100,000: O(n log n)
# n <= 10,000,000: O(n)
# n > 10,000,000: O(log n) 또는 O(1)
```

### 6.5 실수 비교

```python
# 실수 비교 시 오차 고려
a = 0.1 + 0.2
b = 0.3
a == b              # False (부동소수점 오차)

# 올바른 비교
eps = 1e-9
abs(a - b) < eps    # True

# 또는 decimal 모듈 사용
from decimal import Decimal
```

### 6.6 디버깅 팁

```python
# 배열 예쁘게 출력
arr = [[1, 2, 3], [4, 5, 6]]
for row in arr:
    print(*row)

# 변수 이름과 함께 출력 (Python 3.8+)
x = 10
print(f"{x=}")      # x=10

# assert로 검증
assert len(arr) > 0, "배열이 비어있습니다"
```

---

## 7. 종합 예제

### 예제 1: 두 수의 합 (투 포인터)

```python
# 정렬된 배열에서 합이 target인 두 수 찾기
def two_sum(arr, target):
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

arr = [1, 2, 3, 4, 5, 6, 7]
print(two_sum(arr, 10))     # [2, 6] (3 + 7 = 10)
```

### 예제 2: 빈도수 세기 (딕셔너리)

```python
from collections import Counter

def most_frequent(arr):
    counter = Counter(arr)
    return counter.most_common(1)[0][0]

arr = [1, 2, 3, 2, 4, 2, 5]
print(most_frequent(arr))   # 2
```

### 예제 3: 순열 생성 (재귀)

```python
def permute(arr):
    if len(arr) == 1:
        return [arr]

    result = []
    for i in range(len(arr)):
        rest = arr[:i] + arr[i+1:]
        for p in permute(rest):
            result.append([arr[i]] + p)

    return result

print(permute([1, 2, 3]))
```

---

## 8. 다음 주차 예고

다음 주부터는 **알고리즘**을 배웁니다!

**2주차: 시간복잡도와 공간복잡도**
- Big-O 표기법
- 알고리즘 효율성 분석
- 실전 문제 풀이

파이썬 기초를 완벽히 익혔다면 알고리즘 학습을 시작할 준비가 된 것입니다!

---

## 핵심 정리

✅ **기본 문법**: 변수, 연산자, 조건문, 반복문, 함수
✅ **입출력**: `input()`, `print()`, `sys.stdin.readline`
✅ **문자열**: 슬라이싱, 메서드, 형 변환
✅ **자료구조**: 리스트, 튜플, 딕셔너리, 집합, 스택, 큐, 덱, 힙, Counter
✅ **고급**: 리스트 컴프리헨션, 정렬, itertools, 비트 연산
✅ **실전**: 빠른 입출력, 시간복잡도, 자주 쓰는 패턴

---

## 🎯 실습 과제

이번 주 과제:
1. ✅ 모든 예제 코드 직접 실행해보기
2. ✅ 백준 브론즈 문제 10개 풀기
   - 추천: 10950, 10951, 10952, 2741, 2742, 2438, 2439, 1000, 1001, 1008
3. ✅ 각 자료구조를 사용하는 간단한 프로그램 작성
4. ✅ 딕셔너리와 리스트 컴프리헨션 완벽히 익히기

**완료하셨나요?** README.md의 "1주차: 파이썬으로 알고리즘 풀기"에 체크하고 2주차로!

---

**질문이 있거나 도움이 필요하면 언제든 물어보세요!**

**다음 주에 만나요! 알고리즘 정복하러 가자! 🚀**

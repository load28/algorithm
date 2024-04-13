/**
 * 파스칼 삼각형의 n번째 줄의 합을 구하는 동적 프로그래밍 함수를 작성하세요
 * 변수명을 이해할수 있는것으로 지정하세요
 */
function pascalTriangle(n) {
  const arr = [];
  for (let i = 0; i < n; i++) {
    arr[i] = [];
    arr[i][0] = 1;
    for (let j = 1; j < i; j++) {
      arr[i][j] = arr[i - 1][j - 1] + arr[i - 1][j];
    }
    arr[i][i] = 1;
  }
  return arr[n - 1].reduce((acc, cur) => acc + cur, 0);
}

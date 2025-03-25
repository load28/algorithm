/**
 * 동전을 조합하여 원하는 금액을 만들 수 있는 경우의 수를 구하시오.
 */
function countWaysDP(coins, amount) {
  let combinations = Array(amount + 1).fill(0);
  combinations[0] = 1; // 0 센트를 만드는 방법은 아무것도 선택하지 않는 한 가지 방법 뿐이다.

  // 각 동전에 대해 반복
  for (let coin of coins) {
    // 해당 동전을 사용해 만들 수 있는 모든 금액에 대해 반복
    for (let i = coin; i <= amount; i++) {
      // 현재 금액(i)을 만드는 방법의 수는
      // (i-coin)을 만드는 방법의 수에 현재 동전(coin)을 추가하는 방법의 수를 더한 것
      combinations[i] += combinations[i - coin];
    }
  }

  console.log(combinations[amount]);
  // 최종적으로 amount를 만드는 방법의 총 수 반환
  return combinations[amount];
}

// 동전의 종류
const coins = [1, 3, 5];
const amount = 10; // 100센트 (1달러)
console.log(countWaysDP(coins, amount));

# Day_of_the_week_calculator
Day of the week calculator
https://www.youtube.com/watch?v=4LE8CP6QlBA

연도와 월 일을 입력하여 요일을 알아낸다.
## 원리

### 요일 계산
요일을 7로 모듈러 연산한다.

0. 일요일
1. 월요일
2. 화요일
3. 수요일
4. 목요일
5. 금요일
6. 토요일

가령 13이 나오면 13 % 7 = 6 -> 토요일이 된다.
22면 22 % 7 = 1 -> 월요일

### 연도 계산
윤년을 예외처리 해줘야 한다.
윤년은 4년마다 돌아온다.
100년에는 24년의 윤년과 76년의 평년이 있다.
+ 평년: 365일
  + 365일을 7로 나누면 1 남는다.
  + 100년중 평년이 76년이므로 76*1 = 76 
  + 76을 또 7(요일수)로 나눈다.
  + 76 % 7 = 6
+ 윤년: 366일
  + 366일을 7로 나누면 2 남는다.
  + 100년중 윤년이 24년이므로 24*2 = 48
  + 48을 또 7(요일수)로 나눈다.
  + 48 % 7 = 6

**100년 = (6 + 6) % 7 = 5**
200년은 100년(5)x2이므로
200 = 5x2 = 10
10 % 7 = 3

400년은 추가적으로 윤년이므로 하루를 더 해줘야한다.
400년=100년(5)x4+1
400년=21
21 % 7 = 0
0으로 나눠 떨어지므로 400년의 배수는 없는 수 취급해도 된다. (연산의 영향을 주지 않음)

2022년 x월 x일 = 2022년이 진행 중이므로 2021년으로 계산한다.
2021년 -> 2000년은 윤년이므로 빼고 21년 = 21+윤년수(5년) = 21+5=26 
26 % 7 = 5(금요일)

### 월 계산
+ 31일 % 7 = 3
+ 30일 % 7 = 2
+ 28일 % 7 = 0 (윤년)

10월 17일은 10월이 진행중이므로 9월로 계산해야 한다.
> 구현할 때는 연도와 월 둘 다 1씩 빼줘야한다.


10월 = 1월~9월 = 31, 28, 31, 30, .... = (3+0+3+2+3+2+3+2+3) % 7 = 21 % 7 = 0
일요일

### 일 계산 
10월 17일은 10월(0) + 17 = 17
17 % 7 = 3
수요일

### 종합
2022년 10월 17일은 앞서 구한 값들을 더하고 7로 나눈다.
(2021년 + 10월 + 17일) % 7
= (5 + 0 + 3) % 7 = 1 = 월요일
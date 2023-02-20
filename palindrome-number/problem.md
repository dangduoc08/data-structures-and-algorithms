# [9. Palindrome Number](https://leetcode.com/problems/palindrome-number)

Given an integer `x`, return `true` if `x` is a 
`palindrome`, and false otherwise.

<br />

## Example 1:
```
Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
```

<br />

## Example 2:
```
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
```

<br />

## Example 3:
```
Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
```

<br />

## Constraints:
- -2<sup>31</sup> <= x <= 2<sup>31</sup> - 1

<br />

## Follow-up:
Could you solve it without converting the integer to a string?
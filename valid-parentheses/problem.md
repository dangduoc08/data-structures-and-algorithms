# [20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses)

Given a string s containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

<br />

## Example 1:
```
Input: s = "()"
Output: true
```

<br />

## Example 2:
```
Input: s = "()[]{}"
Output: true
```

<br />

## Example 3:
```
Input: s = "(]"
Output: false
```

<br />

## Constraints:
- 1 <= s.length <= 10<sup>4</sup>
- `s` consists of parentheses only `'()[]{}'`.

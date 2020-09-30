[Interlaced Spiral Cipher](https://www.codewars.com/kata/5a24a35a837545ab04001614)
--------------------------
In this kata, your task is to implement what I call **Interlaced Spiral Cipher (ISC)**.

*Encoding a string using ISC is achieved with the following steps*:

1. Form a square large enough to fit all the string characters
2. Starting with the top-left corner, place string characters in the corner cells moving in a clockwise direction
3. After the first cycle is complete, continue placing characters in the cells following the last one in its respective row/column
4. When the outer cells are filled, repeat steps 2 through 4 for the remaining inner squares (refer to the example below for further clarification)
5. Fill up any unused cells with a space character and return the rows joined together.
#### Input
A string comprised of any combination of alphabetic 
characters, the space character, and any of the following 
characters `_!@#$%^&()[]{}+-*/="'<>,.?:;.`

Arguments passed to the `encode`
method will never have any trailing spaces.

#### Output
The `encode` method should return the encoded 
message as a string

The `decode` method should return the 
decoded message as a string with no trailing spaces
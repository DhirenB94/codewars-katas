# codewars-katas

## Game of 3
Given any number, repeatedly divide the number by 3 till you reach 1. Add or subtract 1 whenever division by 3 is not possible. At each stage output the number.

e.g. if you start with 100, it’s not divisible by 3, so subtract 1.

100, 99

Then keep dividing by 3

100, 99, 33, 11

11 is not divisible by 3, so add 1.

100, 99, 33, 11, 12

Keep going till you get to 1.

100, 99, 33, 11, 12, 4, 3, 1

##Rotate
Write a function that rotates a slice by k elements. For example [1,2,3,4,5,6] rotated by 2 becomes [3,4,5,6,1,2]. 

Bonus: Try solving this without creating a copy of the slice. 

## Sqaure or Square Root
Write a method, that will get an integer array as parameter and will process every number from this array.

Return a new array with processing every number of the input-array like this:

If the number has an integer square root, take this, otherwise square the number.

Example

[4,3,9,7,2,1] -> [2,9,3,49,4,1]

Notes
The input array will always contain only positive numbers, and will never be empty or null.

## Merge Lists

Write a function that merges two sorted lists of the same length into a new sorted list.

 [1,4,6],[2,3,5] → [1,2,3,4,5,6]. 

You can do this quicker than concatenating them followed by a sort.

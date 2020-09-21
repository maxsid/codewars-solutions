[Mixbonacci](https://www.codewars.com/kata/5811aef3acdf4dab5e000251)
-----------
#### Task
1. Mix `-nacci` sequences using a given pattern `p`.
2. Return the first n elements of the mixed sequence.
#### Rules
1. The pattern p is given as a list of strings (or 
array of symbols in Ruby) using the pattern 
mapping below (e. g. `['fib', 'pad', 'pel']` means take the next fibonacci, then the next padovan, then the next pell number and so on).
2. When `n` is 0 or `p` is empty return an empty list.
3. If `n` is more than the length of `p` repeat the pattern.
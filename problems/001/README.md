# Problem 1 – Multiples of 3 or 5

**Problem Statement:**

If we list all the natural numbers below $10$ that are multiples of $3$ or $5$, we get $3$, $5$, $6$, and $9$. The sum of these multiples is $23$.

**Task:**  
Find the sum of all the multiples of $3$ or $5$ below $1000$.

## Approach

- Iterate through numbers from $1$ to $999$.
- Check if the number is divisible by $3$ or $5$.
- Accumulate the sum.

## Formula Optimization (Optional)

You can also use arithmetic series summation:

- Sum of multiples of $n$ below $limit$:  

    $$ n \cdot \frac{k(k+1)}{2}, \text{ where } k = \left\lfloor \frac{limit - 1}{n} \right\rfloor $$

Then apply:

$$ \text{sum} = \text{sum}(3) + \text{sum}(5) - \text{sum}(15) $$

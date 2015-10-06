; get factorial of number n
(define (factorial n)
   (if (< n 2)
        1
        (* n (factorial (- n 1)))))

(display (factorial 5))
(newline)

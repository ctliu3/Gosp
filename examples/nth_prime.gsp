; https://projecteuler.net/problem=7
; ans: 104743

(import lib.gsp)

(define prime_?
  (lambda (n k)
    (if (zero? (remainder n k))
         #f
         (if (<= (* k k) n) (prime_? n (+ k 1)) #t))))

(define (nth-prime_ n k)
  (if (zero? n)
       (- k 1)
       (if (prime? k) (nth-prime_ (- n 1) (+ k 1)) (nth-prime_ n (+ k 1)))))

; int -> bool
(define (prime? n) (prime_? n 2))

; int -> int
(define (nth-prime n) (nth-prime_ n 2))

; (display (prime? 9))
(display (nth-prime_ 10000 2))
(newline)

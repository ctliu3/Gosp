; https://projecteuler.net/problem=2
; ans: 4613732

(import lib.gsp)

(define solve
  (lambda (fst snd)
    (let* ((a fst)
           (b snd)
           (c (+ a b)))
      (if (> c 4000000)
        0
        (+ (if (zero? (remainder c 2)) c 0) (solve b c))))))

(display (solve 1 1))
(newline)

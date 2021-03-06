(define PI 3.1415926)

; predicates
(define (int? x) (eqv? (class-of x) 'int))
(define (char? x) (eqv? (class-of x) 'char))
(define (float? x) (eqv? (class-of x) 'float))
(define (bool? x) (eqv? (class-of x) 'bool))
(define (string? x) (eqv? (class-of x) 'string))
(define (list? x) (eqv? (class-of x) 'list))
(define (symbol? x) (eqv? (class-of x) 'symbol))
(define (chan? x) (eqv? (class-of x) 'chan))
(define (vector? x) (eqv? (class-of x) 'vector))

; math methods
(define (zero? n) (if (= n 0) #t #f))
(define (positive? n) (if (> n 0) #t #f))
(define (negative? n) (if (< n 0) #t #f))
(define (abs n) (if (> n 0) n (- 0 n)))

; gcd
(define (gcd a b)
  (if (zero? b)
    a
    (gcd b (remainder a b))))

; lcm
(define (lcm a b)
  (/ (* a b) (gcd a b)))

; even?
(define even?
  (lambda (n)
    (if (zero? n)
         #t
         (odd? (- n 1)))))

; odd?
(define odd?
  (lambda (n)
    (if (zero? n)
         #f
         (even? (- n 1)))))

; list utils
(define (length x)
  (if (null? x)
     0
     (+ 1 (length (cdr x)))))

(define list-tail
  (lambda (x k)
    (if (zero? k)
         x
         (list-tail (cdr x) (- k 1)))))

(define list-ref
  (lambda (x k)
    (if (zero? k)
         (car x)
         (list-ref (cdr x) (- k 1)))))

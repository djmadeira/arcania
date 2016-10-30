# Arcania

## Design principles
Should be flexible in how it's written. Spacing should not matter, nor should line breaks; so it's flexible,
like real language.

Or maybe it should be super rigid? Magic words are supposed to be super rigid, and they just do nothing if you said them wrong.
There is no compiler for magic words.

I think it, ultimately, it shouldn't matter where the words are written (i.e. formatting is unimportant), as long as
they are *exactly* correct and failure results in nothing, just like magic.

## Structure
There are 4 registers of 18 bits each, allowing for 12 characters to be in memory at a time.

## Keywords
### Numerals
aught - zero
an - one
twegen - two
þrie - three
feower - four

### Operators
ægþer - AND gate
ceosan - OR gate
ongean - NOT gate
heah - EXOR gate
ebba - ENOR gate

### Output
acweþan - output the contents of all registers in round-robin order, starting from one (zeroes the registers)
atimbran - Put values in a register (must put store exactly 4 bits at a time)
æftersona - the same thing again
æfter - keep writing, but move to the next register

## Syntax
The first line of every program is considered to be the title and is ignored.

## Basic program (should print "Ƿes hāl middangeard")
Ƿes hāl middangeard

Atimbran an þrie, aught, twegen, onegan ceosan feower twegen, æftersona, aught, ceosan an feower, an, onegan þrie æfter aught, ceosan þrie feower, onegan aught, æftersona, onegen twegen, aught, an, aught, feower æfter, an, twegen, aught, feower, onegan feower, aught, onegan an, feower, æftersona æfter an, onegan ceosan feower twegen, þrie, aught, ceosan an feower, aught, onegan ceosan feower twegen, an, feower

Acweþan

## Letters
Letters are stored in 6-bits, providing 64 possible characters in the language. The characters are:
1  = a
2  = æ
3  = b
4  = c
5  = d
6  = ð
7  = e
8  = f
9  = ᵹ
10 = h
11 = i
12 = l
13 = m
14 = n
15 = o
16 = p
17 = r
18 = ſ
19 = t
20 = þ
21 = u
22 = ƿ
23 = x
24 = y
25 =  (space)
26 = (newline)
27 = A
28 = Æ
29 = B
30 = C
31 = D
32 = Ð
33 = E
34 = F
35 = Ᵹ
36 = H
37 = I
38 = L
39 = M
40 = N
41 = O
42 = P
43 = R
44 = S
45 = T
46 = Þ
47 = U
48 = Ƿ
49 = X
50 = Y
51 = ⁊
52 = ·
53 = ˙
54 = .
55 = &
56 = †
57 = ‡
58 = ♀
59 = ☉
60 = - (next character will be acented if possible)
// 4 are undefined, because I'm an asshole

# finduniqchar
Разработайте алгоритм программы, которая должна выполнять следующее:
- принимать на вход произвольный текст и находить в каждом слове входного текста самый первый символ, который больше не повторяется в анализируемом слове
- далее из полученного набора символов программа должна выбрать первый уникальный (т.е. тот, который больше не встречается в наборе) и вернуть его.

Например, если программе подать следующий текст ниже:

The Tao gave birth to machine language.  Machine language gave birth
to the assembler.
The assembler gave birth to the compiler.  Now there are ten thousand
languages.
Each language has its purpose, however humble.  Each language
expresses the Yin and Yang of software.  Each language has its place within
the Tao.
But do not program in COBOL if you can avoid it.
        -- Geoffrey James, "The Tao of Programming"

то пограмма должна вернуть "m".
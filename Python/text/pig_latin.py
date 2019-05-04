#!/usr/bin/env python3

"""
拉丁猪文游戏
将英语单词的第一个辅音音素的字母移动到词尾, 并加上 -ay 后缀
除了 aeiou 外, 其余字符都是辅音音素

考虑使用 re.search() 方法
该方法扫描字符串, 找到匹配样式的第一个位置

re.compile() 可以将正则表达式的样式编译为一个正则表达式对象
该对象可以用于匹配
"""

import re


word = input("请输入单词: ")
pattern = "[^a|e|i|o|u]"  # [] 内的 ^ 表示取反
prog = re.compile(pattern)
index = prog.search(word).start()
char = word[index]
word = "%s%s%s%s" % (word[:index], word[index+1:], char, "-ay")
print(word)

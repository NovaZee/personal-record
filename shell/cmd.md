# sed

```bash 
sed 's/old/new/' filename
###
s: 表示替换命令（substitute）
old: 是查找的模式
new: 是替换后的内容
filename: 是文件名
```

-  替换:s/
```bash 
###row代表行号,表示只替换第row行,替换所有行使用g
sed 's/old/new/row' filename
###正则例子
sed 's/[0-9]/#/g' filename  # 将所有数字替换为 #
```
- --
-  删除:d
```bash 
###d标识删除, N代表行号,范围只用,隔开
sed 'sed 'Nd' filename' 
sed 'sed '2,5d' filename' 
```
- --
-  插入:i\
-  追加:a\
```bash
###i\表示在指定行前插入, a\表示在指定行后追加
sed 'Ni\新内容' filename
sed 'Na\新内容' filename
```
- 输出
```bash
### -i：在原文件上进行修改，而不输出到终端。
sed -i 's/old/new/g' filename
### -n：只输出经过 sed 修改的行。通常和 p（print） 命令一起使用，来打印出特定内容：
sed -n 's/old/new/p' filename
```
/*
1.使用awk:
	awk 'NR==10' file.txt
	awk '{if(NR==10){print $0}}' file.txt
2.使用grep:
grep -n "" file.txt | grep -w '10' | cut -d: -f2
3.使用sed:
sed -n '10p' file.txt
*/

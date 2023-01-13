# Use this generator to generate all the functions as per the csv file and copy them manually to funcinvoker file
# cd /scripts
# python '.\function generator.py'
import pandas as pd

df = pd.read_csv('./data/cat1_func_k6.csv')
column_list=[]
for col in df.columns:
    column_list.append(col)
print(len(column_list))

text_file = open(r'./data/functions_name.txt', 'w')
template=""
for i in range(0,len(column_list)):
    template = template + "export function " + "f_" + column_list[i] + "() {\ntestK6();\n}\n\n"
text_file.write(template)
text_file.close()
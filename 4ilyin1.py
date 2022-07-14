import pandas as pd
import re

def getTempStudies(data):
    RepetList = pd.DataFrame(columns=data.columns)
    for i in range(data.shape[0]):
        pattern = re.sub(r'[0-9]', '', data.iloc[i]['NAME'])
        print(pattern)
        sub_data = data[data['NAME'].str.contains(pattern)]
        data = data[data['NAME'].str.contains(pattern)==False]
        if len(set(sub_data["Study Temp."])) > 1:
            sub_data = sub_data.sort_values(by=["Study Temp."])
            RepetList = pd.concat([RepetList, sub_data])
            print(RepetList)
    return RepetList

if __name__ == '__main__':
    data = pd.read_csv('data_from_CSDB.csv')
    result = getTempStudies(data)
    result.to_csv("result.csv", encoding='utf-8', index=False)


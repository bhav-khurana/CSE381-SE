import pandas as pd
import json

# Load JSON data from file
with open('exp1/output.json', 'r') as json_file:
    json_data = json.load(json_file)

# Create a DataFrame
df = pd.DataFrame.from_dict(json_data, orient='index').transpose()

# Write to Excel
excel_file_path = 'user_story.xlsx'
with pd.ExcelWriter(excel_file_path, engine='auto', mode='a') as writer:
    # Write to a new sheet
    df.to_excel(writer, sheet_name='userStory' + str(json_data['userStorySerialNumber']), index=False)

print(f"Data has been written to {excel_file_path}")
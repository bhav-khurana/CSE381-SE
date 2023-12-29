# CSE-381 Software Engineering
This repository contains all the files and experiments related to the CSE-381 (Software Engineering) course.

The main scripts are written in `Go` which produce `JSON` output and the script to convert `JSON` to `excel` format is written in python

## Setup - Experiment 1
You need to have an `OPENAI_API_KEY` in order to run the script.

1. Create a `.env` file in the project directory with the following content
    ```
    OPENAI_API_KEY = "<add-your-key>"
    ```
2. Run the `main.go` file using
    ```
    go run "path/to/main.go"
    ```
The output will be stored in the `user_story.xlsx` file with different userStory tables in different sheets within the excel file
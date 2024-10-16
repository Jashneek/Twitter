# Twitter

## Overview

This project shows how to use the Twitter API in the Go programming language. We can post a new tweet and delete an existing tweet. This is a great way to learn about APIs and how to connect to them.

## Requirements

Before we start, we need to have:

- Go installed.
- A Twitter Developer account.
- Twitter API keys and tokens (API Key, API Secret, Access Token, Access Token Secret).

## Setup Instructions

### 1. Get a Twitter Developer Account
Sign up for a Twitter Development Account

### 2. Sign Up for a Twitter Developer Account:

- Go to the Twitter Developer Platform.
- Click on "Apply" and fill out the application form with details about the intended use of the API.
- Provide accurate information about the project, including how we plan to use Twitter's data.
- Once the account is approved, log in to the Twitter Developer Dashboard.

##Creating Project and App  
Click on “Projects & Apps” and select “Create Project.”
Follow the prompts to name the project and app, and fill out any necessary details.

## 3.Generating API Keys, Access Keys and Tokens:

- After creating  app, go to the “Keys and Tokens” tab within your app settings.
- You will find your API Key, API Secret Key, Access Token, and Access Token Secret.
- Generate and Copy these keys as we will need them to authenticate API requests.

## Posting Status:

![image](https://github.com/user-attachments/assets/f30a04a4-3f4f-4176-9a9d-28eae709f841)

![image](https://github.com/user-attachments/assets/b3a76fe8-08f9-4b62-b3da-45c49124ce62)


## Deleting Status:
![image](https://github.com/user-attachments/assets/a6f3f6d6-173f-4f33-8b95-75a9b59b53df)

![image](https://github.com/user-attachments/assets/cfb3a139-6930-4d89-9e70-36e71d1c71bb)


## Authentication and Error Handling

`Authentication`

This project uses OAuth 1.0a for authentication. Ensure that the necessary API keys and tokens are included in the request headers.

`Error Handling`

The program implements error handling to manage various scenarios, such as:
Invalid Credentials: The program checks for authentication issues.
Rate Limiting: The program handles HTTP 429 errors and waits before retrying requests.
Invalid Tweet IDs: If an attempt is made to delete a tweet that does not exist, the program will catch the error and output a meaningful message.

## Invalid Authorization:
![image](https://github.com/user-attachments/assets/d8f26d2d-ab20-41f3-8c04-80630a290a40)

## Tweet Failed:
![image](https://github.com/user-attachments/assets/e8879c60-5ae5-4989-bdd2-f3574f956fcc)

## Wrong ID:
![image](https://github.com/user-attachments/assets/b8173507-0f37-4e62-a835-84e890a12d81)

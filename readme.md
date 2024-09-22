# API Documentation

## Overview

This API provides endpoints to manage goals and goal completions, allowing users to create, delete, and retrieve information about their goals.

## Endpoints

### 1. Create Goal Completion

- **POST** `/completions`
- **Summary**: Create a goal completion
- **Request Body**:
    ```json
    {
        "goalId": "string"
    }
    ```
- **Responses**:
    - **201 Created**: 
      ```json
      {
          "createdAt": "string",
          "goalId": "string",
          "id": "string",
          "updatedAt": "string"
      }
      ```
    - **400 Bad Request**:
      ```json
      {
          "error": "string"
      }
      ```
    - **404 Not Found**:
      ```json
      {
          "error": "string"
      }
      ```
    - **409 Conflict**:
      ```json
      {
          "error": "string"
      }
      ```
    - **500 Internal Server Error**:
      ```json
      {
          "error": "string"
      }
      ```

### 2. Delete Goal Completion

- **DELETE** `/completions`
- **Summary**: Delete a goal completion
- **Request Body**:
    ```json
    {
        "completionId": "string",
        "goalId": "string"
    }
    ```
- **Responses**:
    - **200 OK**:
      ```json
      {
          "message": "string"
      }
      ```
    - **400 Bad Request**:
      ```json
      {
          "error": "string"
      }
      ```
    - **404 Not Found**:
      ```json
      {
          "error": "string"
      }
      ```
    - **500 Internal Server Error**:
      ```json
      {
          "error": "string"
      }
      ```

### 3. Create Goal

- **POST** `/goals`
- **Summary**: Create a goal
- **Request Body**:
    ```json
    {
        "desiredWeeklyFrequency": "integer",
        "title": "string"
    }
    ```
- **Responses**:
    - **201 Created**:
      ```json
      {
          "desiredWeeklyFrequency": "integer",
          "title": "string"
      }
      ```
    - **400 Bad Request**:
      ```json
      {
          "error": "string"
      }
      ```
    - **500 Internal Server Error**:
      ```json
      {
          "error": "string"
      }
      ```

### 4. Get Pending Goals

- **GET** `/pending-goals`
- **Summary**: Retrieve pending goals
- **Responses**:
    - **200 OK**:
      ```json
      [
          {
              "completionCount": "integer",
              "createdAt": "string",
              "desiredWeeklyFrequency": "integer",
              "id": "string",
              "title": "string"
          }
      ]
      ```
    - **400 Bad Request**:
      ```json
      {
          "error": "string"
      }
      ```
    - **500 Internal Server Error**:
      ```json
      {
          "error": "string"
      }
      ```

### 5. Get Week Summary

- **GET** `/summary`
- **Summary**: Retrieve the week's summary
- **Responses**:
    - **200 OK**:
      ```json
      {
          "completed": "integer",
          "goalsPerDay": {
              "day": [
                  {
                      "completedAt": "string",
                      "goalId": "string",
                      "id": "string",
                      "title": "string"
                  }
              ]
          },
          "total": "integer"
      }
      ```
    - **400 Bad Request**:
      ```json
      {
          "error": "string"
      }
      ```
    - **500 Internal Server Error**:
      ```json
      {
          "error": "string"
      }
      ```

## Error Responses

All error responses will follow the format:
```json
{
    "error": "string"
}

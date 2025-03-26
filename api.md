Here's the API documentation in Markdown format based on the `outerRouter` implementation:

# API Documentation

## Health Check
### Ping
- **URL:** `/ping`
- **Method:** `GET`
- **Response:**
```json
{
    "message": "pong"
}
```

## Turtle API
⚠️已废弃

Base path: `/x/turtle`

### List Turtles
- **URL:** `/x/turtle/list`
- **Method:** `GET`
- **Description:** Get a list of turtles

### Set Batch Turtles
- **URL:** `/x/turtle/set/batch`
- **Method:** `POST`
- **Description:** Add multiple turtles in batch

## Idiom API
Base path: `/x/idiom`

### Get Idiom
- **URL:** `/x/idiom/get`
- **Method:** `GET`
- **Query Parameters:**
    - `id`: Idiom ID
    - `room_id`: Room ID
- **Response:** Returns idiom details including answer and image

### Set Batch Idioms
- **URL:** `/x/idiom/set/batch`
- **Method:** `POST`
- **Request Body:**
```json
{
    "idioms": [
        {
            "answer": "string",
            "image": "string",
            "difficulty": "EASY|MEDIUM|HARD",
            "creator": "string"
        }
    ],
    "total": 0
}
```

### Get Ranking
- **URL:** `/x/idiom/ranking`
- **Method:** `GET`
- **Query Parameters:**
    - `roomId`: Room ID
    - `limit`: Number of entries to return
- **Response:** Returns list of top players

### Update Ranking
- **URL:** `/x/idiom/update/ranking`
- **Method:** `POST`
- **Request Body:**
```json
{
    "roomId": 0,
    "uid": 0
}
```

### Start Game
- **URL:** `/x/idiom/start/game`
- **Method:** `POST`
- **Request Body:**
```json
{
    "room_id": "string",
    "up_code_id": 0
}
```
- **Response:** Returns game ID

### End Game
- **URL:** `/x/idiom/end/game`
- **Method:** `POST`
- **Request Body:**
```json
{
    "game_id": "string"
}
```

**Note:** Authentication might be required for the idiom API endpoints (TODO: authentication implementation).
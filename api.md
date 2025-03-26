Here's the updated `api.md` based on the `idiom.proto` file:

```markdown
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
            "id": 0,
            "iid": "string",
            "answer": "string",
            "image": "string",
            "difficulty": "UNDECIDED|EASY|MEDIUM|HARD",
            "creator": "string",
            "state": "string"
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
    "room_id": 123,
    "up_code_id": "string"
}
```
- **Response:** Returns game ID

### End Game
- **URL:** `/x/idiom/end/game`
- **Method:** `POST`
- **Request Body:**
```json
{
    "room_id": 123
}
```

**Note:** Authentication might be required for the idiom API endpoints (TODO: authentication implementation).
```
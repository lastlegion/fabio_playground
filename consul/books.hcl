{
  "service": {
    "id": "books",
    "name": "books",
    "tags": ["urlprefix-/book"],
    "port": 8080,
    "check": {
      "id": "books-http-check",
      "http": "http://127.0.0.1:8080/book",
      "method": "GET",
      "timeout": "15s",
      "interval": "15s"
    }
  }
}

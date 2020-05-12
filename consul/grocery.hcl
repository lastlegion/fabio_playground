{
  "service": {
    "id": "grocery",
    "name": "grocery",
    "tags": ["urlprefix-/grocery"],
    "port": 8081,
    "check": {
      "id": "grocery-http-check",
      "http": "http://127.0.0.1:8081/grocery",
      "method": "GET",
      "timeout": "15s",
      "interval": "15s"
    }
  }
}

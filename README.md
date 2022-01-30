curl http://localhost:8080/albums --header "Content-Type: application/json" --request "GET"

curl http://localhost:8080/albums  --include --header "Content-Type: application/json" --request "POST" --data '{"id": "6","title": "Titulo do meu album de sucesso","artist": "Miguel Muller","price": 158.71}'

curl http://localhost:8080/albums/2
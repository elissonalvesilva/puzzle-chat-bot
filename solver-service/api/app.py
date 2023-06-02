from flask import Flask, request

app = Flask(__name__)

@app.route('/endpoint', methods=['POST'])
def endpoint():
    data = request.json
    
    # Obtém os dados enviados na solicitação POST como JSON
    # Faça o processamento dos dados recebidos aqui
    # ...

    # Retorne uma resposta (opcional)
    return {'message': 'POST received'}
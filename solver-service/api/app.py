from flask import Flask, request, make_response
from solver.solver import puzzle

app = Flask(__name__)

@app.route('/puzzle', methods=['POST'])
def endpoint():
    data = request.json
    puzzle_id = int(data['puzzle_id'])

    # Obtém o campo 'answer' como uma string
    answer = str(data['answer'])

    result, clue = puzzle(puzzle_id, answer)

    if result is True:
        return {'message': {'clue': clue}}
    else:
        message = 'Answer is incorrect'
        response = make_response(message, 404)
        return response
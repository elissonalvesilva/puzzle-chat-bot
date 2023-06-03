import spacy
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import cosine_similarity

nlp = spacy.load('pt_core_news_sm')

def preprocess_text(text):
    doc = nlp(text)
    lemmas = [token.lemma_ for token in doc]
    preprocessed_lemmas = []
    
    for lemma in lemmas:
        if lemma.isalpha():
            preprocessed_lemmas.append(lemma)
        else:
            preprocessed_lemmas.append(lemma.lower())
    
    return ' '.join(preprocessed_lemmas)

def compare_responses(response, response_options):
    preprocessed_response = preprocess_text(response)
    preprocessed_options = [preprocess_text(option) for option in response_options]
    
    vectorizer = TfidfVectorizer()
    response_vectors = vectorizer.fit_transform([preprocessed_response] + preprocessed_options)
    
    similarity_scores = cosine_similarity(response_vectors[0], response_vectors[1:])[0]
    
    return similarity_scores

def get_puzzle(puzzle_id):
    puzzles = [
    {
        "id": 1,
        "answers": ["Vinho da babilônia", "Vinho da prostituição", "vinho da babilonia", "vinho da babilônia"],
        "clue": "",
    },
    {
        "id": 2,
        "answer": ["Mensagens angelicais", "Palavra de Deus"],
        "clue": "",
    },
    {
        "id": 3,
        "answer": ["morada de demônios, espíritos de demônios e feitiçaria", "Morada de demonios, espiritos de demonios e feiticaria"],
        "clue": "",
    },
    {
        "id": 4,
        "answer": ["A doutrina da imortalidade da alma", "imortalidade da alma", "Reencarnação", "Reencarnacao"],
        "clue": "",
    },
    {
        "id": 5,
        "answer": ["constantino", "Constantino", "Imperador constantino", "imperador constantino"],
        "clue": "",
    },
    {
        "id": 6,
        "answer": ["Sol", "Apolo", "Deus do Sol", "Deus sol"],
        "clue": "",
    },
    {
        "id": 7,
        "answer": ["apocalipse 18", "apocalipse18", "ap 18", "ap18"],
        "clue": "",
    }
    ]
    for puzzle in puzzles:
        if puzzle['id'] == puzzle_id:
            return puzzle['answers'], puzzle['clue']
    return [], ""


def puzzle(puzzle_id, answer):
    coeficiente_de_similaridade_min = 0.6

    answers_options, clue = get_puzzle(puzzle_id)

    similarity_scores = compare_responses(answer, answers_options)
    for i, score in enumerate(similarity_scores):
        if score >= coeficiente_de_similaridade_min:
            print(f"Similaridade entre a resposta dada e a opção {i+1}: {score}")
            return True, clue
        print(f"Similaridade entre a resposta dada e a opção {i+1}: {score}")

    return False, None

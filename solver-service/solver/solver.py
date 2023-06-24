import spacy
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import cosine_similarity
from unidecode import unidecode

nlp = spacy.load('pt_core_news_sm')


def preprocess_text(text):
    doc = nlp(text)
    preprocessed_lemmas = []

    for token in doc:
        lemma_with_accent = token.lemma_
        lemma_without_accent = unidecode(lemma_with_accent)

        if lemma_with_accent.isalpha():
            preprocessed_lemmas.append(lemma_with_accent.capitalize())
            preprocessed_lemmas.append(lemma_with_accent.lower())

            if lemma_with_accent != lemma_without_accent:
                preprocessed_lemmas.append(lemma_without_accent.capitalize())
                preprocessed_lemmas.append(lemma_without_accent.lower())
        else:
            preprocessed_lemmas.append(lemma_with_accent)

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
        "answers": [
            "o grande conflito",
            "o livro o grande conflito",
            "o grande controversia",
            "livro do ano",
        ],
        "clue": "",
    },
    {
        "id": 2,
        "answers": [
            "Whore of Babylon",
            "Prostituta da Babilônia",
        ],
        "clue": "",
    },
    {
        "id": 3,
        "answers": [
            "LORENA",
            "lorena",
            "Lorena",
        ],
        "clue": "",
    },
    {
        "id": 4,
        "answers": [
            "Apocalipse 13:8",
            "Ap 13:8",
        ],
        "clue": "",
    },
    {
        "id": 5,
        "answers": [
            "morte de jesus",
            "a morte de jesus",
            "ressureição de jesus",
            "jesus",
        ],
        "clue": "",
    },
    {
        "id": 6,
        "answers": [
            "Mateus 24:42",
            "Mat 24:42",
            "mateus 24:42",
            "mat 24:42"
        ],
        "clue": "",
    },
    {
        "id": 7,
        "answers": [
            "Enquanto as pessoas se apegarem à Bíblia e seguirem o que ela ensina, não serão enganadas na crise final"
        ],
        "clue": "",
    },
    {
        "id": 8,
        "answers": [
            "Mensagem de cristo para sua igreja",
            "Mensagem de cristo para igreja"
        ],
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
        percentage_score = score * 100
        if score >= coeficiente_de_similaridade_min:
            print(f"Similaridade entre a resposta dada e a opção {i+1}: {percentage_score:.2f}%")
            return True, clue, percentage_score
        print(f"Similaridade entre a resposta dada e a opção {i+1}: {percentage_score:.2f}%")

    return False, None, percentage_score
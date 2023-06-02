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
            'id': 1,
            'puzzle': 'Em uma terra vasta e cheia de vida, '
                    'Há habitantes cuja sede é tão intensa e aflita. '
                    'Buscam nas vinhas de Babilônia seu elixir, '
                    'Ignorando que nele há um veneno aferir. '
                    'O líquido rubro escorre em suas gargantas, '
                    'Um engano mortal, em busca de suas ansiadas jantas. '
                    'Quem são esses seres, cegos pela luxúria? '
                    'Deixe-me dizer, sua identidade é obscura.',
            'answers': [
                'Humanos',
                'Bebedores de vinho',
                'Babilônios',
                'Viciados em álcool',
                'Consumidores imprudentes',
                'Aqueles que ignoram os perigos',
                'Amantes da bebida',
                'Pessoas sedentas',
                'Adoradores de Dionísio (deus do vinho na mitologia grega)',
                'Festeiros imprudentes',
                'Vinho da Babilônia'
            ],
            'clue': 'Eu te amo'
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
            return True, clue
        print(f"Similaridade entre a resposta dada e a opção {i+1}: {score}")

    return False, None

    # for i, score in enumerate(similarity_scores):
    #     print(f"Similaridade entre a resposta dada e a opção {i+1}: {score}")

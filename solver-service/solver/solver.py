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
        "answers": ["chifre pequeno", "roma", "roma papal"],
        "clue": "Fé é a vitória",
    },
    {
        "id": 2,
        "answers": ["marca da besta","marca do inimigo", "insignia da besta", "sistema maligno", "testa", "adoração a besta"],
        "clue": "Sim, fé",
    },
    {
        "id": 3,
        "answers": ["sábado", "dia do senhor", "Sábado do senhor", "dia de adoração"],
        "clue": "sempre tem poder",
    },
    {
        "id": 4,
        "answers": ["fé", "resistência espiritual", "fé inabalável", "espírito resiliente"],
        "clue": "Fé é a",
    },
    {
        "id": 5,
        "answers": ["Eis que a sua alma está orgulhosa, não é reta nele; mas o justo pela sua fé viverá.", "o justo viverá pela fé", "...o justo viverá pela fé", '"o justo viverá pela fé"', "habacuque 2:4"],
        "clue": "vitória",
    },
    {
        "id": 6,
        "answers": ["intolerancia religiosa", "intolerância religiosa", "Intolerancia Religiosa"],
        "clue": "ao mundo irá vencer",
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

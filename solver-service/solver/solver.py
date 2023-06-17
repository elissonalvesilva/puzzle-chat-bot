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
            "Corrupção na igreja",
            "Falsos ensinamentos",
            "Desvio doutrinário",
            "Heresias cristãs",
            "Distorção do evangelho",
            "Apostasia religiosa",
            "Engano na fé",
            "Aposta em falsos profetas",
            "Falsas doutrinas",
            "Apostasia",
            "Apostasias",
        ],
        "clue": "Filho vai chegou a hora",
    },
    {
        "id": 2,
        "answers": [
            "Tempo simbólico",
            "Relação entre dias e anos",
            "Um dia equivale a um ano",
            "Calendário profético",
            "Cada dia por um ano",
            "um dia é igual a um ano",
            "um ano é um dia",
            "dia=ano",
        ],
        "clue": "Filho vai sem mais demora",
    },
    {
        "id": 3,
        "answers": [
            "Verdade Absoluta",
            "Testemunho Sagrado",
            "Símbolo da Fidelidade",
            "Marca da Autenticidade",
            "Selo divino",
            "Sábado",
            "Selo de Deus",
            "celo de Deus",
        ],
        "clue": "Vai buscar",
    },
    {
        "id": 4,
        "answers": [
            "No coração",
            "Na mente",
            "Na consciência",
            "Na adoração",
            "Testa e Mão",
            "Na testa e na Mao",
            "Testa Mão"
        ],
        "clue": "os meus amados",
    },
    {
        "id": 5,
        "answers": [
            "Dia de Descanso",
            "Observância do Sábado",
            "Mandamento do Repouso",
            "Lei do Descanso Divino",
            "Santificação do Sábado",
            "Guarda do Dia Sagrado",
            "Sábado",
            "Quarto Mandamento",
            "Dia do Senhor",
        ],
        "clue": "E traz de volta",
    },
    {
        "id": 6,
        "answers": [
            "1260 anos",
            "3 anos e meio",
            "42 meses proféticos",
            "Contagem de tempo profético",
            "1260 dias",
            "1260 anos",
            "purificação do santuário",
        ],
        "clue": "aqueles por quem",
    },
    {
        "id": 7,
        "answers": [
            "Papado",
            "Igreja Católica",
            "Hierarquia papal",
            "Santo Padre",
            "Papa",
            "Roma papal",
        ],
        "clue": "você morreu",
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
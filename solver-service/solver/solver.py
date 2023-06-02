import spacy
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import cosine_similarity

nlp = spacy.load('pt_core_news_sm')

def preprocess_text(text):
    doc = nlp(text)
    lemmas = [token.lemma_ for token in doc]
    
    return ' '.join(lemmas)

def compare_responses(response, response_options):
    preprocessed_response = preprocess_text(response)
    preprocessed_options = [preprocess_text(option) for option in response_options]
    
    vectorizer = TfidfVectorizer()
    response_vectors = vectorizer.fit_transform([preprocessed_response] + preprocessed_options)
    
    similarity_scores = cosine_similarity(response_vectors[0], response_vectors[1:])[0]
    
    return similarity_scores

def process():
    response = "Falso ensinamento"
    response_options = [
        "Falsos ensinamentos",
        "Falso ensinamento.",
        "Falsos cristos",
        "Falso enganos"
    ]

    similarity_scores = compare_responses(response, response_options)
    return similarity_scores

    # for i, score in enumerate(similarity_scores):
    #     print(f"Similaridade entre a resposta dada e a opção {i+1}: {score}")

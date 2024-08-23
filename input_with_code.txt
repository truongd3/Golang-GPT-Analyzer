import os
from dotenv import load_dotenv
from pathlib import Path
env_path = Path(".")/".env"
load_dotenv(dotenv_path = env_path)

import openai
import helper
db = "CustomerList.csv"

system_content = "You will receive a list of Justworks' customers - these companies are using Justworks products. Some Justworks employees want to use their services/products because they trust their customers. You should recommend the related and good services. Be descriptive and helpful."

client = openai.OpenAI(
    api_key = os.environ["AI_API_KEY"],
    base_url = "https://api.aimlapi.com",
)

def getRecommendation(command_text):
    prompt = helper.get_data_to_prompt()
    prompt += f"Please recommend users the companies that do {command_text} services."
    chat_completion = client.chat.completions.create(
        model = "mistralai/Mistral-7B-Instruct-v0.2",
        messages = [
            {"role": "system", "content": system_content},
            {"role": "user", "content": prompt},
        ],
        temperature = 0.7,
        max_tokens = 100,
    )
    return chat_completion.choices[0].message.content
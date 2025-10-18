from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from pymongo import MongoClient
from pydantic import BaseModel
from typing import Optional, List
import os
from datetime import datetime

app = FastAPI(title="Product Service", version="1.0.0")

# CORS
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# MongoDB connection
MONGO_URL = os.getenv("MONGO_URL", "mongodb://localhost:27017")
client = MongoClient(MONGO_URL)
db = client.products_db
products_collection = db.products

# Models
class Product(BaseModel):
    name: str
    description: str
    price: float
    stock: int
    category: str

# Routes
@app.get("/health")
async def health_check():
    return {"status": "healthy", "service": "product-service", "version": "1.0.0"}

@app.get("/api/products")
async def get_products():
    products = list(products_collection.find())
    for product in products:
        product["id"] = str(product.pop("_id"))
    return products

@app.post("/api/products", status_code=201)
async def create_product(product: Product):
    product_dict = product.dict()
    product_dict["created_at"] = datetime.utcnow()
    result = products_collection.insert_one(product_dict)
    product_dict["id"] = str(result.inserted_id)
    return product_dict

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=3002)

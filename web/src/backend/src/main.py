from typing import Union
from pydantic import BaseModel
from fastapi import FastAPI

from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()

origins = [
    "http://localhost.tiangolo.com",
    "https://localhost.tiangolo.com",
    "http://localhost",
    "http://localhost:8080",
    "http://localhost:3000"
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

class UserRegistration(BaseModel):
  userLogin: str
  userPassword: str
  userTerms: str
  userRemember: str
  userEmail: str|None

#login
@app.post("/")
def read_root(user: UserRegistration):
  return {"login": user.userLogin, "password": user.userPassword "rememberME": user.userRemember}

#registre

app.post("/")
def read_root(user: UserRegistration):
  return {"login": user.userLogin, "password": user.userPassword "email": user.userTerms "terms": user.userTerms}
from __future__ import annotations

from pydantic import BaseModel, ConfigDict, Field
from typing import List, Optional


class Person(BaseModel):
    """
    This is a Person struct with camelCase rename
    """
    model_config = ConfigDict(populate_by_name=True)

    first_name: str = Field(alias="firstName")
    last_name: str = Field(alias="lastName")
    age: int
    extra_special_field_1: int = Field(alias="extraSpecialField1")
    extra_special_field_2: Optional[List[str]] = Field(alias="extraSpecialField2", default=None)

class Person2(BaseModel):
    """
    This is a Person2 struct with UPPERCASE rename
    """
    model_config = ConfigDict(populate_by_name=True)

    first_name: str = Field(alias="FIRST_NAME")
    last_name: str = Field(alias="LAST_NAME")
    age: int = Field(alias="AGE")


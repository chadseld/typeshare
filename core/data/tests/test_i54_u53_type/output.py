from __future__ import annotations

from pydantic import BaseModel


class Foo(BaseModel):
    a: int
    b: int


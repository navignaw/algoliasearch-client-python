# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from enum import Enum
from json import loads
from typing import Self


class RelatedModel(str, Enum):
    """
    Related products or similar content model.  This model recommends items that are similar to the item with the ID `objectID`. Similarity is determined from the user interactions and attributes.
    """

    """
    allowed enum values
    """
    RELATED_MINUS_PRODUCTS = "related-products"

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of RelatedModel from a JSON string"""
        return cls(loads(json_str))
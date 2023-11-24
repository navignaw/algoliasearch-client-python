# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

import json
import pprint
from typing import Dict, List, Optional, Union

from pydantic import BaseModel, ValidationError, field_validator
from typing_extensions import Literal

from algoliasearch.search.models.search_for_facet_values_response import (
    SearchForFacetValuesResponse,
)
from algoliasearch.search.models.search_response import SearchResponse

try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

SEARCHRESULT_ONE_OF_SCHEMAS = ["SearchForFacetValuesResponse", "SearchResponse"]


class SearchResult(BaseModel):
    """
    SearchResult
    """

    # data type: SearchResponse
    oneof_schema_1_validator: Optional[SearchResponse] = None
    # data type: SearchForFacetValuesResponse
    oneof_schema_2_validator: Optional[SearchForFacetValuesResponse] = None
    actual_instance: Optional[
        Union[SearchForFacetValuesResponse, SearchResponse]
    ] = None
    one_of_schemas: List[str] = Literal[
        "SearchForFacetValuesResponse", "SearchResponse"
    ]

    model_config = {"validate_assignment": True}

    def __init__(self, *args, **kwargs) -> None:
        if args:
            if len(args) > 1:
                raise ValueError(
                    "If a position argument is used, only 1 is allowed to set `actual_instance`"
                )
            if kwargs:
                raise ValueError(
                    "If a position argument is used, keyword arguments cannot be used."
                )
            super().__init__(actual_instance=args[0])
        else:
            super().__init__(**kwargs)

    @field_validator("actual_instance")
    def actual_instance_must_validate_oneof(cls, v):
        SearchResult.model_construct()
        error_messages = []
        match = 0
        # validate data type: SearchResponse
        if not isinstance(v, SearchResponse):
            error_messages.append(f"""Type '{type(v)}' is not 'SearchResponse'""")
        else:
            match += 1
        # validate data type: SearchForFacetValuesResponse
        if not isinstance(v, SearchForFacetValuesResponse):
            error_messages.append(
                f"""Type '{type(v)}' is not 'SearchForFacetValuesResponse'"""
            )
        else:
            match += 1
        if match > 1:
            # more than 1 match
            raise ValueError(
                "Multiple matches found when setting `actual_instance` in SearchResult with oneOf schemas: SearchForFacetValuesResponse, SearchResponse. Details: "
                + ", ".join(error_messages)
            )
        elif match == 0:
            # no match
            raise ValueError(
                "No match found when setting `actual_instance` in SearchResult with oneOf schemas: SearchForFacetValuesResponse, SearchResponse. Details: "
                + ", ".join(error_messages)
            )
        else:
            return v

    @classmethod
    def from_dict(cls, obj: dict) -> Self:
        return cls.from_json(json.dumps(obj))

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Returns the object represented by the json string"""
        instance = cls.model_construct()
        error_messages = []
        match = 0

        # deserialize data into SearchResponse
        try:
            instance.actual_instance = SearchResponse.from_json(json_str)
            match += 1
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))
        # deserialize data into SearchForFacetValuesResponse
        try:
            instance.actual_instance = SearchForFacetValuesResponse.from_json(json_str)
            match += 1
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))

        if match > 1:
            # more than 1 match
            raise ValueError(
                "Multiple matches found when deserializing the JSON string into SearchResult with oneOf schemas: SearchForFacetValuesResponse, SearchResponse. Details: "
                + ", ".join(error_messages)
            )
        elif match == 0:
            # no match
            raise ValueError(
                "No match found when deserializing the JSON string into SearchResult with oneOf schemas: SearchForFacetValuesResponse, SearchResponse. Details: "
                + ", ".join(error_messages)
            )
        else:
            return instance

    def to_json(self) -> str:
        """Returns the JSON representation of the actual instance"""
        if self.actual_instance is None:
            return "null"

        to_json = getattr(self.actual_instance, "to_json", None)
        if callable(to_json):
            return self.actual_instance.to_json()
        else:
            return json.dumps(self.actual_instance)

    def to_dict(self) -> Dict:
        """Returns the dict representation of the actual instance"""
        if self.actual_instance is None:
            return None

        to_dict = getattr(self.actual_instance, "to_dict", None)
        if callable(to_dict):
            return self.actual_instance.to_dict()
        else:
            # primitive type
            return self.actual_instance

    def to_str(self) -> str:
        """Returns the string representation of the actual instance"""
        return pprint.pformat(self.model_dump())

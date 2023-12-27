# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from json import dumps
from typing import Dict, Optional, Self, Union

from pydantic import BaseModel, ValidationError

from algoliasearch.ingestion.models.source_big_commerce import SourceBigCommerce
from algoliasearch.ingestion.models.source_big_query import SourceBigQuery
from algoliasearch.ingestion.models.source_commercetools import SourceCommercetools
from algoliasearch.ingestion.models.source_csv import SourceCSV
from algoliasearch.ingestion.models.source_docker import SourceDocker
from algoliasearch.ingestion.models.source_json import SourceJSON


class SourceInput(BaseModel):
    """
    SourceInput
    """

    oneof_schema_1_validator: Optional[SourceCommercetools] = None
    oneof_schema_2_validator: Optional[SourceBigCommerce] = None
    oneof_schema_3_validator: Optional[SourceJSON] = None
    oneof_schema_4_validator: Optional[SourceCSV] = None
    oneof_schema_5_validator: Optional[SourceBigQuery] = None
    oneof_schema_6_validator: Optional[SourceDocker] = None
    actual_instance: Optional[
        Union[
            SourceBigCommerce,
            SourceBigQuery,
            SourceCSV,
            SourceCommercetools,
            SourceDocker,
            SourceJSON,
        ]
    ] = None

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

    @classmethod
    def from_dict(cls, obj: dict) -> Self:
        return cls.from_json(dumps(obj))

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Returns the object represented by the json string"""
        instance = cls.model_construct()
        error_messages = []

        try:
            instance.actual_instance = SourceCommercetools.from_json(json_str)

            return instance
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))
        try:
            instance.actual_instance = SourceBigCommerce.from_json(json_str)

            return instance
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))
        try:
            instance.actual_instance = SourceJSON.from_json(json_str)

            return instance
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))
        try:
            instance.actual_instance = SourceCSV.from_json(json_str)

            return instance
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))
        try:
            instance.actual_instance = SourceBigQuery.from_json(json_str)

            return instance
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))
        try:
            instance.actual_instance = SourceDocker.from_json(json_str)

            return instance
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))

        raise ValueError(
            "No match found when deserializing the JSON string into SourceInput with oneOf schemas: SourceBigCommerce, SourceBigQuery, SourceCSV, SourceCommercetools, SourceDocker, SourceJSON. Details: "
            + ", ".join(error_messages)
        )

    def to_json(self) -> str:
        """Returns the JSON representation of the actual instance"""
        if self.actual_instance is None:
            return "null"

        to_json = getattr(self.actual_instance, "to_json", None)
        if callable(to_json):
            return self.actual_instance.to_json()
        else:
            return dumps(self.actual_instance)

    def to_dict(self) -> Dict:
        """Returns the dict representation of the actual instance"""
        if self.actual_instance is None:
            return None

        to_dict = getattr(self.actual_instance, "to_dict", None)
        if callable(to_dict):
            return self.actual_instance.to_dict()
        else:
            return self.actual_instance

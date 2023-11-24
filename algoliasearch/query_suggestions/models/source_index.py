# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

import json
import pprint
from typing import Annotated, Any, ClassVar, Dict, List, Optional

from pydantic import BaseModel, Field, StrictBool, StrictStr

from algoliasearch.query_suggestions.models.facet import Facet

try:
    from typing import Self
except ImportError:
    from typing_extensions import Self


class SourceIndex(BaseModel):
    """
    Configuration of an Algolia index for Query Suggestions.
    """

    index_name: StrictStr = Field(
        description="Name of the Algolia index to use as source for query suggestions.",
        alias="indexName",
    )
    replicas: Optional[StrictBool] = Field(
        default=False,
        description="If true, Query Suggestions uses all replicas of the primary index to find popular searches. If false, only the primary index is used.  ",
    )
    analytics_tags: Optional[List[StrictStr]] = Field(
        default=None,
        description="[Analytics tags](https://www.algolia.com/doc/api-reference/api-parameters/analyticsTags/) for filtering the popular searches. ",
        alias="analyticsTags",
    )
    facets: Optional[List[Facet]] = Field(
        default=None,
        description="Facets to use as top categories with your suggestions.  If provided, Query Suggestions adds the top facet values to each suggestion. ",
    )
    min_hits: Optional[Annotated[int, Field(strict=True, ge=0)]] = Field(
        default=5,
        description="Minimum number of hits required to be included as a suggestion.  A search query must at least generate `minHits` hits to be included in the Query Suggestions index. ",
        alias="minHits",
    )
    min_letters: Optional[Annotated[int, Field(strict=True, ge=0)]] = Field(
        default=4,
        description="Minimum letters required to be included as a suggestion.  A search query must be at least `minLetters` long to be included in the Query Suggestions index. ",
        alias="minLetters",
    )
    generate: Optional[List[List[StrictStr]]] = None
    external: Optional[List[StrictStr]] = Field(
        default=None,
        description="Algolia indices with popular searches to use as query suggestions.  Records of these indices must have these attributes:    - `query`: search query which will be added as a suggestion   - `count`: measure of popularity of that search query  For example, you can export popular searches from an external analytics tool, such as Google Analytics or Adobe Analytics, and feed this data into an external Algolia index. You can use this external index to generate query suggestions until your Algolia analytics has collected enough data. ",
    )
    __properties: ClassVar[List[str]] = [
        "indexName",
        "replicas",
        "analyticsTags",
        "facets",
        "minHits",
        "minLetters",
        "generate",
        "external",
    ]

    model_config = {"populate_by_name": True, "validate_assignment": True}

    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True,
        # exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of SourceIndex from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        _dict = self.model_dump(
            by_alias=True,
            exclude={},
            exclude_none=True,
        )
        # override the default output from pydantic by calling `to_dict()` of
        # each item in facets (list)
        _items = []
        if self.facets:
            for _item in self.facets:
                if _item:
                    _items.append(_item.to_dict())
            _dict["facets"] = _items
        # set to None if analytics_tags (nullable) is None
        # and model_fields_set contains the field
        if self.analytics_tags is None and "analytics_tags" in self.model_fields_set:
            _dict["analyticsTags"] = None

        # set to None if facets (nullable) is None
        # and model_fields_set contains the field
        if self.facets is None and "facets" in self.model_fields_set:
            _dict["facets"] = None

        # set to None if external (nullable) is None
        # and model_fields_set contains the field
        if self.external is None and "external" in self.model_fields_set:
            _dict["external"] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of SourceIndex from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate(
            {
                "indexName": obj.get("indexName"),
                "replicas": obj.get("replicas")
                if obj.get("replicas") is not None
                else False,
                "analyticsTags": obj.get("analyticsTags"),
                "facets": [Facet.from_dict(_item) for _item in obj.get("facets")]
                if obj.get("facets") is not None
                else None,
                "minHits": obj.get("minHits") if obj.get("minHits") is not None else 5,
                "minLetters": obj.get("minLetters")
                if obj.get("minLetters") is not None
                else 4,
                "generate": obj.get("generate"),
                "external": obj.get("external"),
            }
        )
        return _obj

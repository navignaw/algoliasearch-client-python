# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from enum import Enum
from json import loads
from typing import Self


class TriggerType(str, Enum):
    """
    Task trigger, describing when a task should run.  - `onDemand`.   Manually trigger the task with the `/run` endpoint.  - `schedule`.   Regularly trigger the task on a `cron` schedule.  - `subscription`.   Trigger the task after an event is received, such as, a webhook.  - `streaming`.   Run the task continuously.
    """

    """
    allowed enum values
    """
    ONDEMAND = "onDemand"
    SCHEDULE = "schedule"
    SUBSCRIPTION = "subscription"
    STREAMING = "streaming"

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of TriggerType from a JSON string"""
        return cls(loads(json_str))

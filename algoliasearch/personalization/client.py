# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from json import dumps
from typing import Annotated, Any, Dict, List, Optional, Self, Tuple, Union
from urllib.parse import quote

from pydantic import Field, StrictStr

from algoliasearch.http.api_response import ApiResponse
from algoliasearch.http.request_options import RequestOptions
from algoliasearch.http.serializer import bodySerializer
from algoliasearch.http.transporter import Transporter
from algoliasearch.http.verb import Verb
from algoliasearch.personalization.config import PersonalizationConfig
from algoliasearch.personalization.models.delete_user_profile_response import (
    DeleteUserProfileResponse,
)
from algoliasearch.personalization.models.get_user_token_response import (
    GetUserTokenResponse,
)
from algoliasearch.personalization.models.personalization_strategy_params import (
    PersonalizationStrategyParams,
)
from algoliasearch.personalization.models.set_personalization_strategy_response import (
    SetPersonalizationStrategyResponse,
)


class PersonalizationClient:
    """The Algolia 'PersonalizationClient' class.

    Args:
    app_id (str): The Algolia application ID to retrieve information from.
    api_key (str): The Algolia api key bound to the given `app_id`.
    region ("eu" | "us"): The region of your Algolia application.

    Returns:
    The initialized API client.

    Example:
    _client = PersonalizationClient("YOUR_ALGOLIA_APP_ID", "YOUR_ALGOLIA_API_KEY", region="'eu' or 'us'")
    _client_with_named_args = PersonalizationClient(app_id="YOUR_ALGOLIA_APP_ID", api_key="YOUR_ALGOLIA_API_KEY", region="'eu' or 'us'")

    See `PersonalizationClient.create_with_config` for advanced configuration.
    """

    _transporter: Transporter
    _config: PersonalizationConfig
    _request_options: RequestOptions

    def __init__(
        self,
        app_id: Optional[str] = None,
        api_key: Optional[str] = None,
        region: str = None,
        transporter: Optional[Transporter] = None,
        config: Optional[PersonalizationConfig] = None,
    ) -> None:
        if transporter is not None and config is None:
            config = transporter._config

        if config is None:
            config = PersonalizationConfig(app_id, api_key, region)
        self._config = config
        self._request_options = RequestOptions(config)

        if transporter is None:
            transporter = Transporter(config)
        self._transporter = transporter

    def create_with_config(
        config: PersonalizationConfig, transporter: Optional[Transporter] = None
    ) -> Self:
        """Allows creating a client with a customized `PersonalizationConfig` and `Transporter`. If `transporter` is not provided, the default one will be initialized from the given `config`.

        Args:
        config (PersonalizationConfig): The config of the API client.
        transporter (Transporter): The HTTP transporter, see `http/transporter.py` for implementation details.

        Returns:
        The initialized API client.

        Example:
        _client_with_custom_config = PersonalizationClient.create_with_config(config=PersonalizationConfig(...))
        _client_with_custom_config_and_transporter = PersonalizationClient.create_with_config(config=PersonalizationConfig(...), transporter=Transporter(...))
        """
        if transporter is None:
            transporter = Transporter(config)

        return PersonalizationClient(
            app_id=config.app_id,
            api_key=config.api_key,
            region=config.region,
            transporter=transporter,
            config=config,
        )

    async def __aenter__(self) -> None:
        return self

    async def __aexit__(self, exc_type, exc_value, traceback) -> None:
        """Closes the underlying `transporter` of the API client."""
        await self.close()

    async def close(self) -> None:
        """Closes the underlying `transporter` of the API client."""
        return await self._transporter.close()

    async def custom_delete_with_http_info(
        self,
        path: Annotated[
            StrictStr,
            Field(
                description='Path of the endpoint, anything after "/1" must be specified.'
            ),
        ],
        parameters: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Query parameters to apply to the current query."),
        ] = None,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        Send requests to the Algolia REST API.

        This method allow you to send requests to the Algolia REST API.

        :param path: Path of the endpoint, anything after \"/1\" must be specified. (required)
        :type path: str
        :param parameters: Query parameters to apply to the current query.
        :type parameters: Dict[str, object]
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        if path is None:
            raise ValueError(
                "Parameter `path` is required when calling `custom_delete`."
            )

        _query_parameters: List[Tuple[str, str]] = []

        if parameters is not None:
            for _qpkey, _qpvalue in parameters.items():
                _query_parameters.append((_qpkey, _qpvalue))

        return await self._transporter.request(
            verb=Verb.DELETE,
            path="/1{path}".replace("{path}", path),
            data=None,
            request_options=self._request_options.merge(
                query_parameters=_query_parameters,
                user_request_options=request_options,
            ),
            use_read_transporter=False,
        )

    async def custom_delete(
        self,
        path: Annotated[
            StrictStr,
            Field(
                description='Path of the endpoint, anything after "/1" must be specified.'
            ),
        ],
        parameters: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Query parameters to apply to the current query."),
        ] = None,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> object:
        """
        Send requests to the Algolia REST API.

        This method allow you to send requests to the Algolia REST API.

        :param path: Path of the endpoint, anything after \"/1\" must be specified. (required)
        :type path: str
        :param parameters: Query parameters to apply to the current query.
        :type parameters: Dict[str, object]
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'object' result object.
        """
        return (
            await self.custom_delete_with_http_info(path, parameters, request_options)
        ).deserialize(object)

    async def custom_get_with_http_info(
        self,
        path: Annotated[
            StrictStr,
            Field(
                description='Path of the endpoint, anything after "/1" must be specified.'
            ),
        ],
        parameters: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Query parameters to apply to the current query."),
        ] = None,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        Send requests to the Algolia REST API.

        This method allow you to send requests to the Algolia REST API.

        :param path: Path of the endpoint, anything after \"/1\" must be specified. (required)
        :type path: str
        :param parameters: Query parameters to apply to the current query.
        :type parameters: Dict[str, object]
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        if path is None:
            raise ValueError("Parameter `path` is required when calling `custom_get`.")

        _query_parameters: List[Tuple[str, str]] = []

        if parameters is not None:
            for _qpkey, _qpvalue in parameters.items():
                _query_parameters.append((_qpkey, _qpvalue))

        return await self._transporter.request(
            verb=Verb.GET,
            path="/1{path}".replace("{path}", path),
            data=None,
            request_options=self._request_options.merge(
                query_parameters=_query_parameters,
                user_request_options=request_options,
            ),
            use_read_transporter=False,
        )

    async def custom_get(
        self,
        path: Annotated[
            StrictStr,
            Field(
                description='Path of the endpoint, anything after "/1" must be specified.'
            ),
        ],
        parameters: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Query parameters to apply to the current query."),
        ] = None,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> object:
        """
        Send requests to the Algolia REST API.

        This method allow you to send requests to the Algolia REST API.

        :param path: Path of the endpoint, anything after \"/1\" must be specified. (required)
        :type path: str
        :param parameters: Query parameters to apply to the current query.
        :type parameters: Dict[str, object]
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'object' result object.
        """
        return (
            await self.custom_get_with_http_info(path, parameters, request_options)
        ).deserialize(object)

    async def custom_post_with_http_info(
        self,
        path: Annotated[
            StrictStr,
            Field(
                description='Path of the endpoint, anything after "/1" must be specified.'
            ),
        ],
        parameters: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Query parameters to apply to the current query."),
        ] = None,
        body: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Parameters to send with the custom request."),
        ] = None,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        Send requests to the Algolia REST API.

        This method allow you to send requests to the Algolia REST API.

        :param path: Path of the endpoint, anything after \"/1\" must be specified. (required)
        :type path: str
        :param parameters: Query parameters to apply to the current query.
        :type parameters: Dict[str, object]
        :param body: Parameters to send with the custom request.
        :type body: object
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        if path is None:
            raise ValueError("Parameter `path` is required when calling `custom_post`.")

        _query_parameters: List[Tuple[str, str]] = []

        if parameters is not None:
            for _qpkey, _qpvalue in parameters.items():
                _query_parameters.append((_qpkey, _qpvalue))

        _body = {}
        if body is not None:
            _body = body

        return await self._transporter.request(
            verb=Verb.POST,
            path="/1{path}".replace("{path}", path),
            data=dumps(bodySerializer(_body)),
            request_options=self._request_options.merge(
                query_parameters=_query_parameters,
                user_request_options=request_options,
            ),
            use_read_transporter=False,
        )

    async def custom_post(
        self,
        path: Annotated[
            StrictStr,
            Field(
                description='Path of the endpoint, anything after "/1" must be specified.'
            ),
        ],
        parameters: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Query parameters to apply to the current query."),
        ] = None,
        body: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Parameters to send with the custom request."),
        ] = None,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> object:
        """
        Send requests to the Algolia REST API.

        This method allow you to send requests to the Algolia REST API.

        :param path: Path of the endpoint, anything after \"/1\" must be specified. (required)
        :type path: str
        :param parameters: Query parameters to apply to the current query.
        :type parameters: Dict[str, object]
        :param body: Parameters to send with the custom request.
        :type body: object
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'object' result object.
        """
        return (
            await self.custom_post_with_http_info(
                path, parameters, body, request_options
            )
        ).deserialize(object)

    async def custom_put_with_http_info(
        self,
        path: Annotated[
            StrictStr,
            Field(
                description='Path of the endpoint, anything after "/1" must be specified.'
            ),
        ],
        parameters: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Query parameters to apply to the current query."),
        ] = None,
        body: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Parameters to send with the custom request."),
        ] = None,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        Send requests to the Algolia REST API.

        This method allow you to send requests to the Algolia REST API.

        :param path: Path of the endpoint, anything after \"/1\" must be specified. (required)
        :type path: str
        :param parameters: Query parameters to apply to the current query.
        :type parameters: Dict[str, object]
        :param body: Parameters to send with the custom request.
        :type body: object
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        if path is None:
            raise ValueError("Parameter `path` is required when calling `custom_put`.")

        _query_parameters: List[Tuple[str, str]] = []

        if parameters is not None:
            for _qpkey, _qpvalue in parameters.items():
                _query_parameters.append((_qpkey, _qpvalue))

        _body = {}
        if body is not None:
            _body = body

        return await self._transporter.request(
            verb=Verb.PUT,
            path="/1{path}".replace("{path}", path),
            data=dumps(bodySerializer(_body)),
            request_options=self._request_options.merge(
                query_parameters=_query_parameters,
                user_request_options=request_options,
            ),
            use_read_transporter=False,
        )

    async def custom_put(
        self,
        path: Annotated[
            StrictStr,
            Field(
                description='Path of the endpoint, anything after "/1" must be specified.'
            ),
        ],
        parameters: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Query parameters to apply to the current query."),
        ] = None,
        body: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Parameters to send with the custom request."),
        ] = None,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> object:
        """
        Send requests to the Algolia REST API.

        This method allow you to send requests to the Algolia REST API.

        :param path: Path of the endpoint, anything after \"/1\" must be specified. (required)
        :type path: str
        :param parameters: Query parameters to apply to the current query.
        :type parameters: Dict[str, object]
        :param body: Parameters to send with the custom request.
        :type body: object
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'object' result object.
        """
        return (
            await self.custom_put_with_http_info(
                path, parameters, body, request_options
            )
        ).deserialize(object)

    async def delete_user_profile_with_http_info(
        self,
        user_token: Annotated[
            StrictStr,
            Field(
                description="userToken representing the user for which to fetch the Personalization profile."
            ),
        ],
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        Delete a user profile.

        Delete the user profile and all its associated data.  Returns, as part of the response, a date until which the data can safely be considered as deleted for the given user. This means if you send events for the given user before this date, they will be ignored. Any data received after the deletedUntil date will start building a new user profile.  It might take a couple hours for the deletion request to be fully processed.

        :param user_token: userToken representing the user for which to fetch the Personalization profile. (required)
        :type user_token: str
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        if user_token is None:
            raise ValueError(
                "Parameter `user_token` is required when calling `delete_user_profile`."
            )

        return await self._transporter.request(
            verb=Verb.DELETE,
            path="/1/profiles/{userToken}".replace(
                "{userToken}", quote(str(user_token), safe="")
            ),
            data=None,
            request_options=self._request_options.merge(
                user_request_options=request_options,
            ),
            use_read_transporter=False,
        )

    async def delete_user_profile(
        self,
        user_token: Annotated[
            StrictStr,
            Field(
                description="userToken representing the user for which to fetch the Personalization profile."
            ),
        ],
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> DeleteUserProfileResponse:
        """
        Delete a user profile.

        Delete the user profile and all its associated data.  Returns, as part of the response, a date until which the data can safely be considered as deleted for the given user. This means if you send events for the given user before this date, they will be ignored. Any data received after the deletedUntil date will start building a new user profile.  It might take a couple hours for the deletion request to be fully processed.

        :param user_token: userToken representing the user for which to fetch the Personalization profile. (required)
        :type user_token: str
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'DeleteUserProfileResponse' result object.
        """
        return (
            await self.delete_user_profile_with_http_info(user_token, request_options)
        ).deserialize(DeleteUserProfileResponse)

    async def get_personalization_strategy_with_http_info(
        self, request_options: Optional[Union[dict, RequestOptions]] = None
    ) -> ApiResponse[str]:
        """
        Get the current strategy.

        The strategy contains information on the events and facets that impact user profiles and personalized search results.

        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        return await self._transporter.request(
            verb=Verb.GET,
            path="/1/strategies/personalization",
            data=None,
            request_options=self._request_options.merge(
                user_request_options=request_options,
            ),
            use_read_transporter=False,
        )

    async def get_personalization_strategy(
        self, request_options: Optional[Union[dict, RequestOptions]] = None
    ) -> PersonalizationStrategyParams:
        """
        Get the current strategy.

        The strategy contains information on the events and facets that impact user profiles and personalized search results.

        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'PersonalizationStrategyParams' result object.
        """
        return (
            await self.get_personalization_strategy_with_http_info(request_options)
        ).deserialize(PersonalizationStrategyParams)

    async def get_user_token_profile_with_http_info(
        self,
        user_token: Annotated[
            StrictStr,
            Field(
                description="userToken representing the user for which to fetch the Personalization profile."
            ),
        ],
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        Get a user profile.

        Get the user profile built from Personalization strategy.  The profile is structured by facet name used in the strategy. Each facet value is mapped to its score. Each score represents the user affinity for a specific facet value given the userToken past events and the Personalization strategy defined. Scores are bounded to 20. The last processed event timestamp is provided using the ISO 8601 format for debugging purposes.

        :param user_token: userToken representing the user for which to fetch the Personalization profile. (required)
        :type user_token: str
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        if user_token is None:
            raise ValueError(
                "Parameter `user_token` is required when calling `get_user_token_profile`."
            )

        return await self._transporter.request(
            verb=Verb.GET,
            path="/1/profiles/personalization/{userToken}".replace(
                "{userToken}", quote(str(user_token), safe="")
            ),
            data=None,
            request_options=self._request_options.merge(
                user_request_options=request_options,
            ),
            use_read_transporter=False,
        )

    async def get_user_token_profile(
        self,
        user_token: Annotated[
            StrictStr,
            Field(
                description="userToken representing the user for which to fetch the Personalization profile."
            ),
        ],
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> GetUserTokenResponse:
        """
        Get a user profile.

        Get the user profile built from Personalization strategy.  The profile is structured by facet name used in the strategy. Each facet value is mapped to its score. Each score represents the user affinity for a specific facet value given the userToken past events and the Personalization strategy defined. Scores are bounded to 20. The last processed event timestamp is provided using the ISO 8601 format for debugging purposes.

        :param user_token: userToken representing the user for which to fetch the Personalization profile. (required)
        :type user_token: str
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'GetUserTokenResponse' result object.
        """
        return (
            await self.get_user_token_profile_with_http_info(
                user_token, request_options
            )
        ).deserialize(GetUserTokenResponse)

    async def set_personalization_strategy_with_http_info(
        self,
        personalization_strategy_params: PersonalizationStrategyParams,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        Set a new strategy.

        A strategy defines the events and facets that impact user profiles and personalized search results.

        :param personalization_strategy_params: (required)
        :type personalization_strategy_params: PersonalizationStrategyParams
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        if personalization_strategy_params is None:
            raise ValueError(
                "Parameter `personalization_strategy_params` is required when calling `set_personalization_strategy`."
            )

        _body = {}
        if personalization_strategy_params is not None:
            _body = personalization_strategy_params

        return await self._transporter.request(
            verb=Verb.POST,
            path="/1/strategies/personalization",
            data=dumps(bodySerializer(_body)),
            request_options=self._request_options.merge(
                user_request_options=request_options,
            ),
            use_read_transporter=False,
        )

    async def set_personalization_strategy(
        self,
        personalization_strategy_params: PersonalizationStrategyParams,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> SetPersonalizationStrategyResponse:
        """
        Set a new strategy.

        A strategy defines the events and facets that impact user profiles and personalized search results.

        :param personalization_strategy_params: (required)
        :type personalization_strategy_params: PersonalizationStrategyParams
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'SetPersonalizationStrategyResponse' result object.
        """
        return (
            await self.set_personalization_strategy_with_http_info(
                personalization_strategy_params, request_options
            )
        ).deserialize(SetPersonalizationStrategyResponse)
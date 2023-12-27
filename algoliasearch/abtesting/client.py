# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from json import dumps
from typing import Annotated, Any, Dict, List, Optional, Self, Tuple, Union
from urllib.parse import quote

from pydantic import Field, StrictInt, StrictStr

from algoliasearch.abtesting.config import AbtestingConfig
from algoliasearch.abtesting.models.ab_test import ABTest
from algoliasearch.abtesting.models.ab_test_response import ABTestResponse
from algoliasearch.abtesting.models.add_ab_tests_request import AddABTestsRequest
from algoliasearch.abtesting.models.list_ab_tests_response import ListABTestsResponse
from algoliasearch.http.api_response import ApiResponse
from algoliasearch.http.request_options import RequestOptions
from algoliasearch.http.serializer import bodySerializer
from algoliasearch.http.transporter import Transporter
from algoliasearch.http.verb import Verb


class AbtestingClient:
    _transporter: Transporter
    _config: AbtestingConfig
    _request_options: RequestOptions

    def app_id(self) -> str:
        return self._config.app_id

    def __init__(self, transporter: Transporter, config: AbtestingConfig) -> None:
        self._transporter = transporter
        self._config = config
        self._request_options = RequestOptions(config)

    def create_with_config(config: AbtestingConfig) -> Self:
        transporter = Transporter(config)

        return AbtestingClient(transporter, config)

    def create(
        app_id: Optional[str] = None,
        api_key: Optional[str] = None,
        region: Optional[str] = None,
    ) -> Self:
        return AbtestingClient.create_with_config(
            AbtestingConfig(app_id, api_key, region)
        )

    async def close(self) -> None:
        return await self._transporter.close()

    async def add_ab_tests_with_http_info(
        self,
        add_ab_tests_request: AddABTestsRequest,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        Create an A/B test.

        Creates an A/B test.

        :param add_ab_tests_request: (required)
        :type add_ab_tests_request: AddABTestsRequest
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        if add_ab_tests_request is None:
            raise ValueError(
                "Parameter `add_ab_tests_request` is required when calling `add_ab_tests`."
            )

        _query_parameters: List[Tuple[str, str]] = []
        _headers: Dict[str, Optional[str]] = {}

        _body = {}
        if add_ab_tests_request is not None:
            _body = add_ab_tests_request

        response = await self._transporter.request(
            verb=Verb.POST,
            path="/2/abtests",
            data=dumps(bodySerializer(_body)),
            request_options=self._request_options.merge(
                query_parameters=_query_parameters,
                headers=_headers,
                user_request_options=request_options,
            ),
            use_read_transporter=False,
        )

        response.data = response.raw_data

        return response

    async def add_ab_tests(
        self,
        add_ab_tests_request: AddABTestsRequest,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ABTestResponse:
        """
        Create an A/B test.

        Creates an A/B test.

        :param add_ab_tests_request: (required)
        :type add_ab_tests_request: AddABTestsRequest
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'ABTestResponse' result object.
        """

        response = await self.add_ab_tests_with_http_info(
            add_ab_tests_request, request_options
        )

        return response.deserialize(ABTestResponse)

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
        _headers: Dict[str, Optional[str]] = {}

        if parameters is not None:
            for _qpkey, _qpvalue in parameters.items():
                _query_parameters.append((_qpkey, _qpvalue))

        response = await self._transporter.request(
            verb=Verb.DELETE,
            path="/1{path}".replace("{path}", path),
            data=None,
            request_options=self._request_options.merge(
                query_parameters=_query_parameters,
                headers=_headers,
                user_request_options=request_options,
            ),
            use_read_transporter=False,
        )

        response.data = response.raw_data

        return response

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

        response = await self.custom_delete_with_http_info(
            path, parameters, request_options
        )

        return response.deserialize(object)

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
        _headers: Dict[str, Optional[str]] = {}

        if parameters is not None:
            for _qpkey, _qpvalue in parameters.items():
                _query_parameters.append((_qpkey, _qpvalue))

        response = await self._transporter.request(
            verb=Verb.GET,
            path="/1{path}".replace("{path}", path),
            data=None,
            request_options=self._request_options.merge(
                query_parameters=_query_parameters,
                headers=_headers,
                user_request_options=request_options,
            ),
            use_read_transporter=False,
        )

        response.data = response.raw_data

        return response

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

        response = await self.custom_get_with_http_info(
            path, parameters, request_options
        )

        return response.deserialize(object)

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
        _headers: Dict[str, Optional[str]] = {}

        if parameters is not None:
            for _qpkey, _qpvalue in parameters.items():
                _query_parameters.append((_qpkey, _qpvalue))

        _body = {}
        if body is not None:
            _body = body

        response = await self._transporter.request(
            verb=Verb.POST,
            path="/1{path}".replace("{path}", path),
            data=dumps(bodySerializer(_body)),
            request_options=self._request_options.merge(
                query_parameters=_query_parameters,
                headers=_headers,
                user_request_options=request_options,
            ),
            use_read_transporter=False,
        )

        response.data = response.raw_data

        return response

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

        response = await self.custom_post_with_http_info(
            path, parameters, body, request_options
        )

        return response.deserialize(object)

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
        _headers: Dict[str, Optional[str]] = {}

        if parameters is not None:
            for _qpkey, _qpvalue in parameters.items():
                _query_parameters.append((_qpkey, _qpvalue))

        _body = {}
        if body is not None:
            _body = body

        response = await self._transporter.request(
            verb=Verb.PUT,
            path="/1{path}".replace("{path}", path),
            data=dumps(bodySerializer(_body)),
            request_options=self._request_options.merge(
                query_parameters=_query_parameters,
                headers=_headers,
                user_request_options=request_options,
            ),
            use_read_transporter=False,
        )

        response.data = response.raw_data

        return response

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

        response = await self.custom_put_with_http_info(
            path, parameters, body, request_options
        )

        return response.deserialize(object)

    async def delete_ab_test_with_http_info(
        self,
        id: Annotated[StrictInt, Field(description="Unique A/B test ID.")],
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        Delete an A/B test.

        Delete an A/B test. To determine the `id` for an A/B test, use the [`listABTests` operation](#tag/abtest/operation/listABTests).

        :param id: Unique A/B test ID. (required)
        :type id: int
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        if id is None:
            raise ValueError(
                "Parameter `id` is required when calling `delete_ab_test`."
            )

        _query_parameters: List[Tuple[str, str]] = []
        _headers: Dict[str, Optional[str]] = {}

        response = await self._transporter.request(
            verb=Verb.DELETE,
            path="/2/abtests/{id}".replace("{id}", quote(str(id), safe="")),
            data=None,
            request_options=self._request_options.merge(
                query_parameters=_query_parameters,
                headers=_headers,
                user_request_options=request_options,
            ),
            use_read_transporter=False,
        )

        response.data = response.raw_data

        return response

    async def delete_ab_test(
        self,
        id: Annotated[StrictInt, Field(description="Unique A/B test ID.")],
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ABTestResponse:
        """
        Delete an A/B test.

        Delete an A/B test. To determine the `id` for an A/B test, use the [`listABTests` operation](#tag/abtest/operation/listABTests).

        :param id: Unique A/B test ID. (required)
        :type id: int
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'ABTestResponse' result object.
        """

        response = await self.delete_ab_test_with_http_info(id, request_options)

        return response.deserialize(ABTestResponse)

    async def get_ab_test_with_http_info(
        self,
        id: Annotated[StrictInt, Field(description="Unique A/B test ID.")],
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        Get A/B test details.

        Get specific details for an A/B test. To determine the `id` for an A/B test, use the [`listABTests` operation](#tag/abtest/operation/listABTests).

        :param id: Unique A/B test ID. (required)
        :type id: int
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        if id is None:
            raise ValueError("Parameter `id` is required when calling `get_ab_test`.")

        _query_parameters: List[Tuple[str, str]] = []
        _headers: Dict[str, Optional[str]] = {}

        response = await self._transporter.request(
            verb=Verb.GET,
            path="/2/abtests/{id}".replace("{id}", quote(str(id), safe="")),
            data=None,
            request_options=self._request_options.merge(
                query_parameters=_query_parameters,
                headers=_headers,
                user_request_options=request_options,
            ),
            use_read_transporter=False,
        )

        response.data = response.raw_data

        return response

    async def get_ab_test(
        self,
        id: Annotated[StrictInt, Field(description="Unique A/B test ID.")],
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ABTest:
        """
        Get A/B test details.

        Get specific details for an A/B test. To determine the `id` for an A/B test, use the [`listABTests` operation](#tag/abtest/operation/listABTests).

        :param id: Unique A/B test ID. (required)
        :type id: int
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'ABTest' result object.
        """

        response = await self.get_ab_test_with_http_info(id, request_options)

        return response.deserialize(ABTest)

    async def list_ab_tests_with_http_info(
        self,
        offset: Annotated[
            Optional[StrictInt],
            Field(
                description="Position of the starting record. Used for paging. 0 is the first record."
            ),
        ] = None,
        limit: Annotated[
            Optional[StrictInt],
            Field(description="Number of records to return (page size)."),
        ] = None,
        index_prefix: Annotated[
            Optional[StrictStr],
            Field(
                description="Only return A/B tests for indices starting with this prefix."
            ),
        ] = None,
        index_suffix: Annotated[
            Optional[StrictStr],
            Field(
                description="Only return A/B tests for indices ending with this suffix."
            ),
        ] = None,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        List all A/B tests.

        List all A/B tests.

        :param offset: Position of the starting record. Used for paging. 0 is the first record.
        :type offset: int
        :param limit: Number of records to return (page size).
        :type limit: int
        :param index_prefix: Only return A/B tests for indices starting with this prefix.
        :type index_prefix: str
        :param index_suffix: Only return A/B tests for indices ending with this suffix.
        :type index_suffix: str
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        _query_parameters: List[Tuple[str, str]] = []
        _headers: Dict[str, Optional[str]] = {}

        if offset is not None:
            _query_parameters.append(("offset", offset))
        if limit is not None:
            _query_parameters.append(("limit", limit))
        if index_prefix is not None:
            _query_parameters.append(("indexPrefix", index_prefix))
        if index_suffix is not None:
            _query_parameters.append(("indexSuffix", index_suffix))

        response = await self._transporter.request(
            verb=Verb.GET,
            path="/2/abtests",
            data=None,
            request_options=self._request_options.merge(
                query_parameters=_query_parameters,
                headers=_headers,
                user_request_options=request_options,
            ),
            use_read_transporter=False,
        )

        response.data = response.raw_data

        return response

    async def list_ab_tests(
        self,
        offset: Annotated[
            Optional[StrictInt],
            Field(
                description="Position of the starting record. Used for paging. 0 is the first record."
            ),
        ] = None,
        limit: Annotated[
            Optional[StrictInt],
            Field(description="Number of records to return (page size)."),
        ] = None,
        index_prefix: Annotated[
            Optional[StrictStr],
            Field(
                description="Only return A/B tests for indices starting with this prefix."
            ),
        ] = None,
        index_suffix: Annotated[
            Optional[StrictStr],
            Field(
                description="Only return A/B tests for indices ending with this suffix."
            ),
        ] = None,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ListABTestsResponse:
        """
        List all A/B tests.

        List all A/B tests.

        :param offset: Position of the starting record. Used for paging. 0 is the first record.
        :type offset: int
        :param limit: Number of records to return (page size).
        :type limit: int
        :param index_prefix: Only return A/B tests for indices starting with this prefix.
        :type index_prefix: str
        :param index_suffix: Only return A/B tests for indices ending with this suffix.
        :type index_suffix: str
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'ListABTestsResponse' result object.
        """

        response = await self.list_ab_tests_with_http_info(
            offset, limit, index_prefix, index_suffix, request_options
        )

        return response.deserialize(ListABTestsResponse)

    async def stop_ab_test_with_http_info(
        self,
        id: Annotated[StrictInt, Field(description="Unique A/B test ID.")],
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        Stop an A/B test.

        If stopped, the test is over and can't be restarted. There is now only one index, receiving 100% of all search requests. The data gathered for stopped A/B tests is retained. To determine the `id` for an A/B test, use the [`listABTests` operation](#tag/abtest/operation/listABTests).

        :param id: Unique A/B test ID. (required)
        :type id: int
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        if id is None:
            raise ValueError("Parameter `id` is required when calling `stop_ab_test`.")

        _query_parameters: List[Tuple[str, str]] = []
        _headers: Dict[str, Optional[str]] = {}

        response = await self._transporter.request(
            verb=Verb.POST,
            path="/2/abtests/{id}/stop".replace("{id}", quote(str(id), safe="")),
            data=None,
            request_options=self._request_options.merge(
                query_parameters=_query_parameters,
                headers=_headers,
                user_request_options=request_options,
            ),
            use_read_transporter=False,
        )

        response.data = response.raw_data

        return response

    async def stop_ab_test(
        self,
        id: Annotated[StrictInt, Field(description="Unique A/B test ID.")],
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ABTestResponse:
        """
        Stop an A/B test.

        If stopped, the test is over and can't be restarted. There is now only one index, receiving 100% of all search requests. The data gathered for stopped A/B tests is retained. To determine the `id` for an A/B test, use the [`listABTests` operation](#tag/abtest/operation/listABTests).

        :param id: Unique A/B test ID. (required)
        :type id: int
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'ABTestResponse' result object.
        """

        response = await self.stop_ab_test_with_http_info(id, request_options)

        return response.deserialize(ABTestResponse)

<?php

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

namespace Algolia\AlgoliaSearch\Model\Search;

/**
 * Acl Class Doc Comment.
 *
 * @category Class
 *
 * @description API key permissions:  &#x60;addObject&#x60;: required to add or update records, copy or move an index. &#x60;analytics&#x60;: required to access the Analytics API. &#x60;browse&#x60;: required to view records &#x60;deleteIndex&#x60;: required to delete indices. &#x60;deleteObject&#x60;: required to delete records. &#x60;editSettings&#x60;: required to change index settings. &#x60;inference&#x60;: required to access the Inference API. &#x60;listIndexes&#x60;: required to list indices. &#x60;logs&#x60;: required to access logs of search and indexing operations. &#x60;recommendation&#x60;: required to access the Personalization and Recommend APIs. &#x60;search&#x60;: required to search records &#x60;seeUnretrievableAttributes&#x60;: required to retrieve [&#x60;unretrievableAttributes&#x60;](https://www.algolia.com/doc/api-reference/api-parameters/unretrievableAttributes/) for all operations that return records. &#x60;settings&#x60;: required to examine index settings.
 */
class Acl
{
    /**
     * Possible values of this enum.
     */
    public const ADD_OBJECT = 'addObject';

    public const ANALYTICS = 'analytics';

    public const BROWSE = 'browse';

    public const DELETE_OBJECT = 'deleteObject';

    public const DELETE_INDEX = 'deleteIndex';

    public const EDIT_SETTINGS = 'editSettings';

    public const INFERENCE = 'inference';

    public const LIST_INDEXES = 'listIndexes';

    public const LOGS = 'logs';

    public const PERSONALIZATION = 'personalization';

    public const RECOMMENDATION = 'recommendation';

    public const SEARCH = 'search';

    public const SEE_UNRETRIEVABLE_ATTRIBUTES = 'seeUnretrievableAttributes';

    public const SETTINGS = 'settings';

    public const USAGE = 'usage';

    /**
     * Gets allowable values of the enum.
     *
     * @return string[]
     */
    public static function getAllowableEnumValues()
    {
        return [
            self::ADD_OBJECT,
            self::ANALYTICS,
            self::BROWSE,
            self::DELETE_OBJECT,
            self::DELETE_INDEX,
            self::EDIT_SETTINGS,
            self::INFERENCE,
            self::LIST_INDEXES,
            self::LOGS,
            self::PERSONALIZATION,
            self::RECOMMENDATION,
            self::SEARCH,
            self::SEE_UNRETRIEVABLE_ATTRIBUTES,
            self::SETTINGS,
            self::USAGE,
        ];
    }
}

<?php

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

namespace Algolia\AlgoliaSearch\Model\Search;

/**
 * ScopeType Class Doc Comment.
 *
 * @category Class
 */
class ScopeType
{
    /**
     * Possible values of this enum.
     */
    public const SETTINGS = 'settings';

    public const SYNONYMS = 'synonyms';

    public const RULES = 'rules';

    /**
     * Gets allowable values of the enum.
     *
     * @return string[]
     */
    public static function getAllowableEnumValues()
    {
        return [
            self::SETTINGS,
            self::SYNONYMS,
            self::RULES,
        ];
    }
}

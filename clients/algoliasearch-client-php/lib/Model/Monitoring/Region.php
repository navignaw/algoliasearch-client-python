<?php

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

namespace Algolia\AlgoliaSearch\Model\Monitoring;

/**
 * Region Class Doc Comment.
 *
 * @category Class
 *
 * @description The region where the cluster is located.
 */
class Region
{
    /**
     * Possible values of this enum.
     */
    public const AU = 'au';

    public const BR = 'br';

    public const CA = 'ca';

    public const DE = 'de';

    public const EU = 'eu';

    public const HK = 'hk';

    public const IN = 'in';

    public const JP = 'jp';

    public const SG = 'sg';

    public const UAE = 'uae';

    public const UK = 'uk';

    public const USC = 'usc';

    public const _USE = 'use';

    public const USW = 'usw';

    public const ZA = 'za';

    /**
     * Gets allowable values of the enum.
     *
     * @return string[]
     */
    public static function getAllowableEnumValues()
    {
        return [
            self::AU,
            self::BR,
            self::CA,
            self::DE,
            self::EU,
            self::HK,
            self::IN,
            self::JP,
            self::SG,
            self::UAE,
            self::UK,
            self::USC,
            self::_USE,
            self::USW,
            self::ZA,
        ];
    }
}

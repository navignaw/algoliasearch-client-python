<?php

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

namespace Algolia\AlgoliaSearch\Model\QuerySuggestions;

/**
 * GetLogFile200Response Class Doc Comment
 *
 * @category Class
 * @package Algolia\AlgoliaSearch
 */
class GetLogFile200Response extends \Algolia\AlgoliaSearch\Model\AbstractModel implements ModelInterface, \ArrayAccess, \JsonSerializable
{
    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $modelTypes = [
        'timestamp' => 'string',
        'level' => '\Algolia\AlgoliaSearch\Model\QuerySuggestions\LogLevel',
        'message' => 'string',
        'contextLevel' => 'int',
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $modelFormats = [
        'timestamp' => null,
        'level' => null,
        'message' => null,
        'contextLevel' => null,
    ];

    /**
      * Array of attributes where the key is the local name,
      * and the value is the original name
      *
      * @var string[]
    */
    protected static $attributeMap = [
        'timestamp' => 'timestamp',
        'level' => 'level',
        'message' => 'message',
        'contextLevel' => 'contextLevel',
    ];

    /**
      * Array of attributes where the key is the local name,
      * and the value is the original name
      *
      * @return array
      */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function modelTypes()
    {
        return self::$modelTypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function modelFormats()
    {
        return self::$modelFormats;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'timestamp' => 'setTimestamp',
        'level' => 'setLevel',
        'message' => 'setMessage',
        'contextLevel' => 'setContextLevel',
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'timestamp' => 'getTimestamp',
        'level' => 'getLevel',
        'message' => 'getMessage',
        'contextLevel' => 'getContextLevel',
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     */
    public function __construct(array $data = null)
    {
        if (isset($data['timestamp'])) {
            $this->container['timestamp'] = $data['timestamp'];
        }
        if (isset($data['level'])) {
            $this->container['level'] = $data['level'];
        }
        if (isset($data['message'])) {
            $this->container['message'] = $data['message'];
        }
        if (isset($data['contextLevel'])) {
            $this->container['contextLevel'] = $data['contextLevel'];
        }
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }

    /**
     * Gets timestamp
     *
     * @return string|null
     */
    public function getTimestamp()
    {
        return $this->container['timestamp'] ?? null;
    }

    /**
     * Sets timestamp
     *
     * @param string|null $timestamp Timestamp in [ISO-8601](https://wikipedia.org/wiki/ISO_8601) format.
     *
     * @return self
     */
    public function setTimestamp($timestamp)
    {
        $this->container['timestamp'] = $timestamp;

        return $this;
    }

    /**
     * Gets level
     *
     * @return \Algolia\AlgoliaSearch\Model\QuerySuggestions\LogLevel|null
     */
    public function getLevel()
    {
        return $this->container['level'] ?? null;
    }

    /**
     * Sets level
     *
     * @param \Algolia\AlgoliaSearch\Model\QuerySuggestions\LogLevel|null $level level
     *
     * @return self
     */
    public function setLevel($level)
    {
        $this->container['level'] = $level;

        return $this;
    }

    /**
     * Gets message
     *
     * @return string|null
     */
    public function getMessage()
    {
        return $this->container['message'] ?? null;
    }

    /**
     * Sets message
     *
     * @param string|null $message details about this log entry
     *
     * @return self
     */
    public function setMessage($message)
    {
        $this->container['message'] = $message;

        return $this;
    }

    /**
     * Gets contextLevel
     *
     * @return int|null
     */
    public function getContextLevel()
    {
        return $this->container['contextLevel'] ?? null;
    }

    /**
     * Sets contextLevel
     *
     * @param int|null $contextLevel Level indicating the position of a suggestion in a hierarchy of records.   For example, a `contextLevel` of 1 indicates that this suggestion belongs to a previous suggestion with `contextLevel` 0.
     *
     * @return self
     */
    public function setContextLevel($contextLevel)
    {
        $this->container['contextLevel'] = $contextLevel;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param int $offset Offset
     *
     * @return bool
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param int $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param int $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }
}

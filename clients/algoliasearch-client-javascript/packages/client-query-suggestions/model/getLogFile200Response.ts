// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import type { LogLevel } from './logLevel';

export type GetLogFile200Response = {
  /**
   * Timestamp in [ISO-8601](https://wikipedia.org/wiki/ISO_8601) format.
   */
  timestamp?: string;

  level?: LogLevel;

  /**
   * Details about this log entry.
   */
  message?: string;

  /**
   * Level indicating the position of a suggestion in a hierarchy of records.   For example, a `contextLevel` of 1 indicates that this suggestion belongs to a previous suggestion with `contextLevel` 0.
   */
  contextLevel?: number;
};
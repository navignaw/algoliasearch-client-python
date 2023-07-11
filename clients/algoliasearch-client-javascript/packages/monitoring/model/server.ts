// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import type { Region } from './region';
import type { ServerStatus } from './serverStatus';
import type { Type } from './type';

export type Server = {
  /**
   * Server name.
   */
  name?: string;

  region?: Region;

  /**
   * Included to support legacy applications. Do not rely on this attribute being present in the response. Use `is_replica` instead.
   */
  is_slave?: boolean;

  /**
   * Indicates whether this server is a replica of another server.
   */
  is_replica?: boolean;

  /**
   * Name of the cluster to which this server belongs.
   */
  cluster?: string;

  status?: ServerStatus;

  type?: Type;
};

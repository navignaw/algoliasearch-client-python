// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

/**
 * Whether the pattern parameter matches the beginning (`startsWith`) or end (`endsWith`) of the
 * query string, is an exact match (`is`), or a partial match (`contains`).
 */
public enum Anchoring {
  IS("is"),

  STARTS_WITH("startsWith"),

  ENDS_WITH("endsWith"),

  CONTAINS("contains");

  private final String value;

  Anchoring(String value) {
    this.value = value;
  }

  @JsonValue
  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  @JsonCreator
  public static Anchoring fromValue(String value) {
    for (Anchoring b : Anchoring.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }
}

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.recommend;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** Condition */
public class Condition {

  @JsonProperty("pattern")
  private String pattern;

  @JsonProperty("anchoring")
  private Anchoring anchoring;

  @JsonProperty("alternatives")
  private Boolean alternatives;

  @JsonProperty("context")
  private String context;

  public Condition setPattern(String pattern) {
    this.pattern = pattern;
    return this;
  }

  /** Query pattern syntax. */
  @javax.annotation.Nullable
  public String getPattern() {
    return pattern;
  }

  public Condition setAnchoring(Anchoring anchoring) {
    this.anchoring = anchoring;
    return this;
  }

  /** Get anchoring */
  @javax.annotation.Nullable
  public Anchoring getAnchoring() {
    return anchoring;
  }

  public Condition setAlternatives(Boolean alternatives) {
    this.alternatives = alternatives;
    return this;
  }

  /** Whether the pattern matches on plurals, synonyms, and typos. */
  @javax.annotation.Nullable
  public Boolean getAlternatives() {
    return alternatives;
  }

  public Condition setContext(String context) {
    this.context = context;
    return this;
  }

  /** Rule context format: [A-Za-z0-9_-]+). */
  @javax.annotation.Nullable
  public String getContext() {
    return context;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Condition condition = (Condition) o;
    return (
      Objects.equals(this.pattern, condition.pattern) &&
      Objects.equals(this.anchoring, condition.anchoring) &&
      Objects.equals(this.alternatives, condition.alternatives) &&
      Objects.equals(this.context, condition.context)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(pattern, anchoring, alternatives, context);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Condition {\n");
    sb.append("    pattern: ").append(toIndentedString(pattern)).append("\n");
    sb.append("    anchoring: ").append(toIndentedString(anchoring)).append("\n");
    sb.append("    alternatives: ").append(toIndentedString(alternatives)).append("\n");
    sb.append("    context: ").append(toIndentedString(context)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
}

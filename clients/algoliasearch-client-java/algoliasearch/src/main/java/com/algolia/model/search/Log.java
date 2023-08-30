// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** Log */
public class Log {

  @JsonProperty("timestamp")
  private String timestamp;

  @JsonProperty("method")
  private String method;

  @JsonProperty("answer_code")
  private String answerCode;

  @JsonProperty("query_body")
  private String queryBody;

  @JsonProperty("answer")
  private String answer;

  @JsonProperty("url")
  private String url;

  @JsonProperty("ip")
  private String ip;

  @JsonProperty("query_headers")
  private String queryHeaders;

  @JsonProperty("sha1")
  private String sha1;

  @JsonProperty("nb_api_calls")
  private String nbApiCalls;

  @JsonProperty("processing_time_ms")
  private String processingTimeMs;

  @JsonProperty("index")
  private String index;

  @JsonProperty("query_params")
  private String queryParams;

  @JsonProperty("query_nb_hits")
  private String queryNbHits;

  @JsonProperty("inner_queries")
  private List<LogQuery> innerQueries;

  public Log setTimestamp(String timestamp) {
    this.timestamp = timestamp;
    return this;
  }

  /** Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format. */
  @javax.annotation.Nonnull
  public String getTimestamp() {
    return timestamp;
  }

  public Log setMethod(String method) {
    this.method = method;
    return this;
  }

  /** HTTP method of the performed request. */
  @javax.annotation.Nonnull
  public String getMethod() {
    return method;
  }

  public Log setAnswerCode(String answerCode) {
    this.answerCode = answerCode;
    return this;
  }

  /** HTTP response code. */
  @javax.annotation.Nonnull
  public String getAnswerCode() {
    return answerCode;
  }

  public Log setQueryBody(String queryBody) {
    this.queryBody = queryBody;
    return this;
  }

  /** Request body. Truncated after 1,000 characters. */
  @javax.annotation.Nonnull
  public String getQueryBody() {
    return queryBody;
  }

  public Log setAnswer(String answer) {
    this.answer = answer;
    return this;
  }

  /** Answer body. Truncated after 1,000 characters. */
  @javax.annotation.Nonnull
  public String getAnswer() {
    return answer;
  }

  public Log setUrl(String url) {
    this.url = url;
    return this;
  }

  /** Request URL. */
  @javax.annotation.Nonnull
  public String getUrl() {
    return url;
  }

  public Log setIp(String ip) {
    this.ip = ip;
    return this;
  }

  /** IP address of the client that performed the request. */
  @javax.annotation.Nonnull
  public String getIp() {
    return ip;
  }

  public Log setQueryHeaders(String queryHeaders) {
    this.queryHeaders = queryHeaders;
    return this;
  }

  /** Request headers (API key is obfuscated). */
  @javax.annotation.Nonnull
  public String getQueryHeaders() {
    return queryHeaders;
  }

  public Log setSha1(String sha1) {
    this.sha1 = sha1;
    return this;
  }

  /** SHA1 signature of the log entry. */
  @javax.annotation.Nonnull
  public String getSha1() {
    return sha1;
  }

  public Log setNbApiCalls(String nbApiCalls) {
    this.nbApiCalls = nbApiCalls;
    return this;
  }

  /** Number of API calls. */
  @javax.annotation.Nonnull
  public String getNbApiCalls() {
    return nbApiCalls;
  }

  public Log setProcessingTimeMs(String processingTimeMs) {
    this.processingTimeMs = processingTimeMs;
    return this;
  }

  /** Processing time for the query. Doesn't include network time. */
  @javax.annotation.Nonnull
  public String getProcessingTimeMs() {
    return processingTimeMs;
  }

  public Log setIndex(String index) {
    this.index = index;
    return this;
  }

  /** Index targeted by the query. */
  @javax.annotation.Nullable
  public String getIndex() {
    return index;
  }

  public Log setQueryParams(String queryParams) {
    this.queryParams = queryParams;
    return this;
  }

  /** Query parameters sent with the request. */
  @javax.annotation.Nullable
  public String getQueryParams() {
    return queryParams;
  }

  public Log setQueryNbHits(String queryNbHits) {
    this.queryNbHits = queryNbHits;
    return this;
  }

  /** Number of hits returned for the query. */
  @javax.annotation.Nullable
  public String getQueryNbHits() {
    return queryNbHits;
  }

  public Log setInnerQueries(List<LogQuery> innerQueries) {
    this.innerQueries = innerQueries;
    return this;
  }

  public Log addInnerQueries(LogQuery innerQueriesItem) {
    if (this.innerQueries == null) {
      this.innerQueries = new ArrayList<>();
    }
    this.innerQueries.add(innerQueriesItem);
    return this;
  }

  /** Performed queries for the given request. */
  @javax.annotation.Nullable
  public List<LogQuery> getInnerQueries() {
    return innerQueries;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Log log = (Log) o;
    return (
      Objects.equals(this.timestamp, log.timestamp) &&
      Objects.equals(this.method, log.method) &&
      Objects.equals(this.answerCode, log.answerCode) &&
      Objects.equals(this.queryBody, log.queryBody) &&
      Objects.equals(this.answer, log.answer) &&
      Objects.equals(this.url, log.url) &&
      Objects.equals(this.ip, log.ip) &&
      Objects.equals(this.queryHeaders, log.queryHeaders) &&
      Objects.equals(this.sha1, log.sha1) &&
      Objects.equals(this.nbApiCalls, log.nbApiCalls) &&
      Objects.equals(this.processingTimeMs, log.processingTimeMs) &&
      Objects.equals(this.index, log.index) &&
      Objects.equals(this.queryParams, log.queryParams) &&
      Objects.equals(this.queryNbHits, log.queryNbHits) &&
      Objects.equals(this.innerQueries, log.innerQueries)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(
      timestamp,
      method,
      answerCode,
      queryBody,
      answer,
      url,
      ip,
      queryHeaders,
      sha1,
      nbApiCalls,
      processingTimeMs,
      index,
      queryParams,
      queryNbHits,
      innerQueries
    );
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Log {\n");
    sb.append("    timestamp: ").append(toIndentedString(timestamp)).append("\n");
    sb.append("    method: ").append(toIndentedString(method)).append("\n");
    sb.append("    answerCode: ").append(toIndentedString(answerCode)).append("\n");
    sb.append("    queryBody: ").append(toIndentedString(queryBody)).append("\n");
    sb.append("    answer: ").append(toIndentedString(answer)).append("\n");
    sb.append("    url: ").append(toIndentedString(url)).append("\n");
    sb.append("    ip: ").append(toIndentedString(ip)).append("\n");
    sb.append("    queryHeaders: ").append(toIndentedString(queryHeaders)).append("\n");
    sb.append("    sha1: ").append(toIndentedString(sha1)).append("\n");
    sb.append("    nbApiCalls: ").append(toIndentedString(nbApiCalls)).append("\n");
    sb.append("    processingTimeMs: ").append(toIndentedString(processingTimeMs)).append("\n");
    sb.append("    index: ").append(toIndentedString(index)).append("\n");
    sb.append("    queryParams: ").append(toIndentedString(queryParams)).append("\n");
    sb.append("    queryNbHits: ").append(toIndentedString(queryNbHits)).append("\n");
    sb.append("    innerQueries: ").append(toIndentedString(innerQueries)).append("\n");
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
